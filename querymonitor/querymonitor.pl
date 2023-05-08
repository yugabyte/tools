#!/usr/bin/perl

our $VERSION = "1.01";
my $HELP_TEXT = << "__HELPTEXT__";
#    querymonitor.pl  Version $VERSION
#    ===============
# Monitor running queries
# collect gzipped JSON file for offline analysis

__HELPTEXT__
use strict;
use warnings;
use Getopt::Long;
use Fcntl qw(:DEFAULT :flock);
use POSIX qw/setsid/;
use HTTP::Tiny;
{# Forward local-package declarations
  package Analysis::Mode;
  package Web::Interface; # Handles communication with YBA API 
  package MIME::Write::Simple;
  package MIME::Multipart::ParseSimple;
  package JSON::Tiny;
}

my %opt=(
    API_TOKEN                      => $ENV{API_TOKEN},
    YBA_HOST                       => $ENV{YBA_HOST},
    CUST_UUID                      => $ENV{CUST_UUID},
    UNIV_UUID                      => $ENV{UNIV_UUID},

    # Operating variables
    STARTTIME	                  => unixtime_to_printable(time(),"YYYY-MM-DD HH:MM"),
	STARTTIME_TZ                  => unixtime_to_printable(time(),"YYYY-MM-DD HH:MM","Include tz offset"),
    INTERVAL_SEC                  => 5,             # You can put fractions of a second here
    RUN_FOR                       => "4h",  # 4 hours
    ENDTIME_EPOCH                 => 0, # Calculated after options are processed
    CURL                          => "curl",
    FLAGFILE                      => "querymonitor.defaultflags",
	OUTPUT                        => "queries." . unixtime_to_printable(time(),"YYYY-MM-DD") . ".mime.gz",
    # Misc
    DEBUG                         => 0,
    HELP                          => 0,
    DAEMON                        => 1,
    LOCKFILE                      => "/var/lock/querymonitor.lock", # UNIV_UUID will be appended
    LOCK_FH                      => undef,
	MAX_QUERY_LEN                => 2048,
	MAX_ERRORS                   => 10,
	SANITIZE					 => 0,   # Remove PII 
	ANALYZE						 => undef,     # Input File-name ( "..csv.gz" ) to process through sqlite
	DB                           => undef,     # output SQLITE database file name
	SQLITE                       => "sqlite3", # path to Sqlite binary
	UNIVERSE                     => undef,     # Universe detail info (Populated in initialize)
	HTTPCONNECT                  => "tiny",    # How to connect to the YBA : "curl", or "tiny" (HTTP::Tiny)
	USETESTDATA                  => undef,     # TESTING only !!
	TZOFFSET                     => do{my $tz = (localtime time)[8] * 60 - POSIX::mktime(gmtime 0) / 60;
                                      	sprintf "%+03d:%02d", $tz / 60, abs($tz) % 60},
);
my $quit_daemon = 0;
my $loop_count  = 0;
my $error_counter=0;
my $YBA_API;  # Populated in `Initialize` to a Web::Interface object
my $output;   # Populated in `Initialize to a MIME::Write::Simple object

Initialize();

if ($opt{ANALYZE}){
   my $anl = Analysis::Mode::->new();
   $anl->Process_file_and_create_Sqlite();
   exit 0;
}

daemonize();

#------------- M a i n    L o o p --------------------------
while (not ($quit_daemon  or  time() > $opt{ENDTIME_EPOCH} )){  # Infinite loop ...
   $loop_count++;
   Main_loop_Iteration();
   sleep($opt{INTERVAL_SEC});
}
#------------- E n d  M a i n    L o o p ---------------
# Could get here if a SIGNAL is received
warn(unixtime_to_printable(time(),"YYYY-MM-DD HH:MM:SS") ." Program $$ Completed after $loop_count iterations.\n"); 

$opt{LOCK_FH} and close $opt{LOCK_FH} ;  # Should already be closed and removed by sig handler
unlink $opt{LOCKFILE};
$output->Close();

exit 0;
#------------------------------------------------------------------------------
#==============================================================================
#------------------------------------------------------------------------------
sub Main_loop_Iteration{

    my $query_type = "Unknown";

    if ($opt{DEBUG}){
        print "DEBUG: Start main loop iteration $loop_count\n";
    }
    my $queries = $YBA_API->Get("/live_queries");
    my $ts = time();
	if ($queries->{error}){
		$error_counter++ >= $opt{MAX_ERRORS} and $quit_daemon = 1;
	    return $queries->{error};	
	}
    for my $type (qw|ysql ycql|){
       for my $q (@{ $queries->{$type}{queries} }){
		   for my $subquery (split /;/, $q->{query}){
			 #Sanitize PII from each
			 my $sanitized_query = $opt{SANITIZE} ? SQL_Sanitize($subquery) : $subquery;
             #print join(",",$ts, $type,$q->{nodeName},$q->{id},$subquery),"\n";
			 $output->WriteQuery($ts, $type, $q, $sanitized_query);
		   }
       }
    }

  return 0;
}
#------------------------------------------------------------------------------
sub SQL_Sanitize{ # Remove PII 
   my ($q) = @_;
    # Remove leading spaces
   $q=~s/^\s+//;
    # For INSERT, remove VALUES
	# For SELECT, remove WHERE clause
   $q=~s/ values .+//gi;
   $q=~s/( where\s+\S[^=<>!"'\s]+).*/$1/i; # Include Var name after WHERE, but not value 
   return $q;
}
#------------------------------------------------------------------------------


#------------------------------------------------------------------------------
sub Initialize{
  chomp ($opt{HOSTNAME} = qx|hostname|);
  print $opt{STARTTIME_TZ}," Starting $0 version $VERSION  PID $$ on $opt{HOSTNAME}\n";

  my @program_options = qw[ DEBUG! HELP! DAEMON! SANITIZE!
                       API_TOKEN=s YBA_HOST=s CUST_UUID=s UNIV_UUID=s
                       INTERVAL_SEC=i RUN_FOR|RUNFOR=s CURL=s SQLITE=s
                       FLAGFILE=s OUTPUT=s DB=s
					   MAX_QUERY_LEN|MAXQL|MQL=i ANALYZE|PROCESS=s
					   HTTPCONNECT=s USETESTDATA!];
  my %flags_used;
  Getopt::Long::GetOptions (\%flags_used, @program_options)
      or die "ERROR: Bad Option\n";
  $opt{$_} = $flags_used{$_} for keys %flags_used; # Apply cmd-line flags immediately
  if ($opt{HELP}){
    print $HELP_TEXT;
    print "Program Options:\n",
         map {my $x=$_; $x=~s/\W.*//g; "\t--$x\n"} sort @program_options;
    exit 1;
  }
  if (@ARGV){
    die "ERROR: Unknown argument (flag expected) : @ARGV";	  
  }
  if (-f $opt{FLAGFILE}){
    $opt{DEBUG} and print "DEBUG: Reading Flagfile $opt{FLAGFILE}\n";
    open my $ff, "<", $opt{FLAGFILE} or die "ERROR: Could not open $opt{FLAGFILE}:$!";
    chomp (my @flag_options = grep !m/^\s*#/, <$ff>);
    close $ff;
    $opt{DEBUG} and print "DEBUG: Flagfile option:'$_'\n;" for @flag_options;
    my %flagfile_option_value;
    Getopt::Long::GetOptionsFromArray(\@flag_options,\%flagfile_option_value, @program_options)
        or die "ERROR: Bad FlagFile $opt{FLAGFILE} Option\n";
    for my $k (keys %flagfile_option_value){
        if (exists $flags_used{$k}){  # Cmd line overrides flagfile
            $opt{DEBUG} and print "DEBUG: Flagfile option $k=$flagfile_option_value{$k} IGNORED.. overridden by cmd line.\n";
        }elsif ($k eq "FLAGFILE"){
            die "ERROR: Nested flag files are not allowed."; 
        }else{
            $opt{DEBUG} and print "DEBUG: Flagfile option $k=$flagfile_option_value{$k} set.\n";
            $opt{$k} = $flagfile_option_value{$k};
        }
    }
  }

  my ($run_digits,$run_unit) = $opt{RUN_FOR} =~m/^(\d+)([dhms]?)$/i;
  $run_digits or die "ERROR:'RUN_FOR' option incorrectly specified($opt{RUN_FOR}). Use: 1d 3h 30s or just a number of seconds";
  $run_unit ||= "s"; # Default to seconds  
  my %unit_idx= (d=>24*3600,  m => 60 , h => 3600 ,s => 1); 
  $unit_idx{uc $_} = $unit_idx{$_} for keys %unit_idx;
  $opt{ENDTIME_EPOCH} = time() + $run_digits * $unit_idx{$run_unit};

  if ($opt{ANALYZE}){
	  return; # no more initialization needed   
  }

  $YBA_API = Web::Interface::->new();

  # Get universe name ..
  $opt{UNIVERSE} = $YBA_API->Get("");

  $opt{DEBUG} and print "DEBUG: $_\t","=>",$opt{UNIVERSE}{$_},"\n" for sort keys %{$opt{UNIVERSE}};
  #my ($universe_name) =  $json_string =~m/,"name":"([^"]+)"/;
  if ($opt{UNIVERSE}{name}){
	 print "UNIVERSE: ", $opt{UNIVERSE}{name}," on ", $opt{UNIVERSE}{universeDetails}{clusters}[0]{userIntent}{providerType}, " ver ",$opt{UNIVERSE}{universeDetails}{clusters}[0]{userIntent}{ybSoftwareVersion},"\n";
  }else{
     print "WARNING: Universe info not found \n";
  }
  $opt{NODEHASH} = {map{$_->{nodeName} => {%{$_->{cloudInfo}}, uuid=>$_->{nodeUuid} ,state=>$_->{state},isTserver=>$_->{isTserver}}} 
						@{ $opt{UNIVERSE}->{universeDetails}{nodeDetailsSet} } };

  $output = MIME::Write::Simple::->new(UNIVERSE_JSON=>$YBA_API->{json_string}); # No I/O so far 
  if ($opt{USETESTDATA}){
	 print "Writing ONE record of test data, then exiting...\n";
	 $output->WriteQuery(time(), "ycql", {FIVE=>5,SIX=>6,COW=>"Moo"},"SELECT some_junk FROM made_up");
	 $output->Close();
	 exit 2;
  }  
  # Run iteration ONCE, to verify it works...
  return unless $opt{DAEMON} ;
  
  $opt{LOCKFILE} .= ".$opt{UNIV_UUID}"; # one lock per universe 
  print "Testing main loop before daemonizing...\n";

  if (Main_loop_Iteration()){
	  print "ERROR in main loop iteration. quitting...\n";
     exit 2;
  }
  # Close open file handles that may be leftover from main loop outputs
  $output->Close();
  print "End main loop test.\n";  
}
#------------------------------------------------------------------------------
#------------------------------------------------------------------------------
sub daemonize {
    if ( ! $opt{DAEMON}){
      # Non-daemon mode -do this ONE time only ...
      warn "NOTE: Non-daemon mode is intended for TEST/DEBUG only\n";
      return $$;
    }
    my $grandchild_output = "nohup.out";
    sysopen $opt{LOCK_FH}, $opt{LOCKFILE}, O_EXCL | O_RDWR | O_CREAT | O_NONBLOCK
       or die "ERROR (Fatal):$0 is already running: Lockfile $opt{LOCKFILE}:$!";
     # Handle Signals ..
    sub Signal_Handler {   # 1st argument is signal name
        my($sig) = @_;
        #$sig ||= $! ||= "*Unknown*"; # Could  happen if sig within a sig ..
        my $msg =  "Caught a SIG=" . (defined($sig)? $sig:$!) . " --shutting down\n";
        warn $msg;
        close $opt{LOCK_FH}; # Note: LOCK, not Log - closing Since we will quit ...
        unlink $opt{LOCKFILE};
        $opt{LOCK_FH} = undef;
        $quit_daemon = 1; # Tell loop to quit
    }
    $SIG{$_}  = \&Signal_Handler 
        for qw{ INT KILL QUIT SEGV STOP TERM TSTP __DIE__}; # HUP
    $SIG{USR1}=sub{ warn( "Caught signal USR1. Setting DEBUG=" . ($opt{DEBUG} = (1 - $opt{DEBUG}))) 
                    # This sig will interrupt SLEEP, so we should automatically get another cycle.
                };
    my $grandpop_pid = $$; # Used to connect grandkid's pid message.
    #parent process to start another process that will go on its own is to "double fork."
    # 	The child itself forks and it then exits right away, so its child is taken over by init and can't be a zombie.
    warn (unixtime_to_printable(time()) 
	     . " Daemonizing. Expected to run in background until "
		 .  unixtime_to_printable($opt{ENDTIME_EPOCH}). "\n");
    my $pid = fork ();
    if ($pid < 0) {
      die "first fork failed: $!";
    } elsif ($pid) {
	  # This is the parent (grandpop) process. We want to exit() ASAP.
	  sleep 1; # Wait for grandkid to  write nohup.out 
	  exit 0 unless -f $grandchild_output;
	  open my $f, "<", $grandchild_output or exit 0;
	  my ($grandkid_pid) = grep {m/process (\d+) started\(from $grandpop_pid/} <$f>; 
	  close $f;
	  $grandkid_pid or exit 0;
	  ($grandkid_pid) = $grandkid_pid =~ m/process (\d+) started\(from $grandpop_pid/;
	  $grandkid_pid or exit 0;
      warn "To terminate daemon process, enter:\n"
            ."     kill -s QUIT $grandkid_pid\n"
            ."                              ('kill -s USR1 $grandkid_pid'  toggles DEBUG)\n";

      exit 0; # Exit the parent
    }
    # We are now a child process
    #close std fds inherited from parent
	close STDIN;
	close STDOUT;
	close STDERR;
	POSIX::setsid or die "setsid: $!"; # detaches future kids from the controlling terminal
    open (STDIN, "</dev/null"); # Null all file descriptors
    open (STDOUT, ">>$grandchild_output");
    open (STDERR, ">&STDOUT");
    my $grandchild = fork();
    if ($grandchild < 0) {
        die "second  fork failed : $!";
    } elsif ($grandchild) {
	    #print "Running as Daemon PID=$$\n";
		$| = 1;   # set autoflush on
	    print unixtime_to_printable(time(),"YYYY-MM-DD HH:MM:SS")," querymonitor child process $grandchild started(from $grandpop_pid). exiting parent(s)\n";
        exit 0; # Exit child, leaving grandkid running 		
    }
   ## chdir "/"; 
   POSIX::setsid or die "setsid: $!"; # detaches future kids from the controlling terminal
   umask 0; # Clear the file creation mask
   #foreach (0 .. (POSIX::sysconf (&POSIX::_SC_OPEN_MAX) || 1024))
   #   { POSIX::close $_ } # Close all open file descriptors

   return $pid; # Will always return "0", since this is the child process.
 }
#------------------------------------------------------------------------------
sub unixtime_to_printable{
	my ($unixtime,$format, $showTZ) = @_;
	my ($sec,$min,$hour,$mday,$mon,$year,$wday,$yday,$isdst) = localtime($unixtime);
	if (not defined $format  or  $format eq "YYYY-MM-DD HH:MM:SS"){
       return sprintf("%04d-%02d-%02d %02d:%02d:%02d", $year+1900, $mon+1, $mday, $hour, $min, $sec)
	        . ($showTZ? " " . $opt{TZOFFSET} : "");
	}
	if ($format eq "YYYY-MM-DD HH:MM"){
       return sprintf("%04d-%02d-%02d %02d:%02d", $year+1900, $mon+1, $mday, $hour, $min)
	          . ($showTZ? " " . $opt{TZOFFSET} : "");
	}	
	if ($format eq "YYYY-MM-DD"){
		  return sprintf("%04d-%02d-%02d", $year+1900, $mon+1, $mday)
		         . ($showTZ? " " . $opt{TZOFFSET} : "");
	}
	die "ERROR: Unsupported format:'$format' ";
}
#------------------------------------------------------------------------------

#==============================================================================
BEGIN{
package Analysis::Mode; 

sub new{ # Unused
    my ($class) = @_;
    my $self = bless {}, $class;	
	$self->{INPUT}  = MIME::Multipart::ParseSimple::->new();
	$self->{PIECENBR} = 0;
	$self->{HEADER_PRINTED} =[0,0,0,0]; # Index by piece number 
	return $self;
}

sub Handle_MIME_HEADER  {
	my ($self,$dispatch_type) = @_;
	
	return unless $dispatch_type eq "Header"; # Only expecting a Header 

    for ( grep {!/_SECTION_|Content-Type|params$/} sort keys %{$self->{INPUT}{general_header}} ){
		   print {$self->{OUTPUT_FH}} "INSERT INTO kv_store VALUES ('MIMEHDR','$_','$self->{INPUT}{general_header}{$_}');\n";
	}
		
}

sub Handle_UNIVERSE_JSON{
	my ($self,$dispatch_type, $body) = @_;
	$opt{DEBUG} and print "DEBUG: IN: UNIV JSON handler type $dispatch_type\n";
    return unless $dispatch_type eq "Body"; # Only Interested in Body 
    if (length($body) < 10){
       print "ERROR: UNIVERSE Info was not found in '$body'\n";
       return;
    }
    my $bj = JSON::Tiny::decode_json($body);

	my $count=0;
	for my $n (@{  $bj->{universeDetails}{nodeDetailsSet} }){
       $self->{node}[$count] ={map({$_=>$n->{$_}} qw|nodeIdx nodeName nodeUuid azUuid isMaster isTserver ysqlServerHttpPort yqlServerHttpPort |),
	                              map({$_=>$n->{cloudInfo}{$_}} qw|private_ip public_ip az region |) };
	   if ($count== 0){
		     $opt{DEBUG} and print "DEBUG:","CREATE TABLE NODE (",
			      join(",", sort keys %{ $self->{node}[$count]} ), ");\n";
		     print {$self->{OUTPUT_FH}} "CREATE TABLE NODE (",
			      join(",", sort keys %{ $self->{node}[$count]} ), ");\n";
	   }
	   $opt{DEBUG} and print "DEBUG:","INSERT INTO NODE VALUES('",
			      join("','", map{$self->{node}[$count]{$_}} sort keys %{ $self->{node}[$count]} ), "');\n";
	   print {$self->{OUTPUT_FH}} "INSERT INTO NODE VALUES('",
			      join("','", map{$self->{node}[$count]{$_}} sort keys %{ $self->{node}[$count]} ), "');\n";
       $count++;
    }
     print {$self->{OUTPUT_FH}} "----- End of nodes -----\n";    
}

sub Handle_CSVHEADER	{
	my ($self,$dispatch_type) = @_;
	$opt{DEBUG} and print "DEBUG:IN:CSVHEADER handler type $dispatch_type\n";
	return unless $dispatch_type eq "Header";
	my $type = $self->{INPUT}{general_header}{TYPE};
	((my $Ignore_type_field_in_fieldnames),@{$self->{FIELDS}{$type}}) = split /,/, $self->{INPUT}{general_header}{FIELDS};
	$opt{DEBUG} and print "DEBUG:","CREATE TABLE $type (",
                    join(",",@{$self->{FIELDS}{$type}}), ");\n";
	print {$self->{OUTPUT_FH}} "CREATE TABLE $type (",
	                 join(",",@{$self->{FIELDS}{$type}}), ");\n";    
}
sub Handle_MONITOR_DATA {
	my ($self,$dispatch_type,$body ) = @_;
	$opt{DEBUG} and print "DEBUG:IN: MonitorData handler type $dispatch_type\n";
	return unless $dispatch_type eq "Body" and length($body) > 2;
	my ($type,@values) = split /,/,$body;
	s/'/~/g for @values;
	print {$self->{OUTPUT_FH}} "INSERT INTO $type VALUES('",
	        join("','",@values),      "');\n"; 
};

my %Section_Handler =( # Defines Handler subroutines for each Mime piect (_SECTION_) received
	MIME_HEADER		=> \&Handle_MIME_HEADER	,
	UNIVERSE_JSON	=> \&Handle_UNIVERSE_JSON ,
	CSVHEADER		=> \&Handle_CSVHEADER	,
	MONITOR_DATA	=> \&Handle_MONITOR_DATA,
);

sub Parse_Body_Record{
   my ($self, $rec) = @_;

   my $dispatch = $Section_Handler{ $self->{INPUT}{general_header}{_SECTION_} };
   if ( ! $self->{HEADER_PRINTED}[$self->{PIECENBR}]){
      print "HDR $self->{PIECENBR}:",map({"$_=>" . $self->{INPUT}{general_header}{$_} . "; "} 
	     grep {!/params$/} sort keys %{$self->{INPUT}{general_header}}),"\n";
	  $dispatch and $dispatch->($self,"Header"); # Handler is called here 
	  $self->{HEADER_PRINTED}[$self->{PIECENBR}] = 1;
   }

   if (! defined $rec) {
      print "---- PIECE COMPLETE --\n" ;

	  $dispatch and $dispatch->($self,"EOF"); # Handler is called here 
	  $self->{PIECENBR}++;
	  return;
   }   
   $opt{DEBUG} and print "DEBUG:GOT Piece:",substr($rec,0,200),"..\n";
   $dispatch and $dispatch->($self,"Body", $rec); # Handler is called here    
}

sub Process_file_and_create_Sqlite{
	my ($self) = @_;
	
	if ( -f $opt{ANALYZE} ){
		# File exists - fall through and process it
	}else{
	    die "ERROR: 'ANALYZE' file does not exist: No file '$opt{ANALYZE}'.";	
	}
	print "Analyzing $opt{ANALYZE} ...\n";
	# Create INPUT handle to read (possibly gzipped) MIME file ---
	if ($opt{ANALYZE} =~/\.gz$/i){
		open($self->{INPUT_FH}, "-|", "gunzip -c $opt{ANALYZE}") or die "ERROR: Cannot fork gunzip to open $opt{ANALYZE}";
	}else{
		open($self->{INPUT_FH},"<",$opt{ANALYZE}) or die "ERROR opening $opt{ANALYZE}:$!";
	}
	
    $self->Initialize_SQLITE_Output();
	
	$self->{INPUT}->parse($self, \&Parse_Body_Record);
	close $self->{INPUT_FH};
	print {$self->{OUTPUT_FH}} "SELECT 'All input records processed.'";
	# Should create and run VIEWs here ....
	close $self->{OUTPUT_FH};

	return;
}

sub OLD_DEAD_CODE{
	
	my ($sqlite_version) = do {my $vv=qx|$opt{SQLITE} --version|;chomp $vv;$vv=~/([\d\.]+)/};
	$! and die "ERROR: Cannot run $opt{SQLITE} :$!";
	$sqlite_version or die "ERROR: could not get SQLITE version";
	# We do several extra steps to allow for OLD sqlite (< 3.8):
	# (a) It does not support instr(), so we use xx like yy
	# (b) the .import command does not allow --skip, so we pre-skip on input file
	
	# Use Forking and fifo to de-compress, remove first line, 
	# then feed via the fifo into the ".import" command in sqlite
	my $fifo = "/tmp/querymonitor_fifo_$$";
	print "Creating temporary fifo $fifo ...\n";
	qx|mkfifo $fifo|;
	$! and die "ERROR: Fifo failed: $!";
	
	#Fork to provide CSV stream to fifo

    my $pid = fork ();
    if ($pid < 0) {
      die "ERROR: fork failed: $!";
    } elsif ($pid) {
	  # This is the parent --fall through 
    }else{
		# This is the CHILD 
	    #close std fds inherited from parent
		close STDIN;
		close STDOUT;
		close STDERR;		
	####	qx{$extract_cmd $opt{ANALYZE} | sed 1d > $fifo};
		exit 0; # Exit child process and close FIFO
	}
	
	if (! $opt{DB}){
	   # DB name was not specified .. generate it from ANALYZE file name
	   $opt{DB} = $opt{ANALYZE};
	   $opt{DB} =~s/\.gz$//i;  # drop .gz
	   $opt{DB} =~s/\.csv$//i; # drop .csv
	   $opt{DB} .= ".sqlite";  # append .sqlite 
	}
	print "Creating sqlite database $opt{DB} ...\n";
	my $populate_zone_map=""; # Need a better way to do this...
	if ($opt{ANALYZE} =~/SECONDARY/i){
		$populate_zone_map = <<"__SEC_ZONE__";
		update zone_map set zone='DC3';
		update zone_map set zone='DC1' where node like '%n2';
		update zone_map set zone='DC1' where node like '%n3';
		update zone_map set zone='DC1' where node like '%n4';
		update zone_map set zone='DC1' where node like '%n9';
__SEC_ZONE__
    }elsif ( $opt{ANALYZE} =~/PRIMARY/i ){
		$populate_zone_map = <<"__PRI_ZONE__";
		update zone_map set zone='DC1';
		update zone_map set zone='DC3' where node like '%n1';
		update zone_map set zone='DC3' where node like '%n2';
		update zone_map set zone='DC3' where node like '%n3';
		update zone_map set zone='DC3' where node like '%n5';
__PRI_ZONE__
	}
    
	
	open my $sqlfh ,  "|-" , $opt{SQLITE} , $opt{DB} or die "ERROR: Could not run SQLITE";
	print $sqlfh <<"__SQL__";
.version
.header off
CREATE TABLE IF NOT EXISTS run_info(key text, value text);
INSERT INTO run_info VALUES ('data file','$opt{ANALYZE}')
      ,('HOSTNAME','$opt{HOSTNAME}')
	  ,('import date',datetime('now','localtime'));
CREATE TABLE zone_map (node,zone);	  
CREATE TABLE q (ts integer,clientHost,clientPort,elapsedMillis integer,id,keyspace,nodeName,privateIp,query,type);
CREATE VIEW summary as select datetime((ts/600)*600,'unixepoch','-4 Hours') as EDT,
    sum(case when instr(query,' system.')> 0 then 1 else 0 end) as systemq,
        sum(case when instr(query,' system.')=0 then 1 else 0 end) as cqlcount,
        sum(case when instr(query,' system.')>0 and elapsedmillis > 120 then 1 else 0 end) as sys_gt120,
        sum(case when instr(query,' system.')=0 and elapsedmillis > 120 then 1 else 0 end) as cql_gt120,
        sum(case when instr(query,' system.')>0 and zone='DC1' then 1 else 0 end) as sys_dc1,
        sum(case when instr(query,' system.')=0 and zone='DC1' then 1 else 0 END) as cql_dc1,
        sum(case when instr(query,' system.')>0 and zone='DC3' then 1 else 0 end) as sys_dc3,
        sum(case when instr(query,' system.')=0 and zone='DC3' then 1 else 0 END) as cql_dc3,		
         round(sum(case when instr(query,' system.')=0 and elapsedmillis > 120 then 1 else 0 end) *100.0
               / sum(case when instr(query,' system.')=0 then 1 else 0 end)
                   ,2) as breach_pct
FROM q,zone_map
   where q.nodename=zone_map.node 
 group by EDT;

CREATE VIEW slow_queries  as  select query, count(*) as nbr_querys, round(avg(elapsedmillis),1) as avg_milli ,
      sum (CASE when elapsedmillis > 120 then 1 else 0 END)*100 / count(*) as pct_gt120,
          sum ( CASE WHEN zone = 'DC1' THEN 1 ELSE 0 END) as dc1_queries,
		  sum ( CASE WHEN zone = 'DC3' THEN 1 ELSE 0 END) as dc3_queries
          FROM q, zone_map
		   where q.nodename=zone_map.node 
          GROUP BY query
          HAVING nbr_querys > 50 and avg_milli >30  ORDER by avg_milli  desc;
		  
CREATE VIEW NODE_Report AS  select nodename, round(avg(elapsedmillis),1) as avg_ms, 
       count(*), sum(case when instr(query,' system.') > 0 then 1 else 0 end) as sys_count,  
	   sum(case when instr(query,' system.')= 0 then 1 else 0 end) as cql_count,
	   sum(case when instr(query,' system.')= 0 and elapsedmillis > 120 then 1 else 0 end) as cql_gt_120,
	   sum(case when instr(query,' system.')> 0 and elapsedmillis > 120 then 1 else 0 end) as sys_gt_120,
	   zone
  from q,zone_map 
  where nodename=node  
  group by  nodename 
  order by nodename;
  
.mode csv
.import '$fifo' q
.mode column
SELECT 'Imported ' || count(*) ||' rows from $opt{ANALYZE}.' as Imported_count from q;

insert into zone_map select distinct nodename,'UNKNOWN' from q;
$populate_zone_map;
SELECT ''; -- blank line
SELECT '====== Summary Report ====';
.header on
SELECT * from summary;
.header off
SELECT '';
SELECT '======= Slow Queries =======';
.header on
select * from slow_queries;
.q
__SQL__

   close $sqlfh;
   unlink $fifo;
   wait; # For kid 
}

sub Initialize_SQLITE_Output{
    my ($self) = @_;
	my ($sqlite_version) = do {my $vv=qx|$opt{SQLITE} --version|;chomp $vv;$vv=~/([\d\.]+)/};
	$! and die "ERROR: Cannot run $opt{SQLITE} :$!";
	# Create output handle to SQLITE ----
	($self->{SQLITE_FILENAME}) = $opt{ANALYZE} =~m/(.+)\.\w+$/;
	$self->{SQLITE_FILENAME}   =~s/\.mime$//i; # Zap the mime
	$self->{SQLITE_FILENAME}   .= ".sqlite"; # Add sqlite suffix
	$opt{DB}  and    $self->{SQLITE_FILENAME} = $opt{DB}; # DB wins, if specified.
	print "Populating Sqlite($sqlite_version) database '$self->{SQLITE_FILENAME}'...\n";
	
	open($self->{OUTPUT_FH}, "|-", "$opt{SQLITE} $self->{SQLITE_FILENAME}") 
	     or die "ERROR: Cannot fork sqlite to open $self->{SQLITE_FILENAME}";
	print {$self->{OUTPUT_FH}} <<"__SQL1__";
CREATE TABLE IF NOT EXISTS kv_store(type text,key text, value text);
INSERT INTO kv_store VALUES ('GENERAL','data file','$opt{ANALYZE}')
      ,('GENERAL','HOSTNAME','$opt{HOSTNAME}')
	  ,('GENERAL','import date','$opt{STARTTIME_TZ}')
	  ,('GENERAL','processing file','$opt{ANALYZE}')
	  ,('GENERAL','Analysis version','$main::VERSION');
	
__SQL1__

}

1;
} # End of Analysis::Mode
#==============================================================================
BEGIN{
package Web::Interface; # Handles communication with YBA API 

sub new{
	my ($class) = @_;
	for(qw|API_TOKEN YBA_HOST CUST_UUID UNIV_UUID|){
        $opt{$_} or die "ERROR: Required parameter --$_ was not specified.\n";
    }
	my $self =bless {map {$_ => $opt{$_}} qw|HTTPCONNECT UNIV_UUID API_TOKEN YBA_HOST CUST_UUID| }, $class;
	$self->{BASE_URL} = "$opt{YBA_HOST}/api/customers/$opt{CUST_UUID}/universes/$opt{UNIV_UUID}";
	if ($self->{HTTPCONNECT} eq "curl"){
		  $self->{curl_base_cmd} = join " ", $opt{CURL}, 
					 qq|-s --request GET --header 'Content-Type: application/json'|,
					 qq|--header "X-AUTH-YW-API-TOKEN: $opt{API_TOKEN}"|,
					 qq|--url $self->{BASE_URL}|;
		  if ($opt{DEBUG}){
			 print "DEBUG:CURL base CMD: $self->{curl_base_cmd}\n";
		  }
		  return $self;
    }
    
	$self->{HT} = HTTP::Tiny->new( default_headers => {
                         'X-AUTH-YW-API-TOKEN' => $opt{API_TOKEN},
						 'Content-Type'      => 'application/json',
	                  });

    return $self;
}

sub Get{
	my ($self, $endpoint) = @_;
	$self->{json_string}= "";
	if ($self->{HTTPCONNECT} eq "curl"){
		$self->{json_string} = qx|$self->{curl_base_cmd}$endpoint|;
		if ($?){
		   print "ERROR: curl get '$endpoint' failed: $?\n";
		   exit 1;
		}
    }else{ # HTTP::Tiny
	   $self->{raw_response} = $self->{HT}->get( 'http://' . $self->{BASE_URL} . $endpoint );
	   if (not $self->{raw_response}->{success}){
		  print "ERROR: Get '$endpoint' failed with status=$self->{raw_response}->{status}: $self->{raw_response}->{reason}\n";
		  exit 1;
	   }
	   $self->{json_string} = $self->{raw_response}{content};
	}
	#### $self->{json_string}=~s/:/=>/g;
	#### $self->{json_string}=~s/(\W)true(\W)/${1}1$2/g;
	#### $self->{json_string}=~s/(\W)false(\W)/${1}0$2/g;
	#### if (length($endpoint) == 0){
	####     # Special case for UNIVERSE .. clean unparsable...
    ####     $self->{json_string}=~s/"sampleAppCommandTxt"=>.+",//;
	#### }
    #### $@="";
    #### $self->{response} = eval $self->{json_string};
	#### $@ and die "EVAL ERROR getting $endpoint:$@";
    #### #my %univ = $json_string=~m/"(\w[^"]+)":"([^"]+)/g;  # Grab all simple scalars
    #### #%univ = %univ, $json_string=~m/"(\w[^"]+)":\[([^\]]+)\]/g; # Append all arrays (not decoded)
	$self->{response} = JSON::Tiny::decode_json( $self->{json_string} );
	return $self->{response};
}

} # End of  package Web::Interface;  
#==============================================================================
BEGIN{
package MIME::Write::Simple;

sub new{
	my ($class,%att) = @_;
	
	$att{boundary} ||= time() . $$;
	$att{MIME_VERSION_SENT} = 0;
	return bless {%att},$class;
}

sub header{
	my ($self,$content_type,$header_msg) = @_;
	$header_msg ||= "";
	chomp $header_msg;
	$header_msg .= "\n\n";
	my $mime_ver = $self->{MIME_VERSION_SENT} ? "" 
                   : "MIME Version: 1.0\n";
    $self->{MIME_VERSION_SENT} = 1;
	print { $self->{OUTPUT_FH} }   $mime_ver
	      . qq|Content-Type: $content_type; boundary="$self->{boundary}"\n|
		  . $header_msg ;
}

sub boundary{ # getter/setter
   my ($self,$b, $final) = @_;
   $b and $self->{boundary} = $b;
   print { $self->{OUTPUT_FH} }  "--" . $self->{boundary} . ($final ? "--\n" : "\n");
}

sub Open_and_Initialize{
	my ($self) = @_;
	$opt{DEBUG} and print "DEBUG: Opening output file=" , $opt{OUTPUT},"\n";
	my $output_already_exists = -f $opt{OUTPUT};
	open $self->{OUTPUT_FH} , "|-", "gzip -c >> " . $opt{OUTPUT}
	  or die "ERROR: Cannot fork output zip:$!";
	if (! $output_already_exists){
	   # Need MIME headers etc...
	   $self->header("multipart/mixed",
	      join ("\n",
			  "Querymonitor_version: $VERSION",
			  "UNIVERSE: $opt{UNIVERSE}{name}",
			  "UNIV_UUID: $opt{UNIV_UUID}",
			  "STARTTIME: $opt{STARTTIME_TZ}",
			  "Run_host: $opt{HOSTNAME}",
			  "_SECTION_: MIME_HEADER",
			  "SANITIZED: $opt{SANITIZE}"
		  ));
	    $self->boundary();  
        # Insert NODE/ZONE info
        $self->header("application/json", 
                      "Description: Universe detail JSON\nName: universe_info\n"
					  ."_SECTION_: UNIVERSE_JSON");
		print { $self->{OUTPUT_FH} } $self->{UNIVERSE_JSON},"\n";
		$self->boundary();
	}
}

sub Initialize_query_type{
	my ($self,$type,$q) = @_;
	
    if (! $self->{OUTPUT_FH} ){
	  # Need to initialize output
      $self->Open_and_Initialize();
    }
    if ($self->{IN_CSV_SECTION}){
	   $self->boundary(); # Close the CSV section
	}
	$self->header("text/csvheader",
	      join("\n",
		     "TYPE: $type",
			 "_SECTION_: CSVHEADER",
			 "FIELDS: " . join(",","type","ts",sort(keys %$q),"query")
			 ));
	$self->boundary();
	$self->{TYPE_INITIALIZED}{$type} = 1;
	$self->header("text/csv","_SECTION_: MONITOR_DATA");
	$self->{IN_CSV_SECTION} = 1;
}

sub WriteQuery{
  my ($self,$ts, $type, $q, $sanitized_query) = @_;
  
  if ( ! $self->{TYPE_INITIALIZED}{$type} ){
	  $self->Initialize_query_type($type, $q);  
  }
  $sanitized_query =~tr/"/~/; # Zap internal double quotes - which will mess CSV
  if (length($sanitized_query) > $opt{MAX_QUERY_LEN}){
	   $sanitized_query = substr($sanitized_query,0,$opt{MAX_QUERY_LEN}/2 -2) 
	                     . ".." . substr($sanitized_query,-($opt{MAX_QUERY_LEN}/2));
  }
  $q->{query} = qq|"$sanitized_query"|;
  print { $self->{OUTPUT_FH} } join(",", $type, $ts, map( {$q->{$_}} sort keys %$q)),"\n";
}

sub Close{
	my ($self) = @_;
    return unless $self->{OUTPUT_FH};
	$self->boundary(undef,"FINAL");
	$self->{IN_CSV_SECTION} = 0;
    close $self->{OUTPUT_FH};
    $self->{OUTPUT_FH} = undef;	
}	
1;
} # End of MIME::Write::Simple
#==============================================================================
BEGIN{
package MIME::Multipart::ParseSimple;

use strict;
use warnings FATAL => 'all';
use Carp;

our $VERSION = '0.02';


sub new {
  my $p = shift;
  my $c = ref($p) || $p;
  my $o = {};
  bless $o, $c;
  return $o;
}

=head2 parse

takes one argument: a file handle.

returns a listref, each item corresponding to a MIME header in
the document.  The first is the multipart file header itself.
Each header item is stored as key/value.  Additional parameters
are stored $key.params.  e.g. the boundary is at

    $o->[0]->{"Content-Type.params"}->{"boundary"}

The first item may also have {"Preamble"} and {"Epilog"} if these
existed in the file.

The content of each part is stored as {"Body"}.

=cut

sub parse {
  # load a MIME-multipart-style file containing at least one application/x-ptk.markdown
  my ($o,$caller, $callback) = @_;
  $o->{fh}       = $caller->{INPUT_FH};
  $o->{CALLER}   = $caller;
  $o->{callback} = $callback;

  my $mp1 = readline($o->{fh});
  my $mp1e = 'MIME Version: 1.0';
  die "Multipart header line 1 must begin ``$mp1e'' " unless $mp1 =~ /^$mp1e/;
 
  $o->{general_header} = $o->parseHeader();
  croak "no boundary defined" unless $o->{general_header}->{"Content-Type.params"}->{"boundary"};
  my $b = $o->{general_header}->{"Content-Type.params"}->{"boundary"};
  if (length($b)>2 and (my $quote=substr($b,0,1)) eq substr($b,-1,1)){
      if ($quote eq '"'  or $quote eq "'"){
		  $b=substr($b,1,length($b)-2);
	  }
  }
  $o->{boundary} = $b;
  
  $o->parseBody();

  #my @parts = ($general_header);

  while(! (eof($o->{fh}) || $o->{eof})){
    $o->{general_header} = $o->parseHeader();
    $o->parseBody();
    #push @parts, $header;
  }

  #$general_header->{Epilog} = $o->parseBody();

  #return \@parts;

}

sub parseBody {
  my ($o) = @_;
  my $fh = $o->{fh};
  my $boundary = $o->{boundary};
  $o->{recordnumber} = 0;
  while(<$fh>){
    $o->{eof} = 1 if /^--$boundary--/;
    if (/^--$boundary/){
		$o->{callback}->($o->{CALLER},undef); # Indiates Piece completed.
		$o->{general_header} = undef;
		return;
	}
	$o->{recordnumber}++;
	chomp;
    $o->{callback}->($o->{CALLER},$_);
  }
  return undef;
}

sub parseHeader {
  my ($o) = @_;
  my $fh = $o->{fh};
  my %header = ();
  my ($k,$v,$e,$p);
  while(<$fh>){
    last if /^\s*$/; # break on a blank line...
    my @parts = split /;/;
    if(/^\S/){ # non space at start means a new header item
      my $header = shift @parts;
      ($k,$v) = split(/\:/, $header, 2);
	  next unless $v;
      $k =~ s/(?:^\s+|\s+$)//g;
      $v =~ s/(?:^\s+|\s+$)//g;
      $header{$k} = $v;
      $p = $k.'.params';
      $header{$p} = {};
    }
    foreach my $part(@parts){
      my ($l,$w) = split(/=/, $part, 2);
	  next unless $w;
      $l =~ s/(?:^\s+|\s+$)//g;
      $w =~ s/(?:^\s+|\s+$)//g;
      $header{$p}->{$l} = $w;
    }
  }
  return \%header;
}
1;
} # End of  MIME::Multipart::ParseSimple
#==============================================================================
BEGIN{
package JSON::Tiny;

# Minimalistic JSON. Adapted from Mojo::JSON. (c)2012-2015 David Oswald
# License: Artistic 2.0 license.
# http://www.perlfoundation.org/artistic_license_2_0

use strict;
use warnings;
use Carp 'croak';
use Exporter 'import';
use Scalar::Util 'blessed';
use Encode ();
use B;

our $VERSION = '0.58';
our @EXPORT_OK = qw(decode_json encode_json false from_json j to_json true);

# Literal names
# Users may override Booleans with literal 0 or 1 if desired.
our($FALSE, $TRUE) = map { bless \(my $dummy = $_), 'JSON::Tiny::_Bool' } 0, 1;

# Escaped special character map with u2028 and u2029
my %ESCAPE = (
  '"'     => '"',
  '\\'    => '\\',
  '/'     => '/',
  'b'     => "\x08",
  'f'     => "\x0c",
  'n'     => "\x0a",
  'r'     => "\x0d",
  't'     => "\x09",
  'u2028' => "\x{2028}",
  'u2029' => "\x{2029}"
);
my %REVERSE = map { $ESCAPE{$_} => "\\$_" } keys %ESCAPE;

for(0x00 .. 0x1f) {
  my $packed = pack 'C', $_;
  $REVERSE{$packed} = sprintf '\u%.4X', $_ unless defined $REVERSE{$packed};
}

sub decode_json {
  my $err = _decode(\my $value, shift);
  return defined $err ? croak $err : $value;
}

sub encode_json { Encode::encode 'UTF-8', _encode_value(shift) }

sub false () {$FALSE}  ## no critic (prototypes)

sub from_json {
  my $err = _decode(\my $value, shift, 1);
  return defined $err ? croak $err : $value;
}

sub j {
  return encode_json $_[0] if ref $_[0] eq 'ARRAY' || ref $_[0] eq 'HASH';
  return decode_json $_[0];
}

sub to_json { _encode_value(shift) }

sub true () {$TRUE} ## no critic (prototypes)

sub _decode {
  my $valueref = shift;

  eval {

    # Missing input
    die "Missing or empty input\n" unless length( local $_ = shift );

    # UTF-8
    $_ = eval { Encode::decode('UTF-8', $_, 1) } unless shift;
    die "Input is not UTF-8 encoded\n" unless defined $_;

    # Value
    $$valueref = _decode_value();

    # Leftover data
    return m/\G[\x20\x09\x0a\x0d]*\z/gc || _throw('Unexpected data');
  } ? return undef : chomp $@;

  return $@;
}

sub _decode_array {
  my @array;
  until (m/\G[\x20\x09\x0a\x0d]*\]/gc) {

    # Value
    push @array, _decode_value();

    # Separator
    redo if m/\G[\x20\x09\x0a\x0d]*,/gc;

    # End
    last if m/\G[\x20\x09\x0a\x0d]*\]/gc;

    # Invalid character
    _throw('Expected comma or right square bracket while parsing array');
  }

  return \@array;
}

sub _decode_object {
  my %hash;
  until (m/\G[\x20\x09\x0a\x0d]*\}/gc) {

    # Quote
    m/\G[\x20\x09\x0a\x0d]*"/gc
      or _throw('Expected string while parsing object');

    # Key
    my $key = _decode_string();

    # Colon
    m/\G[\x20\x09\x0a\x0d]*:/gc
      or _throw('Expected colon while parsing object');

    # Value
    $hash{$key} = _decode_value();

    # Separator
    redo if m/\G[\x20\x09\x0a\x0d]*,/gc;

    # End
    last if m/\G[\x20\x09\x0a\x0d]*\}/gc;

    # Invalid character
    _throw('Expected comma or right curly bracket while parsing object');
  }

  return \%hash;
}

sub _decode_string {
  my $pos = pos;
  
  # Extract string with escaped characters
  m!\G((?:(?:[^\x00-\x1f\\"]|\\(?:["\\/bfnrt]|u[0-9a-fA-F]{4})){0,32766})*)!gc; # segfault on 5.8.x in t/20-mojo-json.t
  my $str = $1;

  # Invalid character
  unless (m/\G"/gc) {
    _throw('Unexpected character or invalid escape while parsing string')
      if m/\G[\x00-\x1f\\]/;
    _throw('Unterminated string');
  }

  # Unescape popular characters
  if (index($str, '\\u') < 0) {
    $str =~ s!\\(["\\/bfnrt])!$ESCAPE{$1}!gs;
    return $str;
  }

  # Unescape everything else
  my $buffer = '';
  while ($str =~ m/\G([^\\]*)\\(?:([^u])|u(.{4}))/gc) {
    $buffer .= $1;

    # Popular character
    if ($2) { $buffer .= $ESCAPE{$2} }

    # Escaped
    else {
      my $ord = hex $3;

      # Surrogate pair
      if (($ord & 0xf800) == 0xd800) {

        # High surrogate
        ($ord & 0xfc00) == 0xd800
          or pos($_) = $pos + pos($str), _throw('Missing high-surrogate');

        # Low surrogate
        $str =~ m/\G\\u([Dd][C-Fc-f]..)/gc
          or pos($_) = $pos + pos($str), _throw('Missing low-surrogate');

        $ord = 0x10000 + ($ord - 0xd800) * 0x400 + (hex($1) - 0xdc00);
      }

      # Character
      $buffer .= pack 'U', $ord;
    }
  }

  # The rest
  return $buffer . substr $str, pos $str, length $str;
}

sub _decode_value {

  # Leading whitespace
  m/\G[\x20\x09\x0a\x0d]*/gc;

  # String
  return _decode_string() if m/\G"/gc;

  # Object
  return _decode_object() if m/\G\{/gc;

  # Array
  return _decode_array() if m/\G\[/gc;

  # Number
  my ($i) = /\G([-]?(?:0|[1-9][0-9]*)(?:\.[0-9]*)?(?:[eE][+-]?[0-9]+)?)/gc;
  return 0 + $i if defined $i;

  # True
  return $TRUE if m/\Gtrue/gc;

  # False
  return $FALSE if m/\Gfalse/gc;

  # Null
  return undef if m/\Gnull/gc;  ## no critic (return)

  # Invalid character
  _throw('Expected string, array, object, number, boolean or null');
}

sub _encode_array {
  '[' . join(',', map { _encode_value($_) } @{$_[0]}) . ']';
}

sub _encode_object {
  my $object = shift;
  my @pairs = map { _encode_string($_) . ':' . _encode_value($object->{$_}) }
    sort keys %$object;
  return '{' . join(',', @pairs) . '}';
}

sub _encode_string {
  my $str = shift;
  $str =~ s!([\x00-\x1f\x{2028}\x{2029}\\"/])!$REVERSE{$1}!gs;
  return "\"$str\"";
}

sub _encode_value {
  my $value = shift;

  # Reference
  if (my $ref = ref $value) {

    # Object
    return _encode_object($value) if $ref eq 'HASH';

    # Array
    return _encode_array($value) if $ref eq 'ARRAY';

    # True or false
    return $$value ? 'true' : 'false' if $ref eq 'SCALAR';
    return $value  ? 'true' : 'false' if $ref eq 'JSON::Tiny::_Bool';

    # Blessed reference with TO_JSON method
    if (blessed $value && (my $sub = $value->can('TO_JSON'))) {
      return _encode_value($value->$sub);
    }
  }

  # Null
  return 'null' unless defined $value;


  # Number (bitwise operators change behavior based on the internal value type)

  return $value
    if B::svref_2object(\$value)->FLAGS & (B::SVp_IOK | B::SVp_NOK)
    # filter out "upgraded" strings whose numeric form doesn't strictly match
    && 0 + $value eq $value
    # filter out inf and nan
    && $value * 0 == 0;

  # String
  return _encode_string($value);
}

sub _throw {

  # Leading whitespace
  m/\G[\x20\x09\x0a\x0d]*/gc;

  # Context
  my $context = 'Malformed JSON: ' . shift;
  if (m/\G\z/gc) { $context .= ' before end of data' }
  else {
    my @lines = split "\n", substr($_, 0, pos);
    $context .= ' at line ' . @lines . ', offset ' . length(pop @lines || '');
  }

  die "$context\n";
}

# Emulate boolean type
package JSON::Tiny::_Bool;
use overload '""' => sub { ${$_[0]} }, fallback => 1;
1;	

}; #End of JSON::Tiny
#==============================================================================
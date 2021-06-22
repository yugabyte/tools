// Code generated by protoc-gen-ybrpc. DO NOT EDIT.

package tserver

// service: yb.tserver.TabletServerAdminService
// service: TabletServerAdminService

// methods
type TabletServerAdminServiceImpl struct{}

func (s *TabletServerAdminServiceImpl) CreateTablet(ts string, pb *CreateTabletRequestPB) (*CreateTabletResponsePB, error) {
	return &CreateTabletResponsePB{}, nil
}

func (s *TabletServerAdminServiceImpl) DeleteTablet(ts string, pb *DeleteTabletRequestPB) (*DeleteTabletResponsePB, error) {
	return &DeleteTabletResponsePB{}, nil
}

func (s *TabletServerAdminServiceImpl) AlterSchema(ts string, pb *ChangeMetadataRequestPB) (*ChangeMetadataResponsePB, error) {
	return &ChangeMetadataResponsePB{}, nil
}

func (s *TabletServerAdminServiceImpl) GetSafeTime(ts string, pb *GetSafeTimeRequestPB) (*GetSafeTimeResponsePB, error) {
	return &GetSafeTimeResponsePB{}, nil
}

func (s *TabletServerAdminServiceImpl) BackfillIndex(ts string, pb *BackfillIndexRequestPB) (*BackfillIndexResponsePB, error) {
	return &BackfillIndexResponsePB{}, nil
}

func (s *TabletServerAdminServiceImpl) BackfillDone(ts string, pb *ChangeMetadataRequestPB) (*ChangeMetadataResponsePB, error) {
	return &ChangeMetadataResponsePB{}, nil
}

func (s *TabletServerAdminServiceImpl) CopartitionTable(ts string, pb *CopartitionTableRequestPB) (*CopartitionTableResponsePB, error) {
	return &CopartitionTableResponsePB{}, nil
}

func (s *TabletServerAdminServiceImpl) FlushTablets(ts string, pb *FlushTabletsRequestPB) (*FlushTabletsResponsePB, error) {
	return &FlushTabletsResponsePB{}, nil
}

func (s *TabletServerAdminServiceImpl) CountIntents(ts string, pb *CountIntentsRequestPB) (*CountIntentsResponsePB, error) {
	return &CountIntentsResponsePB{}, nil
}

func (s *TabletServerAdminServiceImpl) AddTableToTablet(ts string, pb *AddTableToTabletRequestPB) (*AddTableToTabletResponsePB, error) {
	return &AddTableToTabletResponsePB{}, nil
}

func (s *TabletServerAdminServiceImpl) RemoveTableFromTablet(ts string, pb *RemoveTableFromTabletRequestPB) (*RemoveTableFromTabletResponsePB, error) {
	return &RemoveTableFromTabletResponsePB{}, nil
}

func (s *TabletServerAdminServiceImpl) SplitTablet(ts string, pb *SplitTabletRequestPB) (*SplitTabletResponsePB, error) {
	return &SplitTabletResponsePB{}, nil
}

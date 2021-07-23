/*
Copyright © 2021 Yugabyte Support

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			if len(f.Services) > 0 {
				generate := false
				for _, service := range f.Services {
					if len(service.Methods) > 0 {
						generate = true
					}
				}
				if generate {
					_ = generateFile(gen, f)
				}
			}
		}
		return nil
	})
}

// generateFile generates a _ascii.pb.go file containing gRPC service definitions.
func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + "_ybrpc.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)

	protoDeclaration := getProtoDeclaration(file)
	generateComments(g, generateCommentSet(protoDeclaration), false)

	g.P("// Code generated by protoc-gen-ybrpc. DO NOT EDIT.")

	packageDeclaration := getPackageDeclaration(file)
	generateComments(g, generateCommentSet(packageDeclaration), file.Desc.Options().(*descriptorpb.FileOptions).GetDeprecated())

	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	generateImports(g)

	for _, service := range file.Services {
		generateService(g, service)

		generateServiceImpl(g, service)

		for _, method := range service.Methods {
			generateServiceMethod(g, service, method)
		}
	}
	g.P()

	return g
}

// TODO: This statically outputs imports- how is this _actually_ supposed to be implemened?
func generateImports(g *protogen.GeneratedFile) {
	g.P()
	g.P(`import (`)
	g.P(`    "github.com/go-logr/logr"`)
	g.P(`    "github.com/yugabyte/yb-tools/protoc-gen-ybrpc/pkg/message"`)
	g.P(`)`)
	g.P()
}

func generateService(g *protogen.GeneratedFile, service *protogen.Service) {
	g.P("// service: ", service.Desc.FullName())
	g.P("// service: ", service.GoName)

	generateComments(g, service.Comments, service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated())
	g.P("type ", service.GoName, " interface {")
	for _, method := range service.Methods {
		generateMethodSigniture(g, method)
	}
	g.P("}")
	g.P()
}

func generateMethodSigniture(g *protogen.GeneratedFile, method *protogen.Method) {
	g.P(method.GoName, "(request *", method.Input.GoIdent.GoName+")"+"(*"+method.Output.GoIdent.GoName+", error)")
}

func generateServiceImpl(g *protogen.GeneratedFile, service *protogen.Service) {
	g.P("type ", service.GoName, "Impl struct {")
	g.P("    Log logr.Logger")
	g.P("    Messenger message.Messenger")
	g.P("}")
	g.P()
}

func generateServiceMethod(g *protogen.GeneratedFile, service *protogen.Service, method *protogen.Method) {
	generateComments(g, method.Comments, method.Desc.Options().(*descriptorpb.MethodOptions).GetDeprecated())
	g.P("func (s *" + service.GoName + "Impl)" + method.GoName + "(request *" + method.Input.GoIdent.GoName + ")" + "(*" + method.Output.GoIdent.GoName + ", error) {")
	g.P(`    s.Log.V(1).Info("sending RPC message", "service", "`, string(service.Desc.FullName()), `", "method", "`, string(method.Desc.Name()), `", "message", request)`)
	g.P("    response := &" + method.Output.GoIdent.GoName + "{}")
	g.P()
	g.P(`    err := s.Messenger.SendMessage("`, string(service.Desc.FullName()), `", "`, string(method.Desc.Name()), `", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())`)
	g.P("    if err != nil {")
	g.P("        return nil, err")
	g.P("    }")
	g.P()
	g.P(`    s.Log.V(1).Info("received RPC response", "service", "`, string(service.Desc.FullName()), `", "method", "`, string(method.Desc.Name()), `", "message", response)`)
	g.P()
	g.P("    return response, nil")
	g.P("}")
	g.P("")
}

func generateComments(g *protogen.GeneratedFile, comments protogen.CommentSet, deprecated bool) {
	for _, s := range comments.LeadingDetached {
		g.P(s)
		g.P()
	}
	if s := comments.Leading; s != "" {
		g.P(s)
	}
	if deprecated {
		g.P(protogen.Comments(" Deprecated: Do not use."))
	}
}

func generateCommentSet(location protoreflect.SourceLocation) protogen.CommentSet {
	commentSet := protogen.CommentSet{}
	for _, comment := range location.LeadingDetachedComments {
		commentSet.LeadingDetached = append(commentSet.LeadingDetached, protogen.Comments(comment))
	}

	commentSet.Leading = protogen.Comments(location.LeadingComments)

	return commentSet
}

func getProtoDeclaration(f *protogen.File) protoreflect.SourceLocation {
	return f.Desc.SourceLocations().ByPath(protoreflect.SourcePath{12})
}

func getPackageDeclaration(f *protogen.File) protoreflect.SourceLocation {
	return f.Desc.SourceLocations().ByPath(protoreflect.SourcePath{2})
}

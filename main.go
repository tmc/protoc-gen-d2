package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var (
	flags            flag.FlagSet
	typeSkipEdgeToRe = flags.String("skip-re", `google\.protobuf\.*`, "skip edges to types matching this regexp")
)

func main() {
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generateFile(gen, f, *typeSkipEdgeToRe)
		}
		return nil
	})
}

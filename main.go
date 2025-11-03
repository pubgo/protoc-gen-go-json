package main

import (
	"flag"
	"fmt"
	"log/slog"

	"github.com/pubgo/protoc-gen-go-json/internal"
	"github.com/pubgo/protoc-gen-go-json/logging"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var (
	enumsAsInts  = flag.Bool("enums_as_ints", false, "render enums as integers as opposed to strings")
	emitDefaults = flag.Bool("emit_defaults", false, "render fields with zero values")
	origName     = flag.Bool("orig_name", false, "use original (.proto) name for fields")
	allowUnknown = flag.Bool("allow_unknown", false, "allow messages to contain unknown fields when unmarshaling")
	debug        = flag.Bool("debug", false, "enable debug mode")
)

func main() {
	flag.Parse()

	logging.Setup(debug)

	protogen.Options{ParamFunc: flag.CommandLine.Set}.Run(func(gp *protogen.Plugin) error {
		gp.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		opts := internal.Options{
			EnumsAsInts:        *enumsAsInts,
			EmitDefaults:       *emitDefaults,
			OrigName:           *origName,
			AllowUnknownFields: *allowUnknown,
		}

		for _, name := range gp.Request.FileToGenerate {
			f := gp.FilesByPath[name]

			if len(f.Messages) == 0 {
				slog.Info(fmt.Sprintf("Skipping %s, no messages", name))
				continue
			}

			slog.Info(fmt.Sprintf("Processing %s", name))
			slog.Info(fmt.Sprintf("Generating %s.pb.json.go", f.GeneratedFilenamePrefix))

			gf := gp.NewGeneratedFile(fmt.Sprintf("%s.json.pb.go", f.GeneratedFilenamePrefix), f.GoImportPath)

			err := internal.ApplyTemplate(gf, f, opts)
			if err != nil {
				gf.Skip()
				gp.Error(err)
				continue
			}
		}

		return nil
	})
}

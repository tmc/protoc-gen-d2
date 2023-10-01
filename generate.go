package main

import (
	"bytes"
	"context"

	"google.golang.org/protobuf/compiler/protogen"
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2layouts/d2elklayout"
	"oss.terrastruct.com/d2/d2lib"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
	"oss.terrastruct.com/d2/lib/textmeasure"
)

// generateFile processes a single proto file.
func generateFile(gen *protogen.Plugin, f *protogen.File, skipTypeRe string) {
	// capture to bytes buffer
	buf := new(bytes.Buffer)

	// execute template with protobuf File schema
	if err := tmpl.Execute(buf, f); err != nil {
		gen.Error(err)
		return
	}
	gf := gen.NewGeneratedFile(f.GeneratedFilenamePrefix+".d2", f.GoImportPath)
	if _, err := gf.Write(buf.Bytes()); err != nil {
		gen.Error(err)
		return
	}

	svgf := gen.NewGeneratedFile(f.GeneratedFilenamePrefix+".svg", f.GoImportPath)
	svg, err := renderSvg(buf.String())
	if err != nil {
		gen.Error(err)
		return
	}
	if _, err := svgf.Write(svg); err != nil {
		gen.Error(err)
		return
	}
}

// renderSvg is a helper function that calls the D2 libraries to render the SVG output
func renderSvg(contents string) ([]byte, error) {
	ruler, _ := textmeasure.NewRuler()
	defaultLayout := func(engine string) (d2graph.LayoutGraph, error) {
		return d2elklayout.DefaultLayout, nil
	}
	diagram, _, err := d2lib.Compile(context.Background(), contents, &d2lib.CompileOptions{
		LayoutResolver: defaultLayout,
		Ruler:          ruler,
	}, &d2svg.RenderOpts{})
	if err != nil {
		return nil, err
	}
	return d2svg.Render(diagram, &d2svg.RenderOpts{})
}

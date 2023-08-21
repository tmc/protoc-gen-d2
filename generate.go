package main

import (
	"bytes"
	"context"
	"io/ioutil"

	"google.golang.org/protobuf/compiler/protogen"
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2layouts/d2elklayout"
	"oss.terrastruct.com/d2/d2lib"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
	"oss.terrastruct.com/d2/d2themes/d2themescatalog"
	"oss.terrastruct.com/d2/lib/textmeasure"
	"oss.terrastruct.com/util-go/go2"
)

// generateFile processes a single proto file.
func generateFile(gen *protogen.Plugin, f *protogen.File) {
	// capture to bytes buffer
	buf := new(bytes.Buffer)

	// execute template with protobuf File schema
	if err := tmpl.Execute(buf, f); err != nil {
		gen.Error(err)
		return
	}
	// write out d2 file:
	if err := ioutil.WriteFile(f.GeneratedFilenamePrefix+".erd.d2", buf.Bytes(), 0600); err != nil {
		gen.Error(err)
		return
	}

	// call the helper function to render the SVG
	if err := renderSvg(buf.String(), f.GeneratedFilenamePrefix+".erd.svg"); err != nil {
		gen.Error(err)
		return
	}
}

// renderSvg is a helper function that calls the D2 libraries to render the SVG output
func renderSvg(contents string, outFilePath string) error {
	ruler, _ := textmeasure.NewRuler()
	defaultLayout := func(engine string) (d2graph.LayoutGraph, error) {
		return d2elklayout.DefaultLayout, nil
	}
	diagram, _, err := d2lib.Compile(context.Background(), contents, &d2lib.CompileOptions{
		LayoutResolver: defaultLayout,
		Ruler:          ruler,
	}, &d2svg.RenderOpts{})
	if err != nil {
		return err
	}
	out, err := d2svg.Render(diagram, &d2svg.RenderOpts{
		ThemeID: go2.Pointer(d2themescatalog.GrapeSoda.ID),
	})
	if err != nil {
		return err
	}
	return ioutil.WriteFile(outFilePath, out, 0600)
}

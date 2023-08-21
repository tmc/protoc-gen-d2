package main

import (
	"text/template"

	"google.golang.org/protobuf/compiler/protogen"
)

// tmpl is the template used for generating D2 diagrams from proto files.
var tmpl = template.Must(template.New("d2Diagram").Funcs(template.FuncMap{
	"d2Name": func(msg *protogen.Message) string {
		return msg.GoIdent.GoName
	},
	"fieldType": func(field *protogen.Field) string {
		// Check if the field is a map
		if field.Desc.IsMap() {
			keyType := field.Desc.MapKey().Kind().String()
			valueType := field.Desc.MapValue().Kind().String()
			return "map<" + keyType + ", " + valueType + ">"
		}

		// Handle other field types
		switch {
		case field.Enum != nil:
			return field.Enum.GoIdent.GoName
		case field.Message != nil:
			return field.Message.GoIdent.GoName
		default:
			return field.Desc.Kind().String()
		}
	},
}).Parse(`
{{- $packageName := .Desc.Package }}
# Package - {{$packageName}}
{{- range .Services}}
{{- $service := . }}
# Service - {{.GoName}}
{{.GoName}}: {
  shape: class

  {{- range .Methods}}
	+{{.GoName}}(input: {{if .Desc.IsStreamingClient}}stream {{end}}{{.Input.GoIdent.GoName}}): ({{if .Desc.IsStreamingServer}}stream {{end}}{{.Output.GoIdent.GoName}})
  {{- end}}
}

# Edges from package to service
"{{ $packageName }}" -> {{ .GoName }}

{{- range .Methods}}
{{- $method := . }}
# Edges from service to RPC types
{{$service.GoName}} <- {{.Input.GoIdent.GoName}}: {{$method.GoName}}
{{$service.GoName}} -> {{.Output.GoIdent.GoName}}: {{$method.GoName}}


{{- end}}
{{- end}}

{{- range .Messages}}
{{ $message := . }}
# Class - {{d2Name .}}
{{d2Name .}}: {
  shape: class
  {{- range .Fields}}
  {{.Desc.TextName}}: {{fieldType .}}
  {{- end}}
}
{{- range .Fields}}
  {{- if or .Enum (and .Message (not .Desc.IsMap)) }}
{{d2Name $message}} -> {{fieldType .}}
  {{- end }}
{{- end }}

{{- end}}

{{ range .Enums}}
# Enum - {{.GoIdent.GoName}}
{{.GoIdent.GoName}}: {
  grid-columns: 2
  grid-gap: 0
  {{- range .Values}}
  {{.GoIdent.GoName}}
  {{- end}}
  {{- range .Values}}
  {{.Desc.Number}}
  {{- end}}
}
{{- end}}

`))

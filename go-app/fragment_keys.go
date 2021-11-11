package main

import (
	"net/url"
	"text/template"

	"github.com/emicklei/renderbee"
	"github.com/kramphub/kiya/backend"
)

var TableOfKeys_Template = template.Must(template.New("TableOfKeys").Parse(`<html>
<table style="font-size:large">
	<tr>
		<th>Name</th>
	</tr>
{{ range .Keys }}
	<tr>
		<td><a href="/fetch?name={{.urlescape .Name}}">{{.Name}}</a></td>
	</tr>
{{ end }}
</table>
`))

type TableOfKeys struct {
	Keys []backend.Key
}

func (f TableOfKeys) RenderOn(hc *renderbee.HtmlCanvas) {
	fm := template.FuncMap{}
	fm["urlescape"] = url.QueryEscape
	TableOfKeys_Template.Funcs(fm)
	TableOfKeys_Template.Execute(hc, f)
}

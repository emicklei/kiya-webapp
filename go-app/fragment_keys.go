package main

import (
	"log"
	"net/url"
	"text/template"

	"github.com/emicklei/renderbee"
	"github.com/kramphub/kiya/backend"
)

var TableOfKeys_Template = template.Must(template.New("TableOfKeys").Funcs(template.FuncMap{"urlescape": url.QueryEscape}).Parse(`<html>
<table style="font-size:large">
{{ range .Keys }}
	<tr>
		<td><a href="javascript:fetchAndCopy('{{urlescape .Name}}')">{{.Name}}</a></td
	</tr>
{{ end }}
</table>
`))

type TableOfKeys struct {
	Keys []backend.Key
}

func (f TableOfKeys) RenderOn(hc *renderbee.HtmlCanvas) {
	if err := TableOfKeys_Template.Execute(hc, f); err != nil {
		log.Println(err)
	}
}

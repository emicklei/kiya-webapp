package main

import (
	"bytes"
	"testing"

	"github.com/emicklei/renderbee"
	"github.com/kramphub/kiya/backend"
)

func TestTableOfKeys_RenderOn(t *testing.T) {
	buf := new(bytes.Buffer)
	canvas := renderbee.NewHtmlCanvas(buf)
	table := TableOfKeys{Keys: []backend.Key{{Name: "test/me"}}}
	table.RenderOn(canvas)
	t.Log(buf.String())
}

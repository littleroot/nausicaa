package ui

// Code generated by webgen. DO NOT EDIT.

import (
	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/dom"
	"github.com/gowebapi/webapi/html"
	"github.com/gowebapi/webapi/html/canvas"
	"github.com/gowebapi/webapi/html/media"
)

type (
	_ *webapi.Document // prevent unused import errors
	_ *dom.Element
	_ *html.HTMLDivElement
	_ *canvas.HTMLCanvasElement
	_ *media.HTMLAudioElement
)

var (
	_document = webapi.GetDocument()
)

// source: testdata/standalone/multipleRoots.html

type multipleRoots struct {
	roots []*dom.Element
}

func newMultipleRoots() *multipleRoots {
	li0 := _document.CreateElement("li", nil)
	const stringliteral0 = "hello"
	li0.SetTextContent(&stringliteral0)
	li1 := _document.CreateElement("li", nil)
	const stringliteral1 = "world"
	li1.SetTextContent(&stringliteral1)
	return &multipleRoots{
		roots: []*dom.Element{li0, li1},
	}
}

func (v *multipleRoots) Roots() []*dom.Element {
	return v.roots
}

// source: testdata/include/includeMultipleRoots.html

type includeMultipleRoots struct {
	roots []*dom.Element
}

func newIncludeMultipleRoots() *includeMultipleRoots {
	div0 := _document.CreateElement("div", nil)
	include0 := newMultipleRoots()
	for _, r := range include0.roots {
		div0.AppendChild(&r.Node)
	}
	return &includeMultipleRoots{
		roots: []*dom.Element{div0},
	}
}

func (v *includeMultipleRoots) Roots() []*dom.Element {
	return v.roots
}

// source: testdata/include/multilevel.html

type multilevel struct {
	roots []*dom.Element
}

func newMultilevel() *multilevel {
	div0 := _document.CreateElement("div", nil)
	include0 := newIncludeMultipleRoots()
	for _, r := range include0.roots {
		div0.AppendChild(&r.Node)
	}
	return &multilevel{
		roots: []*dom.Element{div0},
	}
}

func (v *multilevel) Roots() []*dom.Element {
	return v.roots
}

package missed

import (
	"embed"
	"github.com/microcosm-cc/bluemonday"
)

//go:embed index.html
var IndexHtml []byte

//go:embed js/* img/* css/*
var Js embed.FS

var bm = bluemonday.StrictPolicy()

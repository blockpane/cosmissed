package missed

import "embed"

//go:embed index.html
var IndexHtml []byte

//go:embed js/* img/*
var Js embed.FS

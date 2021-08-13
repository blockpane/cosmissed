package missed

import (
	"embed"
	"github.com/microcosm-cc/bluemonday"
	"net/http"
)

//go:embed index.html
var IndexHtml []byte

//go:embed js/* img/* css/*
var Js embed.FS

var (
	bm               = bluemonday.StrictPolicy()
	TClient, CClient *http.Client
	TUrl, CUrl       string
)

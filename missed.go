package missed

import (
	"embed"
	"github.com/microcosm-cc/bluemonday"
	"github.com/oschwald/geoip2-golang"
	"log"
	"mime"
	"net/http"
)

//go:embed index.html
var IndexHtml []byte

//go:embed network.html
var NetHtml []byte

//go:embed js/* img/* css/*
var StaticContent embed.FS

var (
	bm               = bluemonday.StrictPolicy()
	TClient, CClient *http.Client
	TUrl, CUrl       string
	GeoDb            *geoip2.Reader
)

func init() {
	_ = mime.AddExtensionType(".html", "text/html; charset=UTF-8")
	_ = mime.AddExtensionType(".js", "application/javascript")
	_ = mime.AddExtensionType(".css", "text/css")

	var err error
	//FIXME: set flag for file location
	GeoDb, err = geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Println("error opening GeoLite2-City.mmdb, geoip features disabled:", err.Error())
	}
}
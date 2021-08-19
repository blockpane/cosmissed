package missed

import (
	"embed"
	"github.com/microcosm-cc/bluemonday"
	"log"
	"mime"
	"net/http"
	"os"
)

//go:embed index.html
var IndexHtml []byte

//go:embed network.html
var NetHtml []byte

//go:embed missed.html
var MissedHtml []byte

//go:embed js/* img/* css/*
var StaticContent embed.FS

var (
	bm        = bluemonday.StrictPolicy()
	CClient   *http.Client
	TClient   *http.Client
	NetworkId string
	CUrl      string
	TUrl      string
	//GeoDb     *geoip2.Reader
	l         *log.Logger
	peerMap   PeerMap
)

func init() {
	l = log.New(os.Stderr, "cosmissed | ", log.Lshortfile|log.LstdFlags)
	_ = mime.AddExtensionType(".html", "text/html; charset=UTF-8")
	_ = mime.AddExtensionType(".js", "application/javascript")
	_ = mime.AddExtensionType(".css", "text/css")

	//var err error
	////FIXME: set flag for file location
	//GeoDb, err = geoip2.Open("GeoLite2-City.mmdb")
	//if err != nil {
	//	l.Println("error opening GeoLite2-City.mmdb, geoip features disabled:", err.Error())
	//}

	peerMap = make([]PeerSet, 0)

}

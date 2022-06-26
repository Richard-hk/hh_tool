package tool

import (
	"hh_tool/util"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetSiteDoc(url string) *goquery.Document {
	resp, err := http.Get(url)
	util.HandleError(err, "GetSiteStr failed")
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	util.HandleError(err, "GetSiteDoc failed")
	return doc
}

package urlinfo

import (
	"bytes"
	"hh_tool/util"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetUrlDoc(url string) *goquery.Document {
	resp, err := http.Get(url)
	util.HandleError(err, "GetSiteStr failed")
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	util.HandleError(err, "GetUrlDoc failed")
	return doc
}

func GetUrlResp(url string, json_data []byte) *http.Response {
	resp, err := http.Post(url, "application/json",
		bytes.NewBuffer(json_data))
	util.HandleError(err, "GetUrlResp failed")
	return resp
}

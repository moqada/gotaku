// gotaku is a package of Web GYOTAKU Client.

package gotaku

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"golang.org/x/net/html"

	"github.com/PuerkitoBio/goquery"
)

var host = "http://megalodon.jp"
var validURL = regexp.MustCompile(`http://megalodon\.jp/\d{4}-\d{4}-\d{4}-\d+/`)

func getUrls(doc *goquery.Document) []string {
	var urls []string
	doc.Find(".container a[id^=fish]").Each(func(i int, s *goquery.Selection) {
		val, exists := s.Attr("href")
		if exists {
			urls = append(urls, val)
		}
	})
	return urls
}

func noRedirect(req *http.Request, via []*http.Request) error {
	return fmt.Errorf("do redirect")
}

// Take a GYOTAKU of target and return url
func Take(target string) (string, error) {
	v := url.Values{}
	v.Set("url", target)
	client := &http.Client{CheckRedirect: noRedirect}
	res, err := client.PostForm(host+"/pc/get_simple/decide", v)
	if err != nil && err.(*url.Error).Err.Error() != "do redirect" {
		return "", err
	}
	if res.StatusCode == 302 {
		urls, ok := res.Header["Location"]
		if ok && validURL.MatchString(urls[0]) {
			return urls[0], nil
		}
	}
	return "", fmt.Errorf("Redirect URL not found")
}

// List GYOTAKU URLs of target url
func List(target string) ([]string, error) {
	fmt.Println("Listing...")
	v := url.Values{}
	v.Set("url", target)
	doc, err := goquery.NewDocument(host + "/?" + v.Encode())
	if err != nil {
		return nil, err
	}
	return getUrls(doc), nil
}

// ListFromHTML return GYOTAKU URLs from HTML string
func ListFromHTML(content string) ([]string, error) {
	r := strings.NewReader(content)
	node, err := html.Parse(r)
	doc := goquery.NewDocumentFromNode(node)
	if err != nil {
		return nil, err
	}
	return getUrls(doc), nil
}

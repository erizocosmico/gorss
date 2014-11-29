// gorss is a library for easily parse RSS feeds and convert them to Go structs
//
//  package main
//
//  import "github.com/mvader/gorss"
//  import "fmt"
//
//  func main() {
//    rss, _ := gorss.LoadFeed("url to my feed")
//
//    fmt.Println(*rss)
//  }
package rss

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"
)

type Rss struct {
	Rss      xml.Name  `xml:"rss"`
	Channels []Channel `xml:"channel"`
}

type Channel struct {
	Title          string           `xml:"title"`
	Link           string           `xml:"link"`
	Description    string           `xml:"articledescription"`
	Language       string           `xml:"language"`
	Copyright      string           `xml:"copyright"`
	PubDate        string           `xml:"pubDate"`
	ManagingEditor string           `xml:"managingEditor"`
	WebMaster      string           `xml:"webMaster"`
	LastBuildDate  string           `xml:"lastBuildDate"`
	Categories     []string         `xml:"category"`
	Ttl            int64            `xml:"ttl"`
	Generator      string           `xml:"generator"`
	Docs           string           `xml:"docs"`
	Cloud          ChannelCloud     `xml:"cloud"`
	Rating         string           `xml:"rating"`
	TextInput      ChannelTextInput `xml:"textInput"`
	SkipHours      ChannelSkipHours `xml:"skipHours"`
	SkipDays       ChannelSkipDays  `xml:"skipDays"`
	Image          ChannelImage     `xml:"image"`
	Items          []Item           `xml:"item"`
}

type ChannelTextInput struct {
	Title       string `xml:"title"`
	Name        string `xml:"name"`
	Link        string `xml:"link"`
	Description string `xml:"articledescription"`
}

type ChannelSkipHours struct {
	Hours []int64 `xml:"hour"`
}

type ChannelSkipDays struct {
	Days []string `xml:"day"`
}

type ChannelCloud struct {
	Domain            string `xml:"domain,attr"`
	Port              int64  `xml:"port,attr"`
	Path              string `xml:"path,attr"`
	RegisterProcedure string `xml:"registerProcedure,attr"`
	Protocol          string `xml:"protocol,attr"`
}

type ChannelImage struct {
	Title       string `xml:"title"`
	URL         string `xml:"url"`
	Link        string `xml:"link"`
	Width       int64  `xml:"width"`
	Height      int64  `xml:"height"`
	Description string `xml:"articledescription"`
}

type Item struct {
	Title          string         `xml:"title"`
	Description    string         `xml:"articledescription"`
	Author         string         `xml:"author"`
	Creator        string         `xml:"creator"`
	Link           string         `xml:"link"`
	Categories     []ItemCategory `xml:"category"`
	Comments       string         `xml:"comments"`
	MediaContent   []MediaContent `xml:"content"`
	MediaThumbnail MediaThumbnail `xml:"thumbnail"`
	Guid           string         `xml:"guid"`
	PubDate        string         `xml:"pubDate"`
	Source         ItemSource     `xml:"source"`
	Enclosure      ItemEnclosure  `xml:"enclosure"`
}

type ItemEnclosure struct {
	URL    string `xml:"url,attr"`
	Type   string `xml:"type,attr"`
	Length int64  `xml:"length,attr"`
}

type ItemSource struct {
	URL   string `xml:"url,attr"`
	Value string `xml:",chardata"`
}

type ItemCategory struct {
	Value  string `xml:",chardata"`
	Domain string `xml:"domain,attr"`
}

type MediaTypeValue struct {
	Value string `xml:",chardata"`
	Type  string `xml:"type,attr"`
}

type MediaContent struct {
	URL    string         `xml:"url,attr"`
	Medium string         `xml:"medium,attr"`
	Width  int64          `xml:"width,attr"`
	Height int64          `xml:"height,attr"`
	Type   string         `xml:"type,attr"`
	Title  MediaTypeValue `xml:"title"`
}

type MediaThumbnail struct {
	URL    string `xml:"url,attr"`
	Width  int64  `xml:"width,attr"`
	Height int64  `xml:"height,attr"`
}

// DecodeFeed decodes the content of a feed into an Rss struct.
func DecodeFeed(content string) (*Rss, error) {
	// Collides with <media:description> sometimes
	content = strings.Replace(strings.Replace(content, "<description>", "<articledescription>", -1), "</description>", "</articledescription>", -1)
	rss := &Rss{}
	if err := xml.Unmarshal([]byte(content), rss); err != nil {
		return nil, err
	}

	return rss, nil
}

// loadPage loads the response content of a page
func loadPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// LoadFeed loads and decodes an RSS feed.
func LoadFeed(url string) (*Rss, error) {
	// Load the feed contents
	content, err := loadPage(url)
	if err != nil {
		return nil, err
	}

	// Decode them
	rss, err := DecodeFeed(content)
	if err != nil {
		return nil, err
	}

	return rss, nil
}

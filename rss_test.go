package rss

import "testing"
import "reflect"

var expectedRSS = `
<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:atom="http://www.w3.org/2005/Atom" xmlns:media="http://search.yahoo.com/mrss/" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd">
	<channel>
		<title>Title</title>
		<link>Link</link>
		<description>Description</description>
		<copyright>Copyright</copyright>
		<pubDate>PubDate</pubDate>
		<language>Language</language>
		<managingEditor>ManagingEditor</managingEditor>
		<webMaster>WebMaster</webMaster>
		<lastBuildDate>LastBuildDate</lastBuildDate>
		<category>Category 1</category>
		<category>Category 2</category>
		<ttl>20</ttl>
		<generator>Generator</generator>
		<docs>Docs</docs>
		<cloud domain="Domain" port="8080"
			path="/path" registerProcedure="RegisterProcedure"
			protocol="Protocol" />
		<rating>Rating</rating>
		<textInput>
			<title>Title</title>
			<name>Name</name>
			<link>Link</link>
			<description>Description</description>
		</textInput>
		<skipHours>
			<hour>1</hour>
			<hour>2</hour>
			<hour>3</hour>
			<hour>4</hour>
		</skipHours>
		<skipDays>
			<day>monday</day>
			<day>tuesday</day>
		</skipDays>
		<image>
			<title>Title</title>
			<url>URL</url>
			<link>Link</link>
			<width>80</width>
			<height>80</height>
			<description>Description</description>
		</image>

		<item>
			<title>Title</title>
			<description>Description</description>
			<author>Author</author>
			<dc:creator>Creator</dc:creator>
			<link>Link</link>
			<category domain="Domain">Value</category>
			<category domain="Domain">Value</category>
			<comments>Comments</comments>
			<media:content url="URL" medium="Medium" width="80" height="80" type="Type">
				<media:title type="html">Title</media:title>
			</media:content>
			<media:thumbnail url="URL" width="80" height="80" />
			<guid>Guid</guid>
			<pubDate>PubDate</pubDate>
			<source url="URL">Value</source>
			<enclosure url="URL" type="Type" length="80" />
		</item>

		<item>
			<title>Title</title>
			<description>Description</description>
			<author>Author</author>
			<dc:creator>Creator</dc:creator>
			<link>Link</link>
			<category domain="Domain">Value</category>
			<category domain="Domain">Value</category>
			<comments>Comments</comments>
			<media:content url="URL" medium="Medium" width="80" height="80" type="Type">
				<media:title type="html">Title</media:title>
			</media:content>
			<media:thumbnail url="URL" width="80" height="80" />
			<guid>Guid</guid>
			<pubDate>PubDate</pubDate>
			<source url="URL">Value</source>
			<enclosure url="URL" type="Type" length="80" />
		</item>
	</channel>
</rss>
`

func Test_DecodeFeed(t *testing.T) {
	rss, err := DecodeFeed(expectedRSS)
	if err != nil {
		t.Errorf("expected err to be nil")
	}

	expectedTestRSS := buildTestRSS()
	if !reflect.DeepEqual(*expectedTestRSS, *rss) {
		t.Errorf("expected `%+v`\nTo be: `%+v`\n", *rss, *expectedTestRSS)
	}
}

func buildTestRSS() *Rss {
	rss, _ := DecodeFeed(`<rss version="2.0" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:atom="http://www.w3.org/2005/Atom" xmlns:media="http://search.yahoo.com/mrss/" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd"></rss>`)
	rss.Channels = []Channel{buildTestChannel()}
	return rss
}

func buildTestChannel() Channel {
	return Channel{
		Title:          "Title",
		Link:           "Link",
		Description:    "Description",
		Language:       "Language",
		Copyright:      "Copyright",
		PubDate:        "PubDate",
		ManagingEditor: "ManagingEditor",
		WebMaster:      "WebMaster",
		LastBuildDate:  "LastBuildDate",
		Categories:     []string{"Category 1", "Category 2"},
		Ttl:            20,
		Generator:      "Generator",
		Docs:           "Docs",
		Rating:         "Rating",
		Cloud: ChannelCloud{
			Domain:            "Domain",
			Port:              8080,
			Path:              "/path",
			RegisterProcedure: "RegisterProcedure",
			Protocol:          "Protocol",
		},
		TextInput: ChannelTextInput{
			Title:       "Title",
			Name:        "Name",
			Link:        "Link",
			Description: "Description",
		},
		SkipHours: ChannelSkipHours{
			Hours: []int64{1, 2, 3, 4},
		},
		SkipDays: ChannelSkipDays{
			Days: []string{"monday", "tuesday"},
		},
		Image: ChannelImage{
			Title:       "Title",
			URL:         "URL",
			Link:        "Link",
			Width:       80,
			Height:      80,
			Description: "Description",
		},
		Items: []Item{buildTestItem(), buildTestItem()},
	}
}

func buildTestItem() Item {
	return Item{
		Title:       "Title",
		Description: "Description",
		Author:      "Author",
		Creator:     "Creator",
		Link:        "Link",
		Categories:  []ItemCategory{buildItemCategory(), buildItemCategory()},
		Comments:    "Comments",
		MediaContent: []MediaContent{
			MediaContent{
				URL:    "URL",
				Medium: "Medium",
				Width:  80,
				Height: 80,
				Type:   "Type",
				Title: MediaTypeValue{
					Type:  "html",
					Value: "Title",
				},
			},
		},
		MediaThumbnail: MediaThumbnail{
			URL:    "URL",
			Width:  80,
			Height: 80,
		},
		Guid:    "Guid",
		PubDate: "PubDate",
		Source: ItemSource{
			URL:   "URL",
			Value: "Value",
		},
		Enclosure: ItemEnclosure{
			URL:    "URL",
			Type:   "Type",
			Length: 80,
		},
	}
}

func buildItemCategory() ItemCategory {
	return ItemCategory{
		Domain: "Domain",
		Value:  "Value",
	}
}

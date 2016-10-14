package api

import "encoding/xml"

type TaggingAnalysis struct {
	XMLName xml.Name `xml:"cogito"`
	Doc
}

type Doc struct {
	XMLName xml.Name `xml:"doc"`
	Content
	Knowledges []Knowledge `xml:"knowledge"`
}

type Content struct {
	XMLName xml.Name `xml:"content"`
	Text    string   `xml:"text"`
}

type Knowledge struct {
	XMLName xml.Name `xml:"knowledge"`
}

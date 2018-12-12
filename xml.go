package main

import (
	"encoding/xml"
)

type Blogposts struct {
	XMLName   xml.Name   `xml:"output"`
	Blogposts []Blogpost `xml:"blogpost"`
}

type Blogpost struct {
	XMLName  xml.Name    `xml:"blogpost"`
	URL      string      `xml:"url"`
	UserID   string      `xml:"email"`
	Title    string      `xml:"title"`
	GUID     int         `xml:"guid"`
	Comments XMLComments `xml:"comments"`
}

type XMLComments struct {
	XMLName  xml.Name     `xml:"comments"`
	Comments []XMLComment `xml:"comment"`
}

type XMLComment struct {
	XMLName xml.Name `xml:"comment"`
	ID      string   `xml:"id,attr"`
	Pid     string   `xml:"parrentid,attr"`
	Text    string   `xml:"text"`
	Name    string   `xml:"name"`
	Score   int      `xml:"score"`
	Time    string   `xml:"date"`
}

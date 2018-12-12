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
	Pid     string   `xml:"parentid,attr"`
	Text    string   `xml:"text"`
	Name    string   `xml:"name"`
	UserID  string   `xml:"email"`
	IP      string   `xml:"ip"`
	Score   int      `xml:"score"`
	Time    string   `xml:"date"`
}

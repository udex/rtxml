package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	prepURL = "https://radio-t.com/site-api/last/%v?categories=prep"
	podURL  = " https://radio-t.com/site-api/podcast/%v"
)

type Request struct {
	cache  map[int]string
	client http.Client
}

type Entry struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

func defineURL(s string) (string, error) {
	if strings.HasPrefix(s, "Темы для") {
		return "", nil
	} else if strings.HasPrefix(s, "Радио") {
		return "", nil
	}
	return "", errors.New("Ни шмагла")
}

func (r *Request) GetEnries(url string) []Entry {
	res, err := r.client.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var list []Entry
	err = json.Unmarshal(body, &list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

func (r *Request) cacheURLs() {
	next := r.nextEpisode()
	url := fmt.Sprintf(prepURL, next)
	list := r.GetEnries(url)
	for i := range list {
		t := list[i].Title
		n := extractNumber(t)
		r.cache[n] = list[i].URL
	}
}

func (r *Request) nextEpisode() int {
	url := fmt.Sprintf(prepURL, 1)
	list := r.GetEnries(url)
	return extractNumber(list[0].Title)

}

func extractNumber(title string) int {
	f := strings.Fields(title)
	i, err := strconv.Atoi(f[len(f)-1])
	if err != nil {
		log.Fatal(err)
	}
	return i

}

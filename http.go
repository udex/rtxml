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
	podURL  = "https://radio-t.com/site-api/podcast/%v"
)

type Request struct {
	cache  map[int]string
	client http.Client
}

type Entry struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

func newReqest() Request {
	r := Request{
		cache:  map[int]string{},
		client: http.Client{},
	}
	r.cacheURLs()
	return r
}

func defineURL(r *Request, s string) (string, error) {
	if strings.HasPrefix(s, "Темы для") {
		num, _ := extractNumber(s)
		apiURL := fmt.Sprintf(podURL, num)
		e := r.Entry(apiURL)
		return e.URL, nil
	} else if strings.HasPrefix(s, "Радио") {
		num, _ := extractNumber(s)
		e := r.Entry(r.cache[num])
		return e.URL, nil
	}
	return "", errors.New("Ни шмагла")
}

func (r *Request) Entries(url string) []Entry {
	body := r.Body(url)
	var list []Entry
	err := json.Unmarshal(body, &list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

func (r *Request) Entry(url string) Entry {
	body := r.Body(url)
	var e Entry
	err := json.Unmarshal(body, &e)
	if err != nil {
		log.Fatal(err)
	}
	return e
}

func (r *Request) Body(url string) []byte {
	res, err := r.client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func (r *Request) cacheURLs() {
	next := r.nextEpisode()
	url := fmt.Sprintf(prepURL, next)
	list := r.Entries(url)
	for i := range list {
		t := list[i].Title
		n, err := extractNumber(t)
		if err == nil {
			r.cache[n] = list[i].URL
		}

	}
}

func (r *Request) nextEpisode() int {
	url := fmt.Sprintf(prepURL, 1)
	list := r.Entries(url)
	i, err := extractNumber(list[0].Title)
	if err != nil {
		log.Fatal(err)
	}
	return i

}

func extractNumber(title string) (int, error) {
	f := strings.Fields(title)
	last := f[len(f)-1]
	if strings.HasPrefix(last, "#") {
		last = last[1:]
	}
	i, err := strconv.Atoi(last)
	if err != nil {
		return 0, errors.New("Неизвестный пост")
	}
	return i, nil

}

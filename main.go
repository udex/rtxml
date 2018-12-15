package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// XMLname := flag.String("input", "", "путь до xml файла для парсинга")
	prepf := flag.String("prep", "", "путь до json файла с категорией prep, полученно из api сайта radio-t.com")
	podcastf := flag.String("podcast", "", "путь до json файла с категорией podcast, полученно из api сайта radio-t.com")
	//outf := flag.String("output", "", "путь до json файла, в который будет записан результат")
	flag.Parse()

	// Проверка, введены ли все аргументы
	set := []string{*prepf, *podcastf}
	for _, f := range set {
		if f == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}
	}

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var posts Blogposts
	err = xml.Unmarshal(data, &posts)
	if err != nil {
		log.Fatal(err)
	}
	r := newRt(*podcastf, *prepf)

	cs := commgen(posts, &r)
	st := jsongen(cs)
	for line := range st {
		fmt.Println(line)
	}
}

func read(p string) []byte {
	f, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func write(b []byte, p string) {
	f, err := os.Create(p)
	if err != nil {
		log.Fatal(err)
	}
	i, err := f.Write(b)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d bytes has been written\n", i)

}

package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	XMLname := flag.String("input", "", "путь до xml файла для парсинга")
	prepf := flag.String("prep", "", "путь до json файла с категорией prep, полученно из api сайта radio-t.com")
	podcastf := flag.String("podcast", "", "путь до json файла с категорией podcast, полученно из api сайта radio-t.com")
	outf := flag.String("output", "", "путь до json файла, в который будет записан результат")
	flag.Parse()

	// Проверка, введены ли все аргументы
	set := []string{*XMLname, *prepf, *podcastf, *outf}
	for _, f := range set {
		if f == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}
	}
	data := read(*XMLname)
	var posts Blogposts
	err := xml.Unmarshal(data, &posts)
	r := newRt(*podcastf, *prepf)
	b, err := encodeJSON(posts, &r)
	if err != nil {
		log.Fatal(err)
	}
	write(b, *outf)
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
	fmt.Printf("%d bytes has been written\n", i)

}

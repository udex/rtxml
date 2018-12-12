package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	var posts Blogposts
	err = xml.Unmarshal(data, &posts)
	b, err := encodeJSON(posts)
	if err != nil {
		log.Fatal(err)
	}
	s := string(b)
	fmt.Print(s)

}

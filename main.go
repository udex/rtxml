package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	//data, err := ioutil.ReadAll(os.Stdin)
	data := getdata(path.Join("test_data", "input.xml"))
	// if err != nil {
	// 	panic(err)
	// }

	var posts Blogposts
	//var output Users
	err := xml.Unmarshal(data, &posts)
	fmt.Println(err)
	for _, p := range posts.Blogposts {
		for _, c := range p.Comments.Comments {
			fmt.Println(c)
		}
	}
}

func getdata(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	d, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return d
}

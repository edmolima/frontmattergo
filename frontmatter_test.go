package frontmattergo_test

import (
	"fmt"
	"log"

	"github.com/edmolima/frontmattergo"
)

var markdown = []byte(`---
	title: Hello World
	authors:
		- Edmo
		- Lima
	---
	Lorem ipsum dolor sit ammet
`)

type article struct {
	Title   string
	Authors []string
}

func ExampleUnmarshal() {
	var a article

	content, err := frontmattergo.Unmarshal(markdown, &a)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Title: %s\nAuthors: %v\nContent: %s\n", a.Title, a.Authors, content)
}

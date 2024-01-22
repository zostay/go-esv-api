package main

import (
	"fmt"
	"os"

	"github.com/zostay/go-esv-api/pkg/esv"
)

func main() {
	client := esv.New(os.Getenv("ESV_API_KEY"))

	passages, err := client.PassageSearch("resurrection")
	if err != nil {
		panic(err)
	}

	for _, passage := range passages.Results {
		fmt.Printf("%s: %s\n", passage.Reference, passage.Content)
	}
}

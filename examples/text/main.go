package main

import (
	"fmt"
	"os"

	"github.com/zostay/go-esv-api/pkg/esv"
)

func main() {
	client := esv.New(os.Getenv("ESV_API_KEY"))

	passage, err := client.PassageText("John 3:16")
	if err != nil {
		panic(err)
	}

	fmt.Println(passage.Passages[0])
}

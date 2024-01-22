package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/zostay/go-esv-api/pkg/esv"
)

func main() {
	client := esv.New(os.Getenv("ESV_API_KEY"))

	passage, err := client.PassageHtml("John 3:16")
	if err != nil {
		panic(err)
	}

	hf, err := os.Create("john316.html")
	if err != nil {
		panic(err)
	}
	defer hf.Close()

	fmt.Fprint(hf, passage.Passages[0])

	openCmd := ""
	switch runtime.GOOS {
	case "darwin":
		openCmd = "open"
	case "linux":
		openCmd = "xdg-open"
	case "windows":
		openCmd = "start"
	default:
		panic("unsupported platform")
	}

	err = exec.Command(openCmd, "john316.html").Run()
	if err != nil {
		panic(err)
	}
}

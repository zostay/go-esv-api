# ESV API for Go

This is a small SDK for allowing Go programs to easily contact the ESV Bible API
at <https://api.esv.org/>.

# Installation

Before using this library, you will need to get an ESV API key from the [ESV 
API](https://api.esv.org/).

This provides no front-end for the end user. If you are interested in a 
front-end, see the [today](https://github.com/zostay/today) project.

To use this package in your project, import it in the usual Go way:

```shell
go get github.com/zostay/go-esv-api
```

# Usage

To use the API, you will need to construct a client, and then make API calls 
via the returned SDK client.

```go
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
```

# SDK Guide

## Options

The optional parameters for the API calls are provided using a variadic list 
of options at the end of the API calls. All the option functions begin with 
`With` in the option constructor name.

This SDK shares options across endpoints, so you'll need to see the 
[documentation for the ESV API](https://api.esv.org/) to know which options 
apply to each API endpoint.

## Passage Text

The `PassageText` method will return plain text for the queried passage or 
passages.

```go
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
```

## Passage HTML

The `PassageHtml` method will return HTML for the queried passage or passages.

```go
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
```

## Passage Search

The `PassageSearch` method will return a list of passages that match the
provided query.

```go
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
```

## Passage Audio

This is not yet implemented. See below.

# Implemented vs. To Do

The passage text, passage HTML, and search endpoints are implemented.

The audio endpoint is not working. However, that it should be very simple to
provide the audio link given by the redirect. I just haven't written the code
for that into the generator template.

# Copyright & License

Copyright 2020 Andrew Sterling Hanenkamp

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

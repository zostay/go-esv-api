package esv

import (
	_ "embed"
	"strings"
)

//go:embed version.txt
var version string

// Version returns the version of the ESV API client.
func Version() string {
	return strings.TrimSpace(version)
}

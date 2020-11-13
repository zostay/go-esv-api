package esv

import (
	"net/url"
	"strconv"
)

// Location is a reference provided by the ESV API into a Bible passage's
// location.
type Location uint

// Option provides the interface that all method options used by the API must
// follow.
type Option interface {
	// UpdateQuery should modify the given query form to add any arguments
	// necessary to apply this option.
	UpdateQuery(url.Values)
}

type optionBool struct {
	name  string
	value bool
}

func (o optionBool) UpdateQuery(q url.Values) {
	q.Add(o.name, strconv.FormatBool(o.value))
}

type optionInt struct {
	name  string
	value int
}

func (o optionInt) UpdateQuery(q url.Values) {
	q.Add(o.name, strconv.FormatInt(int64(o.value), 10))
}

type optionString struct {
	name  string
	value string
}

func (o optionString) UpdateQuery(q url.Values) {
	q.Add(o.name, o.value)
}

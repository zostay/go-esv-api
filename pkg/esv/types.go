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

type OptionBool struct {
	Name  string
	Value bool
}

func (o OptionBool) UpdateQuery(q url.Values) {
	q.Add(o.Name, strconv.FormatBool(o.Value))
}

type OptionInt struct {
	Name  string
	Value int
}

func (o OptionInt) UpdateQuery(q url.Values) {
	q.Add(o.Name, strconv.FormatInt(int64(o.Value), 10))
}

type OptionString struct {
	Name  string
	Value string
}

func (o OptionString) UpdateQuery(q url.Values) {
	q.Add(o.Name, o.Value)
}

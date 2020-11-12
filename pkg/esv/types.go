package esv

import (
	"net/url"
	"strconv"
)

type Location uint

type Client struct{}

type Option interface {
	UpdateQuery(q url.Values)
}

type OptionBool struct {
	name  string
	value bool
}

func (o OptionBool) UpdateQuery(q url.Values) {
	q.Add(o.name, strconv.FormatBool(o.value))
}

type OptionInt struct {
	name  string
	value int
}

func (o OptionInt) UpdateQuery(q url.Values) {
	q.Add(o.name, strconv.FormatInt(int64(o.value), 10))
}

type OptionString struct {
	name  string
	value string
}

func (o OptionString) UpdateQuery(q url.Values) {
	q.Add(o.name, o.value)
}

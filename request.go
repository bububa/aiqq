package aiqq

import (
	"net/url"
)

type Request interface {
	Path() string
	Method() string
	Values() url.Values
}

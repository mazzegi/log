package log

import "github.com/mazzegi/log/entry"

type Filter struct {
	accept func(e entry.Entry) bool
	next   Writer
}

func NewFilter(accept func(e entry.Entry) bool, next Writer) *Filter {
	return &Filter{
		accept: accept,
		next:   next,
	}
}

func (f *Filter) Write(e entry.Entry) {
	if !f.accept(e) {
		return
	}
	f.next.Write(e)
}

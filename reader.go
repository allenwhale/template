package template

import (
	"bytes"
)

type TemplateReader struct {
	name   string
	text   []byte
	pos    int
	length int
}

func (t *TemplateReader) find(needle []byte, args ...int) int {
	start, index := 0, -1
	if len(args) >= 1 {
		start = args[0]
	}
	start += t.pos
	index = bytes.Index(t.text[start:], needle)
	return index
}

func (t *TemplateReader) consume(args ...int) []byte {
	count := t.length - t.pos
	if len(args) >= 1 {
		count = args[0]
	}
	newPos := t.pos + count
	res := t.text[t.pos:newPos]
	t.pos = newPos
	return res
}

func (t *TemplateReader) remaining() int {
	return t.length - t.pos
}

func (t *TemplateReader) get(pos int) byte {
	return t.text[t.pos+pos]
}

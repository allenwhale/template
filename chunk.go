package template

import (
	"bytes"
	"html"
	"reflect"
	"strings"
)

type TextChunk struct {
	contents string
}

func (c TextChunk) generate(t *Template, writer *bytes.Buffer, data GenerateData) {
	writer.WriteString(c.contents)
}

type ExpressionChunk struct {
	contents string
	raw      bool
}

func (c ExpressionChunk) generate(t *Template, writer *bytes.Buffer, data GenerateData) {
	res := tostring(data[string(c.contents)])
	if c.raw == false {
		res = html.EscapeString(res)
	}
	writer.WriteString(res)
}

type IncludeChunk struct {
	name string
}

func (c IncludeChunk) generate(t *Template, writer *bytes.Buffer, data GenerateData) {
	t.Render(writer, c.name, data)
}

type StatementChunk struct {
	contents string
}

func (c StatementChunk) generate(t *Template, writer *bytes.Buffer, data GenerateData) {
	splited := strings.Split(c.contents, "=")
	left := strings.Split(splited[0], ",")
	right := splited[1]
	res := eval(right, data)
	for idx, l := range left {
		data[string(strip([]byte(l)))] = res[idx]
	}
}

type ControlChunk struct {
	contents string
	body     []Chunk
}

func (c ControlChunk) generate(t *Template, writer *bytes.Buffer, data GenerateData) {
	panic(c.contents)
}

type IfChunk struct {
	condition string
	body      []Chunk
}

func (c IfChunk) generate(t *Template, writer *bytes.Buffer, data GenerateData) {
	cond := eval(c.condition, data)[0]
END_COND:
	for _, chunk := range c.body {
		switch chunk.(type) {
		case ElseIfChunk, ElseIfInitChunk:
			if cond == true {
				break END_COND
			}
			func() {
				defer func() {
					if err := recover(); err == true {
						cond = true
					}
				}()
				cond = false
				chunk.generate(t, writer, data)
			}()
		case ElseChunk:
			if cond == true {
				break END_COND
			}
			cond = true
		default:
			if cond == true {
				chunk.generate(t, writer, data)
			}
		}
	}
}

type IfInitChunk struct {
	init      string
	condition string
	body      []Chunk
}

func (c IfInitChunk) generate(t *Template, writer *bytes.Buffer, data GenerateData) {
	StatementChunk{contents: c.init}.generate(t, writer, data)
	cond := eval(c.condition, data)[0]
END_COND:
	for _, chunk := range c.body {
		switch chunk.(type) {
		case ElseIfChunk, ElseIfInitChunk:
			if cond == true {
				break END_COND
			}
			func() {
				defer func() {
					if err := recover(); err == true {
						cond = true
					}
				}()
				cond = false
				chunk.generate(t, writer, data)
			}()
		case ElseChunk:
			if cond == true {
				break END_COND
			}
			cond = true
		default:
			if cond == true {
				chunk.generate(t, writer, data)
			}
		}
	}
}

type ElseIfChunk struct {
	condition string
}

func (c ElseIfChunk) generate(t *Template, writer *bytes.Buffer, data GenerateData) {
	panic(eval(c.condition, data)[0])
}

type ElseIfInitChunk struct {
	init      string
	condition string
}

func (c ElseIfInitChunk) generate(t *Template, writer *bytes.Buffer, data GenerateData) {
	StatementChunk{contents: c.init}.generate(t, writer, data)
	panic(eval(c.condition, data)[0])
}

type ElseChunk struct{}

func (c ElseChunk) generate(t *Template, writer *bytes.Buffer, data GenerateData) {}

type ForChunk struct {
	init      string
	condition string
	after     string
	body      []Chunk
}

func (c ForChunk) generate(t *Template, writer *bytes.Buffer, data GenerateData) {
	StatementChunk{contents: c.init}.generate(t, writer, data)
END_LOOP:
	for eval(c.condition, data)[0] == true {
		isBreak, isContinue := false, false
		for _, chunk := range c.body {
			func() {
				defer func() {
					if err := recover(); err == "break" {
						isBreak = true
					} else if err == "continue" {
						isContinue = true
					}
				}()
				chunk.generate(t, writer, data)
			}()
			if isBreak {
				break END_LOOP
			} else if isContinue {
				break
			}
		}
		StatementChunk{contents: c.after}.generate(t, writer, data)
	}
}

type ForRangeChunk struct {
	idx  string
	item string
	iter string
	body []Chunk
}

func (c ForRangeChunk) generate(t *Template, writer *bytes.Buffer, data GenerateData) {
	iter := reflect.ValueOf(data[c.iter])
END_LOOP:
	for idx := 0; idx < iter.Len(); idx = idx + 1 {
		if c.idx != "_" {
			data[c.idx] = idx
		}
		if c.item != "_" {
			data[c.item] = iter.Index(idx)
		}
		isBreak, isContinue := false, false
		for _, chunk := range c.body {
			func() {
				defer func() {
					if err := recover(); err == "break" {
						isBreak = true
					} else if err == "continue" {
						isContinue = true
					}
				}()
				chunk.generate(t, writer, data)
			}()
			if isBreak {
				break END_LOOP
			} else if isContinue {
				break
			}
		}
	}
}

type WhileChunk struct {
	condition string
	body      []Chunk
}

func (c WhileChunk) generate(t *Template, writer *bytes.Buffer, data GenerateData) {
END_LOOP:
	for eval(c.condition, data)[0] == true {
		isBreak, isContinue := false, false
		for _, chunk := range c.body {
			func() {
				defer func() {
					if err := recover(); err == "break" {
						isBreak = true
					} else if err == "continue" {
						isContinue = true
					}
				}()
				chunk.generate(t, writer, data)
			}()
			if isBreak {
				break END_LOOP
			} else if isContinue {
				break
			}
		}
	}
}

type Chunk interface {
	generate(t *Template, writer *bytes.Buffer, data GenerateData)
}

package template

import (
	"bytes"
)

func Parse(reader *TemplateReader, inBlock string, inLoop string) []Chunk {
	var body []Chunk
	for {
		curly := 0
		for {
			curly = reader.find([]byte("{"), curly)
			if curly == -1 || curly+1 == reader.remaining() {
				//EOF
				if inBlock != "" {
					panic("inBlock")
				}
				body = append(body, TextChunk{contents: string(reader.consume())})
				return body
			}
			if contain(reader.get(curly+1), []byte("{%#")) == false {
				curly++
				continue
			}
			if curly+2 < reader.remaining() && reader.get(curly+1) == '{' && reader.get(curly+2) == '{' {
				curly++
				continue
			}
			break
		}
		if curly > 0 {
			contents := reader.consume(curly)
			body = append(body, TextChunk{contents: string(contents)})
		}
		startBrace := reader.consume(2)
		//jquery template
		if reader.remaining() > 0 && reader.get(0) == '!' {
			reader.consume(1)
			body = append(body, TextChunk{contents: string(startBrace)})
			continue
		}
		//comment
		if bytes.Compare(startBrace, []byte("{#")) == 0 {
			end := reader.find([]byte("#}"))
			if end == -1 {
				panic("Missing end comment #}")
			}
			reader.consume(end)
			reader.consume(2)
			continue
		}
		//expression
		if bytes.Compare(startBrace, []byte("{{")) == 0 {
			end := reader.find([]byte("}}"))
			if end == -1 {
				panic("Missing end expression }}")
			}
			contents := strip(reader.consume(end))
			reader.consume(2)
			if len(contents) == 0 {
				panic("Empty expression")
			}

			body = append(body, ExpressionChunk{contents: string(contents)})
			continue
		}
		if bytes.Compare(startBrace, []byte("{%")) != 0 {
			panic(startBrace)
		}
		//block
		end := reader.find([]byte("%}"))
		if end == -1 {
			panic("Missing end block %}")
		}
		contents := strip(reader.consume(end))
		reader.consume(2)
		if len(contents) == 0 {
			panic("Empty block tag {% %}")
		}
		operator, _, suffix := partition(contents, []byte(" "))
		suffix = strip(suffix)
		allowedParents := intermediateBlock[string(operator)]
		if len(allowedParents) != 0 {
			if inBlock == "" {
				panic("Outside block")
			}
			if contain(inBlock, allowedParents) == false {
				panic("Now allowed")
			}
			if bytes.Compare(operator, []byte("elif")) == 0 {
				if ifInitRegexp.Match(suffix) {
					r := ifInitRegexp.FindSubmatch(suffix)
					init, condition := r[1], r[2]
					body = append(body, ElseIfInitChunk{
						init:      string(init),
						condition: string(condition),
					})
				} else {
					body = append(body, ElseIfChunk{condition: string(suffix)})
				}
			} else if bytes.Compare(operator, []byte("else")) == 0 {
				body = append(body, ElseChunk{})
			}
			continue
		} else if bytes.Compare(operator, []byte("end")) == 0 {
			//end tag
			if len(inBlock) == 0 {
				panic("Extra {% end %} block")
			}
			return body
		} else if contain(string(operator), []string{"include", "set", "comment", "raw"}) {
			var block Chunk
			if bytes.Compare(operator, []byte("comment")) == 0 {
				continue
			} else if bytes.Compare(operator, []byte("include")) == 0 {
				suffix = strip(strip(suffix, byte('\'')), '"')
				if len(suffix) == 0 {
					panic("Include missing file path")
				}
				block = IncludeChunk{name: string(suffix)}
			} else if bytes.Compare(operator, []byte("set")) == 0 {
				if len(suffix) == 0 {
					panic("Set missing statement")
				}
				if contain(byte('='), suffix) == false {
					panic("Set missing assignment")
				}
				block = StatementChunk{
					contents: string(suffix),
				}
			} else if bytes.Compare(operator, []byte("raw")) == 0 {
				block = ExpressionChunk{
					contents: string(suffix),
					raw:      true,
				}
			}
			body = append(body, block)
			continue
		} else if contain(string(operator), []string{"if", "for"}) {
			var blockBody []Chunk
			if bytes.Compare(operator, []byte("for")) == 0 {
				blockBody = Parse(reader, string(operator), string(operator))
				if forRegexp.Match(suffix) {
					r := forRegexp.FindSubmatch(suffix)
					init, condition, after := r[1], r[2], r[3]
					body = append(body, ForChunk{
						init:      string(init),
						condition: string(condition),
						after:     string(after),
						body:      blockBody,
					})
				} else if forRangeRegexp.Match(suffix) {
					r := forRangeRegexp.FindSubmatch(suffix)
					idx, item, iter := r[1], r[2], r[3]
					body = append(body, ForRangeChunk{
						idx:  string(idx),
						item: string(item),
						iter: string(iter),
						body: blockBody,
					})
				} else {
					body = append(body, WhileChunk{
						condition: string(suffix),
						body:      blockBody,
					})
				}
			} else if bytes.Compare(operator, []byte("if")) == 0 {
				blockBody = Parse(reader, string(operator), inLoop)
				if ifInitRegexp.Match(suffix) {
					r := ifInitRegexp.FindSubmatch(suffix)
					init, condition := r[1], r[2]
					body = append(body, IfInitChunk{
						init:      string(init),
						condition: string(condition),
						body:      blockBody,
					})
				} else {
					body = append(body, IfChunk{
						condition: string(suffix),
						body:      blockBody,
					})
				}
			}
			continue
		} else if contain(string(operator), []string{"continue", "break"}) {
			if len(inLoop) == 0 {
				panic("Not in loop")
			}
			body = append(body, ControlChunk{contents: string(contents)})
		}
	}
}

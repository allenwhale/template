package template

import (
	"fmt"
	"testing"
)

type testdata struct {
	input  string
	output string
}

var testdatas = []testdata{
	//comment
	{
		input:  "{#tt#}",
		output: "[{}]",
	},
	{
		input:  "aaa{#tt#}aaa",
		output: "[{aaa} {aaa}]",
	},
	//Text
	{
		input:  "Hello, World",
		output: "[{Hello, World}]",
	},
	//Statement
	{
		input:  "{%set a = 10%}",
		output: "[{a = 10} {}]",
	},
	{
		input:  "{%set a, b = 10, 10%}",
		output: "[{a, b = 10, 10} {}]",
	},
	//If
	{
		input:  "{% if a = 10 %}if{%end%}",
		output: "[{a = 10 [{if}]} {}]",
	},
	{
		input:  "{%if a=10;a<10%}if{%end%}",
		output: "[{a=10 a<10 [{if}]} {}]",
	},
	//else
	{
		input:  "{%if a=10;a<10%}if{%else%}else{%end%}",
		output: "[{a=10 a<10 [{if} {} {else}]} {}]",
	},
	//elif
	{
		input:  "{%if a=10;a<10%}if{%elif b<10%}elif{%end%}",
		output: "[{a=10 a<10 [{if} {b<10} {elif}]} {}]",
	},
	{
		input:  "{%if a=10;a<10%}if{%elif b=20;b<20%}elif{%end%}",
		output: "[{a=10 a<10 [{if} {b=20 b<20} {elif}]} {}]",
	},
	{
		input:  "{%if a=10;a<10%}if{%elif b<10%}elif{%else%}else{%end%}",
		output: "[{a=10 a<10 [{if} {b<10} {elif} {} {else}]} {}]",
	},
	{
		input:  "{%if a=10;a<10%}if{%elif b=20;b<20%}elif{%else%}else{%end%}",
		output: "[{a=10 a<10 [{if} {b=20 b<20} {elif} {} {else}]} {}]",
	},
	//expression
	{
		input:  "{{a}}xx{{b}}yy{{c}}",
		output: "[{a false} {xx} {b false} {yy} {c false} {}]",
	},
	{
		input:  "{%raw a%}{{a}}{%raw b%}",
		output: "[{a true} {a false} {b true} {}]",
	},
	//for
	{
		input:  "{%for a=10;a<100;a=a+1%}for{%end%}",
		output: "[{a=10 a<100 a=a+1 [{for}]} {}]",
	},
	{
		input:  "{%set a = 0%}{%for a<10%}for{%end%}",
		output: "[{a = 0} {a<10 [{for}]} {}]",
	},
	{
		input:  "{%for _, a := range []int{1,2,3,4,5}%}{{a}}{%end%}",
		output: "[{_, a := range []int{1,2,3,4,5} [{a false}]} {}]",
	},
	//break
	{
		input:  "{%set a = 0%}{%for%}{%break%}{%end%}",
		output: "[{a = 0} { [{break []}]} {}]",
	},
	//continue
	{
		input:  "{%set a = 10%}{%for%}{%continue%}{%end%}",
		output: "[{a = 10} { [{continue []}]} {}]",
	},
	//mix
	{
		input:  "{%set a = 10%}{%for a < 10%}{%if a > 10%}{%for b=0;b<10;b=b+1%}for{%end%}{%else%}{{exp}}{%end%}{%end%}",
		output: "[{a = 10} {a < 10 [{a > 10 [{b=0 b<10 b=b+1 [{for}]} {} {exp false}]}]} {}]",
	},
	{
		input:  "{%for a=0;a<10;a++%}{{a}}{%if a>10%}{%break%}{%end%}{%if a<10%}{%continue%}{%elif u=10;u<10%}else{%end%}{%raw a%}{%end%}",
		output: "[{a=0 a<10 a++ [{a false} {a>10 [{break []}]} {a<10 [{continue []} {u=10 u<10} {else}]} {a true}]} {}]",
	},
	{
		input:  "{%for a < 10%}{%for a < 20%}{%for a < 30%}{%for a < 40%}{%end%}{%end%}{%end%}{%end%}",
		output: "[{a < 10 [{a < 20 [{a < 30 [{a < 40 []}]}]}]} {}]",
	},
	{
		input:  "{%for a < 10%}{%raw a%}{%for a < 20%}{%for a < 30%}{%if y=20;y<10%}{%else%}{%if 1 == 1%}{{a}}{%end%}{%end%}{%for a < 40%}{%end%}{%end%}{%end%}{%end%}",
		output: "[{a < 10 [{a true} {a < 20 [{a < 30 [{y=20 y<10 [{} {1 == 1 [{a false}]}]} {a < 40 []}]}]}]} {}]",
	},
}

func TestParse(t *testing.T) {
	for _, test := range testdatas {
		tr := TemplateReader{
			text:   []byte(test.input),
			length: len([]byte(test.input)),
		}
		p := fmt.Sprint(Parse(&tr, "", ""))
		if p != test.output {
			t.Error(string(test.input), "\nget   :", p, "\nexcept:", test.output)
		}
	}
}
func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testdatas {
			//test := testdatas[0]
			tr := TemplateReader{
				text:   []byte(test.input),
				length: len([]byte(test.input)),
			}
			p := fmt.Sprint(Parse(&tr, "", ""))
			if p != test.output {
				b.Error(string(test.input), "\nget   :", p, "\nexcept:", test.output)
			}
		}
	}
}

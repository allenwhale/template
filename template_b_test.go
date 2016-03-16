package template

import (
	//"fmt"
	"bytes"
	"io/ioutil"
	"testing"
)

func New() *Template {
	templ := NewTemplate()
	templ.LoadTemplate("test", []string{".templ"}, false)
	return templ
}

func BenchmarkTemplateSimple(b *testing.B) {
	b.ReportAllocs()
	filePath := "test/simple"
	templ := New()
	for i := 0; i < b.N; i++ {
		var writer bytes.Buffer
		templ.Render(&writer, filePath+".templ", GenerateData{})
		except, _ := ioutil.ReadFile(filePath + ".html")
		if bytes.Compare(except, writer.Bytes()) != 0 {
			b.Error(filePath + ".templ")
			b.Error("get:   ", len(writer.Bytes()))
			b.Error("except:", len(except))
		}
	}
}
func BenchmarkTemplateLong4k(b *testing.B) {
	b.ReportAllocs()
	filePath := "test/long4k"
	templ := New()
	for i := 0; i < b.N; i++ {
		var writer bytes.Buffer
		templ.Render(&writer, filePath+".templ", GenerateData{})
		except, _ := ioutil.ReadFile(filePath + ".html")
		if bytes.Compare(except, writer.Bytes()) != 0 {
			b.Error(filePath + ".templ")
			b.Error("get:   ", len(writer.Bytes()))
			b.Error("except:", len(except))
		}
	}
}
func BenchmarkTemplateLong128k(b *testing.B) {
	b.ReportAllocs()
	filePath := "test/long128k"
	templ := New()
	for i := 0; i < b.N; i++ {
		var writer bytes.Buffer
		templ.Render(&writer, filePath+".templ", GenerateData{})
		except, _ := ioutil.ReadFile(filePath + ".html")
		if bytes.Compare(except, writer.Bytes()) != 0 {
			b.Error(filePath + ".templ")
			b.Error("get:   ", len(writer.Bytes()))
			b.Error("except:", len(except))
		}
	}
}
func BenchmarkTemplateLongDiscrete4k(b *testing.B) {
	b.ReportAllocs()
	filePath := "test/long_discrete4k"
	templ := New()
	for i := 0; i < b.N; i++ {
		var writer bytes.Buffer
		templ.Render(&writer, filePath+".templ", GenerateData{})
		except, _ := ioutil.ReadFile(filePath + ".html")
		if bytes.Compare(except, writer.Bytes()) != 0 {
			b.Error(filePath + ".templ")
			b.Error("get:   ", len(writer.Bytes()))
			b.Error("except:", len(except))
		}
	}
}
func BenchmarkTemplateLongDiscrete128k(b *testing.B) {
	b.ReportAllocs()
	filePath := "test/long_discrete128k"
	templ := New()
	for i := 0; i < b.N; i++ {
		var writer bytes.Buffer
		templ.Render(&writer, filePath+".templ", GenerateData{})
		except, _ := ioutil.ReadFile(filePath + ".html")
		if bytes.Compare(except, writer.Bytes()) != 0 {
			b.Error(filePath + ".templ")
			b.Error("get:   ", len(writer.Bytes()))
			b.Error("except:", len(except))
		}
	}
}
func BenchmarkTemplateFor(b *testing.B) {
	b.ReportAllocs()
	filePath := "test/for"
	templ := New()
	for i := 0; i < b.N; i++ {
		var writer bytes.Buffer
		templ.Render(&writer, filePath+".templ", GenerateData{})
		except, _ := ioutil.ReadFile(filePath + ".html")
		if bytes.Compare(except, writer.Bytes()) != 0 {
			b.Error(filePath + ".templ")
			b.Error("get:   ", len(writer.Bytes()))
			b.Error("except:", len(except))
		}
	}
}
func BenchmarkTemplateForLarge(b *testing.B) {
	b.ReportAllocs()
	filePath := "test/for_large"
	templ := New()
	for i := 0; i < b.N; i++ {
		var writer bytes.Buffer
		templ.Render(&writer, filePath+".templ", GenerateData{})
		except, _ := ioutil.ReadFile(filePath + ".html")
		if bytes.Compare(except, writer.Bytes()) != 0 {
			b.Error(filePath + ".templ")
			b.Error("get:   ", len(writer.Bytes()))
			b.Error("except:", len(except))
		}
	}
}
func BenchmarkTemplateForNest(b *testing.B) {
	b.ReportAllocs()
	filePath := "test/for_nest"
	templ := New()
	for i := 0; i < b.N; i++ {
		var writer bytes.Buffer
		templ.Render(&writer, filePath+".templ", GenerateData{})
		except, _ := ioutil.ReadFile(filePath + ".html")
		if bytes.Compare(except, writer.Bytes()) != 0 {
			ioutil.WriteFile("tmp", writer.Bytes(), 0644)
			b.Error(filePath + ".templ")
			b.Error("get:   ", len(writer.Bytes()))
			b.Error("except:", len(except))
		}
	}
}
func BenchmarkTemplateForNestLarge(b *testing.B) {
	b.ReportAllocs()
	filePath := "test/for_nest_large"
	templ := New()
	for i := 0; i < b.N; i++ {
		var writer bytes.Buffer
		templ.Render(&writer, filePath+".templ", GenerateData{})
		except, _ := ioutil.ReadFile(filePath + ".html")
		if bytes.Compare(except, writer.Bytes()) != 0 {
			ioutil.WriteFile("tmp", writer.Bytes(), 0644)
			b.Error(filePath + ".templ")
			b.Error("get:   ", len(writer.Bytes()))
			b.Error("except:", len(except))
		}
	}
}
func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := 1
		b := 1
		c := a + b
		_ = c
	}
}

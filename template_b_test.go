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

func TestTemplateSimple(t *testing.T) {
	templ := New()
	var writer bytes.Buffer
	templ.Render(&writer, "test/simple.templ", GenerateData{})
	except, _ := ioutil.ReadFile("test/simple.html")
	if bytes.Compare(except, writer.Bytes()) != 0 {
		t.Error("test/simple.templ")
		t.Error("get:   ", len(writer.Bytes()))
		t.Error("except:", len(except))
	}
}
func BenchmarkTemplateSimple(b *testing.B) {
	filePath := "test/simple"
	templ := New()
	var writer bytes.Buffer
	templ.Render(&writer, filePath+".templ", GenerateData{})
	except, _ := ioutil.ReadFile(filePath + ".html")
	if bytes.Compare(except, writer.Bytes()) != 0 {
		b.Error(filePath + ".templ")
		b.Error("get:   ", len(writer.Bytes()))
		b.Error("except:", len(except))
	}
}
func BenchmarkTemplateLong4k(b *testing.B) {
	filePath := "test/long4k"
	templ := New()
	var writer bytes.Buffer
	templ.Render(&writer, filePath+".templ", GenerateData{})
	except, _ := ioutil.ReadFile(filePath + ".html")
	if bytes.Compare(except, writer.Bytes()) != 0 {
		b.Error(filePath + ".templ")
		b.Error("get:   ", len(writer.Bytes()))
		b.Error("except:", len(except))
	}
}
func BenchmarkTemplateLong128k(b *testing.B) {
	filePath := "test/long128k"
	templ := New()
	var writer bytes.Buffer
	templ.Render(&writer, filePath+".templ", GenerateData{})
	except, _ := ioutil.ReadFile(filePath + ".html")
	if bytes.Compare(except, writer.Bytes()) != 0 {
		b.Error(filePath + ".templ")
		b.Error("get:   ", len(writer.Bytes()))
		b.Error("except:", len(except))
	}
}
func BenchmarkTemplateLongDiscrete4k(b *testing.B) {
	filePath := "test/long_discrete4k"
	templ := New()
	var writer bytes.Buffer
	templ.Render(&writer, filePath+".templ", GenerateData{})
	except, _ := ioutil.ReadFile(filePath + ".html")
	if bytes.Compare(except, writer.Bytes()) != 0 {
		b.Error(filePath + ".templ")
		b.Error("get:   ", len(writer.Bytes()))
		b.Error("except:", len(except))
	}
}
func BenchmarkTemplateLongDiscrete128k(b *testing.B) {
	filePath := "test/long_discrete128k"
	templ := New()
	var writer bytes.Buffer
	templ.Render(&writer, filePath+".templ", GenerateData{})
	except, _ := ioutil.ReadFile(filePath + ".html")
	if bytes.Compare(except, writer.Bytes()) != 0 {
		b.Error(filePath + ".templ")
		b.Error("get:   ", len(writer.Bytes()))
		b.Error("except:", len(except))
	}
}
func BenchmarkTemplateFor(b *testing.B) {
	filePath := "test/for"
	templ := New()
	var writer bytes.Buffer
	templ.Render(&writer, filePath+".templ", GenerateData{})
	except, _ := ioutil.ReadFile(filePath + ".html")
	if bytes.Compare(except, writer.Bytes()) != 0 {
		b.Error(filePath + ".templ")
		b.Error("get:   ", len(writer.Bytes()))
		b.Error("except:", len(except))
	}
}
func BenchmarkTemplateForLarge(b *testing.B) {
	filePath := "test/for_large"
	templ := New()
	var writer bytes.Buffer
	templ.Render(&writer, filePath+".templ", GenerateData{})
	except, _ := ioutil.ReadFile(filePath + ".html")
	if bytes.Compare(except, writer.Bytes()) != 0 {
		b.Error(filePath + ".templ")
		b.Error("get:   ", len(writer.Bytes()))
		b.Error("except:", len(except))
	}
}
func BenchmarkTemplateForNest(b *testing.B) {
	filePath := "test/for_nest"
	templ := New()
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
func BenchmarkTemplateForNestLarge(b *testing.B) {
	filePath := "test/for_nest_large"
	templ := New()
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

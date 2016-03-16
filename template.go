package template

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mohae/deepcopy"
	//"io"
	"io/ioutil"
	"path/filepath"
	//"time"
)

func Generate(body []Chunk, t *Template, writer *bytes.Buffer, data GenerateData) {
	copyData := deepcopy.Iface(data).(GenerateData)
	for _, chunk := range body {
		chunk.generate(t, writer, copyData)
	}
}

type Template struct {
	templ map[string][]Chunk
}

func NewTemplate() *Template {
	t := &Template{templ: map[string][]Chunk{}}
	return t
}

func (t *Template) LoadTemplate(dir string, allowExt []string, args ...interface{}) {
	debug := true
	if len(args) >= 1 {
		debug = args[0].(bool)
	}
	dirs, _ := ioutil.ReadDir(dir)
	for _, path := range dirs {
		if path.IsDir() {
			t.LoadTemplate(dir+"/"+path.Name(), allowExt)
		} else {
			ext := filepath.Ext(path.Name())
			if contain(ext, allowExt) {
				filePath := dir + "/" + path.Name()
				if debug {
					fmt.Println("Parsing", filePath)
				}
				content, _ := ioutil.ReadFile(filePath)
				t.templ[filePath] = Parse(&TemplateReader{
					text:   content,
					length: len(content),
				}, "", "")
			}
		}
	}
}

func (t *Template) Render(writer *bytes.Buffer, filePath string, data GenerateData) {
	Generate(t.templ[filePath], t, writer, data)
}

func (t *Template) GinRender(c *gin.Context, filePath string, data GenerateData) {
	c.Writer.Header().Set("Content-Type", "text/html charset=utf8")
	var writer bytes.Buffer
	t.Render(&writer, filePath, data)
	writer.WriteTo(c.Writer)
}

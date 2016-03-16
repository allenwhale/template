package template

import (
	"bytes"
	"fmt"
	"github.com/xtaci/goeval"
	"reflect"
	"regexp"
)

var intermediateBlock = map[string][]string{
	"else": []string{"if"},
	"elif": []string{"if"},
}

type GenerateData map[string]interface{}

var ifInitRegexp = regexp.MustCompile(`(?P<init>[^;]+)\s*;(?P<condition>[^;]+)\s*`)
var forRegexp = regexp.MustCompile(`(?P<init>[^;]+)\s*;\s*(?P<condition>[^;]+)\s*;\s*(?P<after>[^;]+)`)
var forRangeRegexp = regexp.MustCompile(`(?P<idx>\w+)\s*,\s*(?P<item>\w+)\s*:?=\s*range\s*(?P<iter>\w+)`)

func tostring(obj interface{}) string {
	return fmt.Sprint(obj)
}

func eval(exp string, data GenerateData) []interface{} {
	s := goeval.NewScope()
	for key, value := range data {
		s.Set(key, value)
	}
	if exp == "" {
		exp = "true"
	}
	evalRes, _ := s.Eval("return " + exp)
	var res []interface{}
	switch evalRes.(type) {
	case []interface{}:
		res = evalRes.([]interface{})
	default:
		res = []interface{}{evalRes}
	}
	return res
}

func contain(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	for i := 0; i < targetValue.Len(); i++ {
		if targetValue.Index(i).Interface() == obj {
			return true
		}
	}
	return false
}

func partition(str []byte, sep []byte) ([]byte, []byte, []byte) {
	idx := bytes.Index(str, sep)
	var first, second, third []byte
	if idx == -1 {
		first, second, third = str, []byte(""), []byte("")
	} else {
		first = str[:idx]
		if idx+len(sep) > len(str) {
			second = []byte("")
		} else {
			second = str[idx : idx+len(sep)]
		}
		if idx+len(sep) >= len(str) {
			third = []byte("")
		} else {
			third = str[idx+len(sep):]
		}
	}
	return first, second, third
}

func strip(str []byte, args ...byte) []byte {
	sep := byte(' ')
	if len(args) >= 1 {
		sep = args[0]
	}
	length := len(str)
	start, end := 0, length
	for ; start < end; start++ {
		if str[start] != sep {
			break
		}
	}
	for ; end > start; end-- {
		if str[end-1] != sep {
			break
		}
	}
	return str[start:end]
}

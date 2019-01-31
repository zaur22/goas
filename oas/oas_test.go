package oas

import (
	"io/ioutil"
	"path"
	"strings"
	"testing"
)

var from = "test_files"
var to = "test_files"
var fileName = "openapi"
var extension = "yaml"
var res = path.Join(to, "res.yaml")

func TestConvert(t *testing.T) {

	convert(from, to, fileName, extension)

	convRes, err := ioutil.ReadFile(path.Join(to, fileName+"."+extension))
	if err != nil {
		t.Fatal(err.Error())
	}

	need, err := ioutil.ReadFile(res)
	if err != nil {
		t.Fatal(err.Error())
	}

	strRes := strings.Split(string(convRes), "\n")
	strNeed := strings.Split(string(need), "\n")
	for i, res := range strRes {
		if res != strNeed[i] {
			t.Fatalf("line %v. Converted string is not eq to original. Converted:\n'%v'\n Original:\n'%v'", i+1, res, strNeed[i])
		}
	}
}

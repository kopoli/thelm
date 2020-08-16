package thelm

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func TestSourceReader(t *testing.T) {
	data := fmt.Sprintln("Something here")
	in := strings.NewReader(data)
	out := bytes.Buffer{}

	sr := SourceReader{
		Input: in,
	}
	sr.SetOutput(&out)

	err := sr.Run()
	if err != nil {
		t.Fatal("Run should succeed:", err)
	}

	err = sr.Finish()
	if err != nil {
		t.Fatal("Finish should succeed:", err)
	}

	if data != out.String() {
		t.Fatal("SourceReader should copy input to output")
	}
}

func TestSourcefile(t *testing.T) {
	var out bytes.Buffer

	filename := "../LICENSE"

	fp := SourceFile{
		FileName: filename,
	}

	fp.SetOutput(&out)
	err := fp.Run()
	if err != nil {
		t.Fatal("Run should succeed:", err)
	}

	err = fp.Finish()
	if err != nil {
		t.Fatal("Finish should succeed:", err)
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal("Internal error: reading file contents failed:", err)
	}

	if !reflect.DeepEqual(data, out.Bytes()) {
		t.Fatal("File read output should be equal")
	}
}

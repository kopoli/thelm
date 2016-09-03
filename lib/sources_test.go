package thelm

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
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
	require.NoError(t, err, "Run should succeed")

	err = sr.Finish()
	require.NoError(t, err, "Finish should succeed")

	require.Equal(t, data, out.String(), "SourceReader should copy input to output")
}

func TestSourcefile(t *testing.T) {
	var out bytes.Buffer

	filename := "../LICENSE"

	fp := SourceFile{
		FileName: filename,
	}

	fp.SetOutput(&out)
	err := fp.Run()
	require.NoError(t, err, "Run should succeed'")

	err = fp.Finish()
	require.NoError(t, err, "Finish should succeed'")

	data, err := ioutil.ReadFile(filename)

	require.NoError(t, err, "Internal error: reading file contents failed")
	require.EqualValues(t, data, out.Bytes(), "File read output should be equal")
}

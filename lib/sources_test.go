package thelm

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSourcefile(t *testing.T) {
	var out bytes.Buffer

	filename := "../LICENSE"

	fp := SourceFile{
		FileName: filename,
	}

	fp.SetOutput(&out)
	err := fp.Run()
	require.NoError(t, err, "Run shouldn't fail")

	err = fp.Finish()
	require.NoError(t, err, "Finish shouldn't fail")

	data, err := ioutil.ReadFile(filename)

	require.NoError(t, err, "Internal error: reading file contents failed")
	require.EqualValues(t, data, out.Bytes(), "File read output should be equal")
}

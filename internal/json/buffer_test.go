package json

import (
	"bytes"
	"errors"
	"testing"

	"github.com/i9si-sistemas/assert"
)

type testBuffer struct{
	bytes []byte
}

func (t testBuffer) Bytes() ([]byte, error) {
	return t.bytes, nil
}

func (t testBuffer) Buffer() (*bytes.Buffer, error) {
	return nil, errors.New("error not implemented")
}

func TestBuffer(t *testing.T) {
	tbuf := testBuffer{
		bytes: []byte("{\"message\": \"test\"}"),
	}
	b, err := RWBuffer(tbuf)
	assert.NoError(t, err)
	expectedBytes, err := tbuf.Bytes()
	assert.NoError(t, err)
	assert.Equal(t, b.Bytes(), expectedBytes)
}

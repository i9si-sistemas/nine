package e2e_test

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/i9si-sistemas/assert"
	"github.com/i9si-sistemas/nine"
	i9 "github.com/i9si-sistemas/nine/pkg/client"
	"github.com/i9si-sistemas/stringx"
)

func TestNineClient(t *testing.T) {
	client := nine.New(context.Background())
	res, err := client.Get("https://httpbin.org/get", new(i9.Options))
	if res.StatusCode >= 400 {
		t.Skip("httpbin.org is not available")
	}
	assertNoError := func(err error) {
		if err != nil {
			assert.True(t, stringx.String(err.Error()).Includes("nil"))
		}
	}
	assertNoError(err)
	defer res.Body.Close()

	var payload nine.JSON
	b, err := io.ReadAll(res.Body)
	assertNoError(err)
	err = nine.DecodeJSON(b, &payload)
	assertNoError(err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
	assert.Equal(t, payload["url"], "https://httpbin.org/get")
}

package client

import (
	"context"
	"net/http"

	public "github.com/i9si-sistemas/nine/pkg/client"
)

type client struct {
	ctx    context.Context
	client *http.Client
}

func (c *client) Context() context.Context {
	return c.ctx
}

// New creates a new HTTP client instance.
func New(
	ctx context.Context,
	clientConfig ...http.Client,
) *client {
	cl := &http.Client{}
	if len(clientConfig) > 0 {
		cl = &clientConfig[0]
	}
	return &client{ctx: ctx, client: cl}
}

// Get sends an HTTP GET request to the specified URL with the given options.
func (c *client) Get(url string, options *public.Options) (*http.Response, error) {
	return c.executeRequest(http.MethodGet, url, options)
}

// Post sends an HTTP POST request to the specified URL with the given options.
func (c *client) Post(url string, options *public.Options) (*http.Response, error) {
	return c.executeRequest(http.MethodPost, url, options)
}

// Put sends an HTTP PUT request to the specified URL with the given options.
func (c *client) Put(url string, options *public.Options) (*http.Response, error) {
	return c.executeRequest(http.MethodPut, url, options)
}

// Patch sends an HTTP PATCH request to the specified URL with the given options.
func (c *client) Patch(url string, options *public.Options) (*http.Response, error) {
	return c.executeRequest(http.MethodPatch, url, options)
}

// Delete sends an HTTP DELETE request to the specified URL with the given options.
func (c *client) Delete(url string, options *public.Options) (*http.Response, error) {
	return c.executeRequest(http.MethodDelete, url, options)
}

// Response sends an HTTP request and returns the corresponding response.
// It uses the underlying HTTP client to execute the request.
//
// Parameters:
//   - req: a pointer to the HTTP request that will be sent.
//
// Returns:
//   - *http.Response: the HTTP response received from the server.
//   - error: an error if the request failed, or nil if it was successful.
func (c *client) Response(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}

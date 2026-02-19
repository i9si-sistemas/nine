package nine

import (
	"context"
	"net/http"

	client "github.com/i9si-sistemas/nine/internal/client"
	types "github.com/i9si-sistemas/nine/pkg/client"
)

type Client interface {
	// Get sends an HTTP GET request to the specified URL with the given options.
	Get(url string, options *types.Options) (*http.Response, error)
	// Post sends an HTTP POST request to the specified URL with the given options.
	Post(url string, options *types.Options) (*http.Response, error)
	// Put sends an HTTP PUT request to the specified URL with the given options.
	Put(url string, options *types.Options) (*http.Response, error)
	// Patch sends an HTTP PATCH request to the specified URL with the given options.
	Patch(url string, options *types.Options) (*http.Response, error)
	// Delete sends an HTTP DELETE request to the specified URL with the given options.
	Delete(url string, options *types.Options) (*http.Response, error)
	// Context returns the context associated with the HTTP client.
	// This context can be used to control the lifecycle of HTTP requests,
	// allowing for cancellation, timeout.
	Context() context.Context
}

// New instantiates a new HTTP client with the provided context.
func New(ctx context.Context, clientConfig ...http.Client) Client {
	return client.New(ctx, clientConfig...)
}

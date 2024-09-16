package nine

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRequestBody(t *testing.T) {
	bodyContent := "test body"
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(bodyContent))

	request := &Request{req: req}
	body, err := io.ReadAll(request.Body())
	if err != nil {
		t.Fatalf("error reading request body: %v", err)
	}

	if string(body) != bodyContent {
		t.Errorf("expected '%s', got '%s'", bodyContent, string(body))
	}
}

func TestRequestHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-Test-Header", "header-value")

	request := &Request{req: req}
	value := request.Header("X-Test-Header")
	if value != "header-value" {
		t.Errorf("expected 'header-value', got '%s'", value)
	}
}

func TestRequestQuery(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?key=value", nil)
	request := &Request{req: req}

	queryValue := request.Query("key")
	if queryValue != "value" {
		t.Errorf("expected 'value', got '%s'", queryValue)
	}
}

func TestRequestContext(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	type (
		key   string
		value string
	)
	ctx := context.WithValue(context.Background(), key("message"), value("Hello Context"))
	req = req.WithContext(ctx)

	request := &Request{req: req}
	if request.Context().Value(key("message")) != value("Hello Context") {
		t.Errorf("expected 'value' in context, but got '%v'", request.Context().Value("key"))
	}
}

func TestResponseJSON(t *testing.T) {
	payload := JSON{
		"username": "gabrielluizsf",
	}
	handler := func(req *Request, res *Response) error {
		res.Status(http.StatusCreated).JSON(payload)
		return nil
	}

	h := httpHandler(handler)

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, w.Code)
	}
	var user struct {
		Username string `json:"username"`
	}
	if err := DecodeJSON(w.Body.Bytes(), &user); err != nil {
		t.Fatal(err)
	}
	if user.Username != payload["username"] {
		t.Fatal("invalid body")
	}
}

func TestServerError(t *testing.T) {
	w := httptest.NewRecorder()
	serverErr := &ServerError{
		StatusCode:  http.StatusInternalServerError,
		ContentType: "application/json",
		Err:         errors.New("internal server error"),
	}

	serverErr.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/", nil))

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, w.Code)
	}

	expectedBody := `{"err":"internal server error"}`
	if strings.TrimSpace(w.Body.String()) != expectedBody {
		t.Errorf("expected body '%s', got '%s'", expectedBody, w.Body.String())
	}
}

func TestResponseStatus(t *testing.T) {
	handler := func(req *Request, res *Response) error {
		res.SendStatus(http.StatusInternalServerError)
		return nil
	}

	h := httpHandler(handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, w.Code)
	}
}

func TestMiddleware(t *testing.T) {
	middleware := func(req *Request, res *Response) error {
		res.SetHeader("X-Middleware", "processed")
		return nil
	}
	message := "Hello World"
	handler := func(req *Request, res *Response) error {
		return res.Send([]byte(message))
	}

	finalHandler := httpMiddleware(middleware, httpHandler(handler))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	finalHandler.ServeHTTP(w, req)

	if w.Header().Get("X-Middleware") != "processed" {
		t.Errorf("expected 'X-Middleware' header to be 'processed'")
	}
	if w.Body.String() != message {
		t.Errorf("expected body 'Hello', got '%s'", w.Body.String())
	}
}
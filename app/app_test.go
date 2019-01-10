package app

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	teardown()
	os.Exit(retCode)
}

func setup() {
	gin.SetMode("test")
}

func teardown() {
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

func (s *Server) get(
	url string,
	query map[string]string,
	header map[string]string,
) (*http.Request, *httptest.ResponseRecorder, error) {
	return s.requestWithQuery("GET", url, query, header)
}

func (s *Server) delete(
	url string,
	query map[string]string,
	header map[string]string,
) (*http.Request, *httptest.ResponseRecorder, error) {
	return s.requestWithQuery("DELETE", url, query, header)
}

func (s *Server) putJSON(
	url string,
	body interface{},
	header map[string]string,
) (*http.Request, *httptest.ResponseRecorder, error) {
	return s.requestJSON("PUT", url, body, header)
}

func (s *Server) postJSON(
	url string,
	body interface{},
	header map[string]string,
) (*http.Request, *httptest.ResponseRecorder, error) {
	return s.requestJSON("POST", url, body, header)
}

func (s *Server) requestWithQuery(
	method string,
	url string,
	query map[string]string,
	header map[string]string,
) (*http.Request, *httptest.ResponseRecorder, error) {
	r, err := http.NewRequest(method, url, nil)
	if err != nil {
		return r, nil, err
	}

	if query != nil {
		values := r.URL.Query()
		for k, v := range query {
			values.Add(k, v)
		}
		r.URL.RawQuery = values.Encode()
	}

	if header != nil {
		for k, v := range header {
			r.Header.Set(k, v)
		}
	}

	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)
	return r, w, nil
}

func (s *Server) requestJSON(
	method string,
	url string,
	body interface{},
	header map[string]string,
) (*http.Request, *httptest.ResponseRecorder, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return nil, nil, err
	}

	r, err := http.NewRequest(method, url, &buf)
	if err != nil {
		return r, nil, err
	}

	if header != nil {
		for k, v := range header {
			r.Header.Set(k, v)
		}
	}

	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)
	return r, w, nil
}

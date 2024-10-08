package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/smarty/assertions/should"
)

func ServeRequest(path string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()
	NewHTTPRouter().ServeHTTP(response, request)
	return response
}
func assertResponse(t *testing.T, actual *httptest.ResponseRecorder, code int, contentType, body string) {
	t.Helper()
	should.So(t, actual.Header().Get("Content-Type"), should.Equal, contentType)
	should.So(t, actual.Code, should.Equal, code)
	should.So(t, strings.TrimSpace(actual.Body.String()), should.Equal, body)
}

func Test404(t *testing.T) {
	assertResponse(t, ServeRequest("/asdf"),
		http.StatusNotFound, "text/plain; charset=utf-8", "404 page not found")
}
func TestParametersMalformed(t *testing.T) {
	assertResponse(t, ServeRequest("/add?a=NaN&b=3"),
		http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "invalid parameter: 'a'")

	assertResponse(t, ServeRequest("/add?a=2&b=NaN"),
		http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "invalid parameter: 'b'")
}
func TestHappyPaths(t *testing.T) {
	assertResponse(t, ServeRequest("/add?a=2&b=3"), http.StatusOK, "text/plain; charset=utf-8", "5")
	assertResponse(t, ServeRequest("/mul?a=2&b=3"), http.StatusOK, "text/plain; charset=utf-8", "6")
	assertResponse(t, ServeRequest("/div?a=20&b=4"), http.StatusOK, "text/plain; charset=utf-8", "5")
	assertResponse(t, ServeRequest("/sub?a=6&b=2"), http.StatusOK, "text/plain; charset=utf-8", "4")
}

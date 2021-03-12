package mainimport

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	server := &Server{}
	t.Run("return Paco's score", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := createNewRequest()
		server.ServeHTTP(response, request)

		assertStatus(response.Code, http.StatusOK, t)
		assertResponseBody(t, "{\"message\": \"hello world\"}", response.Body.String())

	})
}

func createNewRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	return request
}

func assertStatus(result int, expected int, t *testing.T) {
	if result != expected {
		t.Errorf("wrong status status result '%d', status expected '%d'", result, expected)
	}
}

func assertResponseBody(t *testing.T, expected string, result string) {
	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}
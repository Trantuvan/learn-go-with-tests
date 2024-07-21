package contexts

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type SpyStore struct {
	response string
}

func (spy SpyStore) Fetch() string {
	return spy.response
}

func TestServer(t *testing.T) {
	data := "hello, world"
	//* notice: how to use interface
	//* Server defines param as Store interface
	//* invoke Server with a struct whose interface is Store
	server := Server(SpyStore{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	server(response, request)

	if response.Body.String() != data {
		t.Errorf("got %q, want %q", response.Body.String(), data)
	}
}

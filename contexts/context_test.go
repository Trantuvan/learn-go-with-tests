package contexts

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (spy *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return spy.response
}

func (spy *SpyStore) Cancel() {
	spy.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		//* notice: how to use interface
		//* Server defines param as Store interface
		//* invoke Server with a struct whose interface is Store
		server := Server(&SpyStore{response: data})

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server(response, request)

		if response.Body.String() != data {
			t.Errorf("got %q, want %q", response.Body.String(), data)
		}
	})
	t.Run("tells store to canel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		reponse := httptest.NewRecorder()

		server(reponse, request)

		if !store.cancelled {
			t.Errorf("store was not told to cancel")
		}
	})
}

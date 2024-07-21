package contexts

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func (spy *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return spy.response
}

func (spy *SpyStore) Cancel() {
	spy.cancelled = true
}

func TestServer(t *testing.T) {
	data := "hello, world"
	t.Run("returns data from store", func(t *testing.T) {
		//* notice: how to use interface
		//* Server defines param as Store interface
		//* invoke Server with a struct whose interface is Store
		store := &SpyStore{response: data, t: t}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server(response, request)

		//* stills make get request 1st
		if response.Body.String() != data {
			t.Errorf("got %q, want %q", response.Body.String(), data)
		}

		//* doesn't canceled request
		store.assertWasNotCancelled()
	})
	t.Run("tells store to canel work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		reponse := httptest.NewRecorder()

		server(reponse, request)

		store.assertWasCancelled()
	})
}

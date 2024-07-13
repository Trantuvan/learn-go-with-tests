package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of server, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(time.Millisecond * 20)
		fastServer := makeDelayedServer(time.Microsecond * 0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(fastURL, slowURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if want != got {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(time.Millisecond * 25)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, time.Millisecond*20)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}
func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

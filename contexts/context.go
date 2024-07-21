package contexts

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data := make(chan string)

		go func() {
			data <- store.Fetch()
		}()

		//* cancel get request in case request send cancel token
		//* select help call Cancel() which write to store struct
		//* and not call get request
		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}

package context_gonotes

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Server The thing our handler is responsible for is making sure it sends a context through to the downstream store,
//and that it handles the error that will come from the Store when it is cancelled.
func Server(store Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data, err := store.Fetch(request.Context())
		if err != nil {
			return //todo: log error
		}
		fmt.Fprint(writer, data)
	}
}

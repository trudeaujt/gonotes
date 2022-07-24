package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("return the fastest server", func(t *testing.T) {
		//this is not great... using external websites in our test can be slow, flaky, and won't let us test edge cases.
		//slowURL := "http://www.facebook.com"
		//fastURL := "http://www.quii.dev"

		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		//By prefixing a function call with defer, go will call that function at the end of the containing function.
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("return an error if a server doesn't respond within the specified time", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

//There is no extra magic here for testing...
//This is how you define a normal server in Go too.
func makeDelayedServer(delay time.Duration) *httptest.Server {
	//NewServer takes a http.HandlerFunc, which we are sending in via an anonymous function.
	return httptest.NewServer(http.HandlerFunc(
		//http.HandlerFunc is a type that looks like this:
		//type HandlerFunc func(ResponseWriter, r *request)
		//All this is really sayng is that the NewServer needs a function that takes a ResponseWrite and a Request.
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(http.StatusOK)
		}))
}

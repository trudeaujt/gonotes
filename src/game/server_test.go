package poker

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := NewGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		AssertResponseBody(t, response.Body.String(), "20")
		AssertStatus(t, response.Code, http.StatusOK)
	})
	t.Run("returns Floyd's score", func(t *testing.T) {
		request := NewGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		AssertResponseBody(t, response.Body.String(), "10")
		AssertStatus(t, response.Code, http.StatusOK)
	})
	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := NewGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		AssertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper"
		request := NewPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusAccepted)

		AssertPlayerWin(t, &store, player)
	})
}

//func TestConcurrency(t *testing.T) {
//	store := InMemoryPlayerStore{
//		map[string]int{},
//		sync.Mutex{},
//	}
//	server := NewPlayerServer(&store)
//
//	t.Run("it handles concurrent connections properly", func(t *testing.T) {
//		player := "Pepper"
//		request := NewPostWinRequest(player)
//		response := httptest.NewRecorder()
//
//		wantedCount := 1000
//
//		var wg sync.WaitGroup
//		wg.Add(wantedCount)
//
//		for i := 0; i < wantedCount; i++ {
//			go func() {
//				server.ServeHTTP(response, request)
//				wg.Done()
//			}()
//		}
//		wg.Wait()
//
//		AssertStatus(t, response.Code, http.StatusAccepted)
//		if server.store.GetPlayerScore(player) != wantedCount {
//			t.Errorf("got %d calls to RecordWin want %d", server.store.GetPlayerScore(player), 1000)
//		}
//	})
//}

func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Art", 10},
			{"Bob", 20},
			{"Chris", 30},
		}
		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		request := NewLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := GetLeagueFromResponse(t, response.Body)

		AssertStatus(t, response.Code, http.StatusOK)
		AssertLeague(t, got, wantedLeague)
		AssertContentType(t, response, jsonContentType)
	})
}

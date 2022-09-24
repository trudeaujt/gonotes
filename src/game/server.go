package poker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const jsonContentType = "application/json"

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

type Player struct {
	Name string
	Wins int
}

type PlayerServer struct {
	store PlayerStore
	//previously this was a named property, 'router http.ServeMux'
	//instead, we are now 'embedding' the http.Handler.
	http.Handler
	//What this means is that our PlayerServer now has all of the methods that http.Handler has, which is just ServeHTTP!
	//to 'fill this in', we assign it to the router down below in NewPlayerServer (1).
	//We can do this because http.ServeMux has the method ServeHTTP.
	//This further means we can delete our own ServeHTTP method, as we are already exposing one via the embedded type.
}

// We have moved the routing creation out of ServeHTTP and into this function so that this setup only has to be done once,
// not per request.
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	//(1)
	p.Handler = router
	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	leagueTable := p.getLeagueTable()
	json.NewEncoder(w).Encode(leagueTable)

	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.ProcessWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) getLeagueTable() []Player {
	return p.store.GetLeague()
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) ProcessWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

package poker

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initPlayerDBFile(file)

	if err != nil {
		return nil, fmt.Errorf("error initializing player db file, %v", err)
	}

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}
	return &FileSystemPlayerStore{
		//We don't need to create a new encoder every time we write.
		//So let's initialize one in the constructor and use that instead.
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}

func initPlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("error getting info from file %s, %v", file.Name(), err)
	}

	//if the file is empty, we write an empty JSON object to it.
	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}
	return nil
}

func (f *FileSystemPlayerStore) GetLeague() (league League) {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)
	if player == nil {
		f.league = append(f.league, Player{name, 1})
	} else {
		player.Wins++
	}

	f.database.Encode(f.league)
}

package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() (league []Player) {
	//without doing this, the file cannot be rewound - after parsing the league once it cannot be parsed again.
	f.database.Seek(0, 0)
	league, _ = NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	for _, p := range f.GetLeague() {
		if p.Name == name {
			return p.Wins
		}
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	for i, p := range league {
		if p.Name == name {
			//we can't do p[i].Wins++ because range returns copies of the element at the index.
			league[i].Wins++
		}
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}

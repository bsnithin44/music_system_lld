package domain

import (
	"spotif/pkg/domain/songs"
	"spotif/pkg/domain/users"
)

type MusicAppstruct struct {
	Name string
}

func MusicApp(name string) *MusicAppstruct {

	return &MusicAppstruct{
		Name: name,
	}

}

func (m MusicAppstruct) AddSong(song songs.Song) {
	sR := songs.GetSongsRepo()

	sR.AddSong(&song)

}

func (m MusicAppstruct) AddUser(user users.User) {
	uR := users.GetUserRepo()

	uR.CreateUser(&user)

}

func (m MusicAppstruct) RecommendSong() songs.Song {

	sR := songs.GetSongsRepo()

	// todo

	return *sR.Songs[0]

}

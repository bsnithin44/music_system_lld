package main

import (
	"fmt"
	"spotif/pkg/domain"
	"spotif/pkg/domain/songs"
	"spotif/pkg/domain/users"
)

func main() {

	mApp := domain.MusicApp("app1")

	s1 := songs.NewSong("stars", "nithin", "folk", 1)
	s2 := songs.NewSong("city life", "raghav", "indie", 2)
	u1 := users.NewUser("nithin", "bsnithin@gmail")

	mApp.AddUser(u1)
	mApp.AddSong(s1)
	mApp.AddSong(s2)

	u1.AddSongToPlaylist(s1)
	u1.AddSongToPlaylist(s2)

	plyst := u1.GetPlaylist()

	// print u1 songs
	for _, v := range plyst {
		fmt.Println(v.Name, v.Singer)
	}

}

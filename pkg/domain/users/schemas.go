package users

import "spotif/pkg/domain/songs"

type User struct {
	Name  string
	email string

	Playlist []songs.Song
}

func NewUser(name, email string) User {

	return User{
		Name:  name,
		email: email,
	}
}

func (usr *User) AddSongToPlaylist(song songs.Song) {

	UserRepo := GetUserRepo()

	usr.Playlist = append(usr.Playlist, song)
	UserRepo.UpdateUser(usr)

}

func (usr *User) GetPlaylist() []songs.Song {

	UserRepo := GetUserRepo()

	userData := UserRepo.GetUser(usr.Name)
	return userData.Playlist

}

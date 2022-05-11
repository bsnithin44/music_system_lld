package songs

type SongsRepo struct {
	isActive bool
	Songs    []*Song
}

type Isong interface {
	AddSong(song *Song)
	GetSong() *Song
	// DeleteSong(song *Song)
}

var sDB *SongsRepo

func GetSongsRepo() *SongsRepo {
	if sDB == nil || sDB.isActive == false {
		sDB = &SongsRepo{
			isActive: true,
		}
		return sDB
	}

	return sDB

}

func (sR SongsRepo) AddSong(song *Song) {

	sR.Songs = append(sR.Songs, song)

}

func (sR SongsRepo) GetSong(name string) *Song {
	for _, s := range sR.Songs {
		if s.Name == name {
			return s
		}
	}
	return nil
}

func (sR SongsRepo) GetSongs(name string) []*Song {

	return sR.Songs
}

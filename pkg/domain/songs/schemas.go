package songs

type Song struct {
	Name   string
	Singer string
	Genre  string
	Tempo  int
}

func NewSong(Name, Singer, Genre string, Tempo int) Song {

	return Song{
		Name:   Name,
		Singer: Singer,
		Genre:  Genre,
		Tempo:  Tempo,
	}
}

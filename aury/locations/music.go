package locations

type Music struct {
	Filepath string
}

func NewMusic(filepath string) *Music {
	return &Music{
		Filepath: filepath,
	}
}

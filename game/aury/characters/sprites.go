package characters

type Sprite struct {
	ID       SpriteID
	Filepath string
}

type SpriteID string

func NewSprite(id SpriteID, filepath string) *Sprite {
	return &Sprite{
		ID:       id,
		Filepath: filepath,
	}
}

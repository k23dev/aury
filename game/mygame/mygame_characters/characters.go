package mygame_characters

import (
	"aury/game/aury/aury_assets"
	"aury/game/aury/characters"
	"aury/game/aury/episodes"
)

func GamesCharactersSetup(episode *episodes.Episode) []characters.Character {

	chAury := characters.NewCharacter("aury", "Dominique Aury")

	// attributes
	chAury.AddAttribute("edad", "int", 50)
	chAury.AddAttribute("maldad", "int", 50)
	chAury.AddAttribute("moral", "bool", false)
	// sprites
	chAury.AddSprite("frente", aury_assets.Characters+"/aury/001.jpg")
	chAury.AddSprite("perfil_izq", aury_assets.Characters+"/aury/002.jpg")
	chAury.AddSprite("perfil_der", aury_assets.Characters+"/aury/002.jpg")
	chAury.SetSprite("perfil_izq")
	episode.AddCharacter(*chAury)

	chO := characters.NewCharacter("o", "Ofelia")
	chO.AddAttribute("edad", "int", 25)
	chO.AddAttribute("maldad", "int", 01)
	chO.AddAttribute("moral", "int", 100)
	episode.UpdateCharacter(*chO)

	chRene := characters.NewCharacter("rene", "René")
	chRene.AddAttribute("edad", "int", 30)
	chRene.AddAttribute("maldad", "uint", 20)
	chRene.AddAttribute("moral", "int", 70)
	episode.UpdateCharacter(*chRene)

	chPrince := characters.NewCharacter("prince", "Prince Charles")
	chPrince.AddAttribute("edad", "int", 33)
	chPrince.AddAttribute("maldad", "uint", 80)
	chPrince.AddAttribute("moral", "int", 0)
	episode.UpdateCharacter(*chPrince)

	return episode.Characters
}

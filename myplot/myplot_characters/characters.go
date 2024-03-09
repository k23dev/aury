package myplot_characters

import (
	"aury/aury/aury_assets"
	"aury/aury/characters"
	"aury/aury/episodes"
)

func CharactersEpisode1Setup(episode *episodes.Episode) []characters.Character {
	chAury, _ := episode.AddCharacter("aury", "Dominique Aury")
	chAury.AddAttribute("edad", "int", 50)
	chAury.AddAttribute("maldad", "uint", 100)
	chAury.AddSprite("frente", aury_assets.Characters+"/aury/001.jpg")
	chAury.AddSprite("perfil_izq", aury_assets.Characters+"/aury/002.jpg")
	chAury.AddSprite("perfil_der", aury_assets.Characters+"/aury/002.jpg")
	chAury.SetSprite("perfil_izq")
	episode.UpdateCharacter(*chAury)

	chO, _ := episode.AddCharacter("o", "Ofelia")
	chO.AddAttribute("edad", "int", 25)
	chO.AddAttribute("maldad", "uint", 01)
	episode.UpdateCharacter(*chO)

	chRene, _ := episode.AddCharacter("rene", "René")
	chRene.AddAttribute("edad", "int", 30)
	chRene.AddAttribute("maldad", "uint", 20)
	episode.UpdateCharacter(*chRene)

	chPrince, _ := episode.AddCharacter("prince", "Prince Charles")
	chPrince.AddAttribute("edad", "int", 33)
	chPrince.AddAttribute("maldad", "uint", 80)
	episode.UpdateCharacter(*chPrince)

	return episode.Characters
}

func CharactersEpisode2Setup(episode *episodes.Episode) []characters.Character {
	episode.Characters = CharactersEpisode1Setup(episode)
	// update characters

	chO, _ := episode.GetCharacter("o")
	valAttr := chO.ReadAttribute("maldad")
	val := valAttr.ToInt()
	chO.EditAttribute("maldad", val+10)
	episode.UpdateCharacter(*chO)

	chRene, _ := episode.GetCharacter("rene")
	valAttr = chRene.ReadAttribute("maldad")
	val = valAttr.ToInt()
	chRene.EditAttribute("maldad", val+10)
	episode.UpdateCharacter(*chRene)

	chPrince, _ := episode.GetCharacter("prince")
	valAttr = chPrince.ReadAttribute("maldad")
	val = valAttr.ToInt()
	chPrince.EditAttribute("maldad", val+25)
	episode.UpdateCharacter(*chPrince)

	return episode.Characters
}

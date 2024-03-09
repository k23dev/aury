package episode1

import (
	"aury/aury/episodes"
	"log"
)

func Episode() episodes.Episode {
	episode := episodes.NewEpisode()
	Setup(episode)
	// set the characters
	chAury, _ := episode.GetCharacter("aury")

	// auryDebug.Character(chAury)

	chO, _ := episode.GetCharacter("o")
	chRene, _ := episode.GetCharacter("rene")
	chPrince, _ := episode.GetCharacter("prince")

	currentScene, err := episode.AddScene("intro")
	if err != nil {
		panic(err)
	}

	currentLocation, _ := episode.GetLocation("castle")
	currentScene.SetLocation(currentLocation)

	currentScene.AddDialog(chAury.ID, "Hola, querida. ¿Cómo estás?")
	currentScene.AddDialog(chO.ID, "Bien, ama.")
	currentScene.AddDialog(chAury.ID, "Mejor así ...")
	currentScene.AddActionOfCharacter("abrazo", "abraza", chAury.ID)
	currentScene.AddActionFromCharacterTo("golpea", Golpear, chAury.ID, chO)
	chO.EditAttribute("maldad", 10)
	episode.UpdateCharacter(*chO)

	currentScene, err = episode.AddScene("el_plan")
	if err != nil {
		panic(err)
	}

	currentLocation, err = episode.GetLocation("cafecito")
	if err != nil {
		log.Println(err)
	}
	currentScene.SetLocation(currentLocation)

	currentLocation, _ = episode.GetLocation("cafe")
	currentScene.SetLocation(currentLocation)

	currentScene.AddDialog(chRene.ID, "Todo va según el plan")
	currentScene.AddDialog(chPrince.ID, "¡Excelentes noticias!")
	currentScene.AddActionOfCharacter("Escupirse la mano", EstrecharMano, chPrince.ID)
	currentScene.AddActionFromCharacterTo("estrechar mano", EstrecharMano, chPrince.ID, chRene.ID)

	return *episode
}

func Escupirse() string {
	return "Escupirse"
}

func EstrecharMano() string {
	return "Estrechar"
}

func Golpear() string {
	return "Golpear"
}

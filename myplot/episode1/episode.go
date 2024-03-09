package episode1

import (
	"aury/aury"
)

func Episode() *aury.Episode {

	// set the characters
	chAury, _ := episode.GetCharacter("aury")
	chO, _ := episode.GetCharacter("o")
	chRene, _ := episode.GetCharacter("rene")
	chPrince, _ := episode.GetCharacter("prince")

	currentScene, err := episode.AddScene("intro")
	if err != nil {
		panic(err)
	}

	currentLocation, _ := episode.GetLocation("castle")
	currentScene.SetLocation(currentLocation)

	currentScene.AddDialog(&chAury.ID, "Hola, querida. ¿Cómo estás?")
	currentScene.AddDialog(&chO.ID, "Bien, ama.")
	currentScene.AddDialog(&chAury.ID, "Mejor así ...")

	currentLocation, _ = episode.GetLocation("cafe")
	currentScene.SetLocation(currentLocation)

	currentScene.AddDialog(&chRene.ID, "Todo va según el plan")
	currentScene.AddDialog(&chPrince.ID, "¡Excelentes noticias!")

	return episode

}

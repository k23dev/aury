package episode1

import (
	"aury/aury/episodes"
)

func CharactersSetup(episode *episodes.Episode) {
	chAury, _ := episode.AddCharacter("aury", "Dominique Aury")
	chAury.AddAttribute("edad", "int", 50)
	chAury.AddAttribute("maldad", "uint", 100)
	episode.UpdateCharacter(*chAury)

	chO, _ := episode.AddCharacter("o", "Ofelia")

	chO.AddAttribute("edad", "int", 25)
	chO.AddAttribute("maldad", "uint", 01)
	episode.UpdateCharacter(*chO)

	episode.AddCharacter("rene", "René")
	episode.AddCharacter("prince", "Prince Charles")
}

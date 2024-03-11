package episode1

import (
	"aury/game/aury/episodes"
	"aury/game/mygame/mygame_characters"
	"aury/game/mygame/mygame_locations"
)

func Setup(episode *episodes.Episode) {
	episode.Characters = mygame_characters.GamesCharactersSetup(episode)
	episode.Locations = mygame_locations.LocationsSetup(episode)
}

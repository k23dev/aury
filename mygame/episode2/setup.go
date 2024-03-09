package episode2

import (
	"aury/aury/episodes"
	"aury/mygame/mygame_characters"
	"aury/mygame/mygame_locations"
)

func Setup(episode *episodes.Episode) {
	episode.Characters = mygame_characters.GamesCharactersSetup(episode)
	episode.Locations = mygame_locations.LocationsSetup(episode)
}

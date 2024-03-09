package episode2

import (
	"aury/aury/episodes"
	"aury/myplot/myplot_characters"
	"aury/myplot/myplot_locations"
)

func Setup(episode *episodes.Episode) {
	episode.Characters = myplot_characters.GamesCharactersSetup(episode)
	episode.Locations = myplot_locations.LocationsSetup(episode)
}

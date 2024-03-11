package mygame_locations

import (
	"aury/game/aury/aury_assets"
	"aury/game/aury/episodes"
	"aury/game/aury/locations"
)

func LocationsSetup(episode *episodes.Episode) []locations.Location {
	castle := locations.New("castle")
	castle.SetName("Castillo negro")
	castle.SetBackground(aury_assets.Locations + "/castillo.jpg")
	episode.AddLocation(*castle)

	cafe := locations.New("cafe")
	cafe.SetName("Caffé de Orleans")
	cafe.SetBackground(aury_assets.Locations + "/cafe_001.jpg")
	episode.AddLocation(*cafe)

	cafecito := locations.New("cafe")
	cafecito.SetName("Castillo vagabundo")
	cafecito.SetBackground(aury_assets.Locations + "cafecit_001_final_finalgo ru.jpg")
	cafecito.AddBgStyle("oscuro")
	cafecito.AddBgStyle("fondoAzul")
	episode.AddLocation(*cafe)

	return episode.Locations
}

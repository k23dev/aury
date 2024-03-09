package myplot_locations

import (
	"aury/aury/aury_assets"
	"aury/aury/episodes"
	"aury/aury/locations"
)

func LocationsSetup(episode *episodes.Episode) []locations.Location {
	castle, _ := episode.AddLocation("castle")
	castle.SetBackground(aury_assets.Locations + "/castillo.jpg")
	episode.UpdateLocation(*castle)

	cafe, _ := episode.AddLocation("cafe")
	cafe.SetBackground(aury_assets.Locations + "/cafe_001.jpg")
	episode.UpdateLocation(*cafe)

	cafecito, _ := episode.AddLocation("cafecito")
	cafecito.SetName("Castillo vagabundo")
	cafecito.SetBackground(aury_assets.Locations + "cafecit_001_final_finalgo ru.jpg")
	cafecito.AddBgStyle("oscuro")
	cafecito.AddBgStyle("fondoAzul")
	episode.UpdateLocation(*cafecito)

	return episode.Locations
}

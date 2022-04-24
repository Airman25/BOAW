package rooms

import "github.com/Airman25/BOAW/load"

//ToDo: maybe make the actual game?...still no?...when??? now? yey

//will be actual game at some point
func GetRoom6() []RoomObject {
	return []RoomObject{
		{0, float64(screenHeight) - 128, buttonwidth, buttonheight, buttonscaling, -2, load.ImageEbiten(`\spoilers\` + "game_gui")},
		{float64(screenWidth) - buttonwidth, float64(screenHeight) - buttonheight, buttonwidth, buttonheight, buttonscaling, 0, load.ImageText("button", load.Localisation["buttonBack"], buttonwidth, buttonheight, "", 0, 0)},
	}
}

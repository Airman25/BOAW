package rooms

import (
	"strconv"

	"github.com/Airman25/BOAW/load"
)

//settings
func GetRoom3(screenWidth, screenHeight int) []RoomObject {
	return []RoomObject{
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 - buttonheight*2 - 8, buttonwidth, buttonheight, buttonscaling, 101, load.ImageText("button", load.Localisation["buttonLanguage"], buttonwidth, buttonheight, "", 0, 0)},
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 - buttonheight - 6, buttonwidth, buttonheight, buttonscaling, 102, load.ImageText("button", load.Localisation["buttonFontsize"]+strconv.Itoa(int(load.DefaultFontSize)), buttonwidth, buttonheight, "", 0, 0)},
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 - 4, buttonwidth, buttonheight, buttonscaling, 103, load.ImageText("button", load.Localisation["buttonVolume"]+strconv.Itoa(load.MusicVolume), buttonwidth, buttonheight, "", 0, 0)},
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 + buttonheight - 2, buttonwidth, buttonheight, buttonscaling, 105, load.ImageText("button", load.Localisation["buttonSpeed"]+strconv.Itoa(load.GameSpeed), buttonwidth, buttonheight, "", 0, 0)},
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 + buttonheight*3, buttonwidth, buttonheight, buttonscaling, 0, load.ImageText("button", load.Localisation["buttonBack"], buttonwidth, buttonheight, "", 0, 0)},
	}
}

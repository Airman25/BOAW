package rooms

import (
	"strconv"

	"github.com/Airman25/BOAW/load"
)

//settings
func GetRoom2(screenWidth, screenHeight int) []RoomObject {
	return []RoomObject{
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 - buttonheight*2, buttonwidth, buttonheight, buttonscaling, 98, load.ImageText("button", load.Localisation["buttonLanguage"], buttonwidth, buttonheight, "", 0, 0)},
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 - buttonheight, buttonwidth, buttonheight, buttonscaling, 97, load.ImageText("button", load.Localisation["buttonFontsize"]+strconv.Itoa(int(load.DefaultFontSize)), buttonwidth, buttonheight, "", 0, 0)},
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 + buttonheight*3, buttonwidth, buttonheight, buttonscaling, 0, load.ImageText("button", load.Localisation["buttonBack"], buttonwidth, buttonheight, "", 0, 0)},
	}
}

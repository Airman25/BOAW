package rooms

import (
	"strconv"

	"github.com/Airman25/BOAW/load"
)

//settings
func GetRoom3(screenWidth, screenHeight int) []RoomObject {
	return []RoomObject{
		//languages
		{float64(screenWidth)/2 - buttonwidth, float64(screenHeight)/2 - buttonheight*3, buttonwidth, buttonheight, buttonscaling, 201, load.ImageText("button", "English", buttonwidth, buttonheight, "", 0, 0)},
		{float64(screenWidth) / 2, float64(screenHeight)/2 - buttonheight*3, buttonwidth, buttonheight, buttonscaling, 202, load.ImageText("button", "Русский", buttonwidth, buttonheight, "", 0, 0)},
		//buttons
		{float64(screenWidth)/2 - buttonwidth, float64(screenHeight)/2 - buttonheight, buttonwidth, buttonheight, buttonscaling, 102, load.ImageText("button", "-", buttonwidth, buttonheight, "", 0, 0)},
		{float64(screenWidth) / 2, float64(screenHeight)/2 - buttonheight, buttonwidth, buttonheight, buttonscaling, 103, load.ImageText("button", "+", buttonwidth, buttonheight, "", 0, 0)},
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight) / 2, buttonwidth, buttonheight, buttonscaling, 105, load.ImageText("button", load.Localisation["buttonSpeed"]+strconv.Itoa(load.GameSpeed), buttonwidth, buttonheight, "", 0, 0)},
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 + buttonheight, buttonwidth, buttonheight, buttonscaling, 110, load.ImageText("button", load.Localisation["buttonSize"]+load.GameSize, buttonwidth, buttonheight, "", 0, 0)},
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 + buttonheight*3, buttonwidth, buttonheight, buttonscaling, 0, load.ImageText("button", load.Localisation["buttonBack"], buttonwidth, buttonheight, "", 0, 0)},
		//text (?)
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 - buttonheight*4, buttonwidth, buttonheight, buttonscaling, -1, load.ImageText("button", load.Localisation["buttonLanguage"], buttonwidth, buttonheight, "", 0, 0)},
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 - buttonheight*2, buttonwidth, buttonheight, buttonscaling, -1, load.ImageText("button", load.Localisation["buttonVolume"]+strconv.Itoa(load.MusicVolume), buttonwidth, buttonheight, "", 0, 0)},
	}
}

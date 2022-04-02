package rooms

import "github.com/Airman25/BOAW/load"

//credits
func GetRoom5(screenWidth, screenHeight int) []RoomObject {
	return []RoomObject{
		{float64(screenWidth)/2 - buttonwidth, float64(screenHeight)/2 - buttonheight*3, buttonwidth * 2, buttonheight, buttonscaling, 106, load.ImageText("", "Programmed by: Airman25", buttonwidth*2, buttonheight, "", 0, 1)},
		{float64(screenWidth)/2 - buttonwidth, float64(screenHeight)/2 - buttonheight*2, buttonwidth * 2, buttonheight, buttonscaling, 107, load.ImageText("", "Plot by: Airman25, Skrimer740", buttonwidth*2, buttonheight, "", 0, 1)},
		{float64(screenWidth)/2 - buttonwidth, float64(screenHeight)/2 - buttonheight, buttonwidth * 2, buttonheight, buttonscaling, 108, load.ImageText("", "Textures by: Airman25, Daniil08m", buttonwidth*2, buttonheight, "", 0, 1)},
		{float64(screenWidth)/2 - buttonwidth, float64(screenHeight) / 2, buttonwidth, buttonheight * 2, buttonscaling, 109, load.ImageText("", "Music by: C-Life", buttonwidth*2, buttonheight, "", 0, 1)},
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 + buttonheight*3, buttonwidth, buttonheight, buttonscaling, 0, load.ImageText("button", load.Localisation["buttonBack"], buttonwidth, buttonheight, "", 0, 0)},
	}
}

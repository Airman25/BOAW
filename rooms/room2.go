package rooms

import (
	"github.com/Airman25/BOAW/load"
)

//file load menu
func GetRoom2() []RoomObject {
	return []RoomObject{
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 + buttonheight*3, buttonwidth, buttonheight, buttonscaling, 0, load.ImageText("button", load.Localisation["buttonBack"], buttonwidth, buttonheight, "", 0, 0)},
	}
}

package manager

import (
	"os"

	"github.com/Airman25/BOAW/rooms"
)

func RoomsManager(room, screenWidth, screenHeight int) []rooms.RoomObject { //gets what objects need to be loaded in a game "room"
	switch room {
	case 0:
		return rooms.GetRoom0(screenWidth, screenHeight) //start menu
	case 1:
		return rooms.GetRoom1(screenWidth, screenHeight) //will be actual game at some point
	case 2:
		return rooms.GetRoom2(screenWidth, screenHeight) //settings
	case 3:
		return rooms.GetRoom3(screenWidth, screenHeight) //will be file load menu
	case 4:
		return rooms.GetRoom4(screenWidth, screenHeight) //just credits
	default:
		return nil
	}
}

func ButtonFunction(function int) int { //in most cases change game "room" to the one requested by the button
	if function == 99 {
		os.Exit(0)
	}
	return function
}

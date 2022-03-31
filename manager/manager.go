package manager

import (
	"os"

	"github.com/Airman25/BOAW/rooms"
)

func RoomsManager(room, screenWidth, screenHeight int) []rooms.RoomObject { //gets what objects need to be loaded in a game "room"
	if room == 0 {
		return rooms.GetRoom0(screenWidth, screenHeight)
	} else if room == 1 {
		return rooms.GetRoom1(screenWidth, screenHeight)
	}
	return nil
}

func ButtonFunction(function int) int {
	if function == 99 {
		os.Exit(0)
	}
	return function
}

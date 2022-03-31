package manager

import (
	"github.com/Airman25/BOAW/rooms"
)

func RoomsManager(room, screenWidth, screenHeight int) []rooms.RoomObject { //gets what objects need to be loaded in a game "room"
	if room == 0 {
		return rooms.GetRoom1(screenWidth, screenHeight)
	}
	return nil
}

func ButtonFunction(function int) int {
	if function == 1 {
		rooms.Start()
		return 1
	}
	return 0
}

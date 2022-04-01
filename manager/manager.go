package manager

import (
	"io/ioutil"
	"os"

	"github.com/Airman25/BOAW/load"
	"github.com/Airman25/BOAW/rooms"
)

//redirect you to the room you need
func RoomsManager(room, screenWidth, screenHeight int) []rooms.RoomObject {
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

var currentLanguage int

//Reads file from src/lang
func LangManager() {
	files, err := ioutil.ReadDir("src/lang")
	if err != nil {
		panic(err)
	}
	if currentLanguage < len(files) {
		load.LoadLang(files[currentLanguage].Name())
		currentLanguage++
	} else {
		load.LoadLang(files[0].Name())
		currentLanguage = 1
	}
}

//in most cases change game "room" to the one requested by the button
func ButtonFunction(function int) int {
	switch function {
	case 99:
		os.Exit(0)
	case 98:
		LangManager()
		return 2
	case 97:
		if load.DefaultFontSize < 30 {
			load.DefaultFontSize++
		} else {
			load.DefaultFontSize = 10
		}
		return 2
	}
	return function
}

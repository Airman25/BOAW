package manager

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/Airman25/BOAW/load"
	"github.com/Airman25/BOAW/rooms"
)

//redirect you to the room you need
func RoomsManager(room, screenWidth, screenHeight int) []rooms.RoomObject {
	switch room {
	case 0:
		return rooms.GetRoom0(screenWidth, screenHeight) //start menu
	case 1:
		return rooms.GetRoom1(screenWidth, screenHeight) //pre-game
	case 2:
		return rooms.GetRoom2(screenWidth, screenHeight) //will be file load menu
	case 3:
		return rooms.GetRoom3(screenWidth, screenHeight) //settings
	case 4:
		return rooms.GetRoom4(screenWidth, screenHeight) //achievements or medals or something else
	case 5:
		return rooms.GetRoom5(screenWidth, screenHeight) //just credits
	case 6:
		return rooms.GetRoom5(screenWidth, screenHeight) //will be actual game at some point
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
	case 100:
		os.Exit(0)
	case 101:
		LangManager()
		return 3
	case 102:
		if load.DefaultFontSize < 30 {
			load.DefaultFontSize++
		} else {
			load.DefaultFontSize = 10
		}
		return 3
	case 103:
		if load.MusicVolume < 100 {
			load.MusicVolume += 10
		} else {
			load.MusicVolume = 0
		}
		return 3
	case 104:
		if load.Difficulty < 4 {
			load.Difficulty++
		} else {
			load.Difficulty = 1
		}
		return 1
	case 105:
		if load.GameSpeed < 3 {
			load.GameSpeed++
		} else {
			load.GameSpeed = 1
		}
		return 3
	case 106:

		return -1
	case 107:

		return -1
	case 108:

		return -1
	case 109:

		return -1
	}
	return function
}

//opens url in the web browser
func OpenUrl(url []string) {
	for _, val := range url {
		err := exec.Command("rundll32", "url.dll,FileProtocolHandler", val).Start()
		if err != nil {
			panic(err)
		}
	}
}

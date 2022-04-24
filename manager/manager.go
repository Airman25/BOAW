package manager

import (
	"os"
	"os/exec"

	"github.com/Airman25/BOAW/battle"
	"github.com/Airman25/BOAW/levels"
	"github.com/Airman25/BOAW/load"
	"github.com/Airman25/BOAW/rooms"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

var Animated = 0

//redirect you to the room you need
func RoomsManager(room int) []rooms.RoomObject {
	switch room {
	case 0:
		return rooms.GetRoom0() //start menu
	case 1:
		return rooms.GetRoom1() //pre-game
	case 2:
		return rooms.GetRoom2() //will be file load menu
	case 3:
		return rooms.GetRoom3() //settings
	case 4:
		return rooms.GetRoom4() //achievements or medals or something else
	case 5:
		return rooms.GetRoom5() //just credits
	case 6:
		anim()
		return rooms.GetRoom6() //will be actual game at some point
	case 7:
		return rooms.GetRoom7() //inside of a battle
	case 201:
		return rooms.Strategy()
	case 202:
		return rooms.Skills()
	case 203:
		return rooms.Summon()
	case 204:
		return rooms.Items()
	default:
		return nil
	}
}

//Reads file from src/lang
func LangManager(lang int) {
	if lang == 501 {
		load.LoadLang("en_us.lang")
	} else {
		load.LoadLang("ru_ru.lang")
	}

}

//in most cases change game "room" to the one requested by the button
func ButtonFunction(function int) int {
	if function > 500 {
		LangManager(function)
		return 3
	}
	switch function {
	case 100:
		os.Exit(0)
	case 101:
		return -1
	case 102:
		if load.MusicVolume > 0 {
			load.MusicVolume -= 10
		}
		PlayMusic("ping-pong")
		return 3
	case 103:
		if load.MusicVolume < 100 {
			load.MusicVolume += 10
		}
		PlayMusic("ping-pong")
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
	case 110:
		fh, fw := ebiten.ScreenSizeInFullscreen()
		h, _ := ebiten.WindowSize()
		switch h {
		case 1024:
			load.SizeChanger(1152, 648)
		case 1152:
			load.SizeChanger(1280, 720)
		case 1280:
			load.SizeChanger(1366, 768)
		case 1366:
			load.SizeChanger(1600, 900)
		case 1600:
			load.SizeChanger(1920, 1080)
		case 1920:
			load.SizeChanger(fh, fw)
			ebiten.SetFullscreen(true)
		case fh:
			ebiten.SetFullscreen(false)
			load.SizeChanger(1024, 648)
		}
		return 3
	case 200: //basic attack
		battle.Skill = 1
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

var audioPlayer *audio.Player

func PlayMusic(filename string) {
	if audioPlayer != nil {
		audioPlayer.Close()
	}
	if load.MusicVolume == 0 {
		return
	}
	audioPlayer = load.GetMusic(filename)
	audioPlayer.SetVolume(float64(load.MusicVolume) / 100)
	audioPlayer.Play()
}

//background for battle
func Background(index int) []*ebiten.Image {
	switch index {
	case 1:
		return []*ebiten.Image{load.ImageEbiten(`\spoilers\` + "BattleLocation1")}
	}
	return nil
}

//heroes, currently 1, planned 3
func Heroes() []battle.BattleObject {
	return battle.Hero1()
}

//should chapter animation be launched
func anim() {
	if load.GameLevel == 0 {
		if load.GameLocationX == 0 {
			Animated = 1
			load.GameLocationX = 1
			load.GameLocationY = 1
		}
	}
}

//manages locations for now
func LocationManager() ([]levels.AnimatedObject, []*ebiten.Image) {
	if load.GameLevel == 0 {
		if load.GameLocationX == 1 && load.GameLocationY == 1 {
			return levels.Location1(), []*ebiten.Image{load.ImageEbiten(`\spoilers\` + "Location1")}
		}
	}
	return nil, nil
}

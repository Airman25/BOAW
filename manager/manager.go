package manager

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strconv"

	"github.com/Airman25/BOAW/battle"
	"github.com/Airman25/BOAW/levels"
	"github.com/Airman25/BOAW/load"
	"github.com/Airman25/BOAW/rooms"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

var Animated = 0
var SkillAwait int

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
	case 205:
		return rooms.Target()
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
	} else if function > 400 && function < 404 {
		battle.Skill = SkillAwait
		battle.Target = function % 10
		return 7
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
		SkillAwait = 1
		return 205
	case 205: //block
		battle.Skill = 2
		return 7
	case 206: //change turn... but we only have 1 hero currently...
		return 7
	case 207: //run from battle
		battle.BattleParticipans = nil
		return 6
	case 208: //capture-delayed for...
		return 7
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
	if index < 10 {
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
	if load.GameLocationX%10 == 0 { //first room in each chapter
		if load.GameLevel == 0 { //pre-story not finished yet
			load.GameLevel = 1
		}
		if load.GameLevel == 1 {
			Animated = 1
			load.GameLocationX = 11
			load.GameLocationY = 11
			levels.Player.X = 200
			levels.Player.Y = 200
		}
	} else {
		Animated = -1
	}
}

//calls appropriative location based on the room player is in
func LocationManager() ([]levels.AnimatedObject, []*ebiten.Image) {
	L := &levels.Location{}
	fmt.Println("LocationX" + strconv.Itoa(load.GameLocationX) + "Y" + strconv.Itoa(load.GameLocationY))
	locationData, err := namedFunction(L, "LocationX"+strconv.Itoa(load.GameLocationX)+"Y"+strconv.Itoa(load.GameLocationY))
	if err != nil {
		return nil, nil
	}
	location, _ := (locationData[0]).Interface().([]levels.AnimatedObject)
	return location, []*ebiten.Image{load.ImageEbiten(`\spoilers\` + "Location1")}
}

func namedFunction(class interface{}, funcName string, params ...interface{}) (data []reflect.Value, err error) {
	//get value of named function in class
	method := reflect.ValueOf(class).MethodByName(funcName)
	//checks if method exist
	if !method.IsValid() {
		return nil, fmt.Errorf("Method not found \"%s\"", funcName)
	}
	/*
		//in case I would ever need to also send data to func
		sendData := make([]reflect.Value, len(params))
		for i, param := range params {
			sendData[i] = reflect.ValueOf(param)
		}
		out = m.Call(in)
	*/
	//gets the data from function
	data = method.Call(nil)
	return
}

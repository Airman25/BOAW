package game

import (
	"log"

	"github.com/Airman25/BOAW/load"
	"github.com/Airman25/BOAW/manager"
	"github.com/Airman25/BOAW/rooms"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type Game struct{}

var ObjectsArr []rooms.RoomObject
var Background []*ebiten.Image
var mouseReleased bool
var animated = 0

const screenWidth = 1280
const screenHeight = 720

func Launch(screenSizeWidth, screenSizeHeight int, musicenabled bool) {
	game := &Game{}
	ebiten.SetWindowSize(screenSizeWidth, screenSizeHeight)
	//To Do: add game icon
	//ebiten.SetWindowIcon(iconimage)
	manager.LangManager(201)
	if musicenabled {
		load.AudioContext = audio.NewContext(44100)
		manager.PlayMusic("ping-pong")
	}
	ObjectsArr = manager.RoomsManager(0, screenWidth, screenHeight)
	Background = rooms.DefaultBackground()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) { // Layout takes the outside size (e.g., the window size) and returns the (logical) screen size. If you don't have to adjust the screen size with the outside size, just return a fixed size.
	return 1280, 720
}

func (g *Game) Update() error { // Game's logical update. Update proceeds the game state. Update is called every tick (1/60 [s] by default).
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) { // Game's rendering.
	//renderBackground(screen)
	renderObjects(screen)
	buttonsClickDetect()
}

//ToDo draw the actual background
func renderBackground(screen *ebiten.Image) {
	animated++
	op := &ebiten.DrawImageOptions{}
	if animated < 12 {
		screen.DrawImage(Background[0], op)
	} else if animated < 24 {
		screen.DrawImage(Background[1], op)
	} else {
		screen.DrawImage(Background[0], op)
		animated = 0
	}
}

func renderObjects(screen *ebiten.Image) { //renders all of the objects in a room
	for i := 0; i < len(ObjectsArr); i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(ObjectsArr[i].Scale, ObjectsArr[i].Scale)
		op.GeoM.Translate(ObjectsArr[i].X, ObjectsArr[i].Y)
		screen.DrawImage(ObjectsArr[i].Image, op)
	}
}

//detects mouse button press event
func buttonsClickDetect() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) &&
		mouseReleased {
		mx, my := ebiten.CursorPosition()
		mouseReleased = false
		changeroom := manager.ButtonFunction(buttonPressed(mx, my))
		if changeroom != -1 {
			ObjectsArr = manager.RoomsManager(changeroom, screenWidth, screenHeight)
		}
	} else if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mouseReleased = true
	}
}

//check the coordinates to know whcih button was clicked
func buttonPressed(mx, my int) int {
	for i := 0; i < len(ObjectsArr); i++ {
		if ObjectsArr[i].Function == -1 { //all non button object will have -1 as function
			return -1
		}
		x := int(ObjectsArr[i].X)
		y := int(ObjectsArr[i].Y)
		if mx > x && mx < x+ObjectsArr[i].W && my > y && my < y+ObjectsArr[i].H {
			return ObjectsArr[i].Function
		}
	}
	return -1
}

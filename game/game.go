package game

import (
	"fmt"
	"log"

	"github.com/Airman25/BOAW/chapters"
	"github.com/Airman25/BOAW/levels"
	"github.com/Airman25/BOAW/load"
	"github.com/Airman25/BOAW/manager"
	"github.com/Airman25/BOAW/rooms"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct{}

var objectsArr []rooms.RoomObject           //containts buttons and other immovable objects
var objectsLevelArr []levels.AnimatedObject //containts only current level objects
var background []*ebiten.Image              //contains background images
var mouseReleased bool

const screenWidth = 1280
const screenHeight = 720

const playerWidth = 64
const playerHeight = 128

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
	objectsArr = manager.RoomsManager(0, screenWidth, screenHeight)
	background = rooms.DefaultBackground()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) { // Layout takes the outside size (e.g., the window size) and returns the (logical) screen size. If you don't have to adjust the screen size with the outside size, just return a fixed size.
	return 1280, 720
}

const Speed = 7

// Game's logical update. Update proceeds the game state. Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	if objectsLevelArr != nil { //if any level is loaded
		for _, key := range inpututil.PressedKeys() {
			if key == ebiten.KeyRight || key == ebiten.KeyD { //moving player right
				levels.Player.X += Speed
				levels.Player.ImageInUse = 0
				collisionCheck(0)
			} else if key == ebiten.KeyLeft || key == ebiten.KeyA { //moving player left
				levels.Player.X -= Speed
				levels.Player.ImageInUse = 1
				collisionCheck(1)
			} else if key == ebiten.KeyUp || key == ebiten.KeyW { //moving player up
				levels.Player.Y -= Speed
				levels.Player.ImageInUse = 2
				collisionCheck(2)
			} else if key == ebiten.KeyDown || key == ebiten.KeyS { //moving player down
				levels.Player.Y += Speed
				levels.Player.ImageInUse = 3
				collisionCheck(3)
			} else if key == ebiten.KeySpace { //interacting with enemies/chests (later)
				enemy := collisionCheck(4)
				fmt.Print(enemy)
			}
		}
		if levels.Player.X >= screenWidth {
			levels.Player.X = -playerWidth
			load.GameLocationX++
			objectsLevelArr = manager.LocationManager()
			background = manager.Background()
		} else if levels.Player.X < -playerWidth {
			levels.Player.X = screenWidth
			load.GameLocationX--
			objectsLevelArr = manager.LocationManager()
			background = manager.Background()
		} else if levels.Player.Y >= screenHeight-128 {
			levels.Player.Y = -playerHeight
			load.GameLocationY--
			objectsLevelArr = manager.LocationManager()
			background = manager.Background()
		} else if levels.Player.Y < -playerHeight {
			levels.Player.Y = screenHeight - 128
			load.GameLocationY++
			objectsLevelArr = manager.LocationManager()
			background = manager.Background()
		}

	}
	return nil
}

//Game's rendering.
func (g *Game) Draw(screen *ebiten.Image) {

	if manager.Animated != 0 {
		renderAnitmation(screen)
	} else {
		renderBackground(screen)
		renderGame(screen)
		renderObjects(screen)
		buttonsClickDetect()
	}

}

//ToDo draw normal background
func renderBackground(screen *ebiten.Image) {
	if background != nil {
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(background[0], op)
	}
}

func renderObjects(screen *ebiten.Image) { //renders all of the objects in a room
	for i := 0; i < len(objectsArr); i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(objectsArr[i].Scale, objectsArr[i].Scale)
		op.GeoM.Translate(objectsArr[i].X, objectsArr[i].Y)
		screen.DrawImage(objectsArr[i].Image, op)
	}
}

//detects mouse button press event
func buttonsClickDetect() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) &&
		mouseReleased {
		mx, my := ebiten.CursorPosition()
		mouseReleased = false
		changeroom := manager.ButtonFunction(buttonPressed(mx, my))
		if changeroom > -1 {
			objectsArr = manager.RoomsManager(changeroom, screenWidth, screenHeight)
			objectsLevelArr = nil
			background = nil
		}
	} else if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mouseReleased = true
	}
}

//check the coordinates to know whcih button was clicked
func buttonPressed(mx, my int) int {
	for i := 0; i < len(objectsArr); i++ {
		if objectsArr[i].Function == -1 { //non button object will have -1 as function
			return -1
		}
		x := int(objectsArr[i].X)
		y := int(objectsArr[i].Y)
		if mx > x && mx < x+objectsArr[i].W && my > y && my < y+objectsArr[i].H {
			return objectsArr[i].Function
		}
	}
	return -1
}

var animation int
var requipment []*ebiten.Image //images needed for certain animation

func renderAnitmation(screen *ebiten.Image) {
	animation += load.GameSpeed
	if manager.Animated == 1 {
		if requipment == nil {
			requipment = chapters.Requipment1()
		}
		chapters.ChapterIAnitmation(&animation, screen, requipment)
	}
	if animation == 0 {
		manager.Animated = 0
		requipment = nil
		objectsLevelArr = manager.LocationManager()
		background = manager.Background()
	}
}

//renders current level objects
func renderGame(screen *ebiten.Image) {
	if objectsLevelArr != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(levels.Player.Scale, levels.Player.Scale)
		op.GeoM.Translate(levels.Player.X, levels.Player.Y)
		screen.DrawImage(levels.Player.Images[levels.Player.ImageInUse], op)
		for i := 0; i < len(objectsLevelArr); i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(objectsLevelArr[i].Scale, objectsLevelArr[i].Scale)
			op.GeoM.Translate(objectsLevelArr[i].X, objectsLevelArr[i].Y)
			screen.DrawImage(objectsLevelArr[i].Images[objectsLevelArr[i].ImageInUse], op)
		}
	}
}

//checks if player colides with any of the enemys
func collisionCheck(mode int) int {
	var near float64
	if mode == 4 {
		near = 10
	}
	for i := 0; i < len(objectsLevelArr); i++ {
		if levels.Player.X+playerWidth+near > objectsLevelArr[i].X && //our x is more than the x coordinate where enemy is
			levels.Player.X+playerWidth-near < objectsLevelArr[i].X+objectsLevelArr[i].W && //enemy is not behind us
			levels.Player.Y+playerHeight+near > objectsLevelArr[i].Y && //We are not higher or lower than enemy
			levels.Player.Y-near < objectsLevelArr[i].Y+objectsLevelArr[i].H {
			if mode == 0 {
				levels.Player.X -= Speed
			} else if mode == 1 {
				levels.Player.X += Speed
			} else if mode == 2 {
				levels.Player.Y += Speed
			} else if mode == 3 {
				levels.Player.Y -= Speed
			} else {
				return objectsLevelArr[i].Data
			}
			break
		}
	}
	return 0
}

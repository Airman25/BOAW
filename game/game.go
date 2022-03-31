package game

import (
	"log"

	"github.com/Airman25/BOAW/manager"
	"github.com/Airman25/BOAW/rooms"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	gameRoom int
	player   *ebiten.Image
}

var ObjectsArr []rooms.RoomObject
var mouseReleased bool

var screenWidth, screenHeight int

func Launch(screenWidthHere, screenHeightHere int) {
	screenWidth = screenWidthHere
	screenHeight = screenHeightHere
	game := &Game{}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	//To Do: add game icon
	//ebiten.SetWindowIcon(iconimage)

	ObjectsArr = manager.RoomsManager(0, screenWidth, screenHeight)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) { // Layout takes the outside size (e.g., the window size) and returns the (logical) screen size. If you don't have to adjust the screen size with the outside size, just return a fixed size.
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error { // Game's logical update. Update proceeds the game state. Update is called every tick (1/60 [s] by default).
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) { // Game's rendering.
	renderObjects(screen)
	makeButtons()
}

func renderObjects(screen *ebiten.Image) { //renders all of the objects in a room
	for i := 0; i < len(ObjectsArr); i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(ObjectsArr[i].Scale, ObjectsArr[i].Scale)
		op.GeoM.Translate(ObjectsArr[i].X, ObjectsArr[i].Y)
		screen.DrawImage(ObjectsArr[i].Image, op)
	}
}

func makeButtons() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) &&
		mouseReleased {
		mx, my := ebiten.CursorPosition()
		mouseReleased = false
		changeroom := manager.ButtonFunction(pressButton(mx, my))
		if changeroom != 0 {
			ObjectsArr = manager.RoomsManager(changeroom, screenWidth, screenHeight)
		}
	} else if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mouseReleased = true
	}
}

func pressButton(mx, my int) int {
	for i := 0; i < len(ObjectsArr); i++ {
		if ObjectsArr[i].Function > 0 {
			x := int(ObjectsArr[i].X)
			y := int(ObjectsArr[i].Y)
			if mx > x && mx < x+ObjectsArr[i].W && my > y && my < y+ObjectsArr[i].H {
				return ObjectsArr[i].Function
			}
		} else {
			return 0
		}
	}
	return 0
}

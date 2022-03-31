package game

import (
	"log"

	"github.com/Airman25/BOAW/rooms"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	gameRoom int
	player   *ebiten.Image
}

var ObjectsArr []rooms.RoomObject

func Launch(screenWidth, screenHeight int) {
	game := &Game{}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	//To Do: add game icon
	//ebiten.SetWindowIcon(iconimage)
	ObjectsArr = rooms.Rooms(0, screenWidth, screenHeight)
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
}

func renderObjects(screen *ebiten.Image) { //renders all of the objects in a room
	for i := 0; i < len(ObjectsArr); i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(ObjectsArr[i].Scale, ObjectsArr[i].Scale)
		op.GeoM.Translate(ObjectsArr[i].X, ObjectsArr[i].Y)
		screen.DrawImage(ObjectsArr[i].Image, op)
	}
}

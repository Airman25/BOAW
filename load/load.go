package load

import (
	"bufio"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type moreData struct {
	fs    float64
	font  string
	color int
}

func ImageEbiten(filename string) *ebiten.Image { //Images for ebiten to render
	imageEbiten, _ := imageGet(filename)
	return imageEbiten
}

func imageNormal(filename string) image.Image { //Images that might be used for other purposes (for example to write text on them before rendering)
	_, imageNormal := imageGet(filename)
	return imageNormal
}

//gets the image from src folder
func imageGet(filename string) (*ebiten.Image, image.Image) {
	ebitenImg, img, err := ebitenutil.NewImageFromFile(`src\textures\` + filename + ".png")
	if err != nil {
		log.Fatalf("Load image error: %v", err)
		return nil, nil
	}
	return ebitenImg, img
}

func ImageText(filename, text string, width, height int, font string, fs float64, color int) *ebiten.Image {
	if filename != "" {
		return ebiten.NewImageFromImage(imageSetText(imageNormal(filename), text, width, height, font, fs, color))
	}
	return ebiten.NewImageFromImage(imageSetText(nil, text, width, height, font, fs, color))
}

var DefaultFontSize = 24.0

//simply adds text to image
func imageSetText(img image.Image, textonimage string, width int, height int, font string, fs float64, color int) image.Image {
	if font == "" { //default case for font and fontsize
		font = "arial"
	}
	if fs == 0 {
		fs = DefaultFontSize
	}
	dc := gg.NewContext(width, height)
	if err := dc.LoadFontFace(`src\fonts\`+font+".ttf", fs); err != nil {
		panic(err)
	}
	if img != nil {
		dc.DrawImage(img, 0, 0)
	}
	if color == 0 { // 0 = black, 1 = white, 2 = yellow, 3 = blue
		dc.SetRGB(0, 0, 0)
	} else if color == 1 {
		dc.SetRGB(1, 1, 1)
	} else if color == 2 {
		dc.SetRGB(246.0/256.0, 246.0/256.0, 78.0/256.0)
	} else if color == 3 {
		dc.SetRGB(45.0/256.0, 119.0/256.0, 219.0/256.0)
	}
	dc.DrawStringAnchored(textonimage, float64(width)/2, float64(height)/2, 0.5, 0.5)
	dc.Clip()
	return dc.Image()
}

var Localisation = make(map[string]string, 18)

//changes all objects names
func LoadLang(filename string) {
	fin, err := os.Open(`src\lang\` + filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fin)
	for scanner.Scan() {
		s := scanner.Text()
		if !strings.HasPrefix(s, "//") {
			Localisation[strings.Split(s, "=")[0]] = strings.Split(s, "=")[1]
		}
	}
	fin.Close()
}

var Difficulty int
var MusicVolume int
var GameSpeed = 1
var GameSize = "1280x720"

func SizeChanger(screenSizeWidth, screenSizeHeight int) {
	ebiten.SetWindowSize(screenSizeWidth, screenSizeHeight)
	GameSize = strconv.Itoa(screenSizeWidth) + "x" + strconv.Itoa(screenSizeHeight)
}

var AudioContext *audio.Context

func GetMusic(filename string) *audio.Player {
	f, err := ebitenutil.OpenFile(`src\audio\` + filename + ".mp3")
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}
	d, err := mp3.Decode(AudioContext, f)
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}
	audioPlay, err := audio.NewPlayer(AudioContext, d)
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}
	return audioPlay
}

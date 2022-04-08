package rooms

import "github.com/Airman25/BOAW/load"

//before game
func GetRoom1(screenWidth, screenHeight int) []RoomObject {
	difficulty := "difficultyEasy"
	switch load.Difficulty { //stored in number difficulty to text
	case 1:
		difficulty = "difficultyNormal"
	case 2:
		difficulty = "difficultyHard"
	case 3:
		difficulty = "difficultyNightmare"
	}
	return []RoomObject{
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 - buttonheight - 2, buttonwidth, buttonheight, buttonscaling, 6, load.ImageText("button", load.Localisation["buttonActualStart"], buttonwidth, buttonheight, "", 0, 0)},
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 - 4, buttonwidth, buttonheight, buttonscaling, 104, load.ImageText("button", load.Localisation["buttonDifficulty"], buttonwidth, buttonheight, "", 0, 0)},

		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 + buttonheight - 6, buttonwidth, buttonheight, buttonscaling, 104, load.ImageText("button", load.Localisation[difficulty], buttonwidth, buttonheight, "", 0, 0)},

		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 + buttonheight*3, buttonwidth, buttonheight, buttonscaling, 0, load.ImageText("button", load.Localisation["buttonBack"], buttonwidth, buttonheight, "", 0, 0)},
	}

}

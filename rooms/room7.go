package rooms

import (
	"github.com/Airman25/BOAW/load"
)

//inside of a battle
func GetRoom7() []RoomObject {
	return []RoomObject{
		{screenWidth/2 - buttonwidth, screenHeight / 2, buttonwidth, buttonheight, buttonscaling, 200, load.ImageText("button", load.Localisation["buttonAttack"], buttonwidth, buttonheight, "", 0, 0)},
		{screenWidth/2 - buttonwidth, screenHeight/2 + buttonheight, buttonwidth, buttonheight, buttonscaling, 201, load.ImageText("button", load.Localisation["buttonStrategy"], buttonwidth, buttonheight, "", 0, 0)},
		{screenWidth/2 - buttonwidth, screenHeight/2 + buttonheight*2, buttonwidth, buttonheight, buttonscaling, 202, load.ImageText("button", load.Localisation["buttonSkills"], buttonwidth, buttonheight, "", 0, 0)},
		{screenWidth/2 - buttonwidth, screenHeight/2 + buttonheight*3, buttonwidth, buttonheight, buttonscaling, 203, load.ImageText("button", load.Localisation["buttonSummon"], buttonwidth, buttonheight, "", 0, 0)},
		{screenWidth/2 - buttonwidth, screenHeight/2 + buttonheight*4, buttonwidth, buttonheight, buttonscaling, 204, load.ImageText("button", load.Localisation["buttonItems"], buttonwidth, buttonheight, "", 0, 0)},
	}
}

func Strategy() []RoomObject {
	return append([]RoomObject{
		{screenWidth/2 - buttonwidth, screenHeight/2 + buttonheight, buttonwidth, buttonheight, buttonscaling, 7, load.ImageText("button", load.Localisation["buttonStrategy"], buttonwidth, buttonheight, "", 0, 0)}, //overrides strategy button
		{screenWidth / 2, screenHeight / 2, buttonwidth, buttonheight, buttonscaling, 205, load.ImageText("button", load.Localisation["buttonBlock"], buttonwidth, buttonheight, "", 0, 0)},
		{screenWidth / 2, screenHeight/2 + buttonheight, buttonwidth, buttonheight, buttonscaling, 206, load.ImageText("button", load.Localisation["buttonTurn"], buttonwidth, buttonheight, "", 0, 0)},
		{screenWidth / 2, screenHeight/2 + buttonheight*2, buttonwidth, buttonheight, buttonscaling, 207, load.ImageText("button", load.Localisation["buttonRun"], buttonwidth, buttonheight, "", 0, 0)},
		{screenWidth / 2, screenHeight/2 + buttonheight*3, buttonwidth, buttonheight, buttonscaling, 208, load.ImageText("button", load.Localisation["buttonCapture"], buttonwidth, buttonheight, "", 0, 0)},
	}, GetRoom7()...)
}

func Skills() []RoomObject { //209-299, but might be changed in future
	return append([]RoomObject{
		{screenWidth/2 - buttonwidth, screenHeight/2 + buttonheight*2, buttonwidth, buttonheight, buttonscaling, 7, load.ImageText("button", load.Localisation["buttonSkills"], buttonwidth, buttonheight, "", 0, 0)},
		{screenWidth / 2, screenHeight / 2, buttonwidth, buttonheight, buttonscaling, 209, load.ImageText("button", load.Localisation["buttonSkillsTest"], buttonwidth, buttonheight, "", 0, 0)},
	}, GetRoom7()...)
}

func Summon() []RoomObject {
	return append([]RoomObject{ //300-399
		{screenWidth/2 - buttonwidth, screenHeight/2 + buttonheight*3, buttonwidth, buttonheight, buttonscaling, 7, load.ImageText("button", load.Localisation["buttonSummon"], buttonwidth, buttonheight, "", 0, 0)},
		{screenWidth / 2, screenHeight / 2, buttonwidth, buttonheight, buttonscaling, 300, load.ImageText("button", load.Localisation["buttonSummonTest"], buttonwidth, buttonheight, "", 0, 0)},
	}, GetRoom7()...)
}

func Items() []RoomObject { //405-499
	return append([]RoomObject{
		{screenWidth/2 - buttonwidth, screenHeight/2 + buttonheight*4, buttonwidth, buttonheight, buttonscaling, 7, load.ImageText("button", load.Localisation["buttonItems"], buttonwidth, buttonheight, "", 0, 0)},
		{screenWidth / 2, screenHeight / 2, buttonwidth, buttonheight, buttonscaling, 405, load.ImageText("button", load.Localisation["buttonItemsTest"], buttonwidth, buttonheight, "", 0, 0)},
	}, GetRoom7()...)
}

func Target() []RoomObject {
	return []RoomObject{
		{screenWidth/2 - buttonwidth, screenHeight / 2, buttonwidth, buttonheight, buttonscaling, 7, load.ImageText("button", load.Localisation["buttonCancel"], buttonwidth, buttonheight, "", 0, 0)},

		{screenWidth/2 + screenWidth/4, screenHeight / 2, buttonwidth, buttonheight, buttonscaling, 401, load.ImageEbiten(`\spoilers\` + "mark_1")},
		{screenWidth/2 + screenWidth/4, screenHeight/2 - screenHeight/4, buttonwidth, buttonheight, buttonscaling, 402, load.ImageEbiten(`\spoilers\` + "mark_1")},
		{screenWidth/2 + screenWidth/4, screenHeight/2 + screenHeight/4, buttonwidth, buttonheight, buttonscaling, 403, load.ImageEbiten(`\spoilers\` + "mark_1")},
	}
}

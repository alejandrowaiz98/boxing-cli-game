package entitys

type Boxer struct {
	Name  string
	Style string
	Stats
}

type Stats struct {
	Speed, SpeedLevel       int
	Strenght, StrengthLevel int
	HP, HPLevel             int
}

var mainBoxer *Boxer

func GetMainBoxer() *Boxer {

	return mainBoxer

}

var opponentBoxer *Boxer

func GetOpponentBoxer() *Boxer {

	return opponentBoxer

}

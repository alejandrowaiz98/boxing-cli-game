package entitys

type Boxer struct {
	Name   string
	Style  string
	Age    int
	Weight int
	Speed  int
}

var mainBoxer *Boxer

func GetMainBoxer() *Boxer {

	return mainBoxer

}

var opponentBoxer *Boxer

func GetOpponentBoxer() *Boxer {

	return opponentBoxer

}

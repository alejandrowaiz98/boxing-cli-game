package entitys

import (
	"math/rand"
	"time"

	"github.com/alejandrowaiz98/boxing-cli-game/constants"
)

type Boxer struct {
	Name    string
	Style   string
	Punches map[int]Punch
	Stats
}

type Stats struct {
	AvoidChance, AvoidingLevel int
	CritChance, CritLevel      int
	Speed, SpeedLevel          int
	Strenght, StrengthLevel    int
	HP, HPLevel                int
	Defense, DefenseLevel      int
}

func (b Boxer) AttackWith(option int) Punch {

	attack := applyOfensiveStats(b, b.Punches[option])

	finalAttack := applyCrit(b, attack)

	return finalAttack

}

func applyOfensiveStats(b Boxer, p Punch) Punch {

	p.DelayTime = p.DelayTime - constants.SpeedLevels[b.Stats.SpeedLevel]

	p.Power = p.Power + b.Stats.Strenght

	return p
}

func applyCrit(b Boxer, p Punch) Punch {

	rand.Seed(time.Now().UnixNano())

	randomNmb := rand.Intn(100) + 1

	if randomNmb <= b.CritChance {

		p.Power = p.Power * 2

		return p

	}

	return p

}

var mainBoxer *Boxer

func GetMainBoxer() *Boxer {

	return mainBoxer

}

var opponentBoxer *Boxer

func GetOpponentBoxer() *Boxer {

	return opponentBoxer

}

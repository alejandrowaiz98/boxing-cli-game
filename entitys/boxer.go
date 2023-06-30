package entitys

import (
	"math/rand"
	"time"

	"github.com/alejandrowaiz98/boxing-cli-game/constants"
)

type Boxer struct {
	Name        string
	Style       string
	BattleCry   string
	DeathRattle string
	Punches     map[int]Punch
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

var DefaultBoxer Boxer = Boxer{

	Stats: Stats{
		AvoidChance:   10,
		AvoidingLevel: 0,
		CritChance:    10,
		CritLevel:     0,
		Speed:         1,
		SpeedLevel:    0,
		Strenght:      1,
		StrengthLevel: 0,
		HP:            20,
		HPLevel:       1,
		Defense:       1,
		DefenseLevel:  0,
	},
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

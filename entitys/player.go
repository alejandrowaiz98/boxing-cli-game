package entitys

import (
	"github.com/alejandrowaiz98/boxing-cli-game/constants"
)

type Player struct {
	Name     string
	Victorys int
	Boxer    Boxer
}

// Unique instance of player
var player Player

func GetPlayer() Player {
	return player
}

func NewPlayer(playerName, boxerName, battlecry, deathrattle, style string) *Player {

	player.Boxer = DefaultBoxer

	player.Name = playerName

	player.Boxer.BattleCry = battlecry

	player.Boxer.DeathRattle = deathrattle

	return &player

}

// TODO: Add calling to applyStatChangeForLeveling and applyStatChangeForItem
func (p *Player) LevelUp(leveledStats []int) *Player {

	for _, choosenStat := range leveledStats {

		applyStatChangeForLeveling(choosenStat, p)

	}

	return p

}

// TODO: Create logic for each stat (if choosenStat == "option" { change player.stat})
func applyStatChangeForLeveling(choosenStat int, p *Player) *Player {

	if choosenStat == 1 {
		p.Boxer.AvoidingLevel++
		p.Boxer.AvoidChance = p.Boxer.AvoidChance + constants.AvoidingLevels[p.Boxer.AvoidingLevel]
	} else if choosenStat == 2 {
		p.Boxer.CritLevel++
		p.Boxer.CritChance = p.Boxer.CritChance + constants.CritLevels[p.Boxer.CritLevel]
	} else if choosenStat == 3 {
		p.Boxer.SpeedLevel++
		p.Boxer.Speed = p.Boxer.Speed + constants.SpeedLevels[p.Boxer.SpeedLevel]
	} else if choosenStat == 4 {
		p.Boxer.StrengthLevel++
		p.Boxer.Strenght = p.Boxer.Strenght + constants.StrenghtLevels[p.Boxer.StrengthLevel]
	} else if choosenStat == 5 {
		p.Boxer.HPLevel++
		p.Boxer.HP = p.Boxer.HP + constants.HPLevels[p.Boxer.HPLevel]
	} else if choosenStat == 6 {
		p.Boxer.Defense++
		p.Boxer.Defense = p.Boxer.Defense + constants.DefenseLevels[p.Boxer.DefenseLevel]
	}

	return p

}

// TODO: Create logic for each item (if choosenStat == "option" { change player.stat})
func applyStatChangeForItem(choosenStat string, p *Player) *Player {

	return p

}

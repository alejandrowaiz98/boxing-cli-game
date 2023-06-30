package entitys

import "log"

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

//TODO: Add calling to applyStatChangeForLeveling and applyStatChangeForItem
func (p *Player) LevelUp(leveledStats []string) *Player {

	for _, choosenStat := range leveledStats {

		log.Println(choosenStat)

	}

	return p

}

//TODO: Create logic for each item (if choosenStat == "option" { change player.stat})
func applyStatChangeForItem(choosenStat string, p *Player) *Player {

	return p

}

//TODO: Create logic for each stat (if choosenStat == "option" { change player.stat})
func applyStatChangeForLeveling(choosenStat string, p *Player) *Player {

	return p

}

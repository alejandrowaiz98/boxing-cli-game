package entitys

import (
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Punch struct {
	Name      string
	Hits      int
	Power     int
	DelayTime int
	Status
}

type Status struct {
}

var AllPunches map[string]Punch

func (p Punch) To(b *Boxer) (*Boxer, Punch) {

	attack := applyDefensiveStats(p, b)

	b.HP = b.HP - attack.Power

	return b, p

}

func applyDefensiveStats(p Punch, b *Boxer) Punch {

	p.Power = p.Power - b.Stats.Defense

	return p

}

func createPunches() error {

	f, err := excelize.OpenFile("Punches.xlsx")

	if err != nil {
		logger.Err(err).Msgf("getAllPunches | openFile: %v", err)
		return err
	}

	rows, err := f.GetRows("Punches")

	if err != nil {
		logger.Err(err).Msgf("getAllPunches | getRows : %v", err)
		return err
	}

	convertToInt := func(s string) int {

		i, err := strconv.Atoi(s)

		if err != nil {
			logger.Err(err).Msg("GetBoxerStats")
			panic(err)
		}

		return i
	}

	for i, columns := range rows {

		var p Punch

		if i == 0 {
			continue
		}

		for position, value := range columns {
			if position == 0 {
				p.Name = value
			} else if position == 1 {
				p.Hits = convertToInt(value)
			} else if position == 2 {
				p.Power = convertToInt(value)
			} else if position == 3 {
				p.DelayTime = convertToInt(value)
			}
		}

		AllPunches[p.Name] = p

	}

	return nil

}

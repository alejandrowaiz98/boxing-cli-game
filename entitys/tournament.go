package entitys

import (
	"strconv"

	excelize "github.com/xuri/excelize/v2"
)

type Tournament struct {
	Name   string
	Boxers []Boxer
}

var MinorTournament Tournament = Tournament{}

var MajorTournament Tournament = Tournament{}

var WorldTournament Tournament = Tournament{}

func getTournament(tournament string) Tournament {

	f, err := excelize.OpenFile(tournament + ".xlsx")

	if err != nil {

		logger.Err(err).Msgf("ERR opening file: %v", err)

		panic(err)

	}

	rows, err := f.GetRows("Info")

	if err != nil {

		logger.Err(err).Msgf("ERR getting boxer info: %v", err)

		panic(err)

	}

	t := getBoxerInfo(rows)

	rows, err = f.GetRows("Stats")

	if err != nil {

		logger.Err(err).Msgf("ERR getting boxer stats: %v", err)

		panic(err)

	}

	t = getBoxerStats(t, rows)

	return *t

}

func getBoxerInfo(rows [][]string) *Tournament {

	var t *Tournament

	for i, columns := range rows {

		var b Boxer

		if i == 0 {

			logger.Info().Msg(columns[0])

		}

		for position, value := range columns {

			if position == 0 {
				b.Info.Name = value
			} else if position == 1 {
				b.Info.Style = value
			} else if position == 2 {
				b.Info.BattleCry = value
			} else if position == 3 {
				b.Info.DeathRattle = value
			}

		}

		t.Boxers = append(t.Boxers, b)

	}

	return t

}

func getBoxerStats(t *Tournament, rows [][]string) *Tournament {

	convertToInt := func(s string) int {

		i, err := strconv.Atoi(s)

		if err != nil {
			logger.Err(err).Msg("GetBoxerStats")
			panic(err)
		}

		return i
	}

	for i, columns := range rows {

		var b Boxer

		if i == 0 {

			logger.Info().Msg(columns[0])

		}

		for position, value := range columns {

			if position == 0 {
				b.Stats.AvoidChance = convertToInt(value)
			} else if position == 1 {
				b.Stats.AvoidingLevel = convertToInt(value)
			} else if position == 2 {
				b.Stats.CritChance = convertToInt(value)
			} else if position == 3 {
				b.Stats.CritLevel = convertToInt(value)
			} else if position == 4 {
				b.Stats.Speed = convertToInt(value)
			} else if position == 5 {
				b.Stats.SpeedLevel = convertToInt(value)
			} else if position == 6 {
				b.Stats.Strenght = convertToInt(value)
			} else if position == 7 {
				b.Stats.StrengthLevel = convertToInt(value)
			} else if position == 8 {
				b.Stats.HP = convertToInt(value)
			} else if position == 9 {
				b.Stats.HPLevel = convertToInt(value)
			} else if position == 10 {
				b.Stats.Defense = convertToInt(value)
			} else if position == 11 {
				b.Stats.DefenseLevel = convertToInt(value)
			}

		}

		t.Boxers[i].Stats = b.Stats

	}

	return t

}

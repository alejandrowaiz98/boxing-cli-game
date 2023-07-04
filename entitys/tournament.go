package entitys

import (
	"strconv"

	excelize "github.com/xuri/excelize/v2"
)

type Tournament struct {
	Name   string
	Boxers []Boxer
}

var Tournaments []Tournament = []Tournament{
	{Name: "Minor Tournament"},
	{Name: "Major Tournament"},
	{Name: "World Tournament"},
}

func createTournaments() error {

	for i, t := range Tournaments {

		tournament, err := getTournament(t.Name)

		if err != nil {
			logger.Err(err).Msgf("createTournament | getTournament: %v", err)
			return err
		}

		Tournaments[i] = tournament

	}

	return nil

}

func getTournament(tournament string) (Tournament, error) {

	f, err := excelize.OpenFile(tournament + ".xlsx")

	if err != nil {

		logger.Err(err).Msgf("getTournament | openfile: %v", err)

		return Tournament{}, err

	}

	var t *Tournament

	rows, err := f.GetRows("Info")

	if err != nil {

		logger.Err(err).Msgf("getTournament | getRows | info : %v", err)

		return Tournament{}, err

	}

	getBoxerInfo(t, rows)

	rows, err = f.GetRows("Stats")

	if err != nil {

		logger.Err(err).Msgf("getTournament | getRows | stats: %v", err)

		panic(err)

	}

	getBoxerStats(t, rows)

	rows, err = f.GetRows("Punches")

	if err != nil {

		logger.Err(err).Msgf("getTournament | getRows | punches: %v", err)

		panic(err)

	}

	getBoxerPunches(t, rows)

	return *t, nil

}

func getBoxerInfo(t *Tournament, rows [][]string) {

	for i, columns := range rows {

		var b Boxer

		if i == 0 {

			logger.Info().Msg(columns[0])

		}

		for index, value := range columns {

			if index == 0 {
				b.Info.Name = value
			} else if index == 1 {
				b.Info.Style = value
			} else if index == 2 {
				b.Info.BattleCry = value
			} else if index == 3 {
				b.Info.DeathRattle = value
			}

		}

		t.Boxers = append(t.Boxers, b)

	}

}

func getBoxerStats(t *Tournament, rows [][]string) {

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

		for index, value := range columns {

			if index == 0 {
				continue
			} else if index == 1 {
				b.Stats.AvoidChance = convertToInt(value)
			} else if index == 2 {
				b.Stats.AvoidingLevel = convertToInt(value)
			} else if index == 3 {
				b.Stats.CritChance = convertToInt(value)
			} else if index == 4 {
				b.Stats.CritLevel = convertToInt(value)
			} else if index == 5 {
				b.Stats.Speed = convertToInt(value)
			} else if index == 6 {
				b.Stats.SpeedLevel = convertToInt(value)
			} else if index == 7 {
				b.Stats.Strenght = convertToInt(value)
			} else if index == 8 {
				b.Stats.StrengthLevel = convertToInt(value)
			} else if index == 9 {
				b.Stats.HP = convertToInt(value)
			} else if index == 10 {
				b.Stats.HPLevel = convertToInt(value)
			} else if index == 11 {
				b.Stats.Defense = convertToInt(value)
			} else if index == 12 {
				b.Stats.DefenseLevel = convertToInt(value)
			}

		}

		t.Boxers[i].Stats = b.Stats

	}

}

func getBoxerPunches(t *Tournament, rows [][]string) {

	for i, columns := range rows {

		if i == 0 {
			continue
		}

		for index, value := range columns {

			if index == 0 {
				continue
			}

			t.Boxers[i].Punches[index] = AllPunches[value]

		}

	}

}

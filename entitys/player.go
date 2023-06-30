package entitys

var DefaultPlayer Boxer = Boxer{

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

var p Boxer = DefaultPlayer

func NewPlayer(name, style string) Boxer {

	if style == "velocidad" {

		p.Stats.SpeedLevel = 2

	} else if style == "fuerza" {

		p.Stats.StrengthLevel = 2

	} else if style == "vida" {

		p.Stats.HPLevel = 2

	}

	return p

}

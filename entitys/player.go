package entitys

var DefaultBoxer Boxer = Boxer{

	Stats: Stats{SpeedLevel: 1, HPLevel: 1, StrengthLevel: 1},
}

func NewPlayer(name, style string) *Boxer {

	var p *Boxer = &DefaultBoxer

	if style == "velocidad" {

		p.Stats.SpeedLevel = 2

	} else if style == "fuerza" {

		p.Stats.StrengthLevel = 2

	} else if style == "vida" {

		p.Stats.HPLevel = 2

	}

	return p

}

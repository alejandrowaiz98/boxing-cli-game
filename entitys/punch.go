package entitys

type Punch struct {
	Name      string
	DelayTime int
	Power     int
}

func (p Punch) To(b *Boxer) (*Boxer, Punch) {

	attack := applyDefensiveStats(p, b)

	b.HP = b.HP - attack.Power

	return b, p

}

func applyDefensiveStats(p Punch, b *Boxer) Punch {

	p.Power = p.Power - b.Stats.Defense

	return p

}

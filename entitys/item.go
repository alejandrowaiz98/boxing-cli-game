package entitys

type Items struct {
	Name   string
	Rarity string
	StatChange
}

type StatChange struct {
	AvoidChange,
	CritChange,
	SpeedChange,
	StrenghtChange,
	HPChange,
	DefenseChange int
}

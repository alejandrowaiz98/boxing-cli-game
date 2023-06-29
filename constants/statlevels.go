package constants

var SpeedLevels map[int]int
var HPLevels map[int]int
var StrenghtLevels map[int]int
var DefenseLevels map[int]int
var AvoidingLevels map[int]int

//Whenever a boxer gains a level it will be automatically added the value
//from that level to his stats

func init() {

	SpeedLevels = make(map[int]int)
	SpeedLevels[1] = 1
	SpeedLevels[2] = 1
	SpeedLevels[3] = 1
	SpeedLevels[4] = 1
	SpeedLevels[5] = 1
	SpeedLevels[6] = 1
	SpeedLevels[7] = 1

	HPLevels = make(map[int]int)
	HPLevels[1] = 5
	HPLevels[2] = 5
	HPLevels[3] = 10
	HPLevels[4] = 10
	HPLevels[5] = 10
	HPLevels[6] = 15
	HPLevels[7] = 20

	StrenghtLevels = make(map[int]int)
	StrenghtLevels[1] = 1
	StrenghtLevels[2] = 2
	StrenghtLevels[3] = 3
	StrenghtLevels[4] = 4
	StrenghtLevels[5] = 5
	StrenghtLevels[6] = 6
	StrenghtLevels[7] = 7

	DefenseLevels = make(map[int]int)
	DefenseLevels[1] = 1
	DefenseLevels[2] = 2
	DefenseLevels[3] = 2
	DefenseLevels[4] = 3
	DefenseLevels[5] = 3
	DefenseLevels[6] = 4
	DefenseLevels[7] = 5

	AvoidingLevels = make(map[int]int)
	AvoidingLevels[1] = 5
	AvoidingLevels[2] = 5
	AvoidingLevels[3] = 5
	AvoidingLevels[4] = 5
	AvoidingLevels[5] = 5
}

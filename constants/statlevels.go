package constants

var SpeedLevels map[int]int
var HPLevels map[int]int
var StrenghtLevels map[int]int

func init() {

	SpeedLevels = make(map[int]int)
	SpeedLevels[1] = 0
	SpeedLevels[2] = 1
	SpeedLevels[3] = 1
	SpeedLevels[4] = 2
	SpeedLevels[5] = 2
	SpeedLevels[6] = 3
	SpeedLevels[7] = 4

	HPLevels = make(map[int]int)
	HPLevels[1] = 10
	HPLevels[2] = 15
	HPLevels[3] = 20
	HPLevels[4] = 25
	HPLevels[5] = 30
	HPLevels[6] = 35
	HPLevels[7] = 40

	StrenghtLevels = make(map[int]int)
	StrenghtLevels[1] = 1
	StrenghtLevels[2] = 2
	StrenghtLevels[3] = 2
	StrenghtLevels[4] = 3
	StrenghtLevels[5] = 3
	StrenghtLevels[6] = 4
	StrenghtLevels[7] = 5

}

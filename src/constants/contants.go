package constants

var PlayerFaceOffResult = struct {
	Win  int64
	Lose int64
	Tie  int64
}{
	Win:  1,
	Lose: -1,
	Tie:  0,
}

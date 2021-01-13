package game

// Opening for the obstacle is 3 lines tall.

// Obstacle is 5 chars wide.

const ObsW = 5

type Obs struct {
	GapStart int
	CurXMin  int
	CurXMax  int
}

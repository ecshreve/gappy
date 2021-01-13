package game

import "math/rand"

// Opening for the obstacle is 3 lines tall.

// Obstacle is 5 chars wide.

const ObsW = 5

type Obs struct {
	GapStart int
	CurXMin  int
	CurXMax  int
}

func (g *Game) GenNewObs() {
	for _, o := range g.Obstacles {
		if GameW-o.CurXMax < 5 {
			return
		}
	}

	gapS := rand.Intn(GameH)
	o := &Obs{
		GapStart: gapS,
		CurXMin:  GameW,
		CurXMax:  GameW + ObsW,
	}

	g.Obstacles = append(g.Obstacles, o)
}

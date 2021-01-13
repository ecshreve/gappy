package main

import (
	"github.com/ecshreve/gappy/internal/game"
)

func main() {
	o1 := &game.Obs{
		CurXMin:  40,
		CurXMax:  45,
		GapStart: 7,
	}

	g := &game.Game{
		Obstacles: []*game.Obs{o1},
		Bird:      &game.Birdy{PosY: 6},
	}

	g.RunTheGame()
}

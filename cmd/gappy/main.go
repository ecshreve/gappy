package main

import (
	"github.com/ecshreve/gappy/internal/game"
)

func main() {
	o1 := &game.Obs{
		CurXMin:  3,
		CurXMax:  8,
		GapStart: 7,
	}

	g := &game.Game{
		Obstacles: []*game.Obs{o1},
	}

	g.PrintCurrentGame()
}

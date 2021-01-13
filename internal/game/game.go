package game

import (
	"fmt"
	"time"
)

func (g *Game) RunTheGame() {
	for i := 0; i < 100; i++ {
		g.PrintCurrentGame()
		time.Sleep(time.Second)
		g.Tick()
	}
}

const (
	GameW = 60
	GameH = 12
)

func PrintYBorder() {
	fmt.Print("+")
	for i := 0; i < GameW; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

type Game struct {
	Obstacles []*Obs
	Bird      *Birdy
}

func (g *Game) PrintCurrentGame() {
	PrintYBorder()
	for i := 0; i < GameH; i++ {
		fmt.Print("|")
		for j := 0; j < GameW; j++ {
			isBird := false
			if i == g.Bird.PosY && j == 5 {
				fmt.Print("*")
				isBird = true
			}
			isObstacle := false
			for _, o := range g.Obstacles {
				if j > o.CurXMin && j < o.CurXMax {
					if i < o.GapStart || i > o.GapStart+2 {
						if isBird {
							fmt.Println("YOU LOSE")
							return
						}
						fmt.Print("#")
						isObstacle = true
						break
					}
				}
			}
			if !isObstacle && !isBird {
				fmt.Print(" ")
			}
		}
		fmt.Println("|")
	}

	PrintYBorder()
}

func (g *Game) Tick() {
	for _, o := range g.Obstacles {
		if o.CurXMin == 0 {
			o.CurXMax -= 1
		}

		o.CurXMin--
		o.CurXMax = o.CurXMin + ObsW
	}
	g.GenNewObs()
}

package game

import (
	"fmt"
	"time"

	"github.com/samsarahq/go/oops"
)

func (g *Game) RunTheGame() {
	for i := 0; i < 100; i++ {
		if g.Bird.PosY <= 0 || g.Bird.PosY >= GameH {
			panic(oops.Errorf("get rekt"))
		}
		if err := g.PrintCurrentGame(); err != nil {
			panic(err)
		}
		time.Sleep(time.Second / 4)
		if err := g.Tick(); err != nil {
			panic(err)
		}
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
	Score     int
}

func (g *Game) PrintCurrentGame() error {
	fmt.Printf("Score: %d\n", g.Score)
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
							return oops.Errorf("YOU LOSE")
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
	return nil
}

func (g *Game) Tick() error {
	for _, o := range g.Obstacles {
		if o.CurXMin == 0 {
			o.CurXMax -= 1
		}

		o.CurXMin--
		o.CurXMax = o.CurXMin + ObsW
	}
	g.GenNewObs()
	g.Bird.PosY++
	return nil
}

package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/samsarahq/go/oops"
)

const (
	GameW = 60
	GameH = 12
)

type Game struct {
	Obstacles []*Obs
	Bird      *Birdy
	Score     int
}

func NewGame() *Game {
	g := &Game{
		Bird: &Birdy{PosY: 6},
	}
	g.GenNewObs()
	return g
}

type Birdy struct {
	PosY int
}

func (g *Game) JumpBird() {
	g.Score++
	g.Bird.PosY -= 3
}

type Obs struct {
	GapStart int
	CurXMin  int
	CurXMax  int
}

func (g *Game) GenNewObs() {
	for _, o := range g.Obstacles {
		if GameW-o.CurXMax < 10 {
			return
		}
	}

	gapS := rand.Intn(GameH-6) + 3
	o := &Obs{
		GapStart: gapS,
		CurXMin:  GameW,
		CurXMax:  GameW + 5,
	}

	g.Obstacles = append(g.Obstacles, o)
}

func (g *Game) RunTheGame() {
	for i := 0; i < 100; i++ {
		if g.Bird.PosY <= 0 || g.Bird.PosY >= GameH {
			panic(oops.Errorf("get rekt"))
		}
		if err := g.PrintCurrentGame(); err != nil {
			panic(err)
		}
		time.Sleep(time.Second / 4)
		g.Tick()
	}
}

func PrintYBorder() {
	fmt.Print("+")
	for i := 0; i < GameW; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
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

func (g *Game) Tick() {
	for _, o := range g.Obstacles {
		o.CurXMin--
		o.CurXMax = o.CurXMin + 5
	}
	g.GenNewObs()
	g.Bird.PosY++
}

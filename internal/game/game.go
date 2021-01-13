package game

import "fmt"

func RunTheGame() {
	PrintInitialGame()
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

func PrintInitialGame() {
	PrintYBorder()

	for i := 0; i < GameH; i++ {
		fmt.Print("|")
		for j := 0; j < GameW; j++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
	}

	PrintYBorder()
}

type Game struct {
	Obstacles []*Obs
}

func (g *Game) PrintCurrentGame() {
	PrintYBorder()
	for i := 0; i < GameH; i++ {
		fmt.Print("|")
		for j := 0; j < GameW; j++ {
			isObstacle := false
			for _, o := range g.Obstacles {
				if j > o.CurXMin && j < o.CurXMax {
					if i < o.GapStart || i > o.GapStart+2 {
						fmt.Print("#")
						isObstacle = true
						break
					}
				}
			}
			if !isObstacle {
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

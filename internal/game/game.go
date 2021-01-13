package game

import "fmt"

func RunTheGame() {
	PrintInitialGame()
}

const (
	GameW = 60
	GameH = 12
)

func PrintInitialGame() {
	fmt.Print("+")
	for i := 0; i < GameW; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")

	for i := 0; i < GameH; i++ {
		fmt.Print("|")
		for j := 0; j < GameW; j++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
	}

	fmt.Print("+")
	for i := 0; i < GameW; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

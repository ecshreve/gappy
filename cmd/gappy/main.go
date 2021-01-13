package main

import (
	"fmt"

	"github.com/ecshreve/gappy/internal/game"
	term "github.com/nsf/termbox-go"
)

func reset() {
	term.Sync() // cosmestic purpose
}

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

	go g.RunTheGame()

	err := term.Init()
	if err != nil {
		fmt.Println("lol who cares")
	}

	defer term.Close()

	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				break
			case term.KeySpace:
				g.Score++
				g.Bird.Jump()
			default:
				// we only want to read a single character or one key pressed event
				reset()
				fmt.Println("ASCII : ", ev.Ch)

			}
		case term.EventError:
			fmt.Println("get rekt")
		}
	}

}

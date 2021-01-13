package main

import (
	"fmt"

	"github.com/ecshreve/gappy/internal/game"
	term "github.com/nsf/termbox-go"
	"github.com/samsarahq/go/oops"
)

func main() {
	g := game.NewGame()
	go g.RunTheGame()

	if err := term.Init(); err != nil {
		panic(oops.Wrapf(err, "lol who cares"))
	}
	defer term.Close()

	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeySpace:
				g.JumpBird()
			default:
				fmt.Println("ASCII : ", ev.Ch)
			}
		case term.EventError:
			fmt.Println("get rekt")
		}
	}

}

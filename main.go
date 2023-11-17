package main

import (
	"game-of-life/board"
	gameRules "game-of-life/rules"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const frameTime = 1

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Bounds: pixel.R(0, 0, board.Cols*board.Size, board.Rows*board.Size),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	rules := &gameRules.Rules{}
	b := board.New(rules)

	last := time.Now()

	for !win.Closed() {
		t := time.Since(last).Seconds()

		if t >= frameTime {
			b.Update()
			last = time.Now()
		}

		b.Draw(win)
	}
}

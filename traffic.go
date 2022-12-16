package main

import (
	"github.com/EngoEngine/engo"
)

type myScene struct{}

func (*myScene) Type() string { return "myGame" }

func (*myScene) Preload() {}

func (*myScene) Setup(engo.Updater) {}

func main() {
	opts := engo.RunOptions{
		Title:          "Hello World",
		Width:          1024,
		Height:         640,
		StandardInputs: true,
	}
	engo.Run(opts, &myScene{})
}

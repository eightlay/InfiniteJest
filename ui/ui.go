package ui

import "github.com/eightlay/InfiniteJest/iternal/graphicsdriver"

func Init() error {
	return graphicsdriver.Init()
}

func Terminate() {
	graphicsdriver.Terminate()
}

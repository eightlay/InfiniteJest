package graphicsdriver

import (
	"fmt"

	"github.com/eightlay/InfiniteJest/graphics"
	"github.com/go-gl/gl/v4.1-core/gl"
)

func InitGL() {
	if err := gl.Init(); err != nil {
		panic(fmt.Errorf("could not create opengl renderer: %v", err))
	}
}

// Set a default background color
func DefaultColor(c graphics.ColorRGBA) {
	gl.ClearColor(c.R, c.G, c.B, c.A)
}

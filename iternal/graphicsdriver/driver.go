package graphicsdriver

var initiated bool = false

func Init() {
	if !initiated {
		InitGL()
		InitGLFW()
		initiated = true
	}
}

package engine_ecs

import main_ecs "github.com/eightlay/InfiniteJest/iternal/ecs"

type EngineSystem interface {
	main_ecs.System
	EngineSystemInit(config *ConfigEngineECS)
}

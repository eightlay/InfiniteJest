package graphics

import "github.com/eightlay/InfiniteJest/iternal/utils"

const (
	// Number of channels in RGBA scheme
	CHANNELS_NUMBER = 4
	// Min channel value in RGBA scheme
	CHANNEL_MIN_VALUE = 0
	// Max channel value in RGBA scheme
	CHANNEL_MAX_VALUE = 255
	// Min channel value in reduced RGBA scheme
	CHANNEL_REDUCED_MIN_VALUE = 0.0
	// Max channel value in reduced RGBA scheme
	CHANNEL_REDUCED_MAX_VALUE = 1.0
)

const (
	// Red channel position in RGBA scheme
	R uint8 = iota
	// Green channel position in RGBA scheme
	G
	// Blue channel position in RGBA scheme
	B
	// Alpha channel position in RGBA scheme
	A
)

// ColorRGBA represents RGBA color
type ColorRGBA struct {
	// Red channel value in [0, 1]
	R float32
	// Green channel value in [0, 1]
	G float32
	// Blue channel value in [0, 1]
	B float32
	// Alpha channel value in [0, 1]
	A float32
}

// Create a color with the given channel values
//
// CreateColor() returns ColorRGBA{0, 0, 0, 1}
//
// CreateColor(255) returns ColorRGBA{1, 1, 1, 1}
//
// CreateColor(255, 127) returns ColorRGBA{1, 1, 1, ~0.5}
//
// CreateColor(255, 0, 0) returns ColorRGBA{1, 0, 0, 1}
//
// CreateColor(255, 0, 0, 127) returns ColorRGBA{1, 0, 0, ~0.5}
func CreateColor(vals ...int) ColorRGBA {
	if len(vals) == 0 {
		return ColorRGBA{
			CHANNEL_MIN_VALUE,
			CHANNEL_MIN_VALUE,
			CHANNEL_MIN_VALUE,
			CHANNEL_REDUCED_MAX_VALUE,
		}
	}

	ts := channelsTransform(vals)

	switch len(ts) {
	case 1:
		v := ts[0]
		return ColorRGBA{v, v, v, CHANNEL_REDUCED_MAX_VALUE}
	case 2:
		v, a := ts[0], ts[1]
		return ColorRGBA{v, v, v, a}
	case 3:
		return ColorRGBA{ts[R], ts[G], ts[B], CHANNEL_REDUCED_MAX_VALUE}
	}

	return ColorRGBA{ts[R], ts[G], ts[B], ts[A]}
}

// Transform channels' values from [0, 255] to [0, 1]
func channelsTransform(vals []int) []float32 {
	transformed := make([]float32, len(vals))

	for i, v := range vals {
		transformed[i] = channelTransform(v)
	}

	return transformed
}

// Transform channel's value from [0, 255] to [0, 1]
func channelTransform(v int) float32 {
	return utils.Constraint(float32(v), CHANNEL_MIN_VALUE, CHANNEL_MAX_VALUE) / CHANNEL_MAX_VALUE
}

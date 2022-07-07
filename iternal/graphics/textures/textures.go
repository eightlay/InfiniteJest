package textures

import (
	util "github.com/eightlay/InfiniteJest/iternal/utilfuncs"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Texture struct {
	id uint32
}

func (t *Texture) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, t.id)
}

func CreateTexture(file string) (*Texture, error) {
	image, err := util.ReadImage(file)
	if err != nil {
		return nil, err
	}
	size := image.Rect.Size()

	var texture uint32

	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	// Set the texture wrapping parameters
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_BORDER)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_BORDER)

	borderColor := []float32{0.0, 0.0, 0.0, 0.0}
	gl.TexParameterfv(gl.TEXTURE_2D, gl.TEXTURE_BORDER_COLOR, &borderColor[0])

	// Set texture filtering parameters
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.TexImage2D(
		gl.TEXTURE_2D, 0, gl.RGBA, int32(size.X), int32(size.Y),
		0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(image.Pix),
	)
	gl.GenerateMipmap(gl.TEXTURE_2D)

	return &Texture{texture}, nil
}

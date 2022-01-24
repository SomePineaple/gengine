package mesh

import (
	"fmt"
	"github.com/go-gl/gl/all-core/gl"
	"image"
	"image/draw"
	"os"
)

type Texture struct {
	textureID uint32
	width     int
	height    int
}

// NewTexture Creates a new texture from an image
func NewTexture(imagePath string) (*Texture, error) {
	tx := &Texture{}

	imgFile, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("texture %q not found on disk: %v", imagePath, err)
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image %q: %v", imagePath, err)
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return nil, fmt.Errorf("unsupported stride in image: %q", imagePath)
	}

	draw.Draw(rgba, rgba.Bounds(), img, image.Point{X: 0, Y: 0}, draw.Src)

	gl.GenTextures(1, &tx.textureID)
	gl.BindTexture(gl.TEXTURE_2D, tx.textureID)

	gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(rgba.Rect.Size().X), int32(rgba.Rect.Size().Y), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))

	gl.BindTexture(gl.TEXTURE_2D, 0)

	tx.width = rgba.Rect.Size().X
	tx.height = rgba.Rect.Size().Y

	return tx, nil
}

// Bind Tells opengl to use this texture as the TEXTURE_2D
func (tx *Texture) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, tx.textureID)
}

// Unbind Tells opengl to stop using this texture as the TEXTURE_2D
func (tx *Texture) Unbind() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

// Destroy Deletes the texture
func (tx *Texture) Destroy() {
	gl.DeleteTextures(1, &tx.textureID)
}

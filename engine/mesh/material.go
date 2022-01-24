package mesh

import "github.com/go-gl/mathgl/mgl64"

type Material struct {
	ambientColor  mgl64.Vec4
	diffuseColor  mgl64.Vec4
	specularColor mgl64.Vec4
	reflectance   float64
	texture       *Texture
}

func NewMaterial() *Material {
	defaultColor := mgl64.Vec4{1.0, 1.0, 1.0, 1.0}

	material := &Material{}

	material.ambientColor = defaultColor
	material.diffuseColor = defaultColor
	material.specularColor = defaultColor

	material.texture = nil

	material.reflectance = 0

	return material
}

// IsTextured Returns weather the material has a texture or uses colors
func (mtl *Material) IsTextured() bool {
	return mtl.texture != nil
}

// SetTexture Sets the materials texture
func (mtl *Material) SetTexture(texture *Texture) {
	mtl.texture = texture
}

// SetAmbientColor Sets the materials ambient color
func (mtl *Material) SetAmbientColor(ambientColor mgl64.Vec4) {
	mtl.ambientColor = ambientColor
}

// SetDiffuseColor Sets the materials diffuse texture
func (mtl *Material) SetDiffuseColor(diffuseColor mgl64.Vec4) {
	mtl.diffuseColor = diffuseColor
}

// SetSpecularColor Sets the color of specular highlights with this material
func (mtl *Material) SetSpecularColor(specularColor mgl64.Vec4) {
	mtl.specularColor = specularColor
}

// SetReflectance Sets the reflectance of the material
func (mtl *Material) SetReflectance(reflectance float64) {
	mtl.reflectance = reflectance
}

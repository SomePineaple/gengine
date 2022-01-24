package items

import "github.com/go-gl/gl/all-core/gl"

type Mesh struct {
	vaoID       uint32
	name        string
	vboIDList   []uint32
	vertexCount int32
	mat         *Material

	indices []int32
}

// NewMesh Creates a new items with the given vertices, texture coordinates, normals, and indices
func NewMesh(vertices []float32, textCoords []float32, normals []float32, indices []int32, name string) *Mesh {
	mesh := &Mesh{}

	mesh.name = name
	mesh.mat = NewMaterial()
	mesh.indices = indices

	mesh.vertexCount = int32(len(indices))

	gl.GenVertexArrays(1, &mesh.vaoID)
	gl.BindVertexArray(mesh.vaoID)

	// Vertices VBO
	var vboID uint32
	gl.GenBuffers(1, &vboID)
	mesh.vboIDList = append(mesh.vboIDList, vboID)
	gl.BindBuffer(gl.ARRAY_BUFFER, vboID)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, nil)

	// Texture coordinates VBO
	gl.GenBuffers(1, &vboID)
	mesh.vboIDList = append(mesh.vboIDList, vboID)
	gl.BindBuffer(gl.ARRAY_BUFFER, vboID)
	gl.BufferData(gl.ARRAY_BUFFER, len(textCoords)*4, gl.Ptr(textCoords), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 0, nil)

	// Vertex normals VBO
	gl.GenBuffers(1, &vboID)
	mesh.vboIDList = append(mesh.vboIDList, vboID)
	gl.BindBuffer(gl.ARRAY_BUFFER, vboID)
	gl.BufferData(gl.ARRAY_BUFFER, len(normals)*4, gl.Ptr(normals), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(2)
	gl.VertexAttribPointer(2, 3, gl.FLOAT, false, 0, nil)

	// Index VBO
	gl.GenBuffers(1, &vboID)
	mesh.vboIDList = append(mesh.vboIDList, vboID)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, vboID)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	return mesh
}

// Render Draws the items to the screen
func (msh *Mesh) Render() {
	if tex := msh.mat.texture; tex != nil {
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, tex.textureID)
	}

	gl.BindVertexArray(msh.vaoID)
	gl.EnableVertexAttribArray(0)
	gl.EnableVertexAttribArray(1)
	gl.EnableVertexAttribArray(2)

	gl.DrawElements(gl.TRIANGLES, msh.vertexCount, gl.UNSIGNED_INT, gl.Ptr(msh.indices))

	gl.DisableVertexAttribArray(2)
	gl.DisableVertexAttribArray(1)
	gl.DisableVertexAttribArray(0)

	gl.BindVertexArray(0)
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

// Destroy Deletes all the connected vertex buffers, and the vertex array.
func (msh *Mesh) Destroy() {
	gl.DisableVertexAttribArray(0)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	for i := 0; i < len(msh.vboIDList); i++ {
		gl.DeleteBuffers(1, &msh.vboIDList[i])
	}

	if msh.mat.texture != nil {
		msh.mat.texture.Destroy()
	}

	gl.BindVertexArray(0)
	gl.DeleteVertexArrays(1, &msh.vaoID)
}

// SetMaterial Set a custom material for this items
func (msh *Mesh) SetMaterial(mat *Material) {
	msh.mat = mat
}

// GetName Returns the name of the items
func (msh *Mesh) GetName() string {
	return msh.name
}

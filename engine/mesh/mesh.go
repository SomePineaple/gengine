package mesh

import "github.com/go-gl/gl/all-core/gl"

type Mesh struct {
	vaoID       uint32
	name        string
	vboIDList   []uint32
	vertexCount int
}

func NewMesh(vertices []float32, textCoords []float32, normals []float32, indices []int32, name string) *Mesh {
	mesh := &Mesh{}

	mesh.vertexCount = len(indices)

	gl.GenVertexArrays(1, &mesh.vaoID)
	gl.BindVertexArray(mesh.vaoID)

	// Position VBO
	var vboID uint32
	gl.GenBuffers(1, &vboID)
	mesh.vboIDList = append(mesh.vboIDList, vboID)
	gl.BindBuffer(gl.ARRAY_BUFFER, vboID)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	return mesh
}

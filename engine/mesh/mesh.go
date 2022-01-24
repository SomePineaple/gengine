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

	// Vertices VBO
	var vboID uint32
	gl.GenBuffers(1, &vboID)
	mesh.vboIDList = append(mesh.vboIDList, vboID)
	gl.BindBuffer(gl.ARRAY_BUFFER, vboID)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 0, nil)

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

	return mesh
}

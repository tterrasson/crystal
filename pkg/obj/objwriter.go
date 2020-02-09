package obj

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ObjWriter write an obj file
type ObjWriter struct {
	lines       strings.Builder
	vertices    []string
	verticesIdx map[string]int
	faces       map[int][][4]int
}

// NewObjWriter create a new writer
func NewObjWriter() *ObjWriter {
	writer := new(ObjWriter)
	writer.faces = make(map[int][][4]int)
	writer.verticesIdx = make(map[string]int)
	return writer
}

// AddMtlLib to obj file
func (writer *ObjWriter) AddMtlLib(filename string) {
	fmt.Fprintf(&writer.lines, "mtllib %s\n", filename)
}

// AddComment add a comment
func (writer *ObjWriter) AddComment(comment string) {
	fmt.Fprintf(&writer.lines, "# %s\n", comment)
}

// AddVertex add a new vertex
func (writer *ObjWriter) AddVertex(x int, y int, z int) int {
	line := fmt.Sprintf("v %d %d %d", x, y, z)

	idx, ok := writer.verticesIdx[line]
	if ok {
		return idx
	}

	writer.vertices = append(writer.vertices, line)
	writer.verticesIdx[line] = len(writer.vertices)

	return len(writer.vertices)
}

// AddCube to export
func (writer *ObjWriter) AddCube(x int, y int, z int, state int) {
	var idx [4]int

	idx[0] = writer.AddVertex(x, y+1, z+1)
	idx[1] = writer.AddVertex(x, y, z+1)
	idx[2] = writer.AddVertex(x+1, y, z+1)
	idx[3] = writer.AddVertex(x+1, y+1, z+1)
	writer.faces[state] = append(writer.faces[state], idx)

	idx[0] = writer.AddVertex(x+1, y+1, z)
	idx[1] = writer.AddVertex(x+1, y, z)
	idx[2] = writer.AddVertex(x, y, z)
	idx[3] = writer.AddVertex(x, y+1, z)
	writer.faces[state] = append(writer.faces[state], idx)

	idx[0] = writer.AddVertex(x+1, y+1, z+1)
	idx[1] = writer.AddVertex(x+1, y, z+1)
	idx[2] = writer.AddVertex(x+1, y, z)
	idx[3] = writer.AddVertex(x+1, y+1, z)
	writer.faces[state] = append(writer.faces[state], idx)

	idx[0] = writer.AddVertex(x, y+1, z)
	idx[1] = writer.AddVertex(x, y+1, z+1)
	idx[2] = writer.AddVertex(x+1, y+1, z+1)
	idx[3] = writer.AddVertex(x+1, y+1, z)
	writer.faces[state] = append(writer.faces[state], idx)

	idx[0] = writer.AddVertex(x, y+1, z)
	idx[1] = writer.AddVertex(x, y, z)
	idx[2] = writer.AddVertex(x, y, z+1)
	idx[3] = writer.AddVertex(x, y+1, z+1)
	writer.faces[state] = append(writer.faces[state], idx)

	idx[0] = writer.AddVertex(x, y, z+1)
	idx[1] = writer.AddVertex(x, y, z)
	idx[2] = writer.AddVertex(x+1, y, z)
	idx[3] = writer.AddVertex(x+1, y, z+1)
	writer.faces[state] = append(writer.faces[state], idx)
}

// WriteToFile of your choice
func (writer *ObjWriter) WriteToFile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	fmt.Fprintf(w, "%s", []byte(writer.lines.String()))

	// Write all vertices
	for _, vertex := range writer.vertices {
		fmt.Fprintf(w, "%s\n", vertex)
	}

	// Write all faces
	for state, faces := range writer.faces {
		fmt.Fprintf(w, "usemtl state%d\n", state)

		for _, face := range faces {
			fmt.Fprintf(w, "f %d %d %d %d\n", face[0], face[1], face[2], face[3])
		}
	}

	w.Flush()
}

package algorithms

const MAX_VERTICES int = 20

type Vertex struct {
	Label string
	WasVisited bool
}

func NewVertex() {
	vertex := &Vertex{Label: "0", WasVisited: false}
	return vertex
}

type Graph struct {
	MaxVertices int
	VertexList []Vertex
	AdjacencyList []int
	NVerts int
}


func (g *Graph) addVertex(label string, visited bool) {
	vertex =
	VertexList = append()
}

func (g *Graph) updateAdjacencyMatrix() {

}

func NewGraph() Graph {
	vertexList := make([]Vertex, 20)
	vertex := &Vertex{Label: "0", WasVisited: false}

	vertexList[0]=vertex
	adjacencyList := make([][], 0)
	graph := &Graph{VertexList:vertexList, AdjacencyList:adjacencyList}

	return graph
}



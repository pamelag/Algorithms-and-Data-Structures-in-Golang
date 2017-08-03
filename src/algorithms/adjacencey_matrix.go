package algorithms

import "fmt"

// A structure to represent an adjacency list node
type AdjListNode struct {
	Dest int
	Next *AdjListNode
}

type AdjList struct {
	Head *AdjListNode
}

// A structure to represent a graph. A graph is an array of adjacency lists.
// Size of array will be V (number of vertices in graph)
type Graph struct {
	V int
	ArrAdjList *[]AdjList
}

// A utility function to create a new adjacency list node
func NewAdjacencyListNode(dest int) *AdjListNode {
	node := &AdjListNode{}
	node.Dest = dest
	node.Next = nil
	return node
}
// A utility function to create a new adjacency list
func NewAdjList(dest int) *AdjList {
	node := NewAdjacencyListNode(dest)
	list := &AdjList{Head: node}
	return list
}

// A utility function that creates a graph of V vertices
func CreateGraph(v int) *Graph {
	graph := &Graph{V : v}

	// Create an array of adjacency lists.  Size of array will be V
	adjList := NewAdjList(0)

	graph.ArrAdjList = append(graph.ArrAdjList, adjList)
	return graph
}

// Adds an edge to an undirected graph
func addEdge(graph *Graph, src int, dest int) {
	// Add an edge from src to dest.  A new node is added to the adjacency
	// list of src.  The node is added at the beginning
	newNode := NewAdjacencyListNode(dest)
	newNode.Next = graph.ArrAdjList[src].Head
	graph.ArrAdjList[src].Head = newNode

	// Since graph is undirected, add an edge from dest to src also
	newNode = NewAdjacencyListNode(dest)
	newNode.Next = graph.ArrAdjList[dest].Head
	graph.ArrAdjList[dest].Head = newNode;
}

// A utility function to print the adjacency list representation of graph
func PrintGraph(graph *Graph) {
	for v := 0; v <graph.V; v++ {
		pCrawl := graph.ArrAdjList[v].Head
		fmt.Println("Adjacency list of vertex")
		for pCrawl != nil {
			fmt.Println("pCrawl dest")
			pCrawl = pCrawl.Next
		}
	}
}

// Driver program to test above functions
func BuildGraph() {
	// create the graph given in above figure
	V := 5;

	graph := CreateGraph(V);
	addEdge(graph, 0, 1);
	addEdge(graph, 0, 4);
	addEdge(graph, 1, 2);
	addEdge(graph, 1, 3);
	addEdge(graph, 1, 4);
	addEdge(graph, 2, 3);
	addEdge(graph, 3, 4);

	// print the adjacency list representation of the above graph
	PrintGraph(graph);
}

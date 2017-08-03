package algorithms

import "fmt"

type Node struct {
	Destination int
	Value int
}

type NodeList struct {
	Source int
	Nodes []Node
}

func createGraph() {
	Graph := make(map[int]NodeList)
	// createVertices
	Graph[0] = &NodeList{}
	Graph[1] = &NodeList{}
	Graph[2] = &NodeList{}
	Graph[3] = &NodeList{}
	Graph[4] = &NodeList{}
	Graph[5] = &NodeList{}
	Graph[6] = &NodeList{}

	// createChildren
	children := &NodeList{Source: 0, Nodes: {{0, 0, 0},{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}}
	Graph[0] = children

	children = &NodeList{Source: 0, Nodes: {{0, 0, 0},{0, 0, 0}}}
	Graph[1] = children

	//createEdges(Graph, 0, 1)
	//createEdges(Graph, 0, 2)
	//createEdges(Graph, 0, 3)
	//createEdges(Graph, 0, 4)
}

func createEdges(graph map[int]NodeList, src int, dest int) {
	nodeList := graph[src]
	nodeList.Source = src
	for i := 0; i < len(nodeList.Nodes); i++ {
		if i == dest {
			node := nodeList.Nodes[i]
			node.Destination = dest
			node.Value = dest
		}
	}
}

func bfs(graph map[int]NodeList, item int) {
	if val, ok := graph[item]; ok {
		source := val.Source
		nodeList := graph[source]
		for key, val := range nodeList.Nodes {
			fmt.Println(key + " : " + val)
		}
	}
}
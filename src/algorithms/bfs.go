package algorithms

import "fmt"

type Node struct {
	Source *NodeList
	Dest int
	Visited bool
}
func NewNode() *Node {
	return &Node{}
}

type NodeList struct {
	List []Node
	Src int
}

func NewNodeList(src int) *Node {
	node := NewNode()
	list := make([]Node, 0)
	list = append(list, node)
	return &NodeList{Src: src}
}

type Graph struct {
	NumberOfNodes int
	ArrNodeList []NodeList
}

func NewGraph(rootSrc int) *Graph {
	arrNodeList := make([]NodeList, 0)
	addSource(rootSrc, arrNodeList)
	return &Graph{ArrNodeList : arrNodeList}
}

func addSource(src int, arrNodeList []NodeList) {
	if arrNodeList != nil {
		nodeList := NewNodeList(src)
		arrNodeList = append(arrNodeList, nodeList)
	}
}

func AddEdge(graph *Graph, src int, dest int) {
	var found bool = false
	for i:=0; i < len(graph.ArrNodeList); i++ {
		list := graph.ArrNodeList[i]
		if list.Src == src {
			found == true
			break
		}
	}
	if found == false {
		addSource(src, graph.ArrNodeList)
	}
	for i:=0; i < len(graph.ArrNodeList); i++ {
		list := graph.ArrNodeList[i]
		node := NewNode()
		node.Dest = dest
		list = append(list, node)
		node.Source = list

		/*node1 := NewNode()
		node1.Dest = src
		list = append(list, node1)*/
	}
}

// A utility function to print the adjacency list representation of graph
func printGraph(graph *Graph) {
	for v := 0; v <len(graph.ArrNodeList); v++ {
		pCrawl := graph.ArrNodeList[v]
		fmt.Println("Adjacency list of vertex")
		if pCrawl != nil {
			for _, value := range pCrawl.List {
				fmt.Println("pCrawl Source ", value.Source)
				fmt.Println("pCrawl Dest ", value.Dest)
			}
		}
	}
}

func BFS(graph *Graph, dest int) {
	src := 0
	for i := 0; i < len(graph.ArrNodeList); i++ {
		if dest == graph.ArrNodeList[i].Src {
			nodeList := graph.ArrNodeList[i]
			src = nodeList.Src
			fmt.Print(src)
		}
	}
	for x := 0; x < len(graph.ArrNodeList); x++ {
		if src == graph.ArrNodeList[x].Src {
			list := graph.ArrNodeList[x].List
			for _, val := range list {
				fmt.Print(val)
			}
		}
	}
}

// Driver program to test above functions
func Build() {

	graph := NewGraph(0);
	graph.NumberOfNodes = 5

	AddEdge(graph, 0, 1);
	AddEdge(graph, 0, 4);
	AddEdge(graph, 1, 2);
	AddEdge(graph, 1, 3);
	AddEdge(graph, 1, 4);
	AddEdge(graph, 2, 3);
	AddEdge(graph, 3, 4);

	// print the adjacency list representation of the above graph
	printGraph(graph);
}






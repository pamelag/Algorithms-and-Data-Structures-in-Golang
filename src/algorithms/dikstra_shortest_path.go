package algorithms

import "fmt"

const (
	uvinf    = 0x7FF0000000000000
)

type Node struct {
	Key string
	Value int
}

func findLowestCostNode(costs map[string]string) Node {
	lowestCost := -1
	fmt.Print(lowestCost)
	lowestCostNode := Node{}
	// Go through each node
	for node, cost := range costs {
		// If it's the lowest cost so far and hasn't been processed yet...
		// Set it as the lowest new cost node
		lowestCost = cost
		lowestCostNode = Node{Key: node, Value: cost}
	}
	return lowestCostNode
}

func main() {
	//the graph
	//making graph as a map of map
	graph := make(map[string]map[string]int)
	graph["start"] = make(map[string]int)
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2

	graph["a"] = make(map[string]int)
	graph["a"]["fin"] = 1

	graph["b"] = make(map[string]int)
	graph["b"]["a"] = 3
	graph["b"]["fin"] = 5

	graph["fin"] = make(map[string]int)

	// the costs table
	costs := make(map[string]int)
	costs["a"] = 6
	costs["b"] = 2
	costs["fin"] = float64(uvinf)

	// the parents table
	parents := make(map[string]int)
	parents["a"] = "start"
	parents["b"] = "start"
	parents["fin"] = ""

	processed := make([]Node, 0)
	node := Node{}

	// Find the lowest-cost node that you haven't processed yet.
	node = findLowestCostNode(costs)
	for  node != nil {
		cost := costs[node.Key]
		// Go through all the neighbors of this node.
		neighbors := graph[node.Key]
		for key, value := range neighbors {
			newCost := cost + value
			// If it's cheaper to get to this neighbor by going through this node...
			if costs[key] > newCost {
				// ... update the cost for this node.
				costs[key] = newCost
				// This node becomes the new parent for this neighbor.
				parents[key] = value
			}
		}
		// Mark the node as processed.
		processed = append(processed, node)
		// Find the next node to process, and loop.
		node = findLowestCostNode(costs)
	}
}







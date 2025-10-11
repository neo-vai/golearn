package graph_test

import (
	"testing"

	"github.com/neo-vai/golearn/chdata/graph"
)

func buildGraph(Graph *graph.Graph) []*graph.Node {
	node1 := Graph.AddNode(graph.NewNode(1))
	node2 := Graph.AddNode(graph.NewNode(2))
	node3 := Graph.AddNode(graph.NewNode(3))
	node4 := Graph.AddNode(graph.NewNode(4))
	node5 := Graph.AddNode(graph.NewNode(5))
	node6 := Graph.AddNode(graph.NewNode(6))
	node7 := Graph.AddNode(graph.NewNode(7))
	node8 := Graph.AddNode(graph.NewNode(8))
	node9 := Graph.AddNode(graph.NewNode(9))

	Graph.AddEdge(graph.NewEdge(node1, node2, 1))
	Graph.AddEdge(graph.NewEdge(node2, node3, 1))

	Graph.AddEdge(graph.NewEdge(node4, node1, 2))
	Graph.AddEdge(graph.NewEdge(node2, node5, 3))
	Graph.AddEdge(graph.NewEdge(node3, node6, 3))

	Graph.AddEdge(graph.NewEdge(node5, node4, 1))
	Graph.AddEdge(graph.NewEdge(node5, node6, 3))

	Graph.AddEdge(graph.NewEdge(node7, node4, 2))
	Graph.AddEdge(graph.NewEdge(node5, node8, 2))
	Graph.AddEdge(graph.NewEdge(node9, node6, 1))

	Graph.AddEdge(graph.NewEdge(node8, node7, 2))
	Graph.AddEdge(graph.NewEdge(node8, node9, 2))

	return []*graph.Node{node1, node2, node3, node4, node5, node6, node7, node8, node9}
}

func TestGraphDirected(t *testing.T) {
	Graph := graph.NewGraph(true)
	nodes := buildGraph(Graph)

	min, i := Graph.Min(nodes[0], nodes[6])
	answer := 5
	if i != answer {
		minP := make([]int, 0, len(min))
		for _, v := range min {
			minP = append(minP, v.Value)
		}
		t.Errorf("expected distance %d, got %d\n%v\n", answer, i, minP)
	}
}
func TestGraphUnDirected(t *testing.T) {
	Graph := graph.NewGraph(false)
	nodes := buildGraph(Graph)

	min, i := Graph.Min(nodes[0], nodes[6])
	answer := 3
	if i != answer {
		minP := make([]int, 0, len(min))
		for _, v := range min {
			minP = append(minP, v.Value)
		}
		t.Errorf("expected distance %d, got %d\n%v\n", answer, i, minP)
	}
}

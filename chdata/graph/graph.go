package graph

import "container/list"

type Graph struct {
	Nodes      []*Node
	isDirected bool
}

func (graph *Graph) IsDirected() bool {
	return graph.isDirected
}

func NewGraph(isDirect bool) *Graph {
	return &Graph{
		Nodes:      []*Node{},
		isDirected: isDirect,
	}
}

func (graph *Graph) AddNode(node *Node) *Node {
	graph.Nodes = append(graph.Nodes, node)
	return node
}

func (graph *Graph) AddEdge(edge *Edge) {
	edge.From.EdgesOut = append(edge.From.EdgesOut, edge)
	edge.To.EdgesIn = append(edge.To.EdgesIn, edge)

	if !graph.isDirected {
		reverse := NewEdge(edge.To, edge.From, edge.Weight)
		edge.From.EdgesIn = append(edge.From.EdgesIn, reverse)
		edge.To.EdgesOut = append(edge.To.EdgesOut, reverse)
	}
}

type Node struct {
	Value    int
	EdgesIn  []*Edge
	EdgesOut []*Edge
}

func NewNode(value int) *Node {
	return &Node{
		Value:    value,
		EdgesIn:  []*Edge{},
		EdgesOut: []*Edge{},
	}
}

type Edge struct {
	From   *Node
	To     *Node
	Weight int
}

func NewEdge(from, to *Node, weight int) *Edge {
	return &Edge{
		From:   from,
		To:     to,
		Weight: weight,
	}
}

// BFS
func (graph *Graph) Min(start, finish *Node) (path []*Node, weight int) {
	if start == finish {
		return []*Node{start}, 0
	}

	visited := make(map[*Node]bool)
	prev := make(map[*Node]*Node)

	queue := list.New()
	queue.PushBack(start)
	visited[start] = true

	for queue.Len() > 0 {
		elem := queue.Front()
		node := elem.Value.(*Node)
		queue.Remove(elem)

		for _, edge := range node.EdgesOut {
			neighbor := edge.To
			if !visited[neighbor] {
				visited[neighbor] = true
				prev[neighbor] = node
				queue.PushBack(neighbor)

				if neighbor == finish {
					p := []*Node{finish}
					cur := finish
					for cur != start {
						cur = prev[cur]
						p = append([]*Node{cur}, p...)
					}
					return p, len(p)
				}
			}
		}
	}
	return
}

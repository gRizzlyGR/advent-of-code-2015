package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type node struct {
	from string
}

type edge struct {
	node
	cost int
}

type graph struct {
	nodes []*node
	edges map[node][]*edge
}

func (n *node) String() string {
	return fmt.Sprintf("%v", n.from)
}

func (e *edge) String() string {
	return fmt.Sprintf("%v: %v", e.node, e.cost)
}

// func (g *graph) String() string {
// 	return fmt.Sprintf("%v", g.edges)
// }

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatalln("Cannot read file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	trip := newGraph()

	for scanner.Scan() {
		node, edge := parseString(scanner.Text())

		trip.addNode(node)
		trip.addEdge(*node, edge)
	}

	// for k, v := range trip.edges {
	// 	fmt.Println(k, ":", v)
	// }
	// fmt.Println(trip.edges[&node{"Faerun"}])
	// fmt.Println(trip.edges)
	travelCost(trip)
}

func parseString(s string) (*node, *edge) {
	fields := strings.Fields(s)

	n := &node{
		from: fields[0],
	}

	to := node{
		from: fields[2],
	}
	distance, _ := strconv.Atoi(fields[4])

	e := &edge{
		node: to,
		cost: distance,
	}

	return n, e
}

func (g *graph) addNode(n *node) {

	if !contains(n, g.nodes...) {
		g.nodes = append(g.nodes, n)
	}
}

func (g *graph) addEdge(n node, e *edge) {
	if g.edges == nil {
		g.edges = make(map[node][]*edge)
	}
	g.edges[n] = append(g.edges[n], e)
}

func newGraph() *graph {
	nodes := make([]*node, 0)
	edges := make(map[node][]*edge)
	return &graph{nodes, edges}
}

func contains(n *node, nodes ...*node) bool {
	for _, val := range nodes {
		if val.from == n.from {
			return true
		}
	}

	return false
}

func travelCost(g *graph) int {
	min := math.MaxInt64
	for node := range g.edges {
		fmt.Println(node)
		fmt.Println(traverse(g, &node, 0))
	}

	return min
}

func traverse(g *graph, n *node, partialCost int) int {
	edges, ok := g.edges[*n]
	if !ok {
		return partialCost
	}

	next := edges[0]
	//fmt.Println("\t", next)
	partialCost += next.cost
	//fmt.Println("\t\tPartial cost:", partialCost)
	return traverse(g, &next.node, partialCost)
}

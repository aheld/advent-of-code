package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Vertex struct {
	Key      string
	Vertices map[string]*Vertex
}

func (v *Vertex) isBigCave() bool {
	return strings.ToUpper(v.Key) == v.Key
}

func NewVertex(key string) *Vertex {
	return &Vertex{
		Key:      key,
		Vertices: map[string]*Vertex{},
	}
}

func (v *Vertex) String() string {
	s := v.Key + ":"
	for _, neighbor := range v.Vertices {
		s += " " + neighbor.Key
	}
	return s
}

type Graph struct {
	Vertices map[string]*Vertex
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: map[string]*Vertex{},
	}
}

func (g *Graph) AddVertex(key string) {
	if _, ok := g.Vertices[key]; ok {
		return
	}
	v := NewVertex(key)
	g.Vertices[key] = v
}

func (g *Graph) AddEdge(k1, k2 string) {
	g.AddVertex(k1)
	g.AddVertex(k2)
	v1 := g.Vertices[k1]
	v2 := g.Vertices[k2]

	if _, ok := v1.Vertices[v2.Key]; ok {
		return
	}

	v1.Vertices[v2.Key] = v2
	v2.Vertices[v1.Key] = v1

	g.Vertices[v1.Key] = v1
	g.Vertices[v2.Key] = v2
}

func (g *Graph) String() string {
	s := ""
	i := 0
	for _, v := range g.Vertices {
		if i != 0 {
			s += "\n"
		}
		s += v.String()
		i++
	}
	return s
}

var allPaths = make([][]string, 0)

//DFS all paths
func (g *Graph) getAllPaths(start *Vertex, end *Vertex, part int) {
	visitedCount := make(map[string]int)
	// numVertex := len(g.Vertices)
	// fmt.Println("Num Vertex", numVertex)
	for name := range g.Vertices {
		visitedCount[name] = 0
	}
	pathList := make([]string, 0)
	pathList = append(pathList, start.Key)
	printAllPathsUtil(start, end, visitedCount, pathList, part)
}

func indexOf(needle *Vertex, haystack []string) int {
	for k, v := range haystack {
		if v == needle.Key {
			return k
		}
	}
	return -1
}

func printAllPathsUtil(start *Vertex, end *Vertex, visitedCount map[string]int, localPathList []string, part int) {
	if start.Key == end.Key {
		allPaths = append(allPaths, localPathList)
		// fmt.Println(localPathList)
		return
	}
	// After reviewing the available paths, you realize you might have time to visit a single small cave twice.
	//Specifically, big caves can be visited any number of times, a single small cave can be visited at most twice,
	// and the remaining small caves can be visited at most once.
	//However, the caves named start and end can only be visited exactly once each:
	//once you leave the start cave, you may not return to it, and
	//once you reach the end cave, the path must end immediately.

	canBeVisited := func(v *Vertex, visitedCount map[string]int, part int) bool {
		if v.isBigCave() {
			return true
		}
		if part == 1 {
			return visitedCount[v.Key] < 1
		}
		if visitedCount[v.Key] == 0 {
			return true
		}

		for k, vv := range visitedCount {
			if vv > 1 && strings.ToUpper(k) != k {
				// fmt.Println("Error for ", v.Key, k, vv, localPathList)
				return false
			}
		}
		return true
	}

	visitedCount[start.Key]++
	for _, v := range start.Vertices {
		// fmt.Printf("At %v, testing to enter %v, result %v: %v\n", start.Key, v.Key, v.isBigCave() || canBeVisited(v, visitedCount, part), visitedCount)
		if canBeVisited(v, visitedCount, part) {
			if v.Key == "start" && len(localPathList) > 1 {
				// fmt.Println("Rejecting Start", localPathList)
				continue //can't go back to start
			}
			localPathList = append(localPathList, v.Key)
			visitedCountCopy := make(map[string]int)
			for k, v := range visitedCount {
				visitedCountCopy[k] = v
			}
			copyLocalPathList := make([]string, len(localPathList))
			copy(copyLocalPathList, localPathList)
			printAllPathsUtil(v, end, visitedCountCopy, copyLocalPathList, part)
			index := indexOf(v, localPathList)
			localPathList = append(localPathList[:index], localPathList[index+1:]...)
		}
	}
	visitedCount[start.Key] = 0
}

func getAnswer(filename string, part int) int {
	allPaths = make([][]string, 0)
	data := strings.Split(loadFile(filename), "\n")
	g := NewGraph()
	for _, line := range data {
		parts := strings.Split(line, "-")
		g.AddEdge(parts[0], parts[1])
	}
	// fmt.Println(g)
	fmt.Println("GO")
	g.getAllPaths(g.Vertices["start"], g.Vertices["end"], part)
	total := len(allPaths)
	if part > 10 {
		for i, p := range allPaths {
			fmt.Println(i, p)
		}
	}
	allPaths = make([][]string, 0)
	return total
}

func part1(filename string) int {
	return getAnswer(filename, 1)
}

func part2(filename string) int {
	return getAnswer(filename, 2)
}

func loadFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(data)
}

func main() {
	fmt.Println("Part1 ", part1("input.txt"))
	fmt.Println("Part2 ", part2("input.txt"))
}

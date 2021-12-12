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
	isVisited := make(map[string]bool)
	numVertex := len(g.Vertices)
	fmt.Println("Num Vertex", numVertex)
	for name, _ := range g.Vertices {
		isVisited[name] = false
	}
	pathList := make([]string, 0)
	pathList = append(pathList, start.Key)
	printAllPathsUtil(start, end, isVisited, pathList, part)
}

func printAllPathsUtil(start *Vertex, end *Vertex, isVisited map[string]bool, localPathList []string, part int) {
	if start.Key == end.Key {
		allPaths = append(allPaths, localPathList)
		// fmt.Println(localPathList)
		return
	}
	if !start.isBigCave() {
		isVisited[start.Key] = true
	}

	for _, v := range start.Vertices {
		if !isVisited[v.Key] {
			localPathList = append(localPathList, v.Key)
			copyIsVisited := make(map[string]bool)
			for k, v := range isVisited {
				copyIsVisited[k] = v
			}
			copyLocalPathList := make([]string, len(localPathList))
			copy(copyLocalPathList, localPathList)
			printAllPathsUtil(v, end, copyIsVisited, copyLocalPathList, part)
			indexOf := func(needle *Vertex, haystack []string) int {
				for k, v := range haystack {
					if v == needle.Key {
						return k
					}
				}
				return -1
			}
			index := indexOf(v, localPathList)
			localPathList = append(localPathList[:index], localPathList[index+1:]...)
		}
	}
	isVisited[start.Key] = false
}

func getAnswer(filename string, part int) int {
	data := strings.Split(loadFile(filename), "\n")
	g := NewGraph()
	for _, line := range data {
		parts := strings.Split(line, "-")
		g.AddEdge(parts[0], parts[1])
	}
	fmt.Println(g)
	fmt.Println("GO")
	g.getAllPaths(g.Vertices["start"], g.Vertices["end"], part)
	total := len(allPaths)
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

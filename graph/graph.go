package graph

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Graph struct {
	V     int // Number of vertices
	E     int // Number of Edges
	table map[int][]int
}

func NewGraph() (graph *Graph) {
	graph = new(Graph)
	graph.table = make(map[int][]int)
	return
}

func (g *Graph) InitFromFile(file string) {
	absPath, _ := filepath.Abs(file)
	f, err := os.Open(absPath)
	if err != nil {
		log.Error(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	g.V, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	g.E, _ = strconv.Atoi(scanner.Text())
	for scanner.Scan() {
		log.Debug(scanner.Text())
		v_pair := strings.Fields(scanner.Text())
		v1, _ := strconv.Atoi(v_pair[0])
		v2, _ := strconv.Atoi(v_pair[1])
		g.addEdge(v1, v2)
	}
	log.Debug(g)
}

func (g *Graph) addEdge(v1, v2 int) {
	g.table[v1] = append(g.table[v1], v2)
	g.table[v2] = append(g.table[v2], v1)
}

func (g Graph) Search(v int) {

}

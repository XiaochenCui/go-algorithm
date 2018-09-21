package graph

import (
	"os"
	"testing"
)

func TestSearch(t *testing.T) {
	g := new(Graph)
	g.InitFromFile("data/tinyG.txt")
	v := os.Args[1]
}

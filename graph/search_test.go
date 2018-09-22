package graph

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func TestSearch(t *testing.T) {
	graph := NewGraph()
	graph.InitFromFile("../data/tinyG.txt")

	search := DFSSearch{graph: *graph}

	v := 0
	vs := search.Search(v)
	assert.Equal(t, vs, []int{0, 1, 2, 3, 4, 5, 6})
	assert.Equal(t, search.count, 6, "Their should be 6 vertices")

	v = 9
	vs = search.Search(v)
	assert.Equal(t, vs, []int{9, 10, 11, 12})
	assert.Equal(t, search.count, 4, "Their should be 4 vertices")
}

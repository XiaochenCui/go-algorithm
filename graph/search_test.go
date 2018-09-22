package graph

import (
	"testing"

	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func init() {
	filenameHook := filename.NewHook()
	log.AddHook(filenameHook)
	log.SetLevel(log.DebugLevel)
}

func TestSearch(t *testing.T) {
	graph := NewGraph()
	graph.InitFromFile("../data/tinyG.txt")

	assert := assert.New(t)

	search := NewDFSSearch(graph)
	v := 0
	connected := search.Search(v)
	expected := map[int]bool{
		0: true,
		1: true,
		2: true,
		3: true,
		4: true,
		5: true,
		6: true,
	}
	assert.Equal(expected, connected)
	assert.Equal(7, search.count, "Their should be 6 vertices")

	search = NewDFSSearch(graph)
	v = 9
	connected = search.Search(v)
	expected = map[int]bool{
		9:  true,
		10: true,
		11: true,
		12: true,
	}
	assert.Equal(expected, connected)
	assert.Equal(4, search.count, "Their should be 4 vertices")
}

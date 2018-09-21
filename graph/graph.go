package graph

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type Graph struct {
	V int // Number of vertices
	E int // Number of Edges
}

func (g Graph) InitFromFile(file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Error(err)
	}
	defer f.Close()
}

func (g Graph) Search(v int) {

}

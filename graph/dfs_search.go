package graph

type DFSSearch struct {
	graph     Graph
	marked    map[int]bool
	count     int
	connected map[int]bool
}

func NewDFSSearch(g *Graph) (search *DFSSearch) {
	search = new(DFSSearch)
	search.graph = *g
	search.marked = make(map[int]bool)
	search.connected = make(map[int]bool)
	return
}

// Search find all vertices connected with s
func (search *DFSSearch) Search(s int) map[int]bool {
	search.dfs(s)
	return search.connected
}

func (search *DFSSearch) dfs(s int) {
	for _, v := range search.graph.table[s] {
		if !search.marked[v] {
			search.marked[v] = true
			search.count++
			search.connected[v] = true
			search.dfs(v)
		}
	}
}

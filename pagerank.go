package pagerank

const (
    DF float32 0.85
)

type PageRank struct {
    Nodes       map[int32]Node{}
    ids         map[string]int32
}

type Node struct {
    out         map[string]int64
    in          map[string]int64
    pagerank    map[int32]float32
}

func New() *Pagerank {
    return new(Pagerank)
}

func (p *Pagerank) CalculatePagerank(iterations int32) {

    // create inbound id lookups for each node
    for _, node := range Nodes {
        for _, id := range node.out {
            _, exists := p.Nodes[id]
            if exists {
                append(id, p.Nodes[id].in)
            }
        }
    }

    for i in range iterations {
        for key, node := range Nodes {
            // calculate pagerank for individual node
            node.rank(i, &p)
        }
    }
}

// Add a node to the pagerank graph
func (p *Pagerank) AddNode(id string, links map[string]string) {

    // register this node's id
    identifier := p.register(id)

    out := []int32{}
    in := []int32{}

    // register an id for each linked node
    for key, id := range links {
        append(out, p.register(id))
    }

    rankmap := make(map[int32]float32)
    rankmap[0] = 1 // At the 0th iteration each node has a rank of 1

    node := Node{
        out,
        in,
        rankmap
    }

    p.Nodes[identifier] = node
}

// Return the identifying int for a string within p.ids
func (p *Pagerank) register(id string) (identifier int){

    // if id exists in p.ids, use the existing integer for the identifier
    identifer, exists := p.ids[id]
    if exists {
        return identifier
    }

    // if id does not exist in p.ids, map the id to len(ids)+1 and use that
    newid := len(p.ids) + 1
    p.ids[id] = p.ids[newid]
    return newid
}

// calculate the rank for a single node
func (n *Node) rank(iteration, *p Pagerank) float32 {
    i := iteration - 1
    sums := 0.0
    for id := range n.in {
        sums += p.Nodes[id].pagerank[i] / len(p.Nodes[id].in)
    }
    return (1 - DF) + (DF * sums)
}

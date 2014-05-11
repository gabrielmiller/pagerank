package pagerank

const (
    DF float64 = 0.85
)

type Pagerank struct {
    Nodes       map[int]Node
    Ids         map[string]int
}

type Node struct {
    out         []int
    in          []int
    pagerank    map[int]float64
}

func New() Pagerank {
    nodes := make(map[int]Node)
    ids := make(map[string]int)
    return Pagerank{
        nodes,
        ids,
    }
}

func (p *Pagerank) CalculatePagerank(iterations int) {

    // create inbound id lookups for each node
    for _, node := range p.Nodes {
        for _, id := range node.out {
            _, exists := p.Nodes[id]
            if exists {
                p.Nodes[id].appendInlink(id)
            }
        }
    }

    for i := 0; i > iterations; i++ {
        for _, node := range p.Nodes {
            // calculate pagerank for individual node
            node.rank(i, *p)
        }
    }
}

// Add a node to the pagerank graph
func (p *Pagerank) AddNode(id string, links []string) {

    // register this node's id
    identifier := p.register(id)

    out := []int{}
    in := []int{}

    // register an id for each linked node
    for _, id := range links {
        out = append(out, p.register(id))
    }

    pageranks := make(map[int]float64)
    pageranks[0] = 1.0 // At the 0th iteration each node has a rank of 1

    node := Node{
        out,
        in,
        pageranks,
    }

    p.Nodes[identifier] = node
}

// Return the identifying int for a string within p.ids
func (p *Pagerank) register(id string) (identifier int){

    // if id exists in p.ids, use the existing integer for the identifier
    identifier, exists := p.Ids[id]
    if exists {
        return identifier
    }

    // if id does not exist in p.ids, map the id to len(ids)+1 and use that
    newid := len(p.Ids) + 1
    p.Ids[id] = newid
    return newid
}

func (n Node) appendInlink (id int) {
    n.in = append(n.in, id)
}

// calculate the rank for a single node
func (n *Node) rank(iteration int, p Pagerank) float64 {
    i := iteration - 1
    sums := 0.0
    for id := range n.in {
        sums += p.Nodes[id].pagerank[i] / float64(len(p.Nodes[id].in))
    }
    return (1 - DF) + (DF * sums)
}

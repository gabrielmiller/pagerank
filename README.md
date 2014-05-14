##About
An implementation of the pagerank algorithm in go.

Pagerank is an algorithm that, in short, analyzes the connections between nodes in a graph and calculates the probability of landing on a particular node. It was/is the algorith behind Google's search engine. In google's implementation, the probability of landing on a node represents the likelihood that a searcher's result is the item for which the searcher was looking.

##Usage

    import "github.com/gmiller2007/pagerank"
    
    p := pagerank.New()

    // Add each node to the graph and point it toward its outbound nodes
    p.AddNode("identifier1", []string{'identifier2', 'identifier3'})

    // Calculate a given number of iterations of pagerank
    p.CalculatePagerank(3)

    // Look up an identifier's id
    id := p.Ids[identifier]

    // Look up the pagerank for an identifier
    pagerankArray := p.Nodes[id].Pagerank

    // Look up the pagerank at a specific iteration for an identifier
    pagerankArray := p.Nodes[id].Pagerank[iteration]

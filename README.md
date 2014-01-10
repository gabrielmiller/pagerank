##About
An implementation of the pagerank algorithm in go.

Pagerank is an algorithm that, in short, analyzes the connections between nodes in a graph and calculates the probability of landing on a particular node. It was/is the algorith behind Google's search engine. In google's implementation, the probability of landing on a node represents the likelihood that a searcher's result is the item for which the searcher was looking.

##Usage

    import "github.com/gmiller2007/pagerank"
    
    pagerank = pagerank.New()

    // Add each node to the graph
    pagerank.AddNode(nodeId)

    // Create 1-directional links between nodes
    pagerank.AddLink(originNodeId, targetNodeId)

    // Calculate the pagerank for a single node
    pagerankForNode = pagerank.CalculateSingle(damping, iterations, node)

    // Calculate the pagerank for every node in the graph
    pagerankArray = pagerank.CalculateFull(damping, iterations)

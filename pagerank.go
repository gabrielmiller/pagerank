package main

type pageRank struct {
    //id          int
    //links_out   int
    //pagerank    int
    ids         []int{}
    links       [][]int{}
    pageranks   []int{}
}

func New() *pageRank {
    pagerank := new(pageRank)
    return pagerank
}

func AddNode(node int) *pageRank {
    currentIds = pageRank.ids
    //pageRank.ids = currentIds + node
}

func AddLink(originLink int, destinationLink int) *pageRank {
    links = pageRank.links
    //pageRank.links[originLink][destinationLink] = true
}

func CalculateSingle(float64, int, int) *pageRank {
}

func CalculateFull(float64, int, int) *pageRank {
}

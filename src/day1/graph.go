package main

type CompleteGraphEdge struct {
	From, To, Weight int
}

type GraphEdge struct {
	To, Weight int
}

type WeightedAdjacencyList [][]GraphEdge
type WeightedAdjacencyMatrix [][]int
type AdjacencyList [][]int
type AdjacencyMatrix [][]int

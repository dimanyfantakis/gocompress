package main

import (
	"container/heap"
	"sort"
)

type Node struct {
	Char  rune
	Freq  int
	Left  *Node
	Right *Node
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].Freq == pq[j].Freq {
		return pq[i].Char < pq[j].Char
	}
	return pq[i].Freq < pq[j].Freq
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
	sort.Sort(pq)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	sort.Sort(pq)
	return item
}

// Function to build the Huffman binary tree
func buildHuffmanTree(freqMap map[rune]int) *Node {
	// Create leaf nodes and push them into the priority queue
	pq := make(PriorityQueue, len(freqMap))
	i := 0
	for char, freq := range freqMap {
		pq[i] = &Node{Char: char, Freq: freq}
		i++
	}
	heap.Init(&pq)

	// Build the Huffman tree.
	for pq.Len() > 1 {
		// Extract the two nodes with the lowest frequencies
		left := heap.Pop(&pq).(*Node)
		right := heap.Pop(&pq).(*Node)

		// Create a new internal node with the sum of frequencies
		internalNode := &Node{
			Char:  0,
			Freq:  left.Freq + right.Freq,
			Left:  left,
			Right: right,
		}

		// Add the new internal node back into the priority queue
		heap.Push(&pq, internalNode)
	}

	// The last remaining node in the priority queue is the root of the Huffman tree
	return heap.Pop(&pq).(*Node)
}

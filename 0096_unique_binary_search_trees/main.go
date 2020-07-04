package main

import (
	"fmt"
	"reflect"
)

/*
Given n, how many structurally unique BST's (binary search trees) that store values 1 ... n?
*/

func numTrees(n int) int {
	if n < 0 {
		panic("illegal n")
	}

	return constructBST(1, n+1, NewBSTCache())
}

func constructBST(start, end int, cache *BSTCache) (ways int) {
	if start > end {
		return 0
	}
	if start == end || (start+1) == end {
		return 1
	}

	if ways, ok := cache.Get(start, end); ok {
		return ways
	}

	// use one node to be root
	for mid := start; mid < end; mid++ {
		leftWay := constructBST(start, mid, cache)
		rightWay := constructBST(mid+1, end, cache)

		ways += leftWay * rightWay
	}

	cache.Set(start, end, ways)
	return ways
}

func NewBSTCache() *BSTCache {
	return &BSTCache{
		cache: make(map[string]int),
	}
}

type BSTCache struct {
	cache map[string]int
}

func (b *BSTCache) getKey(start, end int) string {
	return fmt.Sprintf("%d_%d", start, end)
}

func (b *BSTCache) Set(start, end, value int) {
	key := b.getKey(start, end)
	b.cache[key] = value
}

func (b *BSTCache) Get(start, end int) (value int, ok bool) {
	key := b.getKey(start, end)
	value, ok = b.cache[key]
	return value, ok
}

func main() {
	result := numTrees(3)
	if reflect.DeepEqual(result, 5) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
}

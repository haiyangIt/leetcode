/*
implement LRU algothim
*/

package main

type LinkedNode struct {
	prev  *LinkedNode
	after *LinkedNode
	key   string
}

type LRUStruct struct {
	CacheMap   map[string]interface{}
	LinkedList *LinkedNode
}

func (myself *LinkedNode) InsertAfter(key string) *LinkedNode {

}

func (myself *LinkedNode) InsertBefore(key string) *LinkedNode {

}

func main() {

}

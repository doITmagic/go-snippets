package hashtable

import (
	"fmt"
)

// ArraySize is the size of the array
const ArraySize = 10

// HashTable is a hash table
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket is a linked list in each slot of the hash table
type bucket struct {
	head *bucketNode
}

// bucketNode is a linked list node that holds the key
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search will take in a key and return true if that key is stored in the hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)

}

// Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// Insert will take in a key, create a node with the key and insert the node in the bucket
func (b *bucket) insert(key string) {
	if !b.search(key) {
		newNode := &bucketNode{key: key}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Printf("Key '%s' already exists \n", key)
	}
}

// Search will take in a key and return true if the bucket has that key
func (b *bucket) search(key string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *bucket) delete(key string) {

	var prevNode *bucketNode
	currentNode := b.head

	for currentNode != nil {
		if currentNode.key == key {
			if prevNode == nil {
				b.head = currentNode.next
			} else {
				prevNode.next = currentNode.next
			}
		}
		prevNode = currentNode
		currentNode = currentNode.next
	}
}

func InitHashTable() *HashTable {
	result := &HashTable{}
	//init array with buckets
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

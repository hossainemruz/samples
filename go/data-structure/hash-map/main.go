package main

import "fmt"

const TOTAL_BUCKETS = 10007

type Node struct {
	key   string
	value string
	next  *Node
}

type HashMap struct {
	items []*Node
}

func NewHashMap() *HashMap {
	return &HashMap{
		items: make([]*Node, TOTAL_BUCKETS),
	}
}

func (h *HashMap) Add(key, value string) {
	// Calculate the hash of the key
	index := hash(key)
	// The node pointed by the index.
	node := h.items[index]
	fmt.Printf("Inserting %q at index %d\n", key, index)
	// If the node is empty, then put the data there.
	if node == nil {
		h.items[index] = &Node{key: key, value: value}
	} else {
		// A collision has occurred.
		// We are going to use separate chaining for collision resolution.
		// Ref: https://en.wikipedia.org/wiki/Hash_table#Separate_chaining
		for node != nil {
			if node.key == key {
				// The key already exist, so update the value
				node.value = value
				return
			}
			// Go to the next node
			node = node.next
		}
		// The key does not exist, so add a new node for the key at the end of the linked-list.
		node.next = &Node{
			key:   key,
			value: value,
		}
	}
}

func (h *HashMap) Get(key string) (string, bool) {
	// Calculate the hash of the key
	index := hash(key)
	fmt.Printf("Looking for %q at index %d\n", key, index)
	// The node pointed by the index.
	node := h.items[index]
	// Now, check if the desired key does exist in the linked-list
	for node != nil {
		if node.key == key {
			// We have found the desired key, so return the value
			return node.value, true
		}
		// Go to next node
		node = node.next
	}
	// No node found with the desired key, so return false.
	return "", false
}

func (h *HashMap) Remove(key string) {
	// Calculate the hash of the key
	index := hash(key)
	// The node pointed by the index.
	node := h.items[index]
	prevNode := node

	fmt.Printf("Removing %q from index %d. node: %v\n", key, index, node)
	// Traverse the linked-list and check if the key exist or not.
	for node != nil {
		if node.key == key {
			// Found the desired key.
			// Link previous node with the next node so that this node become unlinked from the linked list.
			prevNode.next = node.next
			// Empty the current node.
			// This handles the case where prevNode == node.
			node = nil
			fmt.Println("h.items[index]: ", h.items[index])
			return
		}
		// Go to the next node
		prevNode = node
		node = node.next
	}
}

func hash(key string) int {
	// We are going to use polynomial rolling hash function to generate the hash.
	prime := 31
	powerOfPrime := 1
	hashValue := 0

	// Loop through each character and calculate polynomial hash
	for i := range key {
		hashValue = (hashValue + (int(key[i])*powerOfPrime)%TOTAL_BUCKETS) % TOTAL_BUCKETS
		powerOfPrime = (prime * powerOfPrime) % TOTAL_BUCKETS
	}
	return hashValue
}

func main() {

	// declare a new hash map
	mp := NewHashMap()

	// insert sample data
	fmt.Println("Inserting 'key1: value1'")
	mp.Add("key1", "value1")
	fmt.Println("Inserting 'key2: value2'")
	mp.Add("key2", "value2")
	fmt.Println("Inserting 'key3: value3")
	mp.Add("key3", "value3")

	// get value
	v, found := mp.Get("key2")
	fmt.Println("Get 'key2' found: ", found, " value: ", v)

	// remove "key2"
	fmt.Println("Removing 'key2'")
	mp.Remove("key2")

	// try getting key2 again
	fmt.Println("Get 'key2' found: ", found, " value: ", v)
}

package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	var firstname, lastname string
	fmt.Scan(&firstname, &lastname)

	fmt.Printf("Hello %s %s, welcome to the first GO program\n", firstname, lastname)

	// Pointers
	i, j := 45, 100
	fmt.Println(i, j)
	fmt.Println(&i, &j)

	p := &i
	fmt.Println(p)
	fmt.Println(*p)

	testHashTable := &HashTable{}
	fmt.Println(testHashTable)

}

// Custom HashTable implimentation
const ArraySize = 7

// Custom HashTable data sturcture implimentation
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket is a linked list in eaach slot of the
type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

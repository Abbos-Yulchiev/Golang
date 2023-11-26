package helper

import "fmt"

func Greeting(firstname string, lastname string) string {

	return "Hello " + firstname + " " + lastname
}

func Welcoming(firstname string) string {
	return "Welcome " + firstname + ", glad to see you!"
}

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

	fmt.Println(Greeting(firstname, lastname))
	fmt.Println(Welcoming(firstname))
}

// Custom HashTable implimentation
const ArraySize = 7

// Custom HashTable data sturcture implimentation
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket is a linked list in each slot of the hash table array
type bucket struct {
	head *bucketNode
}

// bucketNode is a liniked list node that holds the key
type bucketNode struct {
	key  string
	next *bucketNode
}

package main

import "fmt"


func main() {
	// Declare and add kv
	people := map[string]int{"dave": 22, "divine": 22, "solomon": 22}

	for k,v := range people {
		fmt.Println("Key:",k, "Value:", v)
	}

	persons := []string{"David", "Solomon", "Samson", "Naomi"}
	for _, person := range persons {
		fmt.Println(person)
	}
}
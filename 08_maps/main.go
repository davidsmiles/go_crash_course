package main

import "fmt"


func main() {
	// Define map
	emails := make(map[string]string)
	emails["bob"] = "bob@gmail.com"
	emails["david"] = "david@gmail.com"

	fmt.Println(emails["bob"])

	emails["mike"] = "mike@gmail.com"
	delete(emails, "bob")

	fmt.Println(emails)

	// Declare and add kv
	people := map[string]int{"dave": 22, "divine": 22, "solomon": 22}

	for k,v := range people {
		fmt.Println("Key:",k, "Value:", v)
	}

	persons := []string{"David", "Solomon", "Samson", "Naomi"}
	for i:= range persons {
		fmt.Println(persons[i])
	}
}
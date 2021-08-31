package main

import "fmt"


func main(){
	// Arrays

	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and assign
	fruitSlice := []string{"Apple", "Orange", "Mango"}

	fruitSlice = append(fruitSlice, "Gueva", "Banana")
	fmt.Println(fruitSlice)

	// Use make
	var makeSlice = make([]string, 10)
	makeSlice = append(makeSlice, "Hi", "Hello")
	makeSlice[0] = "How far"
	makeSlice[10] = "shit"
	fmt.Println(makeSlice)
}
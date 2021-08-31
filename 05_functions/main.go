package main

import "fmt"

func greeting(name string) string {
	return "Good morning Mr " + name
}

func sum(num ...int) int {
	sum := 0
	for _, n := range num {
		sum += n
	}
	return sum
}

func main(){
	defer fmt.Println(greeting("David"))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println("hELLO")
}
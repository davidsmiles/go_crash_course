package main

import "fmt"


func main(){

	x := 12
	y := 20

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if y < x {
		fmt.Println("yeah")
	} else {
		fmt.Println("shit")
	}

	color := "red"

	switch color {
	case "blue": 
		fmt.Println("Color is blue")
	case "red":
		fmt.Println("Color is red")
	default:
		fmt.Println("This is the default")
	}

}

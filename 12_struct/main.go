package main

import "fmt"

type person struct {
	firstname string
	lastname string
	gender string
}

func (p *person) getMarried(spouseLastname string){
	if p.gender == "m" {
		return
	} else {
	   p.lastname = spouseLastname
	}
}
 
func main() {
	person1 := person{"naomi", "james", "f"}
	person1.getMarried("ugbero")

	fmt.Println(person1)
}
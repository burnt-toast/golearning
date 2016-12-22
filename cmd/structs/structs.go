package main

import (
	"fmt"
	"github.com/burnt-toast/golearning/person"
)

func main() {
	
	fmt.Println(person.Person{"bob", 20})             //create a new person bob
	fmt.Println(person.Person{Name: "sara", Age: 20}) //create with named fields
	fmt.Println(person.Person{Name: "fred"})          // 0 value of age
	
	fmt.Println(&person.Person{"jake", 26}) // print pointer to struct we created
	s := person.Person{"test", 20}
	fmt.Println(s.Name) //access property of struct
	sp := &s
	fmt.Println(sp.Name)
	
	sp.Age = 21
	fmt.Println(sp.Age)
	fmt.Println(s.Age)
	
	
}

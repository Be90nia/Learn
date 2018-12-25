package main

import "fmt"

type person struct {
	name, sex string
	age       int
}

func main() {
	var people person
	people.name = "name"
	people.sex = "male"
	people.age = 1
	fmt.Println(people)
}

package structs

import "fmt"

type person struct {
	name string
	age  int
}

type secretAgent struct {
	person
	ltk bool
}

// PersonStruct tbd
func PersonStruct() {
	fmt.Println("PersonStruct")

	p1 := person{
		name: "james",
		age:  32,
	}

	fmt.Println(p1)

	p2 := person{
		name: "mikita",
		age:  23,
	}

	fmt.Println(p2)

	people := []person{p1, p2}

	fmt.Println(people)

	agent := secretAgent{
		person: person{
			name: "James Bond",
			age:  32,
		},
		ltk: true,
	}

	fmt.Println("embedded struct")
	fmt.Println(agent)
	fmt.Println(agent.ltk, agent.name)
}

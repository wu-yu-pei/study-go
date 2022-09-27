package main

import "fmt"

type PersonInterface interface {
	eat()
	sleep()
	run()
}

type Person struct {
	name string
	age  int
}

func (p Person) eat() {
	fmt.Printf("%v在eat\n", p.name)
}

func (p Person) sleep() {
	fmt.Printf("%v在sleep\n", p.name)
}

func (p Person) run() {
	fmt.Printf("%v在run\n", p.name)
}

func main() {
	var xiaoming Person

	xiaoming.name = "xiaoming"
	xiaoming.age = 12

	aaa := Person{
		name: "AAA",
		age:  12,
	}
	fmt.Printf("aaa: %v\n", aaa)

	xiaoming.eat()
	xiaoming.run()
	xiaoming.sleep()

	fmt.Printf("xiaoming: %T\n", xiaoming)

	var mapOne map[string]string

	mapTwo := map[string]string{
		"name": "age",
	}

	fmt.Printf("mapOne: %v\n", mapOne)
	fmt.Printf("mapTwo: %v\n", mapTwo)
}

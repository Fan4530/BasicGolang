package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {

}

type Cat struct {

}
type Cow struct {

}

func (d Dog) Speak() string{
	return "Wow"
}

func (c Cat) Speak() string{
	return "miao"
}

func (cow Cow) Speak() string{
	return "moo"
}

func main() {
	poodle := Animal(Dog{})
	fmt.Println(poodle.Speak())

	animals := []Animal{Dog{}, Cat{}, Cow{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
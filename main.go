package main

import "fmt"

type Function interface {
	Drive()
}

type Car struct{}

func (c Car) Drive() {
	fmt.Println("Car is driving")
}

type Bike struct{}

func (c Bike) Drive() {
	fmt.Println("Bike is riding")
}

func doOperation(i Function) {
	i.Drive()
}

func main() {
	c := Car{}
	b := Bike{}

	doOperation(c)
	doOperation(b)
}

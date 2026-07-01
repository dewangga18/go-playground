package main

import "fmt"

func main() {
	fmt.Println("\n=== Struct Example ===")
	customer1 := Customer{
		name: "Aaron Evanjulio",
		city: "Malang",
		age:  22,
	}
	customer2 := Customer{"Evanjulio Dewangga", "Kediri", 22}

	fmt.Println(customer1)
	fmt.Println(customer1.name)
	fmt.Println(customer2)
	fmt.Println(customer2.name)

	fmt.Println("\n=== Struct Method Example ===")
	customer1.sayHello()
	customer2.sayHello()
}

type Customer struct {
	name, city string
	age        int
}

func (cust Customer) sayHello() {
	fmt.Println("Hi,", cust.name)
}

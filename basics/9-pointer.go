package main

import "fmt"

// func main() {
// 	mainPointer()
// }

func mainPointer() {
	fmt.Println("\n== Pointer Example ==")
	address1 := Address{"Malang", "East Java", "Indonesia"}
	address2 := address1
	address2.City = "Kediri"

	fmt.Println("\n== Pass by Value ==")
	fmt.Println("address1 → ", address1)
	fmt.Println("address2 → ", address2)

	fmt.Println("\n== Pass by Reference ==") // or use asterisk operator
	address3 := &address1
	address3.City = "Surabaya"
	fmt.Println("address1 → ", address1)
	fmt.Println("address3 → ", address3)
	fmt.Println(address1 == *address3)
	fmt.Println(address1 == address2)

	fmt.Println("\n== New Function ==")
	address4 := new(Address)
	address5 := address4
	address5.Country = "Indonesia"
	fmt.Println(address4)
	fmt.Println(address5)
	fmt.Println(address4 == address5)

	fmt.Println("\n== Function in Pointer ==")
	address6 := &Address{} // declare
	address6.Country = "Konoha"
	fmt.Println("address6 before → ", address6)
	ChangeAddressToIndonesia(address6)
	fmt.Println("address6 after → ", address6)
	fmt.Println(address6.Country)
	
	address7 := Address{} // forgot to declare
	address7.Country = "Konoha"
	fmt.Println("address7 before → ", address7)
	ChangeAddressToIndonesia2(&address7)
	fmt.Println("address7 after → ", address7)

	fmt.Println("\n== Method in Pointer ==")
	man := &Man{"Dewa"}
	fmt.Println("man before → ", man.Name)
	man.Married()
	fmt.Println("man after → ", man.Name)
}

type Address struct {
	City, Province, Country string
}

// function in pointer
func ChangeAddressToIndonesia(a *Address) {
	a.Country = "Indonesia"
}

func ChangeAddressToIndonesia2(a *Address) {
	a.Country = "Indonesia"
}

// method in pointer
type Man struct {
	Name string
}

func (m *Man) Married() {
	m.Name = "Mr. " + m.Name
}
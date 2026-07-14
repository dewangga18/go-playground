package main

import (
	"fmt"
	_ "goplayground/basics/example_package/blank" // blank identifier used to force initialization of package's init function
	"goplayground/basics/example_package/db"
	"goplayground/basics/example_package/helper"
)

// func main() {
// 	mainPackage()
// }

func mainPackage() {
	fmt.Println("\n=== Package Import Example ===")
	fmt.Println(helper.Square(5))

	fmt.Println("\n=== Access Modifier Example ===")
	fmt.Println(helper.ApplicationName)
	// fmt.Println(helper.version) // not accessible

	fmt.Println("\n=== Package Init Example ===")
	fmt.Println("Database is ", db.GetDB())
}

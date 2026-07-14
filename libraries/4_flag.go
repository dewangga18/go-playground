package main

import(
	"fmt"
	"flag"
)

// func main() {
// 	mainFlag()
// }

func mainFlag() {
	// pattern name, value, description
	host := flag.String("host", "localhost", "host description")
	port := flag.Int("port", 8080, "port description")
	user := flag.String("user", "admin", "user description")
	password := flag.String("password", "123456", "password description")
	
	// run with: go run libraries/4-flag.go -host=localhost -port=8080 -user=root -password=123456
	// flag.Parse() // to read argument on terminal

	fmt.Println("Host:", *host)
	fmt.Println("Port:", *port)
	fmt.Println("User:", *user)
	fmt.Println("Password:", *password)
}
package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv" // use `go get github.com/joho/godotenv` or `go mod tidy` to download this package
)

// func main() {
// 	mainOs()
// 	mainEnv()
// 	goDotEnv()
// }

func mainOs() {
	args := os.Args  // run with : go run libraries/3-os.go test argumen1 argumen2

	fmt.Println("Arguments: ", len(args))
	for i, arg := range args {
		fmt.Println("Index: ", i, "Arg: ", arg)
	}

	host, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname: ", err)
	} else {
		fmt.Println("Hostname: ", host)
	}
}

func mainEnv() {
	e := os.Getenv("SAMPLE_ENV")
	fmt.Println("Environment variable SAMPLE_ENV: ", e)
	value, isExist := os.LookupEnv("SAMPLE_ENV")
	fmt.Println("Environment variable SAMPLE_ENV: ", value, " isExist: ", isExist)

	os.Setenv("SAMPLE_ENV", "hi_env")
	value, isExist = os.LookupEnv("SAMPLE_ENV")
	fmt.Println("Environment variable SAMPLE_ENV: ", value, " isExist: ", isExist)

	os.Unsetenv("SAMPLE_ENV")
	value, isExist = os.LookupEnv("SAMPLE_ENV")
	fmt.Println("Environment variable SAMPLE_ENV: ", value, " isExist: ", isExist)

	// envs := os.Environ()
	// fmt.Println("All environment variables local: ", envs)
}

func goDotEnv() {
	err := godotenv.Load(".env") // path .env
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	e := os.Getenv("TEXT")
	fmt.Println("Environment variable TEXT: ", e)
}
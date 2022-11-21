package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument : ")
	fmt.Println(args)
	fmt.Printf("Argument 1: %v .... argument 2 : %v", args[0], args[1])
	fmt.Println()

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", name)
	} else {
		fmt.Println("Error", err.Error())
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	fmt.Println("Username: ", username)
	fmt.Println("Password: ", password)
}

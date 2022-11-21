package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "put your host")
	var user *string = flag.String("user", "THOSIBA", "put your user")
	var password *string = flag.String("password", "", "put your password")
	var name *string = flag.String("name", "akbar", "put your name")
	flag.Parse()

	fmt.Println("Host : ", *host)
	fmt.Println("User : ", *user)
	fmt.Println("Password : ", *password)
	fmt.Println("Name : ", *name)

}

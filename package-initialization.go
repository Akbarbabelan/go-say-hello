package main

import (
	"fmt"

	database "github.com/Akbarbabelan/go-say-hello/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Akbar Babelan", "Akbar"))
	fmt.Println(strings.Contains("Akbar Babelan", "Budi"))

	fmt.Println(strings.Split("Akbar Babelan", " "))

	fmt.Println(strings.ToLower("Akbar Babaelan Musthofa"))
	fmt.Println(strings.ToUpper("Akbar Babaelan Musthofa"))
	fmt.Println(strings.ToTitle("Akbar Babaelan Musthofa"))
	fmt.Println(strings.Trim("      Akbar Babaelan Musthofa    ", " "))
	fmt.Println(strings.ReplaceAll("Akbar Akbar Akbar", "Akbar", "babelan"))

}

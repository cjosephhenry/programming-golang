package say

import "fmt"

func PrintHello() {
	fmt.Println("Hello")
}

func printWorld() {
	fmt.Println("World")
}

func PrintHellWorld() {
	PrintHello()
	printWorld()
}

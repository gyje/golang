package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("time", time.Now().UTC())
	fmt.Print("test")
	fmt.Printf("i")

	test()

	println("test")
	print("test")
	//printf("test")

}
func test() {
	for i := 1; i < 9; i++ {
		print(i, "\n")
	}
}

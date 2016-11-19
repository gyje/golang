package main

import (
	"fmt"
)

func main() {
	fmt.Println("defer的使用1")
	defer test() //defer可以让函数的执行延后
	//当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
	fmt.Println("defer的使用2")
}
func test() {
	fmt.Println("defer的使用3")
}

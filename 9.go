package main

import (
	"fmt"
	"runtime"
)

func say(s string){
	for i:=0;i<5;i++{
		runtime.Gosched()
		fmt.Println(s)
	}
}
func main() {
	var i = []byte{'1', '3', '6'}
	print(i[0:2])
	type class struct{
		name string
		age int32
	}
	var p class
	p.name="xingming"
	p.age=21
	fmt.Println("class name is %s",p.name)
	go say ("world")
	say("hello")
	}

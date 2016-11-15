package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 15; i++ {
		fmt.Println(i)
	}
	f:=0
	FGO:
	if f<15{	
		fmt.Println(f)
		f++	
	}
	goto FGO
}

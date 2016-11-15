package main

import (
	"fmt"
)

func Season(num int) (back string) {
	if num < 1 || num > 12 {
		fmt.Println("num error")
	} else {
		switch num {
		case 1:
			fmt.Println("yi yue fen")
		case 2:
			fmt.Println("er yue fen")
		case 3:
			fmt.Println("san yue fen")
		case 4:
			fmt.Println("si yue fen")
		case 5:
			fmt.Println("wu yue fen")
		case 6:
			fmt.Println("liu yue fen")
		case 7:
			fmt.Println("qi yue fen")
		case 8:
			fmt.Println("ba yue fen")
		case 9:
			fmt.Println("jiu yue fen")
		case 10:
			fmt.Println("shi yue fen")
		case 11:
			fmt.Println("11 yue fen")
		case 12:
			fmt.Println("12 yue fen")
		}
	}
    return back
}
func main() {
	Season(6)
}

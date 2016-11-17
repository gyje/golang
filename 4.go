package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().UTC()
	fmt.Println(t.Day())
	fmt.Println(t.Weekday())
	fmt.Println(t)
	i := "abc"
	m := &i
	fmt.Println("%s 指针地址为：%p", i, &i, len(i))
	fmt.Print(m)
	//关键字continue只能用在for循环中
	for n := 0; n < 10; n++ {
		if n == 5 {
			//continue
			break
		}
		fmt.Println(n)
	}
	print("")
}

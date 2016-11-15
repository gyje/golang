package main

import "fmt"

func max(num1, num2 int) (output1 int,output2 int) {
	if num1 > num2 {
		return num1,num1-num2
	} else {
		return num2,num2-num1
	}
}
func min(arg...int)(int){
	if arg[0]<arg[1]{
		return arg[1]
	}else{
		return arg[0]
	}
}
func main() {
	fmt.Printf("hello,world,你好世界！")
	fmt.Println("\t\nhello,world,你好世界")
	var vname2, vname3 = true, false
	vname1, vname4 := 1, 2
	vname5, vname6 := "test\n", "test2\t"
	fmt.Println(vname1, vname2, vname3, vname4, vname5, vname6)
	const pi = 3.141592653
	const i = "false"
	fmt.Printf(i)
	fmt.Println(pi, "\n")
	str := "string"
	str = "c" + str[6:] //切片从0开始计数
	fmt.Println("%s\n", str)
	doublestring := `hello,test
    多行字符串，不知道行不行，试一试`
	fmt.Println(doublestring)
	const (
		a = iota
		b = iota
		c = iota
	)
	const d = iota
	_, e := 1, 2
	fmt.Println(a, b, c, d, e)
	f := [...]int{1, 3, 2, 4}
	fmt.Println(f)
	g := [...][3]int{{4, 5, 6}, {1, 2, 3}}
	fmt.Println(g)
	h := []float32{1.2, 3.4, 5.6, 7.8} //这里h是切片类型
	fmt.Println(h[0:2])
	m := make(map[string]string)
	m["hello"] = "nohello"
	m1 := m
	m1["hello"] = "test"
	fmt.Println(m)
	x := 10
	y := 20
	x, y = y, x
	fmt.Println(x, y)
	max_num1,max_num2:=max(100, 200)
	fmt.Println(max_num1,max_num2)
	/*output1num:=output1
	output2num:=output2
	fmt.Println(output1num,output2num)*/
	min_num:=min(300,400)
	fmt.Println(min_num)
}

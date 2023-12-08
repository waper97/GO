package main

import "fmt"

func main() {
	// 注释
	/**
	多行注释
	*/
	var today = "星期五"
	var number = 6
	var str string = "hello"
	var num = 123
	myString := "very good"
	myNumber := 123
	var num1 int
	var str1 string
	var flag bool
	const name string = "fuck"

	fmt.Println("hello world")
	fmt.Println("自动换行")
	fmt.Println("hello world" + "自动换行")
	//格式化字符串  %s字符型 %d整数类型
	fmt.Printf("今天是%s,在一周中排第%d", today, number)
	println(str, num)
	println(myString, myNumber)
	println(num1)
	fmt.Println("num1默认值为:", num1)
	fmt.Println("str1默认值为:", str1)
	fmt.Println("bool默认值为:", &flag)
	fmt.Println("name:", name)

}

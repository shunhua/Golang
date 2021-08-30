package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}


func main() {

	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
	// 	1.结构体类似于PHP中的class\对象
	// 	2.结构体与指针结合遇到 不用*就能直接获取结构体元素内容
	//	3.结构体类型的函数能直接当成结构体方法使用（a.foo()） func (a A)foo int{}
	//	4.结构体函数 可以用结构体指针（指针变量）函数内部改变结构体值func (a *A)foo int{}
}

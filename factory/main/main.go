package main

import (
	"fmt"
	"go_code/factory/model"
)

func main() {
	var stu = model.NewStudent("tom", 100)
	//{
	//	Name : "tom",
	//	Score : 78.9,
	//}
	fmt.Println(stu.Name, stu.GetScore())
}

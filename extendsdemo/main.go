package main

import "fmt"

//通过结构体嵌入匿名结构体来实现

type Student struct {
	Name  string
	Age   int
	Score int
}

func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v \n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

func (stu *Student) testing() {
	println("小学生正在考试......")
}

type Pupil struct {
	Student //嵌入匿名结构体
}

type Graduate struct {
	Student //嵌入匿名结构体
}

func main() {
	pupil := &Pupil{}
	pupil.testing() //pupil.Student.testing()的简写
}

package main

import "fmt"

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
	Name string
}

func (b *B) SayOk() {
	fmt.Println("B SayOk", b.Name)
}

func main() {
	var b B
	b.Name = "jack"
	b.age = 100
	b.SayOk() //B SayOk jack
	b.hello() //A hello 	  使用b.A.Name = "jack"才可以将此 A hello jack
}

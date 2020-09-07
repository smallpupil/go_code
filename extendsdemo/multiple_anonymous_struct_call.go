package main

import "fmt"

type A1 struct {
	Name string
	age  int
}
type B1 struct {
	Name  string
	Score int
}

type C struct {
	A1
	B1
}

type D struct {
	a A1
}
type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV1 struct {
	*Goods
	*Brand
}

func main() {
	var c C
	//c.Name = "tom"  //error
	c.A1.Name = "han"

	var d D
	//d.Name = "Tom"  //D中嵌套的非匿名结构体，则直接不往上层找，只在此结构体找
	d.a.Name = "Tom"

	//多个嵌套直接赋值
	tv := &TV{Goods{"changhong", 100}, Brand{"chang", "1"}}
	fmt.Println("=-=", tv)
	tv1 := TV1{&Goods{"changhong", 100}, &Brand{"chang", "1"}}
	fmt.Println("=-=", *tv1.Brand)

}

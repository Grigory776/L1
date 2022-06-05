package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками, которые
представлены в виде структуры Point с инкапсулированными параметрами x,y и
конструктором.
*/

type Point struct{
	x, y float64
}

func (p *Point) Get()(x,y float64){
	return p.x, p.y
}

//конструктор
func (p *Point) Set(x,y float64){
	p.x = x
	p.y = y
}

func Distance(a,b Point)float64{
	x1, y1 := a.Get()
	x2, y2 := b.Get()
	return math.Sqrt((x2 - x1)*(x2 - x1) + (y2 - y1)*(y2 - y1))
}

func main(){
	var(
		A, B Point
	) 
	A.Set(1,1)
	B.Set(2,1)
	fmt.Println(Distance(A,B))
}
package main

import (
	"fmt"
	"math"
)

/*
Реализовать паттерн «адаптер» на любом примере.
*/

//интерфейсу удовлевторяют оъекты, которые имеют расстояние,
//и могут суммироваться по трем переменным
type DistSum interface{
	Distance()float64
	Sum(x,y,z float64)
}

//vector3d реализует интерфейс DistSum
type vector3d struct{
	x,y,z float64
}

func (v vector3d) Distance()float64{
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v *vector3d) Sum(k1,k2,k3 float64){
	v.x += k1
	v.y += k2
	v.z += k3
}

//vector2d не реализует интерфейс DistSum (метод sum принимает две переменные)
type vector2d struct{
	x,y float64
}

func (v vector2d) Distance()float64{
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *vector2d) Sum(k1,k2 float64){
	v.x += k1
	v.y += k2
}

//адптер структуры vector2d с интерфейсом DistSum
type adapterVector2d struct{
	*vector2d
}

func (a *adapterVector2d) Sum(k1,k2,k3 float64){
	a.vector2d.Sum(k1,k2)
}

func main(){
	var tmp DistSum
	//tmp = &vector2d{x:1, y:1} !не ок
	tmp = &adapterVector2d{&vector2d{x: 1, y:1}} // !ок
	fmt.Println(tmp.Distance()) 
	tmp.Sum(1,1,0)
	fmt.Println(tmp.Distance())
}
package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func swipe (a,b *int){
	*a,*b = *b,*a
}

func main(){
	a := 3
	b := 4
	swipe(&a,&b)
	fmt.Printf("a = %d, b = %d\n",a,b)
}
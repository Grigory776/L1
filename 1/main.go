package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры
Human (аналог наследования).
*/

// Родительская структура
type Human struct{
	Name string
	Surname string
	Age int
}

func (h Human) String() string{
	return fmt.Sprintf("%s %s, %d years", h.Name, h.Surname, h.Age) 
}

func (h *Human) Birthday(){
	h.Age++
}

// Дочерняя структура
type Action struct{
	Human
}

func main(){
	var h = Action{Human{Name: "Grigorii", Surname: "Zolotarev", Age: 21}}
	fmt.Println(h.String())
	h.Birthday()
	fmt.Println(h.String())
}
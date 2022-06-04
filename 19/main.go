package main

import "fmt"

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func reverse(str string)(string){
	r := []rune(str)
	left, right := 0, len(r) - 1
	for left < right{
		r[left], r[right] = r[right], r[left]
		left++
		right--
	}
	return string(r)	
}

func main(){
	fmt.Println(reverse("異абырвалг 異體字"))
}
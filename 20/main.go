package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/
func reverseText(str string)string{
	arr := strings.Split(str," ")
	left, right := 0, len(arr) - 1
	for left < right{
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return strings.Join(arr," ")
}



func main(){
	fmt.Println(reverseText("snow dog sun"))

}
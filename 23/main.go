package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

func removeSliceElement (arr []string, ind int) []string{
	return append(arr[:ind], arr[ind+1:]...)
}

func main(){
	arr := []string{"a", "b", "c", "d"}
	fmt.Println(removeSliceElement(arr,3))
}
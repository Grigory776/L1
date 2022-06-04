package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
собственное множество.
*/

//функция принимает слайс, возвращает собственное множество на основе этого слайса
func ownSet(arr []string) map[string]struct{}{
	res := make(map[string]struct{})
	for _, val := range arr{
		res[val] = struct{}{}
	}
	return res
}

func main(){
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(ownSet(arr))
}
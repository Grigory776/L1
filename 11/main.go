package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

//на вход функция принимает два целочисленных множества
//итерируется по одному из них и если элемент есть также и во втором множестве
//добавляет его в результатирующее множество
func intersectionSet (s1, s2 map[int]struct{})map[int]struct{}{
	res := make(map[int]struct{})
	for val := range s1 {
		if _, ok := s2[val]; ok{
			res[val] = struct{}{}
		}
	}
	return res
}

//вариация на случай если на вход подаются слайсы
func intersectionSlise(a, b []int) []int {
	counter := make(map[int]int)
	var result []int

	for _, elem := range a {
		if _, ok := counter[elem]; !ok {
			counter[elem] = 1
		} else {
			counter[elem] += 1
		}
	}
	for _, elem := range b {
		if count, ok := counter[elem]; ok && count > 0 {
			counter[elem] -= 1	
			result = append(result, elem)
		}
	}
	return result
}

func main(){
	set1 := make(map[int]struct{}) //множество натуральных чисел от 0 до 50
	for i:=0; i <= 50; i++{
		set1[i] = struct{}{}
	}
	set2 := make(map[int]struct{}) // множество четных чисел от 0 до 50
	for i:=0; i <= 50; i+=2{
		set2[i] = struct{}{}
	}
	fmt.Println(intersectionSet(set1, set2))

	var arr1 []int = []int{1,1,1}
	var arr2 []int = []int{3,4,1,1,7,8,9}
	fmt.Println(intersectionSlise(arr1,arr2))
}
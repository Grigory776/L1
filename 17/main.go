package main

import "fmt"

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func binarySearch(arr []int, num int) (res int, ok bool){
	res = -1
	if len(arr) == 0{
		return 
	}
	if len(arr) == 1{
		return arr[0], true
	}
	left := 0
	right := len(arr)-1
	var middle int
	for left <= right{
		//выбираем средний элемент диапозона и сравниваем его с искомым числом,
		//если искомый элемент больше, то откидываем левую часть диапозона,
		//в противном случае - правую
		middle = (left + right) / 2
		if (num == arr[middle]){

			return arr[middle], true

		} else if (num < arr[middle]){
			right = middle -1
		} else {
			left = middle + 1
		}
	}
	return
}

func main(){
	arr := []int{1,2,3,4,5,10,12,227,500}
	x, ok := binarySearch(arr,12)
	fmt.Printf("%v - %v\n", x, ok)
}
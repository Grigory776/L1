package main

import "fmt"

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами
языка.
*/

func QuickSort(arr []string){
	quickSort(arr,0,len(arr) - 1)
}

func quickSort(arr []string, l, r int){
	if l >= r {
		return
	}
	//выбираем опорный элемент q и распологаем его так, чтобы
	//меньшие числа были левее, а большие - правее 		  
	q := partition(arr, l, r)
	//рекурсивно вызываем quicksort для левого и правого диапозона
	quickSort(arr,l,q-1) 
	quickSort(arr,q+1,r)
}

func partition(arr []string, l,r int)int{
	q := l 
	for i:=l; i < r; i++{
		//изначально опорный элемент - это последний элемент промежутка
		//если текущий элемент меньше опорного, то свайпаем этот элемент с q-ым
		if arr[i] <= arr[r]{
			arr[i], arr[q] = arr[q], arr[i]
			q++
		}
	}
	//после цикла q будет указывать на место опорного элемента
	arr[q], arr[r] = arr[r], arr[q]
	return q
}

func main(){
	arr := []string{"bcd", "aaa", "c", "Aa", "aaa", "dog", "sss", "foo"}
	QuickSort(arr)
	fmt.Println(arr)
}
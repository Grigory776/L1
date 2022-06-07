package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func PrintSliceSqr (arr []int){
	if len(arr) == 0{
		return 
	}
	var wg sync.WaitGroup
	for i := 0; i < len(arr); i++{
		wg.Add(1) //на каждой итерации добавляем в список одну горутину
		go func (i int){
			defer wg.Done()//по завершении убираем одну горутину из списка
			fmt.Println(arr[i]*arr[i])
		}(i)
	}
	wg.Wait()//ждем завершение все горутин
}

func main(){
	PrintSliceSqr([]int{2,4,6,8,10})
}
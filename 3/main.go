package main

import "fmt"

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(4+16+36….) с использованием конкурентных вычислений.
*/

func SumSquares (arr []int)(sum int){
	ch := make(chan int)
	defer close(ch)
	go func(){
		for _, val := range arr{
			go func(i int){
				ch <- (i*i)
			}(val)
		}
	}()
	for i:=0; i < len(arr); i++ {
		tmp := <- ch
		sum += tmp
	}
	return sum
}

func main(){
	fmt.Println(SumSquares([]int{2,4,6,8,10}))
}
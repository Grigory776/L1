package main

import (
	"fmt"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(4+16+36….) с использованием конкурентных вычислений.
*/

func SumSquares (arr []int)(sum int){
	ch := make(chan int)
	go func(){
		for _, val := range arr{
			go func(val int){
				ch <- (val*val) //отдельно запускаем все горутины, которые производят вычисление
			}(val)
		}
	}()
	for i:=0; i < len(arr); i++{ //получаем значение от горутин len(arr) раз
		sum+= <-ch
	}
	close(ch)
	return sum
}

func main(){
	fmt.Println(SumSquares([]int{2,4,6,8,10}))
}
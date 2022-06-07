package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в
канал, а с другой стороны канала — читать. По истечению N секунд программа
должна завершаться.
*/

func imitation(N time.Duration){
	t := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), N) //устанавливаем таймер
	ch := make(chan string)
	var i = 1
	go func(i int){
		for { //бесконечный цикл отправки данынх в канал
			ch <- fmt.Sprintf("data number %d was sent to the channel\n",i)
			i++
			time.Sleep(1 * time.Second)
		}
	}(i)
	for {
		select{
		case data := <-ch:
			fmt.Print(data)
		case <-ctx.Done()://по истечении времени программа остановится
			cancel()
			fmt.Printf("the program ran for %v second\n", time.Since(t))
			return
		}
	}	
}
func main(){
	if len(os.Args) < 2{
		log.Println("to run, you need to pass the number of seconds")
	}
	secStr := os.Args[1]
	sec, _ := strconv.Atoi(secStr)
	imitation(time.Duration(sec) * time.Second)
}
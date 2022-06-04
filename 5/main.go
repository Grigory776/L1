package main

import (
	"math/rand"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в
канал, а с другой стороны канала — читать. По истечению N секунд программа
должна завершаться.
*/

func Imitation(N int){
	ch := make(chan int)
	rand.Seed(time.Now().UnixNano())
	go func(){
		for {
			ch <- rand.Intn(100)
		}
	}()
	go func(){
		for {
			<-ch
		}
	}()
	time.Sleep(time.Duration(N) * time.Second)
}
func main(){

}
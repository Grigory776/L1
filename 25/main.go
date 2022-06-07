package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

//в небуфферизированный канал приходят данные по истечению времени t
func Sleep(t time.Duration) { 
	<-time.After(t)
}

func main(){
	t := time.Now()
	fmt.Println("Start!")
	Sleep(2 * time.Second)
	fmt.Printf("passed - %v\n",time.Since(t))
}
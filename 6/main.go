package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

//Завершение горутины через запись в кнаал
func Goroutine1(wg *sync.WaitGroup ,chSignal chan struct{}){
	//иммитация работы (например, здесь может быть запуск горутины с бесконечным циклом)
	t:=time.Now()
	defer wg.Done()

	<- chSignal
	fmt.Printf("The first goroutine terminated later - %v\n", time.Since(t))
	return
}

//Заверешние горутины через отмену контекста
func Goroutine2(wg *sync.WaitGroup, ctx context.Context){
	t:=time.Now()
	defer wg.Done()
	
	<-ctx.Done()
	fmt.Printf("The second goroutine terminated later - %v\n", time.Since(t))
	return
}

//Заверешние горутины через таймаут в контексте (аналогично с дедлайном)
func Goroutine3(wg *sync.WaitGroup, ctx context.Context){
	t:=time.Now()
	defer wg.Done()

	<- ctx.Done()
	fmt.Printf("The third goroutine terminated later - %v\n", time.Since(t))
	return
}


func main(){
	var wg sync.WaitGroup
	wg.Add(3)
	// подождем 1 секунду и отправим сигнал о завершении в 1 горутину
	ch := make(chan struct{})
	go Goroutine1(&wg, ch)
	time.Sleep(1 * time.Second)
	ch <- struct{}{}
	// подождем две секунды и отменим контекст для второй горутины
	ctx, cancel := context.WithCancel(context.Background())
	go Goroutine2(&wg,ctx)
	time.Sleep(2 * time.Second)
	cancel()
	// создадим контекст с таймаутом в три секунды
	ctx2, cancel := context.WithTimeout(context.Background(),3*time.Second)
	go Goroutine3(&wg,ctx2)

	wg.Wait()
	
}
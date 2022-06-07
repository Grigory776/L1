package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	//"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при
старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
способ завершения работы всех воркеров.
*/

type workers struct {
	amount int
	channel chan string
}

func (w *workers) Start(ctx context.Context){ //запуск amount воркеров и их чтение данных из канала
	for i:=1; i <= w.amount; i++{
		go func(c context.Context, i int){
			for {
				select{
				case d := <-w.channel:
					fmt.Printf("worker number %d read the entry: %s\n",i,d)
				case <-ctx.Done():
					fmt.Printf("worker number %d finished work\n",i)
					return
				}
			}
		}(ctx, i)
	}
}

func main(){
	ch := make(chan string)
	w := workers{amount: 10, channel: ch}
	ctx, cancel := context.WithCancel(context.Background())
	w.Start(ctx)
	go func() { //отправка данных в канал
		i:=1
		for{
			ch <- fmt.Sprintf("data number %d",i)
			i++
			time.Sleep(1 * time.Second)
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){ //горутина ждет сигнала для корректного завершения программы
		sign := make(chan os.Signal, 1)
		signal.Notify(sign, syscall.SIGINT)
		defer signal.Stop(sign)
		<-sign //если сигнал приходит, то передаем отмену через контекст
		cancel()
		time.Sleep(2 * time.Millisecond) //даем время воркерам завершиться
		close(ch)
		wg.Done()
	}()
	wg.Wait()
}
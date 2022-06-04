package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговоезначение счетчика.
*/

type counter struct{
	i int
	sync.Mutex
	sync.WaitGroup
}

func (c *counter) String() string{
	return fmt.Sprintf("count = %v",c.i)
}

func (c *counter) increment(){
	c.Lock()
	c.i++
	c.Unlock()
}

//взаимодействие счетчика с функцией 
//(будет посчитано количество горутин запущенных внутри функции)
func foo(c *counter){

	for i:=1; i<=50; i++{
		c.Add(1)
		go func(i int){
			defer c.Done()
			fmt.Println(i) //иммитация работы
			c.increment()
		}(i)
	}
	c.Wait()

}

func main(){
	counter := &counter{i: 0, WaitGroup: sync.WaitGroup{}, Mutex: sync.Mutex{}}
	foo(counter)
	fmt.Println(counter)
}
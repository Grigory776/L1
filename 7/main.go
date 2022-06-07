package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

//определим потокобезопасное отображение
type ThreadSafeMap struct{
	mp map[string]int
	sync.RWMutex
}

func InitThreadSafeMap()*ThreadSafeMap{
	return &ThreadSafeMap{mp: make(map[string]int), RWMutex: sync.RWMutex{}}
}

func (t* ThreadSafeMap) Store(key string, val int){
	t.Lock()
	defer t.Unlock()
	t.mp[key] = val
}

func (t* ThreadSafeMap) Load(key string)(int, bool){
	t.RLock()
	defer t.RUnlock()
	val, ok := t.mp[key]
	return val, ok
}

func (t *ThreadSafeMap)String()(res string){
	for key, val := range t.mp{
		res += fmt.Sprintf("%s:%d\n",key,val)
	}
	return
}

//конкурентно запишем данные в отображение
func writeRandom(t *ThreadSafeMap){
	var wg sync.WaitGroup
	for i:=0; i < 50; i++{
		wg.Add(1)
		go func(i int){
			t.Store(fmt.Sprint(i), rand.Int()%10)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main(){
	t := InitThreadSafeMap()
	writeRandom(t)
	fmt.Println(t)
	fmt.Println(len(t.mp) == 50)
}
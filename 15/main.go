package main

import "fmt"

/*
К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации.
var justString string
func someFunc() {
v := createHugeString(1 << 10)
justString = v[:100]
}
func main() {
someFunc()
}
*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10)

	//justString = v[:100]  - ожидается записать в justString первые сто символов большой строки
	//но string - это массив байтов => если символ занимает больше байта, то 
	//в justString попадут некорректные данные

	tmp := []rune(v) // преобразовываем строку в массив рун
	justString = string(tmp[:100]) // берем первые 100 рун и преобразовывем в стринг

	// Проверка
	fmt.Printf("1:\n%v\n", v[:100])
	fmt.Printf("2:\n%v\nsize justString = %v\n", justString, len(justString))
}

func createHugeString(size uint64)(res string){
	for i := 0; uint64(i) < size; i++{
		res += "異" // символ занимает 3 байта
	}
	return
}

func main() {
	someFunc()
}

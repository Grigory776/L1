package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false etc). Функция проверки должна быть
регистронезависимой.
Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func uniqueCharacters1(str string) bool{
	str = strings.ToLower(str)//делаем все буквы маленькими
	b := []rune(str)
	sort.Slice(b, func(i,j int)bool{ //сортируем строку
		return b[i] < b[j]
	})
	for i:=1; i<len(b); i++{
		if b[i] == b[i-1]{ //если найдутся соседние одинаковые символы, строка не уникальна 
			return false
		}
	}
	return true
}

func uniqueCharacters2(str string) bool{
	str = strings.ToLower(str)//делаем все буквы маленькими
	mp := make(map[string]bool)
	for _, val := range str{ //итерируемся по строке
		ok := mp[string(val)]
		if ok{ //проверяем добавлялась ли руна ранее
			return false
		} else {
			mp[string(val)] = true
		}
	}
	return true
}

func main(){
	fmt.Println(uniqueCharacters1("abcd"))
	fmt.Println(uniqueCharacters1("abCdefAaf"))
	fmt.Println(uniqueCharacters1("aabcd"))
	fmt.Println()
	fmt.Println(uniqueCharacters2("abcd"))
	fmt.Println(uniqueCharacters2("abCdefAaf"))
	fmt.Println(uniqueCharacters2("aabcd"))
}



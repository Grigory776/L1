package main

import (
	"errors"
	"fmt"
	"log"
	"math/big"
)
/*
Разработать программу, которая перемножает, делит, складывает, вычитает две
числовых переменных a,b, значение которых > 2^20.
*/

// sign = '+','-','*','/'
func bigCalcul(x,y *big.Float, sign string)(*big.Float,error){
	switch sign{
	case "+":
		return big.NewFloat(0).Add(x, y), nil
	case "-":
		return big.NewFloat(0).Sub(x, y), nil
	case "*":
		return big.NewFloat(0).Mul(x, y), nil
	case "/":
		return big.NewFloat(0).Quo(x, y), nil
	}
	return nil, errors.New("unknown sign transferred to a bigCalcul")
}

func main(){
	for {
		fmt.Println("введите арифметическое выражение: (через пробел, пример: 1 + 2)")
		var (
			strX, strY, sign string
		)
		fmt.Scanf("%v %v %v",&strX, &sign, &strY)
		x,y := new(big.Float), new(big.Float)
		_, ok := x.SetString(strX)
		if !ok {
			log.Println("passed not a number")
			break
		}
		_, ok = y.SetString(strY)
		if !ok {
			log.Println("passed not a number")
			break
		}
		res,er := bigCalcul(x,y,sign)
		if er != nil{
			log.Println(er)
			break
		}
		fmt.Printf("%v %v %v = %s\n",strX, sign, strY, res.String())
	}
}
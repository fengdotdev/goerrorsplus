package main

import (
	"fmt"

	"github.com/fengdotdev/goerrorsplus/e"
)




func division(n1, n2 int ) (val float64, err error) {

	defer func() {
		if r := recover(); r != nil {
		val = 0
		 err= fmt.Errorf("recover: %s", r)
		}}()

	if n2 == 0 {
		panic("div by zero")
	}
	return float64(n1)/float64(n2), nil
}


func main() {

	result, err := division(1,0)
	if err != nil {
		errp := e.E(err,"div failed",[]string{"math-err"},division,1,0)	
		fmt.Println(errp.V())
	}

	if err == nil {
		fmt.Println("result is ",result)
	}
}
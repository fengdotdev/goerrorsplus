package main

import (
	"fmt"
	"log"

	"github.com/fengdotdev/goerrorsplus/e"
)




func division(n1, n2 int ) (val float64, err error) {

	defer func() {
		if r := recover(); r != nil {
		val = 0
		 er:= fmt.Errorf("recover: %s", r)
		 err = e.E(er,"div failed",[]string{"math-err"},division,n1,n2)
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
		log.Println(errp.V())
		fmt.Println("something went wrong at the division")
	}

	if err == nil {
		fmt.Println("result is ",result)
	}
}
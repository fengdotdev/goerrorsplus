package main

import (
	"errors"
	"fmt"

	"github.com/fengdotdev/goerrorsplus/e"
)

func failDiv(n1, n2 int ) (float64,error) {

	if n2 == 0 {
		return 0, errors.New("div by zero")
	}
	return float64(n1)/float64(n2), nil
}


func main() {
	result, err := failDiv(1,0)
	if err != nil {
		errp := e.E(err,"div failed",[]string{"div"},failDiv,1,0)
		fmt.Println(errp)
		fmt.Println(err)
		fmt.Println(errp.V())
	}

	fmt.Println(result)
}
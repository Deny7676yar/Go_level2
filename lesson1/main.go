package main

import (
	"errors"
	"fmt"
	"log"
	"time"
	""


)

/*
Функция простая в ней содержится ошибка,
в цикле <= вместо просто <
 */
func main() {
	var (
		ErrComparison = errors.New("error in work with comparison")
	)
	defer func() {
		if r := recover();r !=nil {
			dt := time.Now()
			log.Println(fmt.Errorf("%w: %s", ErrComparison, r),dt)
		}
		planB()
	}()

	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77

	var total float64 = 0
	for i := 0; i <= 5; i++ {
		total += x[i]
	}
	fmt.Println(total / 5)



}

func planB() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(total / 5)

}


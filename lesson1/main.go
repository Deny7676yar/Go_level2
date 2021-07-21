package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// ErrorDate тип ошибки хранящий время возникновения
type ErrorDate struct {
	time.Time
}

// NewErrTime создаем экземпляр ошибки
func NewErrTime() *ErrorDate {
	return &ErrorDate{
		Time: time.Now(),
	}
}
//реализация фукции Error из итерфейса
func (e *ErrorDate) Error()string  {
	return fmt.Sprintf("an error occurred %s", e.Format("2006-01-02 15:04:05"))
}


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

			log.Println(fmt.Errorf("%w: %s", ErrComparison, NewErrTime()))
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
//фукция без неверого знака
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

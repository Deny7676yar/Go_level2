package main

import (
	"fmt"
	"log"
	"sync"
)

func main(){
	mm := [][]int {
		{1,2,3},
		{2,3,4},
		{5,6,7},
	}
	a := 8
	// wg используется для ожидания завершения программы.
	//WaitGroup - это счетный семафор, используем его для ведения записи о выполнении goroutines
	var wg = sync.WaitGroup{}
	log.Println ("Start goroutines")
	for i := range mm {
		for j := range mm[i] {
			//добавляем счетчик по одному на каждую горутину
			wg.Add(1)
			// обьявляем анонимную функцию и создаем горутину
			go func(i, j int) {
				//заплонируем вызов Done, что бы сообщить main, что мы закончили
				defer wg.Done()
				mm[i][j] *= a
			}(i,j)
		}
	}
	// Ждем завершения горутин.
	log.Println("Waiting To Finish")
	wg.Wait()


	for _, m := range mm{
		fmt.Println(m)
	}
	fmt.Println ("\nTerminating Program")

}

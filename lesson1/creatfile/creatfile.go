package creatfile

import (
	"errors"
	"fmt"
	"log"
	"os"
)


var (
	ErrNoname = errors.New("error name unknown")
	ErrCreate = errors.New("it is not possible to create a file ")
	ErrClose = errors.New("do not close the file")
)

func creatFile() {
	defer func(){
		if r := recover();r !=nil {
			log.Println(fmt.Errorf("%w: %s", ErrNoname, r))
		}
	}()

	file, err := os.Create("./lesson1/text.txt")
	if err != nil {
		log.Fatalf("%v: %s", ErrCreate, err.Error())
	}

	defer func() {
		errC := file.Close()
		if errC !=nil {
			log.Println(fmt.Errorf("%w: %s", ErrClose, errC.Error()))
		}
	}()


}
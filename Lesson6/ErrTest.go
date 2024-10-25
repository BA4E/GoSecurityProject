package main

import (
	"errors"
	"fmt"
)

var Err = MyError{404}

var err = errors.New("division by zero")

type MyError struct {
	i int
}

func (e MyError) Error() string {
	return fmt.Sprintf("MyError %d", e.i)
}

func division(x int, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("MyError2 %d", 403)
	}
	return x / y, nil
}

func main() {
	_, err := division(7, 0)
	if err != nil {
		fmt.Println(err)
	}

}

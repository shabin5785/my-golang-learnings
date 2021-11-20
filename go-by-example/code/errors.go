package gobyex

import (
	"fmt"
)

func Errors() {
	s, e := checkError(1)

	if e != nil {
		fmt.Printf("Error : %v", e)
	} else {
		fmt.Println(s)
	}
}

func checkError(n int) (string, error) {
	if n == 0 {
		return "", &customError{1, "Zero"}
	} else {
		return fmt.Sprintf("%d : ok", n), nil
	}
}

type customError struct {
	id      int
	message string
}

func (e *customError) Error() string {
	return fmt.Sprintf("%d - %s", e.id, e.message)
}

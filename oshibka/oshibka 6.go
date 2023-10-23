package main

import (
	"fmt"
	"log"
)

type CustError struct {
	code int
	err  string
}

func (a *CustError) Error() string {
	return fmt.Sprintf("code: %d: err %v", a.code, a.err)
}

func req(p int, q string) error {
	return &CustError{
		code: p,
		err:  q,
	}
}

func main() {
	err := req(666, "try")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("good")
}

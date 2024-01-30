package main

import (
	"fmt"
	"mini_flow/flow"
	"mini_flow/ierr"
)

func raiseError() {
	ierr.SystemError.Raise()
}

func main() {
	flow.Try(func() {
		raiseError()
	}).Catch(func(err interface{}) {
		fmt.Printf("%v\n", err)
	})
	fmt.Println("continue execution")
}

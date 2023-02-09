package main

import (
	// "errors"
	"fmt"
)

func mainStartWebServer(paramUsed int) (int, error) {
	fmt.Println("starting ...")

	fmt.Println("runing on", paramUsed)
	// return errors.New("Something went wrong", port)
	return paramUsed, nil

}

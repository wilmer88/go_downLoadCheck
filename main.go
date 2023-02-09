package main

import "fmt"

func main() {
	port := 3000
	// msg, err := mainStartWebServer(port)
	// fmt.Println(msg,err)
	_, err := mainStartWebServer(port)
	fmt.Println(err)

}
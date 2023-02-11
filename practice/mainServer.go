package main

import (
	"net/http"

	"github.com/wilmer88/go_downLoadCheck/controllers"
)

func mainServer() {

	//  creates http server
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}







// package main

// type HTTPRequest struct {
// 		Method string
// }

// func main() {
// 	r := HTTPRequest{Method: "GET"}

// 	switch r.Method {
// // implicit break cases use fallthrough if need to check next case below
// 	case "GET":
// 		println("GET request")
// 		//fallthrough
// 	case "POST":
// 		println("POST request")
// 	case "DELETE":
// 		println("DELETE request")
// 	case "PUT":
// 		println("PUT request")
// 	default:
// 		println("unhandled method")

// 	}

// }

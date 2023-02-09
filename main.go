package main

type HTTPRequest struct {
		Method string
}

func main() {
	r := HTTPRequest{Method: "GET"}

	switch r.Method {
// implicit break cases use fallthrough if need to check next case below
	case "GET":
		println("GET request")
		//fallthrough
	case "POST":
		println("POST request")
	case "DELETE":
		println("DELETE request")
	case "PUT":
		println("PUT request")
	default:
		println("unhandled method")

	}



}

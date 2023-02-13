package main

import (
	// "fmt"
	"net/http"
	

	"github.com/wilmer88/go_downLoadCheck/controllers"
)




func main() {

	controllers.RegisterControllers()
	
	http.ListenAndServe(":3000", nil)

}
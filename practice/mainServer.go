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
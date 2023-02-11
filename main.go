package main

import (
	"os"
	"net/http"
	"github.com/wilmer88/go_downLoadCheck/controllers"
)

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		handler.ServeHTTP(w, r)
	})
}

func main() {

	//  creates http server
	port:= os.Getenv("PORT")

if port == ""{

	port="3000"
}
	controllers.RegisterControllers()
	http.ListenAndServe(":" + port, nil)

}


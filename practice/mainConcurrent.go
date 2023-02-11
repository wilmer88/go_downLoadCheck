package main

import (
	"fmt"
	"math/rand"
	// "net/http"
	"time"

	// "github.com/wilmer88/go_downLoadCheck/controllers"
)
var cache = map[int]Person{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i:=0; i < 3; i++ {
		id := rnd.Intn(3)+1
		if p, ok := queryCache(id); ok {
			fmt.Println("from Cache")
			fmt.Println(p)
			continue
		}
		if p, ok := queryDatabase(id); ok {
			fmt.Println("from Database")
			fmt.Println(p)
			continue
		}
		fmt.Printf("Person not found with id: '%v'", id)
		time.Sleep(150 * time.Millisecond)
	}



		//  creates http server 
	// controllers.RegisterControllers()
	// http.ListenAndServe(":3000", nil)




}
func queryCache(id int) (Person, bool) {
	p, ok := cache[id]
	return p, ok
}

func queryDatabase(id int) (Person, bool) {
	time.Sleep(100* time.Millisecond)
	for _, p := range persons {
		if p.ID == id {
			cache[id] = p
			return p, true
		}
	}
	return Person{},false
}




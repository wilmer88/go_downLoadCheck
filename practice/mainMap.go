package main

import (
	"fmt"
)

const (
	userRoll = "worker"
	userId   = iota
)

func mainMap() {
	m := map[string]int{userRoll: userId}
	fmt.Println(m)
	fmt.Println(m[userRoll])
	m["worker"] = 32
	fmt.Println(m)

	fmt.Println("-----------------")

	sliceMap := []int{m["worker"]}
	sliceMap = append(sliceMap, 32, 65)
	fmt.Println(sliceMap)

	delete(m, userRoll)
	fmt.Println(m)

}

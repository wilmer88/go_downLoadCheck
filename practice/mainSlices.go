package main

import (
	"fmt"
)

func mainSlices() {
	arr := [3]int{1,2,3}
	slice := arr[:]
	arr[1] = 42 
	slice[2] =27
	fmt.Println(arr,slice)
	

	arrSlice := []int{1,2,3}
	fmt.Println(arrSlice)
	arrSlice = append(arrSlice, 9, 10,21)
	fmt.Println(arrSlice)
	s4 := arrSlice[2:]
	s5 := arrSlice[:4]
	s6 := arrSlice[3:6]
	fmt.Println(s4,s5,s6)

}
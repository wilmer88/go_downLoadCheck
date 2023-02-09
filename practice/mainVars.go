package main

import (
	"fmt"
)

func mainVars() {

	var i int32 = 21
	fmt.Println(i)

	complexVar1 := complex(3, 4)
	fmt.Println(complexVar1)

	r, im := real(complexVar1), imag(complexVar1)
	fmt.Println(r, im)

}

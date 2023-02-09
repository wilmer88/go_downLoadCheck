package main

import (
// "net/http"
"fmt"

// "github.com/wilmer88/go_downLoadCheck/controllers"

)

func main() {
	// loop terminates based on condition
	var i int
	for i < 5 {
		println("Ive ran this times:", i)
		i++
		if i == 4 {
		fmt.Println("stoped at 3")

			break
			// continue
		}

		fmt.Println("running....")

	}

	// loop with post clauses

	for pci := 0; pci < 5; pci++ {
		fmt.Println("pci/ post clause iterator ran this amount of time:", pci)
		continue

	}

	// ugly infinit loops
	var infiniI int
	for ; ; {

		if infiniI == 5 {
			break

		}
		fmt.Println("ugly will run till infinity but now only untill 4", infiniI)


		infiniI++
		continue


	}

	// purty infinit loops
		var purtyI int
		for {
	
			if purtyI == 5 {
				break
	
			}
			fmt.Println("pretty  will run till infinity but now only untill 4", purtyI)
	
			purtyI++
			
	
		}

	fmt.Println("------------------")
	fmt.Println("starting slice collection loops")


	// looping over collections
	sliceArr := []int{1,2,3}
	for sliI := 0; sliI < len(sliceArr); sliI++ {
		println(sliceArr[sliI])
	}
	
	  arrS2  := []int{1,2,3}
	for sliI2, v := range arrS2 {
		println("slice values:",sliI2, v)
	}

	go_downLoadCheckCollection  := map[string]int{"http": 80, "https": 443}
	// for _, v := range go_downLoadCheckCollection {
	for k, v := range go_downLoadCheckCollection {
		println("slice values:", k, v)
		println("slice values:", v)
	}

}

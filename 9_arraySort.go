package main

import (
	"fmt"
)

func main() {
	arrayList := [6]int{3, 12, 4, 5, 8, 9}

	tmpInt := 0
	for i := 0; i < len(arrayList); i++ {
		for j := i + 1; j < len(arrayList); j++ {
			if arrayList[j] < arrayList[i] {
				tmpInt = arrayList[j]
				arrayList[j] = arrayList[i]
				arrayList[i] = tmpInt
			}
		}
	}
	fmt.Println(arrayList)
}

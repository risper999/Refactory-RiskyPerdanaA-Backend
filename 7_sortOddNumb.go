package main

import "fmt"

func main() {
	arrayList := [8]int{9, 4, 2, 4, 1, 5, 3, 0}
	tmpInt := 0
	for i := 0; i < len(arrayList); i += 2 {
		for j := i + 2; j < len(arrayList)-1; j += 2 {
			if arrayList[j] < arrayList[i] {
				tmpInt = arrayList[j]
				arrayList[j] = arrayList[i]
				arrayList[i] = tmpInt
			}
		}
	}
	fmt.Println(arrayList)
}

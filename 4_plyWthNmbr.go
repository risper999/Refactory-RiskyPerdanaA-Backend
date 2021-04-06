package main

import "fmt"

func prmNmbExc(lenA int, even []int, prmNumb []int) {
	for i := 0; i < lenA; i++ {
		if even[i] == 0 || even[i] == 1 {
			continue
		} else if even[i] == 2 || even[i] == 3 {
			prmNumb = append(prmNumb, even[i])
		} else if even[i] > 3 && even[i] <= 9 {
			if (even[i]%2) == 1 && (even[i]%3) != 0 {
				prmNumb = append(prmNumb, even[i])
			}
		} else if even[i] >= 10 && even[i] <= 13 {
			if ((even[i] % 2) != 0) && ((even[i] % 3) != 0) && ((even[i] % 5) != 0) {
				prmNumb = append(prmNumb, even[i])
			}
		} else if even[i] > 14 && even[i] <= 23 {
			if ((even[i] % 2) != 0) && ((even[i] % 3) != 0) && ((even[i] % 5) != 0) && ((even[i] % 7) != 0) {
				prmNumb = append(prmNumb, even[i])
			}
		} else if even[i] > 23 && even[i] <= 29 {
			if ((even[i] % 2) != 0) && ((even[i] % 3) != 0) && ((even[i] % 5) != 0) && ((even[i] % 7) != 0) && ((even[i] % 11) != 0) && ((even[i] % 13) != 0) && ((even[i] % 17) != 0) && ((even[i] % 19) != 0) && ((even[i] % 23) != 0) {
				prmNumb = append(prmNumb, even[i])
			}
		} else if even[i] > 29 && even[i] <= 31 {
			if ((even[i] % 2) != 0) && ((even[i] % 3) != 0) && ((even[i] % 5) != 0) && ((even[i] % 7) != 0) && ((even[i] % 11) != 0) && ((even[i] % 13) != 0) && ((even[i] % 17) != 0) && ((even[i] % 19) != 0) && ((even[i] % 23) != 0) && ((even[i] % 29) != 0) {
				prmNumb = append(prmNumb, even[i])
			}
		} else if even[i] > 31 {
			if ((even[i] % 2) != 0) && ((even[i] % 3) != 0) && ((even[i] % 5) != 0) && ((even[i] % 7) != 0) && ((even[i] % 11) != 0) && ((even[i] % 13) != 0) && ((even[i] % 17) != 0) && ((even[i] % 19) != 0) && ((even[i] % 23) != 0) && ((even[i] % 29) != 0) && ((even[i] % 31) != 0) {
				prmNumb = append(prmNumb, even[i])
			}
		} else {
			continue
		}
	}
	fmt.Println(prmNumb)
}

func genNmb(min, max int) []int {
	even := make([]int, max-min+1)
	for i := range even {
		even[i] = min + i
	}
	return even
}

func main() {
	even := genNmb(0, 1000)
	fmt.Println("1) even")
	fmt.Println(even)

	l := len(even) / 2
	if float32(len(even)%2) != 0 {
		l += 1
	}

	var odd []int
	for i := 0; i < l; i++ {
		if i == 0 {
			odd = append(odd, even[i])
		} else {
			odd = append(odd, even[i+i])
		}
	}
	fmt.Println("2) odd")
	fmt.Println(odd)

	fmt.Println("3) numbers multiplies by 5")
	var mltpBy5 []int
	for i := 1; i < len(even); i++ {
		if even[i]%5 == 0 {
			mltpBy5 = append(mltpBy5, even[i])
		}
	}
	fmt.Println(mltpBy5)

	fmt.Println("4) prime numbers")
	var prmNumb []int
	prmNmbExc(len(even), even, prmNumb)
	fmt.Println("5) prime numbers less than 100")
	prmNmbExc(100, even, prmNumb)

}

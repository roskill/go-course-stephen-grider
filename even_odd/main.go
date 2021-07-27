package main

import "fmt"

func main() {
	s := []int{}
	for i := 0; i < 11; i++ {
		s = append(s, i)
	}
	// s := make([]int, 11)
	// for i := 0; i <= 10; i++ {
	// 	s[i] = i
	// }
	fmt.Println(s)
	for _, n := range s {
		if n%2 == 0 {
			fmt.Println(n, "is even")
		} else {
			fmt.Println(n, "is odd")
		}
	}
}

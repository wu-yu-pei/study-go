package utils

import "fmt"

func GetSum(a int, b int) (c int) {
	c = a + b
	fmt.Printf("c: %v\n", c)
	return c
}

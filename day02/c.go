package main

import (
	"fmt"
	"strings"
)

func main() {
	// 找出frientd中a开头的元素
	friends := [...]string{"abcde", "bbcdd", "aaddc"}
	friendsWidthA := []string{}

	// for循环
	for _, val := range friends {
		one := strings.Split(val, "")[0] == "a"
		if one {
			friendsWidthA = append(friendsWidthA, val)
		}
	}

	// 输出
	fmt.Printf("one: %v\n", friendsWidthA)
}

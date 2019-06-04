package main

import (
	"fmt"
	"strconv"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	ans := ""
	n, d := 0, 0 //divisor 除数
	fmt.Scanln(&n, &d)
	//dividend := n  // 被除数
	m := make(map[int]int)

	t := int(n / d)
	ans += strconv.Itoa(t) + "."
	n = n % d

	if n == 0 {
		ans += "0"
	}
	n = n * 10
	for i := 1; n > 0; n, i = n*10, i+1 {
		if m[n] != 0 {
			str := ans[len(ans)-i+m[n]:]
			ans = ans[:len(ans)-i+m[n]] + "(" + str + ")"
			break
		}
		m[n] = i
		t = int(n / d)
		n = n % d
		//fmt.Println(strconv.Itoa(t), ans)
		ans += strconv.Itoa(t)

	}
	for j := 0; j < len(ans); j += 76 {
		fmt.Println(ans[j:min(j+76, len(ans))])
	}

}

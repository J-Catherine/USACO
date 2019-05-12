package main

import "fmt"

var (
	n, c, num int
)

func matrixMulti(m1, m2 [8][8]int) [8][8]int {
	m := [8][8]int{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			for k := 0; k < 8; k++ {
				if m1[i][k]*m2[k][j] != 0 {
					m[i][j] = 1
					break
				}
			}
		}
	}
	return m
}

func matrixPow(m [8][8]int, cnt int) [8][8]int {
	res := [8][8]int{}
	for i := 0; i < 8; i++ {
		res[i][i] = 1
	}
	for cnt > 0 {
		if cnt&1 > 0 {
			res = matrixMulti(res, m)
			//fmt.Println(res)
		}
		m = matrixMulti(m, m)
		cnt >>= 1
	}
	return res
}

func changeState(a int) int {
	if a == 1 {
		return 0
	} else {
		return 1
	}
}

func operation(cur [6]int, num int) [6]int {
	switch num {
	case 1:
		for i := 0; i < 6; i++ {
			cur[i] = changeState(cur[i])
		}
	case 2:
		for i := 0; i < 6; i += 2 {
			cur[i] = changeState(cur[i])
		}
	case 3:
		for i := 1; i < 6; i += 2 {
			cur[i] = changeState(cur[i])
		}
	case 4:
		for i := 0; i < 6; i += 3 {
			cur[i] = changeState(cur[i])
		}
	}
	return cur
}

func main() {

	_, _ = fmt.Scanln(&n)
	_, _ = fmt.Scanln(&c)
	finalOn := make([]int, 6)
	finalOff := make([]int, 6)
	for _, _ = fmt.Scan(&num); num != -1; _, _ = fmt.Scan(&num) {
		t := (num - 1) % 6
		finalOn[t] = 1
	}

	for _, _ = fmt.Scan(&num); num != -1; _, _ = fmt.Scan(&num) {
		t := (num - 1) % 6
		if finalOn[t] == 1 {
			fmt.Println("IMPOSSIBLE")
			return
		} else {
			finalOff[t] = 1
		}
	}

	maskOn := 0
	for i := 0; i != 6; i++ {
		maskOn = maskOn*2 + finalOn[i]
	}
	maskOff := 0
	for i := 0; i != 6; i++ {
		maskOff = maskOff*2 + finalOff[i]
	}
	//fmt.Println(maskOn,maskOff)

	vals := [8]int{0, 14, 21, 27, 36, 42, 49, 63}
	str := [8]string{
		"000000",
		"001110",
		"010101",
		"011011",
		"100100",
		"101010",
		"110001",
		"111111",
	}
	matrix := [8][8]int{}
	cur := [8][6]int{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 0, 1},
		{0, 1, 1, 0, 1, 1},
		{1, 0, 0, 1, 0, 0},
		{1, 0, 1, 0, 1, 0},
		{1, 1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1},
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			for k := 1; k < 5; k++ {
				if operation(cur[i], k) == cur[j] {
					matrix[i][j] = 1
				}
			}
		}

	}
	//fmt.Println(matrix)
	start := [8]int{0, 0, 0, 0, 0, 0, 0, 1}
	mp := matrixPow(matrix, c)
	res := [8]int{}
	for i := 0; i < 8; i++ {
		for k := 0; k < 8; k++ {
			if start[k]*mp[k][i] != 0 {
				res[i] = 1
			}
		}
	}
	//fmt.Println(res)

	flag := 0
	for i := 0; i < 8; i++ {
		ans := ""
		if res[i] == 1 && (vals[i]&maskOn == maskOn) && (vals[i]&maskOff == 0) {
			flag = 1
			for j := 0; j < n; j += 6 {
				if n-j < 6 {
					ans += (str[i])[0 : n-j]
				} else {
					ans += str[i]
				}
			}
			fmt.Println(ans)
		}
	}
	if flag == 0 {
		fmt.Println("IMPOSSIBLE")
	}
}

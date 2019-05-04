package main

import "fmt"

//func exchange(arr *[]int, i, j int) int {
//	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
//	return 1
//}

func mini(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	n := 0
	ans := 0
	fmt.Scanln(&n)
	arr := make([]int, n)
	num := make([]int, 4)
	for i := 0; i < n; i++ {
		fmt.Scanln(&arr[i])
		num[arr[i]] += 1
	}
	num[2] += num[1]
	num[3] += num[2]
	cnt := [4][4]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}

	for i := 0; i < n; i++ {
		if i < num[1] {
			if arr[i] != 1 {
				cnt[1][arr[i]] += 1
			}
		} else if i < num[2] {
			if arr[i] != 2 {
				cnt[2][arr[i]] += 1
			}
		} else if i < num[3] {
			if arr[i] != 3 {
				cnt[3][arr[i]] += 1
			}
		}
	}
	fmt.Println(cnt)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3 && j != i; j++ {
			if cnt[i][j] != 0 && cnt[j][i] != 0 {
				mi := mini(cnt[i][j], cnt[j][i])
				cnt[i][j] -= mi
				cnt[j][i] -= mi
				ans += mi
			}
		}
	}
	ans += (cnt[1][2] + cnt[1][3]) * 2
	fmt.Println(ans)
}

// 交换
//l, k, r := 0, 0, n-1
//cnt := 0
//for ; l < r && arr[l] == 1; l++ {}
//for ; l < r && arr[r] == 3; r-- {}
//fmt.Println(l, k, r)
//for k = l + 1; l <= k && k <= r; {
//	if arr[k] == 2 {
//		k += 1
//	} else if arr[k] == 3 {
//		cnt += exchange(&arr, k, r)
//		for r -= 1; arr[r] == 3; r-- {
//		}
//	} else if arr[k] == 1 {
//		cnt += exchange(&arr, k, l)
//		for l += 1; arr[l] == 1; l++ {
//		}
//	}
//}
//fmt.Println(cnt, arr)

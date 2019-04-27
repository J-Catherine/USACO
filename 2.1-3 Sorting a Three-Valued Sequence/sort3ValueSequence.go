package main

import "fmt"

func exchange(arr *[]int, i, j int) int {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	return 1
}

func main() {
	n := 0
	fmt.Scanln(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&arr[i])
	}
	l, k, r := 0, 0, n-1
	cnt := 0
	for ; l < r && arr[l] == 1; l++ {
	}
	for ; l < r && arr[r] == 3; r-- {
	}
	fmt.Println(l, k, r)
	for k = l + 1; l <= k && k <= r; {
		if arr[k] == 2 {
			k += 1
		} else if arr[k] == 3 {
			cnt += exchange(&arr, k, r)
			for r -= 1; arr[r] == 3; r-- {
			}
		} else if arr[k] == 1 {
			cnt += exchange(&arr, k, l)
			for l += 1; arr[l] == 1; l++ {
			}
		}
	}
	fmt.Println(cnt, arr)

}

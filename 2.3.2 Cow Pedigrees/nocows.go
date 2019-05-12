package main

import "fmt"

func main() {
	n, k := 0, 0
	_, _ = fmt.Scanf("%d %d\n", &n, &k)
	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, k+1)
	}
	for i := 1; i <= k; i++ {
		f[1][i] = 1
	}

	//F(n,k) = sigema{F(i,k-1)*F(n-1-i,k-1)}
	for h := 1; h <= k; h++ {
		for p := 3; p <= n; p += 2 {
			for l := 1; l < p; l += 2 {
				f[p][h] += f[l][h-1] * f[p-1-l][h-1]
				//fmt.Println(p, h, f[p][h])
			}
			f[p][h] %= 9901
		}
	}

	fmt.Println((f[n][k] - f[n][k-1] + 9901) % 9901)
}

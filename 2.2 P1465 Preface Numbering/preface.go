package main

import "fmt"

func main() {
	n := 0
	_, err := fmt.Scanln(&n)
	if err != nil {
		println("Scan n false!")
		return
	}
	s := [4][10]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "M", "MM", "MMM"}}
	alphabet := "IVXLCDM"
	cnt := make([][][]int, 4)
	for i := 0; i < 4; i++ {
		cnt[i] = make([][]int, 10)
		for j := 0; j < 10; j++ {
			cnt[i][j] = make([]int, 300)
			for k := 0; k < len(s[i][j]); k++ {
				cnt[i][j][uint(s[i][j][k])] += 1
			}
		}
	}
	res := make([]int, 300)
	for i := 1; i <= n; i++ {
		for j, tmp := 0, i; tmp > 0; j, tmp = j+1, tmp/10 {
			for k := 0; k < 7; k++ {
				ci := uint((alphabet[k]))
				res[ci] += cnt[j][tmp%10][ci]
			}
		}
	}
	for i := 0; i < 7; i++ {
		ci := uint((alphabet[i]))
		if res[ci] > 0 {
			fmt.Printf("%c %d\n", alphabet[i], res[ci])
		}
	}
}

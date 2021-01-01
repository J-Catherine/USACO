package main

import "fmt"

var (
	v, g                   int
	scoop                  [][]int
	min, res               []int
	ansTypeNum, typeNum    []int
	ansTypeCnt, ansTypeSum int
)

func dfs(scoopType int, add bool, typeCnt int) {
	if scoopType >= g {
		i := 0
		for ; i < v && res[i] >= min[i]; i++ {
		}
		if i == v && ansTypeCnt >= typeCnt {
			typeSum := 0
			for j := 0; j < g; j++ {
				if typeNum[j] == 1 {
					typeSum += j + 1
				}
			}
			//fmt.Println(ansTypeCnt, typeCnt, ansTypeSum,typeSum, ansTypeNum, res)
			if ansTypeCnt == typeCnt && ansTypeSum < typeSum {
				return
			}
			ansTypeSum = typeSum
			ansTypeCnt = typeCnt
			for j := 0; j < g; j++ {
				ansTypeNum[j] = typeNum[j]
			}
		}
		return
	}
	if add {
		for i := 0; i < v; i++ {
			res[i] += scoop[scoopType][i]
		}
		typeNum[scoopType] = 1
	}
	//fmt.Println(scoopType, typeNum, res)
	dfs(scoopType+1, false, typeCnt)
	dfs(scoopType+1, true, typeCnt+1)
	if scoopType+1 < g {
		for i := 0; i < v; i++ {
			res[i] -= scoop[scoopType+1][i]
		}
		typeNum[scoopType+1] = 0
	}

}

func main() {
	_, _ = fmt.Scanln(&v)
	min = make([]int, v)
	for i := 0; i < v; i++ {
		_, _ = fmt.Scan(&min[i])
		//if err != nil{
		//	fmt.Println("scan type fault!")
		//	return
		//}
	}
	_, _ = fmt.Scanln(&g)
	//if err != nil{
	//	fmt.Println("scan G fault!")
	//	return
	//}
	scoop = make([][]int, g)
	for i := 0; i < g; i++ {
		scoop[i] = make([]int, v)
		for j := 0; j < v; j++ {
			_, _ = fmt.Scan(&scoop[i][j])
		}
	}
	res = make([]int, v)
	ansTypeNum, typeNum = make([]int, g), make([]int, g)
	ansTypeCnt, ansTypeSum = g, (g+1)*g
	dfs(0, false, 0)
	dfs(0, true, 1)

	fmt.Printf("%d", ansTypeCnt)
	for i := 0; i < g; i++ {
		if ansTypeNum[i] == 1 {
			fmt.Printf(" %d", i+1)
		}
	}
	fmt.Printf("\n")
}

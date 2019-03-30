package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	ans int
	t,state [][]int
)

func max(x,y int) int {
	if x>y{
		return x
	}
	return y
}

func main()  {
	var inputReader *bufio.Reader
	inputReader = bufio.NewReader(os.Stdin)
	str,_ := inputReader.ReadString('\n')
	r,_ := strconv.Atoi(str[:len(str)-1])
	t=make([][]int,r)
	state = make([][]int,r)
	for i:=0; i<r; i++{
		t[i]=make([]int,i+1)
		state[i] = make([]int, i+1)


		str,_ := inputReader.ReadString('\n')
		ints := strings.Split(str[:len(str)-1], " ")
		for j:=0; j<=i; j++{
			t[i][j], _ = strconv.Atoi(ints[j])
		}

	}

	state[0][0]=t[0][0]
	for i :=1; i<r; i++{
		for j:=0; j<=i; j++{
			if j == 0{
				state[i][j] = state[i-1][j] + t[i][j]
			}else {
				if j == i {
					state[i][j] = state[i-1][j-1] + t[i][j]
				}else {
					state[i][j]= max(state[i-1][j-1],state[i-1][j]) + t[i][j]
				}
			}


		}
	}
	ans = state[r-1][0]
	for _,x :=range state[r-1]{
		ans = max(x,ans)
	}
	fmt.Println(ans)

}

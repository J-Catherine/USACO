package P3864

import "fmt"

var (
	num,i,length,sum,cnt int
	m map[uint8] int
	s string
)
//思路：
//第一种（简单）：把名字化为数字。
//第二种（困难）：把数字化为名字。

func main()  {
	fmt.Scanln(&num)
	m = map[uint8]int{
		'A':2, 'B':2, 'C':2,
		'D':3, 'E':3, 'F':3,
		'G':4, 'H':4, 'I':4,
		'J':5, 'K':5, 'L':5,
		'M':6, 'N':6, 'O':6,
		'P':7, 'R':7, 'S':7,
		'T':8, 'U':8, 'V':8,
		'W':9, 'X':9, 'Y':9,
	}
	for i =0;i<4617;i++{
		fmt.Scanln(&s)
		sum = 0
		for _,c:= range(s){
			sum *= 10
			sum += m[uint8(c)]
		}
		if sum == num{
			fmt.Println(s)
			cnt += 1
		}
	}
	if cnt == 0{
		fmt.Println("NONE")
	}

}

package main

import "fmt"

var (
	N,B,D int
	end uint
	res []uint

)

func main()  {
	fmt.Scanln(&N, &B, &D)
	end = 1<<uint(B) - 1
	res = make([]uint,N)
	res[0] = uint(0)
	ptr := 1
	for i:=uint(1); i<=end; i++{
		flag := true
		for j:=0; j< ptr; j++{
			if distance(i,res[j]) <D{
				flag = false
				break
			}
		}
		if flag {
			res[ptr]=i
			ptr += 1
		}
		if ptr == N{
			break
		}
	}
	for i:=0; i<ptr; i++{
		fmt.Print(res[i])
		if ((i+1)%10)==0||i==(ptr-1) {
			fmt.Println()
		} else {
			fmt.Printf(" ")
		}

	}
}


func distance(a,b uint) int{
	x := a^b
	x=(x & 0x55555555)+((x>>1)& 0x55555555)
	x=(x & 0x33333333)+((x>>2)& 0x33333333)
	x=(x & 0x0F0F0F0F)+((x>>4)& 0x0F0F0F0F)
	x=(x & 0x00FF00FF)+((x>>8)& 0x00FF00FF)
	x=(x & 0x0000FFFF)+((x>>16)& 0x0000FFFF)
	return int(x)
}
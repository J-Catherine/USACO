package main

import (
	"fmt"
)

type Money struct {
	in, out, own int
}

var (
	num, own, recNum int
	giver, receiver  string
)
//var m map[string]Money

func main() {
	fmt.Scanln(&num) //人数
	monies := make(map[string]*Money)
	names := make([]string, num)

	for i := 0; i < num; i++ {
		fmt.Scanln(&names[i])
		monies[names[i]] = &Money{0, 0, 0}
	}
	for i := 0; i < num; i++ {
		fmt.Scanln(&giver)
		fmt.Scanln(&own, &recNum)

		if recNum == 0 {
			continue
		}
		recVal := own / recNum
		for j := 0; j < recNum; j++ {

			fmt.Scanln(&receiver)
			monies[receiver].in += recVal
		}
		monies[giver].out = recVal * recNum
	}
	for _, name := range names {
		fmt.Println(name, monies[name].own-monies[name].out+monies[name].in)
	}
}

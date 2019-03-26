package main

import "fmt"

var i,years,days int
var weeks, months []int

func main(){
	fmt.Scanln(&years)

	// init days, weeks, months
	weeks = make([]int, 7)
		months = make([]int, 12)
		days = 13 - 31
		for i = 0; i <12; i++ {
			if i == 3 || i == 5 || i == 8 || i == 10 {
				months[i] = 30
			} else{
				if i == 1 {
					months[i] = 28
				} else {
					months[i] = 31
				}
			}
	}

	for i = 0; i < years; i++ {
		for month := range months{
			if month == 2 && i % 4 == 0{
				if (i + 1900) % 100 == 0 { //世纪年
					if (i+1900)%400 == 0 { //世纪年只有mod 400 = 0 才为闰年
						days += 1
					}
				}else{
					days += 1
				}
			}
			days += months[(month-1 + 12)%12]
			weeks[days % 7] += 1
		}
	}


	for i=0; i <7; i++ {
		fmt.Print(weeks[(i+6)%7], " ")
	}
}
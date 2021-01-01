package main

import "fmt"

var (
	dx                     = [4]int{-1, 0, 1, 0} //上、右、下、左四个方向
	dy                     = [4]int{0, 1, 0, -1}
	cx, cy, cd, fx, fy, fd = -1, -1, 0, -1, -1, 0
	visited                = [12][12][4][12][12][4]bool{} // default is false
	mmap                   = make([][]bool, 12)
)

func dfs(t int) {
	if fx == cx && fy == cy {
		fmt.Println(t)
		return
	}
	if visited[cx][cy][cd][fx][fy][fd] == true {
		fmt.Println(0)
		return
	}
	visited[cx][cy][cd][fx][fy][fd] = true
	// cow move
	if mmap[cx+dx[cd]][cy+dy[cd]] == true {
		cx, cy = cx+dx[cd], cy+dy[cd]
	} else {
		cd = (cd + 1) % 4
	}

	// farmer move
	//fmt.Println(fx+dx[fd],fy+dy[fd],fd)
	if mmap[fx+dx[fd]][fy+dy[fd]] == true {
		fx, fy = fx+dx[fd], fy+dy[fd]
	} else {
		fd = (fd + 1) % 4
	}
	dfs(t + 1)
}

func main() {
	s := ""
	mmap[0] = make([]bool, 12)
	mmap[11] = make([]bool, 12)
	for i := 1; i <= 10; i++ {
		mmap[i] = make([]bool, 12)
		fmt.Scanln(&s)
		for j := 1; j <= 10; j++ {
			if s[j-1] == '.' {
				mmap[i][j] = true
			}
			if s[j-1] == 'C' {
				cx, cy = i, j
				mmap[i][j] = true
			}
			if s[j-1] == 'F' {
				fx, fy = i, j
				mmap[i][j] = true
			}
		}
	}

	//fmt.Println(cx,cy,fx,fy)
	//fmt.Println(dx,dy,mmap)
	dfs(0)
}

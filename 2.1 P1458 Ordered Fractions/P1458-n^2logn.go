package main

import "fmt"

type Fraction struct {
	numerator, denominator int  // 分子 分母
}

type Heap struct {
	data []Fraction
	size int
}

func gcd(a, b int) int {
	d := a % b
	if d == 1 || d == 0{
		return d
	} else {
		return gcd(b,d)
	}
}

func (h Heap) Less(i,j int) bool {
	return h.data[i].numerator * h.data[j].denominator < h.data[i].denominator * h.data[j].numerator
}


func (h *Heap) down(i int) {
	for j := i*2+1 ; j < h.size; i, j = j, j*2+1 {
		if j+1 < h.size && h.Less(j+1,j) {
			j+= 1
		}
		if h.Less(j,i) {
			h.data[i], h.data[j] = h.data[j], h.data[i]
		}
	}
}

func buildHeap(size int) *Heap {
	heap := &Heap{make([]Fraction,size), size}
	for i:=0; i<size; i++{
		heap.data[i].numerator = 1
		heap.data[i].denominator = i+1
	}
	for i:=size/2-1; i>=0; i--{
		heap.down(i)
	}
	//fmt.Println(heap.data)
	return heap
}

func (h *Heap) popAndAdd()  {
	now := h.data[0]
	fmt.Printf("%d/%d\n",now.numerator,now.denominator)
	i:=now.numerator+1
	for ; i<now.denominator; i++ {
		if gcd(now.denominator, i) == 1 {
			h.data[0].numerator = i
			h.down(0)
			return
		}
	}
	if i >= now.denominator{
		h.size -= 1
		h.data[0] = h.data[h.size]
		h.down(0)
		return
	}
}

func main()  {
	n:=5000
	fmt.Scanln(&n)
	heap := buildHeap(n)
	fmt.Println("0/1")
	for ; heap.size > 0; {
		heap.popAndAdd()
	}
}
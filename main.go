package main

import "fmt"

func reverse(values []int, start int, stop int) []int {
	n := stop - start
	count := 0
	for i := start; i < start+n/2; i++ {
		j := stop - 1 - count
		values[i], values[j] = values[j], values[i]
		count++
	}
	return values
}

func permute(values []int) []int {
	n := len(values)
	pivot := n - 2
	last := false
	for ; pivot >= 0; pivot-- {
		if values[pivot] < values[pivot+1] {
			last = true
			break
		}
	}
	if pivot < 0 {
		pivot = 0
	}
	if !last {
		values = reverse(values, 0, n)
		return values
	}
	for j := n - 1; j >= pivot; j-- {
		if values[pivot] < values[j] {
			values[pivot], values[j] = values[j], values[pivot]
			values = reverse(values, pivot+1, n)
			break
		}
	}
	return values
}
func main() {
	x := []int{4, 3, 2, 1}
	for i := 0; i < 25; i++ {
		fmt.Println(permute(x))
	}
}

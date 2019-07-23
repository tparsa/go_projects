package main

import "fmt"

const MAXN = 1000000

func merge(left, right []int) []int {
	var idx = []int {0, 0, 0}
	ret := make([]int, len(left) + len(right))
	for idx[0] < len(left) && idx[1] < len(right){
		if left[idx[0]] < right[idx[1]]{
			ret[idx[2]] = left[idx[0]]
			idx[0]++
			idx[2]++
		} else {
			ret[idx[2]] = right[idx[1]]
			idx[1]++
			idx[2]++
		}
	}
	for idx[1] < len(right) {
		ret[idx[2]] = right[idx[1]]
		idx[1]++
		idx[2]++
	}
	for idx[0] < len(left){
		ret[idx[2]] = left[idx[0]]
		idx[0]++
		idx[2]++
	}
	return ret
}


func mergeSort(arr []int) []int {
	array_len := len(arr)
	if array_len == 1 {
		return arr
	} else {
		middle := int(array_len / 2)
		var (
			left = make([]int, middle)
			right = make([]int, array_len - middle)
		)
		left = mergeSort(arr[:middle])
		right = mergeSort(arr[middle:])
		return merge(left, right)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	a = mergeSort(a)
	for _, v := range a {
		fmt.Println(v)
	}
}

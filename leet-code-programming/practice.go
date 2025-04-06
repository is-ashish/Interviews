package main

import "fmt"

func main_() {
	// array should be sorted
	arr1 := []int{2, 6, 7, 8, 9, 12, 43, 56, 78, 89}
	arr2 := []int{1, 3, 4, 5, 8}
	mergeTwoArraySorted2(arr1, arr2)
}

// with temp array
func mergeTwoArraySorted1(arr1 []int, arr2 []int) {
	res := []int{}
	j := 0
	k := 0
	for i := 0; j < len(arr1) && k < len(arr2); i++ {
		if arr1[j] <= arr2[k] {
			res = append(res, arr1[j])
			j++
		}
		if arr1[j] > arr2[k] {
			res = append(res, arr2[k])
			k++
		}
	}
	for a := j; a < len(arr1); a++ {
		res = append(res, arr1[a])
	}
	for b := k; b < len(arr2); b++ {
		res = append(res, arr2[b])
	}
	fmt.Println("Final sorted array---->>>", res)
}

// without taking extra array ------------ Pemding
func mergeTwoArraySorted2(arr1 []int, arr2 []int) {
	res := []int{}
	j := 0
	k := 0
	for i := 0; j < len(arr1) && k < len(arr2); i++ {
		if arr1[j] <= arr2[k] {
			res = append(res, arr1[j])
			j++
		}
		if arr1[j] > arr2[k] {
			res = append(res, arr2[k])
			k++
		}
	}
	for a := j; a < len(arr1); a++ {
		res = append(res, arr1[a])
	}
	for b := k; b < len(arr2); b++ {
		res = append(res, arr2[b])
	}
	fmt.Println("Final sorted array---->>>", res)
}

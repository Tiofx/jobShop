package util

import (
	"fmt"
	"github.com/fighterlyt/permutation"
)

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MinMax(arr []int) (min, max int) {
	min = arr[0]
	max = arr[0]

	for _, v := range arr {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}

	return
}

func MaxOf(arr []int) (max int) {
	max = arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	return
}

func MakeArrayOf(element, size int) []int {
	array := make([]int, size)

	for i := range array {
		array[i] = element
	}

	return array
}

func Combination(upBound int) <-chan []int {
	c := make(chan []int)
	initArr := make([]int, upBound)

	for i := 0; i < upBound; i++ {
		initArr[i] = i
	}

	go func(consumer chan<- []int) {
		defer close(consumer)

		p, err := permutation.NewPerm(initArr, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		for container, err := p.Next(); err == nil; container, err = p.Next() {
			consumer <- container.([]int)
		}

		return
	}(c)

	return c
}

func CompareIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func FillIntsWith(a []int, value int) {
	for i := range a {
		a[i] = value
	}
}

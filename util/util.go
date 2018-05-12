package util

import (
	"fmt"
	"github.com/fighterlyt/permutation"
)

func Max(x, y uint64) uint64 {
	if x > y {
		return x
	}
	return y
}

func Min(x, y uint64) uint64 {
	if x < y {
		return x
	}
	return y
}

func MaxInt64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func MinInt64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func MinMax(arr []uint64) (min, max uint64) {
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

func MaxOf(arr []uint64) (max uint64) {
	max = arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	return
}

func MakeArrayOf(element, size uint64) []uint64 {
	array := make([]uint64, size)

	for i := range array {
		array[i] = element
	}

	return array
}

func Combination(upBound uint64) <-chan []uint64 {
	c := make(chan []uint64)
	initArr := make([]uint64, upBound)

	for i := uint64(0); i < upBound; i++ {
		initArr[i] = i
	}

	go func(consumer chan<- []uint64) {
		defer close(consumer)

		p, err := permutation.NewPerm(initArr, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		for container, err := p.Next(); err == nil; container, err = p.Next() {
			consumer <- container.([]uint64)
		}

		return
	}(c)

	return c
}

func CompareIntSlices(a, b []uint64) bool {
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

func FillIntsWith(a []int64, value int64) {
	for i := range a {
		a[i] = value
	}
}

func FillUintsWith(a []uint64, value uint64) {
	for i := range a {
		a[i] = value
	}
}

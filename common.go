package main

import (
	"math/rand"
)

// --------- start common -------------

func randomElement[T any](slice []T) T {
	ri := rand.Intn(len(slice))
	return slice[ri]
}

func removeElement[T comparable](slice []T, rem T) []T {
	res := []T{}
	for i := range slice {
		if slice[i] != rem {
			res = append(res, slice[i])
		}
	}
	return res
}

func filter[T any](slice []T, pred func(T) bool) []T {
	res := []T{}

	for _, item := range slice {
		if pred(item) {
			res = append(res, item)
		}
	}

	return res
}

type Int interface {
	~int8 | ~uint8 | ~int16 | ~uint16 | ~int | ~uint
}

func maxValueFromMap[K comparable, V Int](m map[K]V) (K, V) {
	var maxKey K
	var maxValue V
	firstRun := true

	for k, v := range m {
		if firstRun {
			maxKey = k
			maxValue = v
			firstRun = false
		} else if v > maxValue {
			maxKey = k
			maxValue = v
		}
	}

	return maxKey, maxValue
}

// ----------- end common ------------

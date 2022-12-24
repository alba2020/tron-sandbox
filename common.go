package main

import "math/rand"

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

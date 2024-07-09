package main

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	res := make([]T2, len(s))

	for i, v := range s {
		res[i] = f(v)
	}

	return res
}

func Filter[T any](s []T, f func(T) bool) []T {
	var res []T

	for _, v := range s {
		if f(v) {
			res = append(res, v)
		}
	}

	return res
}

func Reduce[T1, T2 any](s []T1, base T2, f func(T2, T1) T2) T2 {
	res := base

	for _, v := range s {
		res = f(res, v)
	}

	return res
}

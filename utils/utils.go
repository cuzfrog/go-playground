package utils

import (
	"golang.org/x/exp/constraints"
	"math/rand"
)

func Min[T constraints.Ordered](l, r T) (m T) {
	if l < r {
		m = l
	} else {
		m = r
	}
	return
}

func Max[T constraints.Ordered](l, r T) (m T) {
	if l > r {
		m = l
	} else {
		m = r
	}
	return
}

func Suffle[T any](a []T) {
	n := len(a)
	var j int
	for i := range a {
		j = rand.Intn(n)
		a[i], a[j] = a[j], a[i]
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const lettersLength = int64(len(letterBytes))

func RandAlphabet(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%lettersLength]
	}
	return string(b)
}

func Abs[T constraints.Signed](n T) T {
	if n >= 0 {
		return n
	}
	return -n
}

func Hash(h int, m int) int {
	h = 31*h + 17
	h = 31 * h
	return Abs(h) % m
}

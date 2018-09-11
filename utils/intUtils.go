package utils

import "math/rand"

func IntMin(l, r int) (m int) {
	if l < r {
		m = l
	} else {
		m = r
	}
	return
}

func IntMax(l, r int) (m int) {
	if l > r {
		m = l
	} else {
		m = r
	}
	return
}

func Suffle(a []int) {
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
		b[i] = letterBytes[rand.Int63() % lettersLength]
	}
	return string(b)
}
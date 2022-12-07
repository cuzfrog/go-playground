package utils

import (
	"golang.org/x/exp/constraints"
	"math/rand"
	"strconv"
	"strings"
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

func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func IsUpperCaseLetter(c uint8) bool {
	return c >= 65 && c <= 90
}

func SplitString2(s, sep string) (string, string) {
	x := strings.Split(s, sep)
	return x[0], x[1]
}

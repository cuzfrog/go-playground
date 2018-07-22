package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13reader rot13Reader) Read(buf []byte) (int, error) {
	cnt, err := r13reader.r.Read(buf)
	for i := 0; i < cnt; i++ {
		c := &buf[i]
		if (*c >= 65 && *c <= 77) || (*c >= 97 && *c <= 109) {
			*c += 13
		} else if (*c >= 78 && *c <= 90) || (*c >= 110 && *c <= 122) {
			*c -= 13
		}
	}
	return cnt, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

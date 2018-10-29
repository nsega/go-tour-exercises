package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (n int, e error) {
	n, e = rot.r.Read(b)
	if e == nil {
		for i, v := range b {
			switch {
			case v >= 'A' && v <= 'Z':
				b[i] = (v-'A'+13)%26 + 'A'
			case v >= 'a' && v <= 'z':
				b[i] = (v-'a'+13)%26 + 'a'
			}
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case 65 <= b && b <= 90:
		if b += 13; 90 < b {
			b -= 26
		}
	case 97 <= b && b <= 122:
		if b += 13; 122 < b {
			b -= 26
		}
	}
	return b
}

func (rot rot13Reader) Read(buf []byte) (int, error) {

	i, err := rot.r.Read(buf)
	if err != nil {
		return 0, err
	}

	for j := 0; j < i; j++ {
		buf[j] = rot13(buf[j])
	}

	return len(buf), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

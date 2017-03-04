package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	const UPPER = 0
	const LOWER = 1
	const PUNCT = 2
	toConvert := os.Args[1]
	last := [3]int{0, 0, 0}
	idx := UPPER
	out := ""
	for _, r := range toConvert {
		c := int(r)
		f := 7
		switch {
		case c >= 'a' && c <= 'z':
			if idx == UPPER {
				out += ">>>"
			} else if idx == PUNCT {
				out += "<"
			} else {
				out += ">"
			}
			idx = LOWER
		case c >= 'A' && c <= 'Z':
			if idx == LOWER {
				out += "<"
			} else if idx == PUNCT {
				out += "<<<"
			} else {
				out += ">"
			}
			idx = UPPER
		default:
			if idx == UPPER {
				out += ">>>>>"
			} else if idx == LOWER {
				out += ">>>"
			} else {
				out += ">"
			}
			idx = PUNCT
		}
		for ; f > 0; f-- {
			r := int(math.Abs(float64(last[idx]-c))) % f
			if r == 0 {
				break
			}
			if r < 4 {
				out += "<"
				if last[idx] < c {
					for ; r > 0; r-- {
						out += "+"
					}
				} else {
					for ; r > 0; r-- {
						out += "-"
					}
				}
				out += ">"
				break
			}
		}

		for i := f; i > 0; i-- {
			out += "+"
		}
		out += "[<"

		l := last[idx]

		switch {
		case l < c:
			for i := (c - l) / f; i > 0; i-- {
				out += "+"
			}
		case l > c:
			for i := (l - c) / f; i > 0; i-- {
				out += "-"
			}
		}

		last[idx] = c
		out += ">-]<."
	}

	for i, c := range out {
		fmt.Printf("%c", c)
		if i%100 == 99 {
			fmt.Println()
		}
	}
}

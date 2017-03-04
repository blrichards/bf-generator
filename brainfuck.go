package main

import (
	"fmt"
	"math"
	"os"
)

// UPPER uppercase letters counter index
const UPPER = 0

// LOWER lowercase letters counter index
const LOWER = 1

// PUNCT punctuation counter index
const PUNCT = 2

// OUTPUT brainfuck code to be outputed
var OUTPUT = ""

func getIdx(idx, c int) int {
	switch {
	case c >= 'a' && c <= 'z':
		if idx == UPPER {
			OUTPUT += ">>>"
		} else if idx == PUNCT {
			OUTPUT += "<"
		} else {
			OUTPUT += ">"
		}
		return LOWER
	case c >= 'A' && c <= 'Z':
		if idx == LOWER {
			OUTPUT += "<"
		} else if idx == PUNCT {
			OUTPUT += "<<<"
		} else {
			OUTPUT += ">"
		}
		return UPPER
	default:
		if idx == UPPER {
			OUTPUT += ">>>>>"
		} else if idx == LOWER {
			OUTPUT += ">>>"
		} else {
			OUTPUT += ">"
		}
		return PUNCT
	}
}

func main() {
	toConvert := os.Args[1]
	last := [3]int{0, 0, 0}
	idx := UPPER

	for _, r := range toConvert {
		c := int(r)
		f := 7

		idx = getIdx(idx, c)

		if int(math.Abs(float64(last[idx]-c))) <= 10 {
			OUTPUT += "<"
			for last[idx] < c {
				OUTPUT += "+"
				last[idx]++
			}
			for last[idx] > c {
				OUTPUT += "-"
				last[idx]--
			}
			OUTPUT += "."
		} else {
			for ; f > 0; f-- {
				r := int(math.Abs(float64(last[idx]-c))) % f
				if r == 0 {
					break
				}
				if r < 4 {
					OUTPUT += "<"
					if last[idx] < c {
						for ; r > 0; r-- {
							OUTPUT += "+"
						}
					} else {
						for ; r > 0; r-- {
							OUTPUT += "-"
						}
					}
					OUTPUT += ">"
					break
				}
			}

			for i := f; i > 0; i-- {
				OUTPUT += "+"
			}
			OUTPUT += "[<"

			l := last[idx]

			switch {
			case l < c:
				for i := (c - l) / f; i > 0; i-- {
					OUTPUT += "+"
				}
			case l > c:
				for i := (l - c) / f; i > 0; i-- {
					OUTPUT += "-"
				}
			}

			last[idx] = c
			OUTPUT += ">-]<."
		}
	}

	for i, c := range OUTPUT {
		fmt.Printf("%c", c)
		if i%60 == 59 {
			fmt.Println()
		}
	}
}

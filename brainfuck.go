package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	toConvert := os.Args[1]
	last := 0
	for _, r := range toConvert {
		c := int(r)
		f := 30
		for ; f > 0; f-- {
			if int(math.Abs(float64(last-c)))%f == 0 {
				break
			}
		}

		fmt.Print("> ")
		for i := f; i > 0; i-- {
			fmt.Print("+")
		}
		fmt.Print("\n[ < ")

		switch {
		case last < c:
			for i := (c - last) / f; i > 0; i-- {
				fmt.Print("+")
			}
		case last > c:
			for i := (last - c) / f; i > 0; i-- {
				fmt.Print("-")
			}
		}

		fmt.Println(" > - ]\n< .")
		last = c
	}
}

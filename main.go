package main

import (
	"log"
	"os"
	"strconv"
)

func Add(a, c int64) int64 {
	return a + c
}

func main() {
	if len(os.Args) != 3 {
		panic("provide 2 number arguments please")
	}

	var (
		aStr = os.Args[1]
		bStr = os.Args[2]
	)

	a, err := strconv.ParseInt(aStr, 10, 64)
	if err != nil {
		panic(err)
	}

	b, err := strconv.ParseInt(bStr, 10, 64)
	if err != nil {
		panic(err)
	}

	log.Printf("%d + %d is %d", a, b, Add(a, b))
}

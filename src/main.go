package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	param := os.Args[1]
	number, err := strconv.Atoi(param)
	if err != nil {
		return
	}
	runThroughNumbers(number)
}

func runThroughNumbers(input int){
	current := 1000
	for current > 0 {
		ok, m, c, d, u := calcDecimals(current)
		if ok {
			fmt.Println(current)
		}
	
		current = nextNumber(input, m, c, d, u)
	}
}

func nextNumber(input, m, c, d, u int) int {
	if u < input {
		u++
		return getNext(m,c, d, u)
	} else {
		u = 0
	}

	if d < input {
		d++
		return getNext(m,c, d, u)
	} else {
		d = 0
	}

	if c < input {
		c++
		return getNext(m,c, d, u)
	} else {
		c = 0
	}

	if m < input {
		m++
		return getNext(m,c, d, u)
	} else {
		return -1
	}
}

func getNext(m, c, d, u int) int {
	return m * 1000 + c * 100 + d * 10 + u
}

func calcDecimals(currentNumber int) (bool, int, int, int, int) {
	m := (currentNumber - (currentNumber % 1000)) / 1000
	c := ((currentNumber % 1000) - ((currentNumber % 1000) % 100)) / 100
	d := (((currentNumber % 1000) % 100) - ((currentNumber % 1000) % 100) % 10) / 10
	u := currentNumber - m * 1000 - c * 100 - d * 10

	if m + c + d + u == 21 {
		return true, m, c, d, u
	}
	return false, m, c, d, u
}
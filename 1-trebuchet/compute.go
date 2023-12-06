package trebuchet

import (
	"strconv"
	"sync"
)

func traverseRight(line string, c chan<- int) {
	ptr := 0
	lineLength := len(line)
	for {
		if ptr >= lineLength {
			c <- 0
			close(c)
			break
		}
		digit, err := strconv.Atoi(line[ptr : ptr+1])
		if err == nil {
			c <- digit
			close(c)
			break
		}
		ptr++
	}
}

func traverseLeft(line string, c chan<- int) {
	ptr := len(line) - 1
	for {
		if ptr < 0 {
			c <- 0
			close(c)
			break
		}
		digit, err := strconv.Atoi(line[ptr : ptr+1])
		if err == nil {
			c <- digit
			close(c)
			break
		}
		ptr--
	}
}

func compute(line string, c chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	leftDigitChan := make(chan int)
	rightDigitChan := make(chan int)
	go traverseRight(line, leftDigitChan)
	go traverseLeft(line, rightDigitChan)
	leftDigit := <-leftDigitChan
	rightDigit := <-rightDigitChan
	num := leftDigit*10 + rightDigit
	c <- num
}

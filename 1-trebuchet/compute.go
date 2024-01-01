package trebuchet

import (
	"strconv"
	"strings"
	"sync"
)

type NumberInLineMeta struct {
	pos   int
	digit int
}

func traverseRight(line string, c chan<- NumberInLineMeta, wg *sync.WaitGroup) {
	defer wg.Done()
	ptr := 0
	lineLength := len(line)
	for {
		if ptr >= lineLength {
			c <- NumberInLineMeta{
				digit: 0,
				pos:   ptr,
			}
			break
		}
		digit, err := strconv.Atoi(line[ptr : ptr+1])
		if err == nil {
			c <- NumberInLineMeta{
				digit: digit,
				pos:   ptr,
			}
			break
		}
		ptr++
	}
}

func traverseLeft(line string, c chan<- NumberInLineMeta, wg *sync.WaitGroup) {
	defer wg.Done()
	ptr := len(line) - 1
	for {
		if ptr < 0 {
			c <- NumberInLineMeta{
				digit: 0,
				pos:   ptr,
			}
			break
		}
		digit, err := strconv.Atoi(line[ptr : ptr+1])
		if err == nil {
			c <- NumberInLineMeta{
				digit: digit,
				pos:   ptr,
			}
			break
		}
		ptr--
	}
}

var numberWords map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getNumberWords() map[string]int {
	return numberWords
}

func traverseLeftWord(line string, c chan<- NumberInLineMeta, wg *sync.WaitGroup) {
	defer wg.Done()
	posMax := -1
	digit := 0
	for k, v := range getNumberWords() {
		if pos := strings.LastIndex(line, k); pos >= 0 {
			if pos > posMax {
				posMax = pos
				digit = v
			}
		}
	}
	c <- NumberInLineMeta{
		pos:   posMax,
		digit: digit,
	}
}

func traverseRightWord(line string, c chan<- NumberInLineMeta, wg *sync.WaitGroup) {
	defer wg.Done()
	posMin := len(line)
	digit := 0
	for k, v := range getNumberWords() {
		if pos := strings.Index(line, k); pos >= 0 {
			if pos < posMin {
				posMin = pos
				digit = v
			}
		}
	}
	c <- NumberInLineMeta{
		pos:   posMin,
		digit: digit,
	}
}

func compute(line string, c chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	leftDigitChan := make(chan NumberInLineMeta)
	rightDigitChan := make(chan NumberInLineMeta)
	var wgInner sync.WaitGroup
	wgInner.Add(4)
	go traverseRight(line, leftDigitChan, &wgInner)
	go traverseLeft(line, rightDigitChan, &wgInner)
	go traverseRightWord(line, leftDigitChan, &wgInner)
	go traverseLeftWord(line, rightDigitChan, &wgInner)
	go func() {
		wgInner.Wait()
		close(leftDigitChan)
		close(rightDigitChan)
	}()
	leftDigit1, leftDigit2 := <-leftDigitChan, <-leftDigitChan
	rightDigit1, rightDigit2 := <-rightDigitChan, <-rightDigitChan
	var leftDigit int
	var rightDigit int
	if leftDigit1.pos < leftDigit2.pos {
		leftDigit = leftDigit1.digit
	} else {
		leftDigit = leftDigit2.digit
	}

	if rightDigit1.pos > rightDigit2.pos {
		rightDigit = rightDigit1.digit
	} else {
		rightDigit = rightDigit2.digit
	}

	num := leftDigit*10 + rightDigit

	c <- num
}

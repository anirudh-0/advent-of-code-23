package trebuchet

import (
	"bufio"
	"fmt"
	"sync"
)

func Solve() {

	reader, err := fetchInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(reader)
	var wg sync.WaitGroup
	c := make(chan int)
	for scanner.Scan() {
		wg.Add(1)
		go compute(scanner.Text(), c, &wg)
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	sum := 0
	for numberFound := range c {
		sum += numberFound
	}
	fmt.Println(sum)
}

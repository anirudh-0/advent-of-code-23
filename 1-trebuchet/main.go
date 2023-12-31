package trebuchet

import (
	"bufio"
	"fmt"
	"sync"

	adventUtils "github.com/anirudh-0/advent-of-code-23/advent-utils"
)

func Solve() {

	reader, err := adventUtils.FetchInput("./1-trebuchet/lines.txt")
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

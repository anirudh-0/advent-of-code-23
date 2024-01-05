package cubeconundrum

import (
	"bufio"
	"fmt"
	"sync"

	adventUtils "github.com/anirudh-0/advent-of-code-23/advent-utils"
)

func Solve() {
	reader, err := adventUtils.FetchInput("./2-cube-conundrum/games.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(reader)
	var wg sync.WaitGroup
	c := make(chan GameMeta)
	for scanner.Scan() {
		wg.Add(1)
		go compute(scanner.Text(), c, &wg)
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	sum := 0
	sumOfPowers := 0
	for gameMeta := range c {
		sum += gameMeta.gameID
		sumOfPowers += gameMeta.red * gameMeta.blue * gameMeta.green
	}
	fmt.Println(sum)
	fmt.Println(sumOfPowers)
}

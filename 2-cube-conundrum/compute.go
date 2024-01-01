package cubeconundrum

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

var colourCountMap map[string]int = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type GameMeta struct {
	DrawMeta
	power  int
	gameID int
}

type DrawMeta struct {
	red        int
	blue       int
	green      int
	isPossible bool
}

func compute(line string, c chan<- GameMeta, wg *sync.WaitGroup) {
	defer wg.Done()
	gameData := strings.Split(line, ":")
	gameID, _ := strconv.Atoi(strings.TrimSpace(strings.Split(gameData[0], " ")[1]))
	cubeDraws := strings.Split(gameData[1], "; ")
	drawMetaChan := make(chan DrawMeta)
	var wgInner sync.WaitGroup
	for _, draw := range cubeDraws {
		wgInner.Add(1)
		go processDraw(draw, drawMetaChan, &wgInner)
	}

	go func() {
		wgInner.Wait()
		close(drawMetaChan)
	}()

	gameMeta := GameMeta{
		gameID: gameID,
	}
	isOverallPossible := true
	for drawMeta := range drawMetaChan {
		isOverallPossible = drawMeta.isPossible && isOverallPossible
		gameMeta.red = int(math.Max(float64(drawMeta.red), float64(gameMeta.red)))
		gameMeta.blue = int(math.Max(float64(drawMeta.blue), float64(gameMeta.blue)))
		gameMeta.green = int(math.Max(float64(drawMeta.green), float64(gameMeta.green)))
	}

	if !isOverallPossible {
		gameMeta.gameID = 0
	}
	c <- gameMeta
}

func processDraw(draw string, isPossibleChan chan DrawMeta, wg *sync.WaitGroup) {
	defer wg.Done()
	colourCounts := strings.Split(draw, ", ")
	var wgInner sync.WaitGroup
	for _, v := range colourCounts {
		wgInner.Add(1)
		go processColourCount(strings.TrimSpace(v), isPossibleChan, &wgInner)
	}
	wgInner.Wait()
}

func processColourCount(v string, c chan<- DrawMeta, wg *sync.WaitGroup) {
	defer wg.Done()
	colourCount := strings.Split(v, " ")
	count, _ := strconv.Atoi(strings.TrimSpace(colourCount[0]))
	colour := strings.TrimSpace(colourCount[1])
	fmt.Println(v, colourCount, count, colour)
	if count > colourCountMap[colour] {
		c <- false
	} else {
		c <- true
	}
}

package main

import (
	"errors"
	"fmt"
	"os"

	trebuchet "github.com/anirudh-0/advent-of-code-23/1-trebuchet"
	cubeConundrum "github.com/anirudh-0/advent-of-code-23/2-cube-conundrum"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("%v\n", errors.New("invalid cli args - needs a cli arg of advent day #"))
		return
	}
	switch os.Args[1] {
	case "1":
		trebuchet.Solve()
	case "2":
		cubeConundrum.Solve()
	default:
		panic("nothing to run")
	}

}

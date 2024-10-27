package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/charmbracelet/log"
)

type Treebuilder struct {
	Treesize int `arg:"" help:"width of tree"`
}

func (t *Treebuilder) Run() error {
	err := buildtree(t.Treesize)
	if err != nil {
		log.Error("buildtree failed: ", "err", err)
	}
	return nil
}

func buildtree(treesize int) error {
	var windowsize float64 = float64(treesize)
	starttree := 0
	for i := 1; i <= treesize; i++ {
		var space float64 = (windowsize - float64(i)) / 2
		if math.Mod(space, 1) != 0 && i-1 != 0 {
			// create stump analogous to top of tree
			if starttree == 0 {
				starttree = i
			}
			fmt.Print(strings.Repeat(" ", int(space)) + strings.Repeat("*", i-1) + "\n")
		}
	}
	// stump
	fmt.Print(strings.Repeat(" ", (int(windowsize)-starttree)/2) + strings.Repeat("*", starttree-1) + "\n")
	return nil
}

var cli struct {
	Debug       bool        `help:"activate debug loggin" `
	Treebuilder Treebuilder `cmd:"" help:"print an ASCII tree"`
}

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run()
	if err != nil {
		fmt.Println("Error!")
	}

}

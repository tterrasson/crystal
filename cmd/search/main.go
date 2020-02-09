package main

import (
	"flag"
	"fmt"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/tterrasson/crystal/pkg/rule"
	"github.com/tterrasson/crystal/pkg/world"
)

func main() {
	fmt.Println("Starting...")
	iterationArg := flag.Int("iteration", 64, "Number of iteration")
	stateArg := flag.Int("state", 3, "Number of states")
	fillSeedArg := flag.Float64("fillseed", 0.25, "Chance to fill seed")
	fillRuleArg := flag.Float64("fillrule", 0.25, "Chance to fill rule")
	worldSizeArg := flag.Int("worldsize", 100, "World size (SxSxS)")
	outputPathArg := flag.String("output", "explore", "Output path")
	nbArg := flag.Int("nb", 8, "Number of random CA to generate")

	flag.Parse()

	rand.Seed(time.Now().Unix())

	ruleset := rule.NewRandomRuleSet(float32(*fillRuleArg), *stateArg)
	world := world.NewWorld(*worldSizeArg, ruleset)
	world.RandomSeed(float32(*fillSeedArg), *stateArg)

	for n := 0; n < *nbArg; n++ {
		for i := 0; i < *iterationArg; i++ {
			fmt.Printf("Step %d ...\n", i)
			world.Iterate()
		}

		fname := fmt.Sprintf("output-%f-%f.obj", rand.Float64())
		fpath := filepath.Join(*outputPathArg, fname)
		fmt.Printf("Exporting to %s ...\n", fpath)
		world.ExportToFile(fpath)
	}
}

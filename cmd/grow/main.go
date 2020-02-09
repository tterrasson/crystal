package main

import (
	"flag"
	"fmt"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/tterrasson/crystal/pkg/obj"
	"github.com/tterrasson/crystal/pkg/rule"
	"github.com/tterrasson/crystal/pkg/world"
)

func main() {
	fmt.Println("Starting...")
	iterationArg := flag.Int("iteration", 128, "Number of iteration")
	randomSeedArg := flag.Bool("randseed", false, "Use a random seed")
	inputArg := flag.String("input", "output.obj", "Input OBJ file")
	fillSeedArg := flag.Float64("fillseed", 0.25, "Chance to fill seed")
	outputLastOnlyArg := flag.Bool("lastonly", false, "Output only last iteration")
	worldSizeArg := flag.Int("worldsize", 100, "World size (SxSxS)")
	outputPathArg := flag.String("output", "explore", "Output path")
	seedOffsetArg := flag.Int("seedoffset", 0, "Seed offset")

	flag.Parse()

	rand.Seed(time.Now().Unix())

	reader := obj.Reader{*inputArg}
	rules, maxStates := reader.ExtractRuleSet()
	ruleset := rule.Set{rules, maxStates}
	world := world.NewWorld(*worldSizeArg, &ruleset)

	if *randomSeedArg {
		world.RandomSeed(float32(*fillSeedArg), maxStates)
	} else {
		seed := reader.ExtractSeed(*worldSizeArg, *seedOffsetArg)
		world.SetSeed(seed)
	}

	for i := 0; i < *iterationArg; i++ {
		if !*outputLastOnlyArg {
			fname := fmt.Sprintf("output-%06d.obj", i)
			fpath := filepath.Join(*outputPathArg, fname)
			fmt.Printf("Exporting to %s ...\n", fpath)
			world.ExportToFile(fpath)
		}

		fmt.Printf("Step %d ...\n", i)
		world.Iterate()
	}

	if *outputLastOnlyArg {
		fname := fmt.Sprintf("explore/output-%06d.obj", *iterationArg)
		fmt.Printf("Exporting to %s ...\n", fname)
		world.ExportToFile(fname)
	}
}

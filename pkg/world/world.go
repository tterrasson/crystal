package world

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/shomali11/util/xconditions"
	"github.com/tterrasson/crystal/pkg/obj"
	"github.com/tterrasson/crystal/pkg/rule"
)

// World store the voxels
type World struct {
	Size        int
	Voxels      [][][]int
	ActiveCells uint
	Set         *rule.Set
	seed        [][][]int
}

// NewWorld create a new world
func NewWorld(size int, Set *rule.Set) *World {
	world := new(World)
	world.Size = size
	world.Set = Set
	world.ActiveCells = 0
	world.Voxels = make([][][]int, size)

	for x := range world.Voxels {
		world.Voxels[x] = make([][]int, size)
		for y := range world.Voxels[x] {
			world.Voxels[x][y] = make([]int, size)
		}
	}

	return world
}

func (world *World) copyVoxels() [][][]int {
	newVoxels := make([][][]int, world.Size)

	for x := 0; x < world.Size; x++ {
		newVoxels[x] = make([][]int, world.Size)
		copy(newVoxels[x], world.Voxels[x])

		for y := 0; y < world.Size; y++ {
			newVoxels[x][y] = make([]int, world.Size)
			copy(newVoxels[x][y], world.Voxels[x][y])
		}
	}

	return newVoxels
}

// SaveSeed save the current seed
func (world *World) SaveSeed() {
	world.seed = world.copyVoxels()
}

// SeedToJSON export the seed in a compact JSON view
func (world *World) SeedToJSON() ([]byte, error) {
	result := make(map[string]int)

	for x := 0; x < world.Size; x++ {
		for y := 0; y < world.Size; y++ {
			for z := 0; z < world.Size; z++ {
				state := world.seed[x][y][z]
				if state == 0 {
					continue
				}

				h := fmt.Sprintf("%d,%d,%d", x, y, z)
				result[h] = state
			}
		}
	}

	return json.Marshal(result)
}

// RandomSeed generate a random seed
func (world *World) RandomSeed(chance float32, maxStates int) {
	start := int(world.Size / 2)
	end := start + 6

	for x := start; x < end; x++ {
		for y := start; y < end; y++ {
			for z := start; z < end; z++ {
				if rand.Float32() < chance {
					world.Voxels[x][y][z] = rand.Intn(maxStates) + 1
				}
			}
		}
	}

	world.SaveSeed()
}

// RandomSymSeed generate a random symetric seed
func (world *World) RandomSymSeed(maxStates int) {
	pattern := [][]int{
		{2, 0, 0, 0, 1},
		{3, 2, 1, 1, 3},
		{3, 0, 0, 0, 3},
		{3, 1, 1, 2, 3},
		{1, 0, 0, 0, 2},
	}

	i := 0
	a := 0

	for x := 50; x < 55; x++ {
		for y := 50; y < 55; y++ {
			a = 0
			for z := 50; z < 55; z++ {
				world.Voxels[x][y][z] = pattern[i][a]
				a++
			}
		}

		i++
	}

	/* for x := start; x < end; x++ {
		for y := start; y < end; y++ {
			i := 0
			for z := start; z < end; z++ {


			}
		}
	} */

	world.SaveSeed()
}

// SetSeed set a new seed
func (world *World) SetSeed(seed [][][]int) {
	world.Voxels = seed
	world.SaveSeed()
}

// Iterate the world
func (world *World) Iterate() {
	newVoxels := world.copyVoxels()

	for x := 0; x < world.Size; x++ {
		for y := 0; y < world.Size; y++ {
			for z := 0; z < world.Size; z++ {
				state := world.Voxels[x][y][z]

				// FIXME
				if state > 0 {
					continue
				}

				faces, edges, corners := world.checkNeighbors(x, y, z)
				if faces == 0 || faces == -1 {
					continue
				}

				newState := world.Set.Process(state, faces, edges, corners)

				if newState != state {
					newVoxels[x][y][z] = newState

					if newState > 0 {
						world.ActiveCells++
					} else {
						world.ActiveCells--
					}
				}
			}
		}
	}

	world.Voxels = newVoxels
}

func (world *World) checkNeighbors(x int, y int, z int) (int, int, int) {
	var faces, edges, corners int

	// Checks boundaries
	if x+1 >= world.Size || x-1 < 0 {
		return -1, -1, -1
	}
	if y+1 >= world.Size || y-1 < 0 {
		return -1, -1, -1
	}
	if z+1 >= world.Size || z-1 < 0 {
		return -1, -1, -1
	}

	faces = 0
	faces += xconditions.IfThenElse(world.Voxels[x][y][z-1] > 0, 1, 0).(int)
	faces += xconditions.IfThenElse(world.Voxels[x][y][z+1] > 0, 1, 0).(int)
	faces += xconditions.IfThenElse(world.Voxels[x][y-1][z] > 0, 1, 0).(int)
	faces += xconditions.IfThenElse(world.Voxels[x][y+1][z] > 0, 1, 0).(int)
	faces += xconditions.IfThenElse(world.Voxels[x+1][y][z] > 0, 1, 0).(int)
	faces += xconditions.IfThenElse(world.Voxels[x-1][y][z] > 0, 1, 0).(int)

	edges = 0
	edges += xconditions.IfThenElse(world.Voxels[x][y-1][z-1] > 0, 1, 0).(int)
	edges += xconditions.IfThenElse(world.Voxels[x-1][y][z-1] > 0, 1, 0).(int)
	edges += xconditions.IfThenElse(world.Voxels[x+1][y][z-1] > 0, 1, 0).(int)
	edges += xconditions.IfThenElse(world.Voxels[x][y+1][z-1] > 0, 1, 0).(int)
	edges += xconditions.IfThenElse(world.Voxels[x-1][y-1][z] > 0, 1, 0).(int)
	edges += xconditions.IfThenElse(world.Voxels[x+1][y-1][z] > 0, 1, 0).(int)
	edges += xconditions.IfThenElse(world.Voxels[x-1][y+1][z] > 0, 1, 0).(int)
	edges += xconditions.IfThenElse(world.Voxels[x+1][y+1][z] > 0, 1, 0).(int)
	edges += xconditions.IfThenElse(world.Voxels[x][y-1][z+1] > 0, 1, 0).(int)
	edges += xconditions.IfThenElse(world.Voxels[x-1][y][z+1] > 0, 1, 0).(int)
	edges += xconditions.IfThenElse(world.Voxels[x+1][y][z+1] > 0, 1, 0).(int)
	edges += xconditions.IfThenElse(world.Voxels[x][y+1][z+1] > 0, 1, 0).(int)

	corners = 0
	corners += xconditions.IfThenElse(world.Voxels[x-1][y-1][z-1] > 0, 1, 0).(int)
	corners += xconditions.IfThenElse(world.Voxels[x+1][y-1][z-1] > 0, 1, 0).(int)
	corners += xconditions.IfThenElse(world.Voxels[x-1][y+1][z-1] > 0, 1, 0).(int)
	corners += xconditions.IfThenElse(world.Voxels[x+1][y+1][z-1] > 0, 1, 0).(int)
	corners += xconditions.IfThenElse(world.Voxels[x-1][y-1][z+1] > 0, 1, 0).(int)
	corners += xconditions.IfThenElse(world.Voxels[x+1][y-1][z+1] > 0, 1, 0).(int)
	corners += xconditions.IfThenElse(world.Voxels[x-1][y+1][z+1] > 0, 1, 0).(int)
	corners += xconditions.IfThenElse(world.Voxels[x+1][y+1][z+1] > 0, 1, 0).(int)

	return faces, edges, corners
}

// ExportToFile as obj file
func (world *World) ExportToFile(filename string) {
	writer := obj.NewWriter()

	ruleBytes, err := json.Marshal(world.Set)
	if err != nil {
		log.Fatal(err)
	}

	seedBytes, err := world.SeedToJSON()
	if err != nil {
		log.Fatal(err)
	}

	writer.AddComment("State : " + strconv.Itoa(world.Set.MaxStates))
	writer.AddComment("Rule : " + string(ruleBytes))
	writer.AddComment("Seed : " + string(seedBytes))
	writer.AddMtlLib("colors.mtl")

	for x := 0; x < world.Size; x++ {
		for y := 0; y < world.Size; y++ {
			for z := 0; z < world.Size; z++ {
				state := world.Voxels[x][y][z]
				if state == 0 {
					continue
				}

				writer.AddCube(x, y, z, state)
			}
		}
	}

	writer.WriteToFile(filename)
}

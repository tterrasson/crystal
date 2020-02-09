package obj

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Reader struct {
	Filename string
}

func (reader *Reader) initSeed(size int) [][][]int {
	seed := make([][][]int, size)

	for x := range seed {
		seed[x] = make([][]int, size)
		for y := range seed[x] {
			seed[x][y] = make([]int, size)
		}
	}

	return seed
}

func (reader *Reader) initRuleSet(maxStates int) [][][][]int {
	states := make([][][][]int, maxStates+1)

	for state := 0; state <= maxStates; state++ {
		states[state] = make([][][]int, 7)

		for face := 0; face < 7; face++ {
			states[state][face] = make([][]int, 13)

			for edge := 0; edge < 13; edge++ {
				states[state][face][edge] = make([]int, 9)
			}
		}
	}

	return states
}

// ExtractSeed extract the seed from obj file
func (reader *Reader) ExtractSeed(size int, offset int) [][][]int {
	file, err := os.Open(reader.Filename)
	if err != nil {
		log.Panicf("failed reading file: %s", err)
	}
	defer file.Close()

	out := reader.initSeed(size)
	rd := bufio.NewReader(file)

	for {
		line, err := rd.ReadString('\n')

		if strings.HasPrefix(line, "# Seed : ") {
			jseed := strings.Replace(line, "# Seed : ", "", 1)
			var seed map[string]interface{}
			json.Unmarshal([]byte(jseed), &seed)

			for k, v := range seed {
				idx := strings.Split(k, ",")
				x, _ := strconv.Atoi(idx[0])
				y, _ := strconv.Atoi(idx[1])
				z, _ := strconv.Atoi(idx[2])
				out[x+offset][y+offset][z+offset] = int(v.(float64))
			}

			return out
		}

		if err == io.EOF {
			break
		}
	}

	return nil
}

// ExtractRuleSet extract rule from obj file
func (reader Reader) ExtractRuleSet() ([][][][]int, int) {
	file, err := os.Open(reader.Filename)
	if err != nil {
		log.Panicf("failed reading file: %s", err)
	}
	defer file.Close()

	rd := bufio.NewReader(file)

	maxStates := 0

	for {
		line, err := rd.ReadString('\n')

		if strings.HasPrefix(line, "# State : ") {
			sub := strings.Replace(line, "# State : ", "", 1)
			sub = strings.TrimSuffix(sub, "\n")
			maxStates, _ = strconv.Atoi(sub)
		} else if strings.HasPrefix(line, "# Rule : ") {
			out := reader.initRuleSet(maxStates)

			jrule := strings.Replace(line, "# Rule : ", "", 1)
			var rule map[string]interface{}
			json.Unmarshal([]byte(jrule), &rule)
			stateSet := rule["states"].([]interface{})

			for state := 0; state <= maxStates; state++ {
				faceSet := stateSet[state].([]interface{})

				for face := 0; face < 7; face++ {
					edgeSet := faceSet[face].([]interface{})

					for edge := 0; edge < 13; edge++ {
						cornerSet := edgeSet[edge].([]interface{})

						for corner := 0; corner < 9; corner++ {
							out[state][face][edge][corner] = int(cornerSet[corner].(float64))
						}
					}
				}
			}

			return out, maxStates
		}

		if err == io.EOF {
			break
		}
	}

	return nil, 0
}

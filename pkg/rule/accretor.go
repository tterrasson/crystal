package rule

import (
	"math/rand"
)

// RuleSet structure
type RuleSet struct {
	States [][][][]int `json:"states"`
}

// NewRandomRuleSet generate a new random rule set
func NewRandomRuleSet(chance float32, maxStates int) *RuleSet {
	ruleset := new(RuleSet)
	ruleset.States = make([][][][]int, maxStates+1)

	for state := 0; state <= maxStates; state++ {
		ruleset.States[state] = make([][][]int, 7)

		for face := 0; face < 7; face++ {
			ruleset.States[state][face] = make([][]int, 13)

			for edge := 0; edge < 13; edge++ {
				ruleset.States[state][face][edge] = make([]int, 9)

				for corner := 0; corner < 9; corner++ {
					if rand.Float32() < chance {
						ruleset.States[state][face][edge][corner] = rand.Intn(maxStates) + 1
					} else {
						ruleset.States[state][face][edge][corner] = 0
					}
				}
			}
		}
	}

	return ruleset
}

// Process the ruleset
func (ruleset *RuleSet) Process(state int, faces int, edges int, corners int) int {
	return ruleset.States[state][faces][edges][corners]
}

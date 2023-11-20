package mutation

import (
	"math"
	"math/rand"
)

const pMutate = 0.05

var mutations = []Attribute{
	{
		Name:        "Lonely Cell",
		Probability: 0.2,
	},
	{
		Name:        "Friendly Cell",
		Probability: 0.07,
	},
	{
		Name:        "Pregnant Cell",
		Probability: 0.02,
	},
}

type Cell struct {
	State    string
	Mutation Attribute
}

type Attribute struct {
	Name        string
	Probability float64
	Fn          func()
}

func CanMutate(c Cell, mutantParent int) bool {
	return c.State == "ALIVE" && c.Mutation.Name == "" && rand.Intn(100) <= int(pModifier(mutantParent)*pMutate*100)
}

func FindMutation(c *Cell, mutantParent int) Attribute {
	for _, m := range mutations {
		p := rand.Intn(100)
		if p < int(pModifier(mutantParent)*m.Probability*100) {
			c.Mutation = m
			break
		}
	}
	return Attribute{}
}

func pModifier(p int) float64 {
	return math.Pow(2, float64(p))
}

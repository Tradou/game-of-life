package mutation

import (
	"math/rand"
)

const pMutate = 0.05

var mutations = []Attribute{
	{
		Name:        "Lonely Cell",
		Probability: 0.1,
	},
	{
		Name:        "Friendly Cell",
		Probability: 0.07,
	},
}

type Cell struct {
	State    string
	Mutation Attribute
}

type Attribute struct {
	Name        string
	Probability float32
	Fn          func()
}

func CanMutate(c Cell) bool {
	return c.State == "ALIVE" && c.Mutation.Name == "" && rand.Intn(100) <= pMutate*100
}

func FindMutation(c *Cell) Attribute {
	for _, m := range mutations {
		p := rand.Intn(100)
		if p < int(m.Probability*100) {
			c.Mutation = m
			break
		}
	}
	return Attribute{}
}

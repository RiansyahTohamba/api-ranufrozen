package drink

import "fmt"

type cli struct {
	drepo drinkRepository
}

func NewCli(drepo drinkRepository) *cli {
	return &cli{drepo}
}

func (cl cli) Show(id string) Drink {
	drink, err := cl.drepo.findOne(id)
	if err != nil {
		fmt.Println(err)
		return Drink{}
	}
	return drink
}

func (cl cli) List() []Drink {
	drinks, err := cl.drepo.findAll()
	if err != nil {
		fmt.Println(err)
		return []Drink{}
	}
	return drinks
}

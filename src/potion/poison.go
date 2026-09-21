package potion

import "fmt"

func NouvellePotionDePoison(nom string, degats int) Potion {
	return Potion{
		Nom:         nom,
		Description: fmt.Sprintf("Inflige %d points de dégâts par tour", degats),
		Effet:       degats,
		Type:        "poison",
	}
}

var PotionsDePoisonDisponibles = []Potion{
	NouvellePotionDePoison("Poison à explosion différée", 90),
	NouvellePotionDePoison("Poison à stacks", 75),
	NouvellePotionDePoison("Poison paralysant", 50),
	NouvellePotionDePoison("Poison affaiblissant", 35),
	NouvellePotionDePoison("Poison à dégâts sur la durée", 10),
}

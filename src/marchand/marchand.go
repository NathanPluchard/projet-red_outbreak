package marchand

import (
	"fmt"
)

type ShopItem struct {
	Name string
	Cost int
}

func merchant(c *Classe) {
	items := []ShopItem{
		{"Kit de chirurgie de fortune", 40},
		{"Fibre de carbone", 30}, 
		{"Acier", 15},
		{"Bois", 10},
		{"Tissu", 5},
		{"Plume", 2},
		{"Poche de perfusion", 20},
		{"Bandage militaire", 10},
		{"Pansement compressif", 7},
		{"Poche d'adrélanine", 5},
		{"Poison à explosion différée", 30 },
		{"Poison à stacks", 20},
		{"Poison paralysant", 15}, 
		{"Poison affaiblissant", 10},
		{"Poison à dégâts sur la durée", 5},
	}

	for {
		fmt.Println("----Le Troqueur----")
		fmt.Println("Or disponible :", c.Gold)
		for i, it := range items {
			fmt.Printf("%d. %s - %d pièces d'or\n", i+1, it.Name, it.Cost)
		}
		fmt.Println("0. Retourner au menu")
		fmt.Print("Choix :")

		choix := readInt()

		if choix == 0 {
			return
		}

		if choix < 1 || choix > len(items) {
			fmt.Println("Choix Invalide.")
			continue
		}
		choisi := items[choix-1]
		if c.Gold < choisi.Cost {
			fmt.Print("Vous n'avez pas assez d'or pour cet item.")
			continue
		}
		if !canAddItem(c) {
			fmt.Println("Vous avez plus de place dans votre inventaire.")
			continue
		}

		c.Gold -= choisi.Cost
		addInventory(c, choisi.Name)

	}

}

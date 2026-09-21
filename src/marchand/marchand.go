package marchand

import (
	"fmt"

	"outbreak/classes"
)

type ShopItem struct {
	Name string
	Cost int
}

func Merchant(c *classes.Classe) {
	items := []ShopItem{
		{"Kit de chirurgie de fortune", 40},
		{"Poche de perfusion", 20},
		{"Bandage militaire", 10},
		{"Pansement compressif", 7},
		{"Seringue d'adrénaline", 5},
		{"Poison à explosion différée", 30},
		{"Poison à stacks", 20},
		{"Poison paralysant", 15},
		{"Poison affaiblissant", 10},
		{"Poison à dégâts sur la durée", 5},
		{"Acier", 15},
		{"Bois", 10},
		{"Tissu", 5},
		{"Plume", 2},
		{"Sac à dos militaire supplémentaire", 30},
	}

	for {
		fmt.Println("----Le Troqueur----")
		fmt.Println("Or disponible :", c.Gold)
		for i, it := range items {
			fmt.Printf("%d. %s - %d pièces d'or\n", i+1, it.Name, it.Cost)
		}
		fmt.Println("0. Retourner au menu")
		fmt.Print("Choix : ")

		choix := classes.ReadInt()
		if choix == 0 {
			return
		}
		if choix < 1 || choix > len(items) {
			fmt.Println("Choix invalide.")
			continue
		}
		choisi := items[choix-1]
		if c.Gold < choisi.Cost {
			fmt.Println("Vous n'avez pas assez d'or pour cet item.")
			continue
		}
		if !classes.CanAddItem(c) {
			fmt.Println("Vous n'avez plus de place dans votre inventaire.")
			continue
		}
		c.Gold -= choisi.Cost
		classes.AddInventory(c, choisi.Name)
	}
}
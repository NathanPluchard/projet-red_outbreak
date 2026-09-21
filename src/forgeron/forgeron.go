package forgeron

import (
	"fmt"
)

type ArmorTiers struct {
	Level int
	TierName string
	Torso string
	Head string
	Feet string
	Materials map[string]int
	Cost int
}

var armorTiers = []ArmorTiers {
	{1, "Civil", "Casquette renforcée", "Veste de travail", "Chaussures de randonnée", map[string]int{"Tissu" : 4, "Bois" : 2, "Plume" : 1}, 10},
	{2, "Survivant", "Casque de chantier renforcé", "Veste tactique légère", "Bottes de randonnée renforcées", map[string]int{"Tissu" : 5, "Bois" : 3, "Acier" : 1, "Plume" : 2}, 20},
	{3, "Soldat", "Casque tactique", "Gilet tactique renforcé", "Bottes tactiques", map[string]int{"Tissu" : 4, "Bois" : 2, "Acier" : 4, "Plume" : 2}, 35},
	{4, "Tank", "Casque balistique", "Armure à plaques", "Bottes blindées", map[string]int{"Tissu" : 3, "Bois" : 1, "Acier" : 8, "Plume" : 3}, 55},
	{5, "Légendaire", "Casque intégral composite", "Armure anti-zombies", "Bottes renforcées composites", map[string]int{"Tissu" : 2, "Acier" : 12, "Plume" : 4}, 90},
}

func blacksmithTiers (c *Classe) {
	for {
		fmt.Println("---- Le Bricoleur : Atelier pour tout type de protection ----")
		fmt.Println("Votre or :", c.Gold)
		for _, t := range armorTiers {
			fmt.Println("%d. Palier %s (%s / %s / %s) - %d pièces d'or\n", t.Level, t.TierName, t.Head, t.Torso, t.Head, t.Cost)
		}
		fmt.Println("0. Retour")
		fmt.Println("Choix :")

		choix := readInt()
		if choix == 0 {
			return
		}
		if choix < 0 || choix > len(armorTiers) {
			fmt.Println("Choix invalide")
			continue
		}
		tier := armorTiers[choix-1]

		if c.Gold < tier.Cost {
			fmt.Println("Vous n'avez pas assez d'argent pour ce palier")
			continue
		}
		missing := false
		for mat, qty := range tier.Materials {
			if countItem(c, mat) < qty {
				fmt.Println("Il vous manque : %s (x%d requis)\n", mat, qty)
				missing = true
			}
		}
		if missing {
			continue
		}

		if len(c.Inventory) + 3 > c.MaxInventory {
			fmt.Println("Pas assez de place pour ces 3 équipements")
			continue
		}
		for mat, qty := range tier.Materials {
			for i := 0; i < qty; i++ {
				removeInventory(c, mat)
			}
		}
		c.Gold -= tier.Cost

		addInventory(c, tier.Head)
		addInventory(c, tier.Torso)
		addInventory(c, tier.Feet)

		fmt.Println("Set", tier.TierName, "fabriqué avec succès !!")
	}
}
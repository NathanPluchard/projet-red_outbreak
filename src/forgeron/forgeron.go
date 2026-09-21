package forgeron

import (
	"fmt"

	"outbreak/classes"
)

type ArmorTier struct {
	Level     int
	TierName  string
	Head      string
	Torso     string
	Feet      string
	Materials map[string]int
	Cost      int
}

var ArmorTiers = []ArmorTier{
	{1, "Civil", "Casquette renforcée", "Veste de travail", "Chaussures de randonnée", map[string]int{"Tissu": 4, "Bois": 2, "Plume": 1}, 10},
	{2, "Survivant", "Casque de chantier renforcé", "Veste tactique légère", "Bottes de randonnée renforcées", map[string]int{"Tissu": 5, "Bois": 3, "Acier": 1, "Plume": 2}, 20},
	{3, "Soldat", "Casque tactique", "Gilet tactique renforcé", "Bottes tactiques", map[string]int{"Tissu": 4, "Bois": 2, "Acier": 4, "Plume": 2}, 35},
	{4, "Tank", "Casque balistique", "Armure à plaques", "Bottes blindées", map[string]int{"Tissu": 3, "Bois": 1, "Acier": 8, "Plume": 3}, 55},
	{5, "Légendaire", "Casque intégral composite", "Armure anti-zombies", "Bottes renforcées composites", map[string]int{"Tissu": 2, "Acier": 12, "Plume": 4}, 90},
}

func BlacksmithTiers(c *classes.Classe) {
	for {
		fmt.Println("---- Le Bricoleur : Atelier pour tout type de protection ----")
		fmt.Println("Votre or :", c.Gold)
		for _, t := range ArmorTiers {
			fmt.Printf("%d. Palier %s (%s / %s / %s) - %d pièces d'or\n", t.Level, t.TierName, t.Head, t.Torso, t.Feet, t.Cost)
		}
		fmt.Println("0. Retour")
		fmt.Print("Choix : ")

		choix := classes.ReadInt()
		if choix == 0 {
			return
		}
		if choix < 1 || choix > len(ArmorTiers) {
			fmt.Println("Choix invalide")
			continue
		}
		tier := ArmorTiers[choix-1]

		if c.Gold < tier.Cost {
			fmt.Println("Vous n'avez pas assez d'argent pour ce palier")
			continue
		}
		missing := false
		for mat, qty := range tier.Materials {
			if classes.CountItem(c, mat) < qty {
				fmt.Printf("Il vous manque : %s (x%d requis)\n", mat, qty)
				missing = true
			}
		}
		if missing {
			continue
		}
		if len(c.Inventory)+3 > c.MaxInventory {
			fmt.Println("Pas assez de place pour ces 3 équipements")
			continue
		}

		for mat, qty := range tier.Materials {
			for i := 0; i < qty; i++ {
				classes.RemoveInventory(c, mat)
			}
		}
		c.Gold -= tier.Cost
		classes.AddInventory(c, tier.Head)
		classes.AddInventory(c, tier.Torso)
		classes.AddInventory(c, tier.Feet)

		fmt.Println("Set", tier.TierName, "fabriqué avec succès !!")
	}
}

func FindTierSlot(item string) (tier *ArmorTier, slot string) {
	for i := range ArmorTiers {
		t := &ArmorTiers[i]
		if t.Head == item {
			return t, "head"
		}
		if t.Torso == item {
			return t, "torso"
		}
		if t.Feet == item {
			return t, "feet"
		}
	}
	return nil, ""
}

func HpBonus(slot string, level int) int {
	switch slot {
	case "head":
		return level * 5
	case "torso":
		return level * 10
	case "feet":
		return level * 7
	}
	return 0
}

func EquipItem(c *classes.Classe, item string) {
	tier, slot := FindTierSlot(item)
	if tier == nil {
		fmt.Println("Cet objet ne peut pas être équipé.")
		return
	}
	classes.RemoveInventory(c, item)
	bonus := HpBonus(slot, tier.Level)

	switch slot {
	case "head":
		if c.Equip.Head != "" {
			if oldTier, _ := FindTierSlot(c.Equip.Head); oldTier != nil {
				c.PVBase -= HpBonus("head", oldTier.Level)
			}
			classes.AddInventory(c, c.Equip.Head)
		}
		c.Equip.Head = item
	case "torso":
		if c.Equip.Torso != "" {
			if oldTier, _ := FindTierSlot(c.Equip.Torso); oldTier != nil {
				c.PVBase -= HpBonus("torso", oldTier.Level)
			}
			classes.AddInventory(c, c.Equip.Torso)
		}
		c.Equip.Torso = item
	case "feet":
		if c.Equip.Feet != "" {
			if oldTier, _ := FindTierSlot(c.Equip.Feet); oldTier != nil {
				c.PVBase -= HpBonus("feet", oldTier.Level)
			}
			classes.AddInventory(c, c.Equip.Feet)
		}
		c.Equip.Feet = item
	}

	c.PVBase += bonus
	fmt.Println("Vous équipez :", item, "( +", bonus, "PV max )")
}
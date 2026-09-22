package potion

import (
	"fmt"

	"outbreak/classes"
)

type Potion struct {
	Nom         string
	Description string
	Effet       int
	Type        string
}

func trouverPotion(nom string) (Potion, bool) {
	for _, p := range PotionsDeVieDisponibles {
		if p.Nom == nom {
			return p, true
		}
	}
	for _, p := range PotionsDePoisonDisponibles {
		if p.Nom == nom {
			return p, true
		}
	}
	return Potion{}, false
}

func UtiliserPotion(c *classes.Classe, nomPotion string) {
	index := -1
	for i, nom := range c.Inventory {
		if nom == nomPotion {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Printf("%s n'a pas de %s dans son inventaire.\n", c.Nom, nomPotion)
		return
	}

	p, trouve := trouverPotion(nomPotion)
	if !trouve {
		fmt.Printf("Potion inconnue : %s\n", nomPotion)
		return
	}

	switch p.Type {
	case "vie":
		soin := p.Effet
		// Spécialité Médecin de fortune : efficacité de soin augmentée (+50%)
		if c.Type == "Médecin de fortune" {
			soin = soin * 150 / 100
		}
		c.PV += soin
		if c.PV > c.PVBase {
			c.PV = c.PVBase
		}
		fmt.Printf("%s%s utilise %s et regagne %d PV !%s (PV : %s %d/%d)\n", classes.BrightGreen, c.Nom, p.Nom, soin, classes.Reset, classes.HPBar(c.PV, c.PVBase), c.PV, c.PVBase)
	case "poison":
		c.PV -= p.Effet
		if c.PV < 0 {
			c.PV = 0
		}
		fmt.Printf("%s%s subit %s et perd %d PV !%s (PV : %s %d/%d)\n", classes.BrightRed, c.Nom, p.Nom, p.Effet, classes.Reset, classes.HPBar(c.PV, c.PVBase), c.PV, c.PVBase)
	}

	c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
}

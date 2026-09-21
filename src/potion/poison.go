package potion

func NouvellePotionDePoison(nom string, degats int, description string) Potion {
	return Potion{
		Nom:         nom,
		Description: description,
		Effet:       degats,
		Type:        "poison",
	}
}

var PotionsDePoisonDisponibles = []Potion{
	NouvellePotionDePoison("Poison à explosion différée", 90, "Explose après quelques tours et inflige énormément de dégâts."),
	NouvellePotionDePoison("Poison à stacks", 75, "Chaque application ajoute un stack et augmente les dégâts."),
	NouvellePotionDePoison("Poison paralysant", 50, "Inflige des dégâts et peut empêcher la cible d'agir pendant un tour."),
	NouvellePotionDePoison("Poison affaiblissant", 35, "Inflige des dégâts et réduit une statistique de la cible."),
	NouvellePotionDePoison("Poison à dégâts sur la durée", 10, "Inflige simplement des dégâts pendant plusieurs tours."),
}

func trouverPotion(nom string) (potion.Potion, bool) {
	for _, p := range potion.PotionsDeVieDisponibles {
		if p.Nom == nom {
			return p, true
		}
	}
	for _, p := range potion.PotionsDePoisonDisponibles {
		if p.Nom == nom {
			return p, true
		}
	}
	return potion.Potion{}, false
}


func (c *Classe) UtiliserPotion(nomPotion string) {
	
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
		c.PV += p.Effet
		if c.PV > c.PVBase {
			c.PV = c.PVBase
		}
		fmt.Printf("%s utilise %s et regagne %d PV ! (PV: %d/%d)\n",
			c.Nom, p.Nom, p.Effet, c.PV, c.PVBase)

	case "poison":
		c.PV -= p.Effet
		if c.PV < 0 {
			c.PV = 0
		}
		fmt.Printf("%s subit %s et perd %d PV ! (PV: %d/%d)\n",
			c.Nom, p.Nom, p.Effet, c.PV, c.PVBase)
	}


	c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
}

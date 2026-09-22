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

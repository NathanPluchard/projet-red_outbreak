package sort

type Sort struct {
	Nom         string
	Description string
	Degats      int
	CoutMana    int
}

func NouveauSort(nom string, degats int, coutMana int, description string) Sort {
	return Sort{
		Nom:         nom,
		Description: description,
		Degats:      degats,
		CoutMana:    coutMana,
	}
}

var SortsDisponibles = []Sort{
	NouveauSort("Frappe nucléaire improvisée", 100, 50,
		"Un engin de fortune bricolé avec des matériaux radioactifs, dévaste une large zone."),

	NouveauSort("Tir de sniper longue distance", 80, 30,
		"Un tir précis et létal depuis une position éloignée, redoutable contre une cible isolée."),

	NouveauSort("Rafale de mitrailleuse", 65, 25,
		"Une pluie de balles continue qui inflige de lourds dégâts mais consomme beaucoup de munitions."),

	NouveauSort("Cocktail Molotov", 45, 10,
		"Projette une bouteille enflammée qui inflige des dégâts de zone importants."),

	NouveauSort("Lance-pierre renforcé", 30, 5,
		"Une arme de fortune simple mais efficace contre les ennemis faibles."),

	NouveauSort("Coup de crosse", 20, 0,
		"Frappe rapide et fiable avec la crosse de son arme, ne coûte aucune ressource."),

	NouveauSort("Coup de pied de survie", 10, 0,
		"Une frappe désespérée quand on n'a plus rien d'autre sous la main."),
}

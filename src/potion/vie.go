package potion

import "fmt"

func NouvellePotionDeVie(nom string, soin int, description string) Potion {
	_ = fmt.Sprintf // (fmt reste utilisé plus bas si besoin ; sinon retire l'import)
	return Potion{
		Nom:         nom,
		Description: description,
		Effet:       soin,
		Type:        "vie",
	}
}

var PotionsDeVieDisponibles = []Potion{
	NouvellePotionDeVie("Kit de chirurgie de fortune", 100, "Réparation quasi-totale des blessures, agit instantanément."),
	NouvellePotionDeVie("Poche de perfusion", 75, "Stabilise et régénère progressivement une grande partie des PV."),
	NouvellePotionDeVie("Bandage militaire", 50, "Soigne une blessure moyenne et réduit les saignements."),
	NouvellePotionDeVie("Pansement compressif", 25, "Arrête rapidement une hémorragie légère et restaure un peu de vie."),
	NouvellePotionDeVie("Seringue d'adrénaline", 15, "Redonne un peu d'énergie et quelques PV, effet immédiat mais faible."),
}

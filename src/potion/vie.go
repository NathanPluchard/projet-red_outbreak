package potion

import "fmt"


func NouvellePotionDeVie(nom string, soin int) Potion {
    return Potion{
        Nom:         nom,
        Description: fmt.Sprintf("Restaure %d points de vie", soin),
        Effet:       soin,
        Type:        "vie",
    }
}

var PotionsDeVieDisponibles = []Potion{
    NouvellePotionDeVie("Kit de chirurgie de fortune", 100),
    NouvellePotionDeVie("Poche de perfusion", 75),
    NouvellePotionDeVie("Bandage militaire", 50),
    NouvellePotionDeVie("Pansement compressif", 25),
    NouvellePotionDeVie("Seringue d'adrénaline", 15),
}


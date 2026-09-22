package main

import (
	"fmt"
	"strings"

	"outbreak/classes"
	"outbreak/forgeron"
	"outbreak/inventaire"
	"outbreak/marchand"
	"outbreak/monstre"
)

func isAlpha(s string) bool {
	for _, r := range s {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
			return false
		}
	}
	return true
}

func formatName(s string) string {
	s = strings.ToLower(s)
	return strings.ToUpper(s[:1]) + s[1:]
}

func characterCreation() classes.Classe {
	fmt.Print("Entrez votre nom (lettres uniquement) : ")
	var nom string
	for {
		input, _ := classes.Reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if isAlpha(input) && input != "" {
			nom = formatName(input)
			break
		}
		fmt.Print("Nom invalide, réessayez : ")
	}

	fmt.Println("Choisissez votre classe :")
	fmt.Println("1. Survivant")
	fmt.Println("2. Punk")
	fmt.Println("3. Médecin de fortune")
	fmt.Println("4. Sauveur")

	var nomClasse string
	for {
		switch classes.ReadInt() {
		case 1:
			nomClasse = "Survivant"
		case 2:
			nomClasse = "Punk"
		case 3:
			nomClasse = "Médecin de fortune"
		case 4:
			nomClasse = "Sauveur"
		default:
			fmt.Println("Choix invalide, réessayez.")
			continue
		}
		break
	}

	perso := classes.Classes[nomClasse]
	perso.Type = nomClasse
	perso.Nom = nom
	perso.PV = perso.PVBase
	perso.Level = 1
	perso.Attacks = []string{"Attaque basique"}
	perso.Inventory = []string{}

	return perso
}

func showArtists() {
	fmt.Println("Les artistes cachés dans les slides sont les icônes fournies par les créateurs du support visuel (illustrations libres de droit).")
}

func mainMenu(c *classes.Classe) {
	for {
		fmt.Println("\n===== MENU =====")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Marchand (Le Troqueur)")
		fmt.Println("4. Forgeron (Le Bricoleur)")
		fmt.Println("5. Entrainement")
		fmt.Println("6. Qui sont-ils")
		fmt.Println("0. Quitter")
		fmt.Print("Choix : ")

		switch classes.ReadInt() {
		case 1:
			classes.DisplayInfo(c)
		case 2:
			inventaire.AccessInventory(c)
		case 3:
			marchand.Merchant(c)
		case 4:
			forgeron.BlacksmithTiers(c)
		case 5:
			monstre.TrainingFight(c)
		case 6:
			showArtists()
		case 0:
			fmt.Println("A bientôt l'ami j'espère qu'on se reverra.....")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func main() {
	fmt.Println("====Bienvenue cher survivant veillez choisir un nom====")
	fmt.Println("=====DEPART=======")
	c1 := characterCreation()
	mainMenu(&c1)
}
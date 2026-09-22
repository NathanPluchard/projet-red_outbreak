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

	fmt.Println(classes.BrightCyan + "Choisissez votre classe :" + classes.Reset)
	fmt.Printf("%s1.%s Survivant\n", classes.BrightCyan, classes.Reset)
	fmt.Printf("%s2.%s Punk\n", classes.BrightCyan, classes.Reset)
	fmt.Printf("%s3.%s Médecin de fortune\n", classes.BrightCyan, classes.Reset)
	fmt.Printf("%s4.%s Sauveur\n", classes.BrightCyan, classes.Reset)

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
			fmt.Println(classes.BrightRed + "Choix invalide, réessayez." + classes.Reset)
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
		fmt.Println()
		fmt.Println(classes.TitleBox("MENU PRINCIPAL"))
		fmt.Printf("%s1.%s Afficher les informations du personnage\n", classes.BrightCyan, classes.Reset)
		fmt.Printf("%s2.%s Accéder à l'inventaire\n", classes.BrightCyan, classes.Reset)
		fmt.Printf("%s3.%s Marchand (Le Troqueur)\n", classes.BrightCyan, classes.Reset)
		fmt.Printf("%s4.%s Forgeron (Le Bricoleur)\n", classes.BrightCyan, classes.Reset)
		fmt.Printf("%s5.%s %sEntrainement%s\n", classes.BrightCyan, classes.Reset, classes.BrightRed, classes.Reset)
		fmt.Printf("%s6.%s Qui sont-ils\n", classes.BrightCyan, classes.Reset)
		fmt.Printf("%s0.%s Quitter\n", classes.Dim, classes.Reset)
		fmt.Print(classes.Bold + "Choix : " + classes.Reset)

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
			fmt.Println(classes.BrightYellow + "A bientôt l'ami j'espère qu'on se reverra....." + classes.Reset)
			return
		default:
			fmt.Println(classes.BrightRed + "Choix invalide." + classes.Reset)
		}
	}
}

func main() {
	fmt.Println(classes.Bold + classes.BrightRed + `
   ____  _   _ _____ ____  ____  _____    _    _  __
  / __ \| | | |_   _|  _ \|  _ \| ____|  / \  | |/ /
 | |  | | | | | | | | |_) | |_) |  _|   / _ \ | ' / 
 | |__| | |_| | | | |  _ <|  _ <| |___ / ___ \| . \ 
  \____/ \___/  |_| |_| \_\_| \_\_____/_/   \_\_|\_\
` + classes.Reset)
	fmt.Println(classes.Dim + "====Bienvenue cher survivant, veillez choisir un nom====" + classes.Reset)
	fmt.Println(classes.Dim + "=====DEPART=======" + classes.Reset)
	c1 := characterCreation()
	mainMenu(&c1)
}
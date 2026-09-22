package main

import (
"fmt"
"strings"


"outbreak/casino"
"outbreak/classes"
"outbreak/forgeron"
"outbreak/inventaire"
"outbreak/marchand"
"outbreak/monstre"


)



func isAlpha(s string) bool {


if s == "" {
	return false
}

for _, r := range s {

	if !((r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z')) {

		return false
	}
}

return true


}

func formatName(s string) string {


s = strings.ToLower(s)

if len(s) == 0 {
	return s
}

return strings.ToUpper(s[:1]) + s[1:]


}



func characterCreation() classes.Classe {


classes.ClearScreen()

fmt.Println(
	classes.TitleBox(
		"☣ OUTBREAK • CRÉATION DU SURVIVANT ☣",
	),
)

fmt.Println()

fmt.Println(
	classes.Panel(
		"IDENTITÉ",
		"Choisissez le nom de votre survivant.",
		"Utilisez uniquement des lettres.",
	),
)

fmt.Print(
	"\n" +
		classes.BrightCyan +
		"Nom > " +
		classes.Reset,
)

var nom string

for {

	input, _ :=
		classes.Reader.ReadString('\n')

	input =
		strings.TrimSpace(input)

	if isAlpha(input) {

		nom =
			formatName(input)

		break
	}

	fmt.Println(
		classes.BrightRed +
			"✖ Nom invalide." +
			classes.Reset,
	)

	fmt.Print(
		"Nom > ",
	)
}



classes.ClearScreen()

fmt.Println(
	classes.TitleBox(
		"☣ CHOIX DE VOTRE CLASSE ☣",
	),
)

fmt.Println()

fmt.Println(
	classes.Panel(
		"CLASSES DISPONIBLES",

		classes.BrightGreen+
			"[1] Survivant"+
			classes.Reset+
			"  • équilibre et régénération",

		classes.BrightMagenta+
			"[2] Punk"+
			classes.Reset+
			"  • coups critiques",

		classes.BrightCyan+
			"[3] Médecin de fortune"+
			classes.Reset+
			"  • soins améliorés",

		classes.BrightYellow+
			"[4] Sauveur"+
			classes.Reset+
			"  • dégâts subis réduits",
	),
)

fmt.Print(
	"\n" +
		classes.BrightCyan +
		"Classe > " +
		classes.Reset,
)

var nomClasse string

for {

	choix :=
		classes.ReadInt()

	switch choix {

	case 1:
		nomClasse = "Survivant"

	case 2:
		nomClasse = "Punk"

	case 3:
		nomClasse = "Médecin de fortune"

	case 4:
		nomClasse = "Sauveur"

	default:

		fmt.Println(
			classes.BrightRed +
				"✖ Choix invalide." +
				classes.Reset,
		)

		fmt.Print(
			"Classe > ",
		)

		continue
	}

	break
}



perso :=
	classes.Classes[nomClasse]

perso.Nom =
	nom

perso.Type =
	nomClasse

perso.Level =
	1

perso.PV =
	perso.PVBase

perso.Mana =
	perso.ManaMax

perso.Attacks =
	[]string{
		"Attaque basique",
	}

perso.Inventory =
	[]string{}

classes.ClearScreen()

fmt.Println(
	classes.TitleBox(
		"✔ SURVIVANT CRÉÉ",
	),
)

fmt.Println()

fmt.Printf(
	"%sNom :%s %s\n",
	classes.Bold,
	classes.Reset,
	perso.Nom,
)

fmt.Printf(
	"%sClasse :%s %s\n",
	classes.Bold,
	classes.Reset,
	perso.Type,
)

fmt.Printf(
	"%sPV :%s %d\n",
	classes.Bold,
	classes.Reset,
	perso.PV,
)

fmt.Printf(
	"%sMana :%s %d\n",
	classes.Bold,
	classes.Reset,
	perso.Mana,
)

classes.Pause()

return perso


}



func showArtists() {


classes.ClearScreen()

fmt.Println(
	classes.TitleBox(
		"QUI SONT-ILS ?",
	),
)

fmt.Println()

fmt.Println(
	classes.Panel(
		"À PROPOS DU JEU",

		"Bienvenue dans Outbreak.",

		"Votre objectif est de survivre",
		"aux vagues d'infectés.",

		"Explorez, combattez, améliorez",
		"votre équipement et gérez vos ressources.",
	),
)

classes.Pause()


}



func mainMenu(c *classes.Classe) {


for {

	classes.ClearScreen()

	fmt.Println(
		classes.TitleBox(
			"☣ OUTBREAK • REFUGE CENTRAL ☣",
		),
	)

	fmt.Println()

	

	fmt.Printf(
		"%s%s%s\n",
		classes.Bold,
		c.Nom,
		classes.Reset,
	)

	fmt.Printf(
		"Niveau %d\n",
		c.Level,
	)

	fmt.Printf(
		"PV    %s %d/%d\n",
		classes.HPBar(
			c.PV,
			c.PVBase,
		),
		c.PV,
		c.PVBase,
	)

	fmt.Printf(
		"Mana  %s %d/%d\n",
		classes.ManaBar(
			c.Mana,
			c.ManaMax,
		),
		c.Mana,
		c.ManaMax,
	)

	fmt.Printf(
		"XP    %s %d/%d\n",
		classes.ExpBar(
			c.Exp,
			c.MaxExp,
		),
		c.Exp,
		c.MaxExp,
	)

	fmt.Printf(
		"\n%s💰 %d%s\n\n",
		classes.BrightYellow,
		c.Gold,
		classes.Reset,
	)

	
	fmt.Println(
		classes.Panel(
			"CENTRE DE COMMANDE",

			classes.BrightCyan+
				"[1] "+
				classes.Reset+
				"Fiche du survivant",

			classes.BrightCyan+
				"[2] "+
				classes.Reset+
				"Inventaire",

			classes.BrightYellow+
				"[3] "+
				classes.Reset+
				"Le Troqueur",

			classes.BrightBlue+
				"[4] "+
				classes.Reset+
				"Le Bricoleur",

			classes.BrightRed+
				"[5] "+
				classes.Reset+
				"Simulation de combat",

			classes.BrightMagenta+
				"[6] "+
				classes.Reset+
				"Casino",

			classes.Dim+
				"[7] "+
				classes.Reset+
				"Qui sont-ils",

			classes.Dim+
				"[0] "+
				classes.Reset+
				"Quitter",
		),
	)

	fmt.Print(
		"\n" +
			classes.BrightCyan +
			"Action > " +
			classes.Reset,
	)

	choix :=
		classes.ReadInt()

	switch choix {

	

	case 1:

		classes.ClearScreen()

		classes.DisplayInfo(c)

		classes.Pause()

	

	case 2:

		inventaire.AccessInventory(c)

	

	case 3:

		marchand.Merchant(c)

	-

	case 4:

		forgeron.BlacksmithTiers(c)

	

	case 5:

		monstre.SimulationFight(c)

	

	case 6:

		casino.Casino(c)

	

	case 7:

		showArtists()

	

	case 0:

		classes.ClearScreen()

		fmt.Println(
			classes.TitleBox(
				"☣ FIN DE LA SESSION ☣",
			),
		)

		fmt.Printf(
			"\n%sÀ bientôt, survivant %s.%s\n\n",
			classes.BrightYellow,
			c.Nom,
			classes.Reset,
		)

		return

	default:

		fmt.Println(
			classes.BrightRed +
				"\n✖ Choix invalide." +
				classes.Reset,
		)

		classes.Pause()
	}
}


}



func main() {


personnage :=
	characterCreation()

mainMenu(
	&personnage,
)


}

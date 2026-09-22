package marchand

import (
	"fmt"

	"outbreak/classes"
)

// ============================================================
// MARCHAND
// ============================================================

type Item struct {
	Name        string
	Description string
	Price       int
}

// Objets disponibles chez le marchand.
var Items = []Item{
	{
		Name:        "Potion de soin",
		Description: "Restaure 30 PV.",
		Price:       25,
	},
	{
		Name:        "Grande potion de soin",
		Description: "Restaure 70 PV.",
		Price:       50,
	},
	{
		Name:        "Potion de mana",
		Description: "Restaure 25 mana.",
		Price:       30,
	},
	{
		Name:        "Grande potion de mana",
		Description: "Restaure 60 mana.",
		Price:       55,
	},
	{
		Name:        "Kit de survie",
		Description: "Restaure PV et mana.",
		Price:       80,
	},
}

// ============================================================
// MENU PRINCIPAL DU MARCHAND
// ============================================================

func Merchant(c *classes.Classe) {

	for {

		classes.ClearScreen()

		fmt.Println(
			classes.TitleBox(
				"🛒 LE TROQUEUR",
			),
		)

		fmt.Printf(
			"\n%sVotre or : %d 💰%s\n\n",
			classes.BrightYellow,
			c.Gold,
			classes.Reset,
		)

		fmt.Println(
			classes.Panel(
				"BOUTIQUE",
				"[1] Consulter les objets",
				"[2] Acheter un objet",
				"[0] Quitter la boutique",
			),
		)

		fmt.Print(
			"\n" +
				classes.BrightCyan +
				"Action > " +
				classes.Reset,
		)

		choice := classes.ReadInt()

		switch choice {

		case 1:
			displayItems(c)

		case 2:
			buyItem(c)

		case 0:
			return

		default:

			fmt.Println(
				classes.BrightRed +
					"✖ Choix invalide." +
					classes.Reset,
			)

			classes.Pause()
		}
	}
}

// ============================================================
// AFFICHAGE DES OBJETS
// ============================================================

func displayItems(c *classes.Classe) {

	classes.ClearScreen()

	fmt.Println(
		classes.TitleBox(
			"🛒 OBJETS DISPONIBLES",
		),
	)

	fmt.Printf(
		"\n%sVotre portefeuille : %d 💰%s\n\n",
		classes.BrightYellow,
		c.Gold,
		classes.Reset,
	)

	for i, item := range Items {

		fmt.Printf(
			"%s[%d] %s%s\n",
			classes.BrightCyan,
			i+1,
			item.Name,
			classes.Reset,
		)

		fmt.Printf(
			"    %s%s%s\n",
			classes.Dim,
			item.Description,
			classes.Reset,
		)

		fmt.Printf(
			"    Prix : %s%d 💰%s\n\n",
			classes.BrightYellow,
			item.Price,
			classes.Reset,
		)
	}

	classes.Pause()
}

// ============================================================
// ACHAT
// ============================================================

func buyItem(c *classes.Classe) {

	classes.ClearScreen()

	fmt.Println(
		classes.TitleBox(
			"🛒 ACHETER UN OBJET",
		),
	)

	fmt.Printf(
		"\n%sVotre or : %d 💰%s\n\n",
		classes.BrightYellow,
		c.Gold,
		classes.Reset,
	)

	for i, item := range Items {

		fmt.Printf(
			"%s[%d]%s %-28s %s%d 💰%s\n",
			classes.BrightCyan,
			i+1,
			classes.Reset,
			item.Name,
			classes.BrightYellow,
			item.Price,
			classes.Reset,
		)
	}

	fmt.Println(
		"\n" +
			classes.Dim +
			"[0] Annuler" +
			classes.Reset,
	)

	fmt.Print(
		"\nObjet > ",
	)

	choice := classes.ReadInt()

	if choice == 0 {
		return
	}

	if choice < 1 || choice > len(Items) {

		fmt.Println(
			classes.BrightRed +
				"✖ Objet invalide." +
				classes.Reset,
		)

		classes.Pause()

		return
	}

	item := Items[choice-1]

	// --------------------------------------------------------
	// VÉRIFICATION DE L'ARGENT
	// --------------------------------------------------------

	if c.Gold < item.Price {

		fmt.Println(
			"\n" +
				classes.BrightRed +
				"✖ Vous n'avez pas assez d'or." +
				classes.Reset,
		)

		classes.Pause()

		return
	}

	// --------------------------------------------------------
	// PAIEMENT
	// --------------------------------------------------------

	c.Gold -= item.Price

	// --------------------------------------------------------
	// APPLICATION DE L'OBJET
	// --------------------------------------------------------

	switch item.Name {

	case "Potion de soin":

		c.PV += 30

		if c.PV > c.PVBase {
			c.PV = c.PVBase
		}

	case "Grande potion de soin":

		c.PV += 70

		if c.PV > c.PVBase {
			c.PV = c.PVBase
		}

	case "Potion de mana":

		c.Mana += 25

		if c.Mana > c.ManaMax {
			c.Mana = c.ManaMax
		}

	case "Grande potion de mana":

		c.Mana += 60

		if c.Mana > c.ManaMax {
			c.Mana = c.ManaMax
		}

	case "Kit de survie":

		c.PV += 50

		if c.PV > c.PVBase {
			c.PV = c.PVBase
		}

		c.Mana += 40

		if c.Mana > c.ManaMax {
			c.Mana = c.ManaMax
		}
	}

	// --------------------------------------------------------
	// CONFIRMATION
	// --------------------------------------------------------

	fmt.Println()

	fmt.Println(
		classes.BrightGreen +
			"╔════════════════════════════════════════╗" +
			classes.Reset,
	)

	fmt.Println(
		classes.BrightGreen +
			"║          ✔ ACHAT EFFECTUÉ             ║" +
			classes.Reset,
	)

	fmt.Println(
		classes.BrightGreen +
			"╚════════════════════════════════════════╝" +
			classes.Reset,
	)

	fmt.Printf(
		"\n%sVous avez acheté : %s%s\n",
		classes.Bold,
		item.Name,
		classes.Reset,
	)

	fmt.Printf(
		"%s-%d 💰%s\n",
		classes.BrightRed,
		item.Price,
		classes.Reset,
	)

	fmt.Printf(
		"Or restant : %s%d 💰%s\n",
		classes.BrightYellow,
		c.Gold,
		classes.Reset,
	)

	classes.Pause()
}

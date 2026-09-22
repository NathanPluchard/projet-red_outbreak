package forgeron

import (
	"fmt"

	"outbreak/classes"
)

// ============================================================
// AMÉLIORATIONS DU FORGERON
// ============================================================

type Upgrade struct {
	Name        string
	Description string
	Cost        int
	Level       int
	Damage      int
	Defense     int
}

// Les différents niveaux d'amélioration.
var Upgrades = []Upgrade{
	{
		Name:        "Arme renforcée",
		Description: "Augmente les dégâts de base.",
		Cost:        75,
		Level:       2,
		Damage:      3,
		Defense:     0,
	},
	{
		Name:        "Armure renforcée",
		Description: "Améliore la défense.",
		Cost:        100,
		Level:       3,
		Damage:      0,
		Defense:     3,
	},
	{
		Name:        "Arme militaire",
		Description: "Une amélioration importante des dégâts.",
		Cost:        175,
		Level:       5,
		Damage:      7,
		Defense:     0,
	},
	{
		Name:        "Armure tactique",
		Description: "Une protection supérieure.",
		Cost:        225,
		Level:       7,
		Damage:      0,
		Defense:     7,
	},
	{
		Name:        "Équipement d'élite",
		Description: "Améliore fortement l'équipement.",
		Cost:        400,
		Level:       10,
		Damage:      12,
		Defense:     12,
	},
}

// ============================================================
// FORGERON
// ============================================================

func BlacksmithTiers(c *classes.Classe) {

	for {

		classes.ClearScreen()

		fmt.Println(
			classes.TitleBox(
				"🔨 LE BRICOLEUR",
			),
		)

		fmt.Printf(
			"\n%sNiveau : %d%s\n",
			classes.BrightCyan,
			c.Level,
			classes.Reset,
		)

		fmt.Printf(
			"%sOr disponible : %d 💰%s\n\n",
			classes.BrightYellow,
			c.Gold,
			classes.Reset,
		)

		fmt.Printf(
			"%sDégâts :%s %d\n",
			classes.Bold,
			classes.Reset,
			c.AttaqueBase,
		)

		fmt.Printf(
			"%sDéfense :%s %d\n\n",
			classes.Bold,
			classes.Reset,
			c.DefenseBase,
		)

		fmt.Println(
			classes.Panel(
				"ATELIER",
				"[1] Voir les améliorations",
				"[2] Améliorer l'équipement",
				"[0] Quitter",
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
			displayUpgrades(c)

		case 2:
			upgradeEquipment(c)

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
// AFFICHAGE
// ============================================================

func displayUpgrades(c *classes.Classe) {

	classes.ClearScreen()

	fmt.Println(
		classes.TitleBox(
			"🔨 AMÉLIORATIONS DISPONIBLES",
		),
	)

	fmt.Println()

	for i, upgrade := range Upgrades {

		fmt.Printf(
			"%s[%d] %s%s\n",
			classes.BrightCyan,
			i+1,
			upgrade.Name,
			classes.Reset,
		)

		fmt.Printf(
			"    %s%s%s\n",
			classes.Dim,
			upgrade.Description,
			classes.Reset,
		)

		fmt.Printf(
			"    Niveau requis : %d\n",
			upgrade.Level,
		)

		fmt.Printf(
			"    Prix : %s%d 💰%s\n",
			classes.BrightYellow,
			upgrade.Cost,
			classes.Reset,
		)

		if upgrade.Damage > 0 {
			fmt.Printf(
				"    Dégâts : %s+%d%s\n",
				classes.BrightGreen,
				upgrade.Damage,
				classes.Reset,
			)
		}

		if upgrade.Defense > 0 {
			fmt.Printf(
				"    Défense : %s+%d%s\n",
				classes.BrightGreen,
				upgrade.Defense,
				classes.Reset,
			)
		}

		fmt.Println()
	}

	classes.Pause()
}

// ============================================================
// ACHAT D'UNE AMÉLIORATION
// ============================================================

func upgradeEquipment(c *classes.Classe) {

	classes.ClearScreen()

	fmt.Println(
		classes.TitleBox(
			"🔨 ATELIER DU BRICOLEUR",
		),
	)

	fmt.Println()

	for i, upgrade := range Upgrades {

		available := ""

		if c.Level >= upgrade.Level {
			available =
				classes.BrightGreen +
					"DISPONIBLE" +
					classes.Reset
		} else {
			available =
				classes.BrightRed +
					"VERROUILLÉ" +
					classes.Reset
		}

		fmt.Printf(
			"%s[%d]%s %-25s %s%d 💰%s  %s\n",
			classes.BrightCyan,
			i+1,
			classes.Reset,
			upgrade.Name,
			classes.BrightYellow,
			upgrade.Cost,
			classes.Reset,
			available,
		)
	}

	fmt.Println(
		"\n" +
			classes.Dim +
			"[0] Annuler" +
			classes.Reset,
	)

	fmt.Print(
		"\nAmélioration > ",
	)

	choice := classes.ReadInt()

	if choice == 0 {
		return
	}

	if choice < 1 || choice > len(Upgrades) {

		fmt.Println(
			classes.BrightRed +
				"✖ Amélioration invalide." +
				classes.Reset,
		)

		classes.Pause()

		return
	}

	upgrade := Upgrades[choice-1]

	// --------------------------------------------------------
	// NIVEAU
	// --------------------------------------------------------

	if c.Level < upgrade.Level {

		fmt.Printf(
			"\n%s✖ Niveau insuffisant.%s\n",
			classes.BrightRed,
			classes.Reset,
		)

		fmt.Printf(
			"Niveau requis : %d\n",
			upgrade.Level,
		)

		classes.Pause()

		return
	}

	// --------------------------------------------------------
	// ARGENT
	// --------------------------------------------------------

	if c.Gold < upgrade.Cost {

		fmt.Printf(
			"\n%s✖ Vous n'avez pas assez d'or.%s\n",
			classes.BrightRed,
			classes.Reset,
		)

		fmt.Printf(
			"Prix : %d 💰\n",
			upgrade.Cost,
		)

		classes.Pause()

		return
	}

	// --------------------------------------------------------
	// PAIEMENT
	// --------------------------------------------------------

	c.Gold -= upgrade.Cost

	// --------------------------------------------------------
	// AMÉLIORATIONS
	// --------------------------------------------------------

	if upgrade.Damage > 0 {
		c.AttaqueBase += upgrade.Damage
	}

	if upgrade.Defense > 0 {
		c.DefenseBase += upgrade.Defense
	}

	// --------------------------------------------------------
	// CONFIRMATION
	// --------------------------------------------------------

	classes.ClearScreen()

	fmt.Println(
		classes.TitleBox(
			"✔ ÉQUIPEMENT AMÉLIORÉ",
		),
	)

	fmt.Println()

	fmt.Printf(
		"%sAmélioration :%s %s\n",
		classes.Bold,
		classes.Reset,
		upgrade.Name,
	)

	if upgrade.Damage > 0 {

		fmt.Printf(
			"%sDégâts : +%d%s\n",
			classes.BrightGreen,
			upgrade.Damage,
			classes.Reset,
		)
	}

	if upgrade.Defense > 0 {

		fmt.Printf(
			"%sDéfense : +%d%s\n",
			classes.BrightGreen,
			upgrade.Defense,
			classes.Reset,
		)
	}

	fmt.Printf(
		"\n%s-%d 💰%s\n",
		classes.BrightRed,
		upgrade.Cost,
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

// ============================================================
// COMPATIBILITÉ AVEC L'INVENTAIRE
// ============================================================

type TierSlot struct {
	Name string
	Slot string
}

// Equipment conserve les équipements disponibles dans la forge.
// Il est utilisé par l'inventaire pour reconnaître un équipement.
var Equipment = []string{
	"Casque renforcé",
	"Gilet pare-balles",
	"Bottes tactiques",
	"Casque militaire",
	"Armure militaire",
	"Bottes militaires",
}

var tierSlots = map[string]TierSlot{
	"Casque renforcé":   {Name: "Casque renforcé", Slot: "head"},
	"Gilet pare-balles": {Name: "Gilet pare-balles", Slot: "torso"},
	"Bottes tactiques":  {Name: "Bottes tactiques", Slot: "feet"},
	"Casque militaire":  {Name: "Casque militaire", Slot: "head"},
	"Armure militaire":  {Name: "Armure militaire", Slot: "torso"},
	"Bottes militaires": {Name: "Bottes militaires", Slot: "feet"},
}

func FindTierSlot(item string) (*TierSlot, bool) {
	tier, ok := tierSlots[item]
	if !ok {
		return nil, false
	}
	return &tier, true
}

func EquipItem(c *classes.Classe, item string) bool {
	tier, ok := FindTierSlot(item)
	if !ok {
		return false
	}

	switch tier.Slot {
	case "head":
		c.Equip.Head = item
	case "torso":
		c.Equip.Torso = item
	case "feet":
		c.Equip.Feet = item
	default:
		return false
	}

	classes.RemoveInventory(c, item)
	fmt.Printf("%s%s équipé !%s\n", classes.BrightGreen, item, classes.Reset)
	return true
}

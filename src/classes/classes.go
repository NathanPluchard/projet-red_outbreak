package classes

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Reader = bufio.NewReader(os.Stdin)

type Equipement struct {
	Head  string
	Torso string
	Feet  string
}

type Classe struct {
	Nom               string
	Description       string
	Type              string
	Specialite        string
	PVBase            int
	AttaqueBase       int
	DefenseBase       int
	ManaMax           int
	PV                int
	Attaque           int
	Defense           int
	Mana              int
	Level             int
	Exp               int
	MaxExp            int
	Gold              int
	Inventory         []string
	MaxInventory      int
	InventoryUpgrades int
	Attacks           []string
	Equip             Equipement
	Initiative        int
}

var Classes = map[string]Classe{
	"Survivant": {
		Type: "Survivant", Nom: "Survivant", Description: "Généraliste équilibré, doué pour durer sur la longueur.",
		PVBase: 200, Gold: 15, Exp: 0, MaxExp: 100, AttaqueBase: 8, DefenseBase: 8,
		Inventory: []string{}, MaxInventory: 10, Mana: 30, ManaMax: 60, Specialite: "Résistance équilibrée",
	},
	"Punk": {
		Type: "Punk", Nom: "Punk", Description: "Combattant agressif et instable, frappe fort mais encaisse mal.",
		PVBase: 150, Gold: 40, Exp: 0, MaxExp: 300, AttaqueBase: 13, DefenseBase: 4,
		Inventory: []string{}, MaxInventory: 10, Mana: 80, ManaMax: 160, Specialite: "Bonus de dégâts critiques",
	},
	"Médecin de fortune": {
		Type: "Médecin de fortune", Nom: "Médecin de fortune", Description: "Support de l'équipe, moins offensif mais indispensable en soin.",
		PVBase: 100, Gold: 20, Exp: 0, MaxExp: 50, AttaqueBase: 6, DefenseBase: 6,
		Inventory: []string{}, MaxInventory: 10, Mana: 10, ManaMax: 20, Specialite: "Efficacité de soin augmentée",
	},
	"Sauveur": {
		Type: "Sauveur", Nom: "Sauveur", Description: "Protecteur robuste, pensé pour encaisser et défendre le groupe.",
		PVBase: 80, Gold: 30, Exp: 0, MaxExp: 200, AttaqueBase: 9, DefenseBase: 9,
		Inventory: []string{}, MaxInventory: 10, Mana: 50, ManaMax: 100, Specialite: "Réduction de dégâts subis",
	},
}

func NewClasse(nom string) (*Classe, error) {
	template, ok := Classes[nom]
	if !ok {
		return nil, fmt.Errorf("classe inconnue : %q", nom)
	}

	c := template
	c.PV = c.PVBase
	c.Attaque = c.AttaqueBase
	c.Defense = c.DefenseBase
	c.Level = 1
	c.Inventory = make([]string, 0, c.MaxInventory)
	c.Attacks = append([]string(nil), template.Attacks...)

	return &c, nil
}

func ReadInt() int {
	input, _ := Reader.ReadString('\n')
	input = strings.TrimSpace(input)
	n, err := strconv.Atoi(input)
	if err != nil {
		return -1
	}
	return n
}

func ReadIntPrompt(prompt string) int {
	for {
		fmt.Print(prompt)
		input, _ := Reader.ReadString('\n')
		input = strings.TrimSpace(input)
		n, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Entrée invalide, merci de saisir un nombre.")
			continue
		}
		return n
	}
}

func CanAddItem(c *Classe) bool {
	return len(c.Inventory) < c.MaxInventory
}

func AddInventory(c *Classe, item string) bool {
	if !CanAddItem(c) {
		fmt.Println("Inventaire plein ! Impossible d'ajouter :", item)
		return false
	}
	c.Inventory = append(c.Inventory, item)
	fmt.Println(item, "ajouté à l'inventaire.")
	return true
}

func RemoveInventory(c *Classe, item string) bool {
	for i, it := range c.Inventory {
		if it == item {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return true
		}
	}
	return false
}

func CountItem(c *Classe, item string) int {
	count := 0
	for _, it := range c.Inventory {
		if it == item {
			count++
		}
	}
	return count
}

func DisplayInfo(c *Classe) {

	titre := "☣ FICHE DE " + strings.ToUpper(c.Nom) + " ☣"

	fmt.Println(TitleBox(titre))

	fmt.Println()

	fmt.Println(
		Panel(
			"IDENTITÉ",

			StatLine("Nom", BrightWhite+c.Nom+Reset),
			StatLine("Classe", BrightCyan+c.Type+Reset),
			StatLine("Niveau", BrightYellow+fmt.Sprintf("%d", c.Level)+Reset),
			StatLine("Or", BrightYellow+fmt.Sprintf("%d 💰", c.Gold)+Reset),
		),
	)

	fmt.Println()

	fmt.Println(
		Panel(
			"RESSOURCES",

			fmt.Sprintf("PV    %s %d/%d", HPBar(c.PV, c.PVBase), c.PV, c.PVBase),
			fmt.Sprintf("Mana  %s %d/%d", ManaBar(c.Mana, c.ManaMax), c.Mana, c.ManaMax),
			fmt.Sprintf("Exp   %s %d/%d", ExpBar(c.Exp, c.MaxExp), c.Exp, c.MaxExp),
		),
	)

	fmt.Println()

	fmt.Println(
		Panel(
			"ÉQUIPEMENT",

			StatLine("Casque", displayOrEmpty(c.Equip.Head)),
			StatLine("Torse", displayOrEmpty(c.Equip.Torso)),
			StatLine("Pieds", displayOrEmpty(c.Equip.Feet)),
		),
	)

	fmt.Println()

	attaques := strings.Join(c.Attacks, ", ")

	if attaques == "" {
		attaques = Dim + "Aucune" + Reset
	} else {
		attaques = Green + attaques + Reset
	}

	fmt.Println(
		Panel(
			"ATTAQUES CONNUES",
			attaques,
		),
	)

}

func displayOrEmpty(s string) string {
	if s == "" {
		return Dim + "Aucun" + Reset
	}
	return BrightWhite + s + Reset
}
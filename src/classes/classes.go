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
	PVBase            int
	PV                int
	Gold              int
	Exp               int
	MaxExp            int
	AttaqueBase       int
	DefenseBase       int
	Inventory         []string
	MaxInventory      int
	InventoryUpgrades int
	Mana              int
	ManaMax           int
	Level             int
	Attacks           []string
	Equip             Equipement
	Initiative        int
	Specialite        string
	Type              string
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

func ReadInt() int {
	input, _ := Reader.ReadString('\n')
	input = strings.TrimSpace(input)
	n, err := strconv.Atoi(input)
	if err != nil {
		return -1
	}
	return n
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
	fmt.Println(TitleBox("FICHE DE SURVIVANT"))
	fmt.Printf("%sNom       :%s %s%s%s\n", Bold, Reset, BrightWhite, c.Nom, Reset)
	fmt.Printf("%sClasse    :%s %s%s%s\n", Bold, Reset, BrightCyan, c.Type, Reset)
	fmt.Printf("%sNiveau    :%s %s%d%s\n", Bold, Reset, BrightYellow, c.Level, Reset)
	fmt.Printf("%sPV        :%s %s %d/%d\n", Bold, Reset, HPBar(c.PV, c.PVBase), c.PV, c.PVBase)
	fmt.Printf("%sMana      :%s %s %d/%d\n", Bold, Reset, ManaBar(c.Mana, c.ManaMax), c.Mana, c.ManaMax)
	fmt.Printf("%sExp       :%s %s %d/%d\n", Bold, Reset, ExpBar(c.Exp, c.MaxExp), c.Exp, c.MaxExp)
	fmt.Printf("%sOr        :%s %s%d 💰%s\n", Bold, Reset, BrightYellow, c.Gold, Reset)
	fmt.Printf("%sAttaques  :%s %s%s%s\n", Bold, Reset, Green, strings.Join(c.Attacks, ", "), Reset)
	fmt.Printf("%sCasque    :%s %s\n", Bold, Reset, displayOrEmpty(c.Equip.Head))
	fmt.Printf("%sTorse     :%s %s\n", Bold, Reset, displayOrEmpty(c.Equip.Torso))
	fmt.Printf("%sPieds     :%s %s\n", Bold, Reset, displayOrEmpty(c.Equip.Feet))
	fmt.Println(Separator())
}

func displayOrEmpty(s string) string {
	if s == "" {
		return Dim + "Aucun" + Reset
	}
	return BrightWhite + s + Reset
}
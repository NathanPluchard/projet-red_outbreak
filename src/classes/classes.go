package classes

type Classe struct {
	Nom         string
	Description string
	PVBase      int
	MaxHP       int
	Gold 		int
	Exp 		int
	MaxExp 		int 
	AttaqueBase int
	DefenseBase int
	Inventory          []string
	MaxInventory       int
	Specialite  string // ce qui rend la classe unique
}

var Classes = map[string]Classe{
	"Survivant": {
		Nom:         "Survivant",
		Description: "Généraliste équilibré, doué pour durer sur la longueur.",
		PVBase:      200,
		MaxHP:       200,
		Gold:         15,
		Exp:		 0,
		MaxExp:      100,
		AttaqueBase: 8,
		DefenseBase: 8,
		Inventory:    [],
		MaxInventory:  10,
		Specialite:  "Résistance équilibrée",
	},
	"Punk": {
		Nom:         "Punk",
		Description: "Combattant agressif et instable, frappe fort mais encaisse mal.",
		PVBase:      150,
		MaxHP:       150,
		Gold:        40,
		Exp:         0,
		MaxExp:      300,
		AttaqueBase: 13,
		DefenseBase: 4,
		Inventory:    [],
		MaxInventory:  10,
		Specialite:  "Bonus de dégâts critiques",
	},
	"Médecin de fortune": {
		Nom:         "Médecin de fortune",
		Description: "Support de l'équipe, moins offensif mais indispensable en soin.",
		PVBase:      100,
		MaxHP:       100,
		Gold:         20,
		Exp:          0,
		MaxExp:       50,
		AttaqueBase: 6,
		DefenseBase: 6,
		Inventory:    [],
		MaxInventory:  10,
		Specialite:  "Efficacité de soin augmentée",
	},
	"Sauveur": {
		Nom:         "Sauveur",
		Description: "Protecteur robuste, pensé pour encaisser et défendre le groupe.",
		PVBase:      80,
		MaxHP:       80,
		Gold:         30,
		Exp:          0,
		MaxExp:       200,
		AttaqueBase: 9,
		DefenseBase: 9,
		Inventory:    [],
		MaxInventory:  10,
		Specialite:  "Réduction de dégâts subis par l'équipe",
	},
}

<<<<<<< HEAD
=======

func trouverPotion(nom string) (potion.Potion, bool) {
	for _, p := range potion.PotionsDeVieDisponibles {
		if p.Nom == nom {
			return p, true
		}
	}
	for _, p := range potion.PotionsDePoisonDisponibles {
		if p.Nom == nom {
			return p, true
		}
	}
	return potion.Potion{}, false
}


func (c *Classe) UtiliserPotion(nomPotion string) {
	
	index := -1
	for i, nom := range c.Inventory {
		if nom == nomPotion {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Printf("%s n'a pas de %s dans son inventaire.\n", c.Nom, nomPotion)
		return
	}

	
	p, trouve := trouverPotion(nomPotion)
	if !trouve {
		fmt.Printf("Potion inconnue : %s\n", nomPotion)
		return
	}

	
	switch p.Type {
	case "vie":
		c.PV += p.Effet
		if c.PV > c.PVBase {
			c.PV = c.PVBase
		}
		fmt.Printf("%s utilise %s et regagne %d PV ! (PV: %d/%d)\n",
			c.Nom, p.Nom, p.Effet, c.PV, c.PVBase)

	case "poison":
		c.PV -= p.Effet
		if c.PV < 0 {
			c.PV = 0
		}
		fmt.Printf("%s subit %s et perd %d PV ! (PV: %d/%d)\n",
			c.Nom, p.Nom, p.Effet, c.PV, c.PVBase)
	}


	c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
}
>>>>>>> dd8218059ba14bba7b12a82c74841a9837625ad8

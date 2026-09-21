package classes

type Classe struct {
	Nom         string
	Description string
	PVBase      int
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

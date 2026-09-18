package classes

type Classe struct {
	Nom         string
	Description string
	PVBase      int
	AttaqueBase int
	DefenseBase int
	Specialite  string // ce qui rend la classe unique
}

var Classes = map[string]Classe{
	"Survivant": {
		Nom:         "Survivant",
		Description: "Généraliste équilibré, doué pour durer sur la longueur.",
		PVBase:      100,
		AttaqueBase: 8,
		DefenseBase: 8,
		Specialite:  "Résistance équilibrée",
	},
	"Punk": {
		Nom:         "Punk",
		Description: "Combattant agressif et instable, frappe fort mais encaisse mal.",
		PVBase:      90,
		AttaqueBase: 13,
		DefenseBase: 4,
		Specialite:  "Bonus de dégâts critiques",
	},
	"Médecin de fortune": {
		Nom:         "Médecin de fortune",
		Description: "Support de l'équipe, moins offensif mais indispensable en soin.",
		PVBase:      80,
		AttaqueBase: 6,
		DefenseBase: 6,
		Specialite:  "Efficacité de soin augmentée",
	},
	"Sauveur": {
		Nom:         "Sauveur",
		Description: "Protecteur robuste, pensé pour encaisser et défendre le groupe.",
		PVBase:      110,
		AttaqueBase: 9,
		DefenseBase: 9,
		Specialite:  "Réduction de dégâts subis par l'équipe",
	},
}

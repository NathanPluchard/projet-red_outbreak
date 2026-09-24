package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"outbreak/casino"
	"outbreak/classes"
)

type Weapon struct {
	Name   string
	Damage int
	Crit   int
	Price  int
	Scrap  int
	Level  int
}

type Enemy struct {
	Name    string
	HP      int
	MaxHP   int
	Attack  int
	Defense int
	XP      int
	Gold    int
	Scrap   int
	Boss    bool
}

type Zone struct {
	Name    string
	Icon    string
	Level   int
	Danger  int
	Waves   int
	Reward  int
	Enemies []Enemy
	Boss    Enemy
}

type Survivor struct {
	Name        string
	Role        string
	Level       int
	Description string
}

type Game struct {
	Player         *classes.Classe
	Weapon         Weapon
	Scrap          int
	Components     int
	Medicine       int
	Fuel           int
	Food           int
	Day            int
	Night          bool
	RefugeLevel    int
	WorkshopLevel  int
	GeneratorLevel int
	StorageLevel   int
	DefenseLevel   int
	Reputation     int
	Survivors      []Survivor
	Unlocked       map[string]bool
	Cleared        map[string]bool
	WeaponLevels   map[string]int
	MissionsDone   int
	Threat         int
	Morale         int
	DayObjective   string
	ObjectiveDone  bool
}

var weapons = []Weapon{
	{"Couteau de survie", 8, 5, 20, 0, 1},
	{"Pistolet M9", 16, 10, 90, 8, 2},
	{"Shotgun", 28, 8, 190, 18, 4},
	{"SMG Viper", 24, 14, 280, 25, 5},
	{"AK-47", 38, 15, 480, 35, 7},
	{"Fusil de précision", 62, 25, 700, 50, 10},
	{"Lance-roquettes", 100, 20, 1300, 80, 14},
}

var zones = []Zone{
	{
		"Centre-ville", "🏙️", 1, 25, 4, 100,
		[]Enemy{{"Rôdeur", 45, 45, 7, 1, 25, 12, 4, false}, {"Coureur", 65, 65, 11, 2, 35, 18, 6, false}, {"Pillard infecté", 75, 75, 10, 3, 45, 25, 8, false}},
		Enemy{"Mutant du centre", 260, 260, 24, 5, 180, 120, 30, true},
	},
	{
		"Hôpital abandonné", "🏥", 3, 40, 5, 200,
		[]Enemy{{"Infecté médical", 80, 80, 13, 3, 45, 22, 8, false}, {"Coureur malade", 95, 95, 17, 2, 55, 28, 10, false}, {"Infirmier infecté", 120, 120, 15, 6, 70, 35, 14, false}},
		Enemy{"Chirurgien infecté", 420, 420, 30, 8, 300, 220, 45, true},
	},
	{
		"Usine", "🏭", 6, 55, 6, 340,
		[]Enemy{{"Ouvrier infecté", 120, 120, 20, 6, 70, 35, 14, false}, {"Brute", 180, 180, 25, 9, 100, 50, 20, false}, {"Mutant toxique", 150, 150, 23, 5, 110, 60, 22, false}},
		Enemy{"Titan industriel", 600, 600, 38, 12, 500, 400, 70, true},
	},
	{
		"Base militaire", "🪖", 9, 70, 7, 520,
		[]Enemy{{"Soldat contaminé", 170, 170, 28, 10, 110, 65, 20, false}, {"Lourd infecté", 240, 240, 32, 14, 150, 80, 28, false}, {"Traqueur", 130, 130, 35, 7, 135, 75, 24, false}},
		Enemy{"Commandant Zéro", 800, 800, 45, 16, 750, 650, 100, true},
	},
	{
		"Laboratoire", "🧪", 13, 85, 8, 850,
		[]Enemy{{"Mutant alpha", 240, 240, 36, 12, 180, 100, 30, false}, {"Abomination", 320, 320, 42, 18, 240, 140, 40, false}, {"Prédateur", 200, 200, 50, 10, 220, 130, 35, false}},
		Enemy{"Roi des infectés", 1200, 1200, 58, 20, 1500, 1200, 180, true},
	},
}

var survivors = []Survivor{
	{"Milo", "Mécanicien", 1, "Réduit le coût en ferraille des améliorations."},
	{"Sarah", "Médecin", 1, "Améliore les soins et produit des médicaments."},
	{"Nora", "Éclaireuse", 1, "Augmente les chances de trouver du butin."},
	{"Eli", "Gardien", 1, "Renforce les défenses du refuge."},
}

func main() {
	rand.Seed(time.Now().UnixNano())
	c := characterCreation()
	g := &Game{
		Player:         &c,
		Weapon:         weapons[0],
		Scrap:          12,
		Components:     2,
		Medicine:       3,
		Fuel:           8,
		Food:           8,
		Day:            1,
		RefugeLevel:    1,
		WorkshopLevel:  1,
		GeneratorLevel: 1,
		StorageLevel:   1,
		DefenseLevel:   1,
		Reputation:     0,
		Survivors:      []Survivor{},
		Unlocked:       map[string]bool{zones[0].Name: true},
		Cleared:        map[string]bool{},
		WeaponLevels:   map[string]int{},
	}
	g.Player.WeaponName = g.Weapon.Name
	g.Player.WeaponDamage = g.Weapon.Damage
	g.Player.WeaponCrit = g.Weapon.Crit
	g.loop()
}

func characterCreation() classes.Classe {
	classes.ClearScreen()
	fmt.Println(classes.Logo())
	fmt.Println(classes.TitleBox("☣ CRÉATION DU SURVIVANT ☣"))
	var name string
	for {
		fmt.Print("\nNom > ")
		s, _ := classes.Reader.ReadString('\n')
		s = strings.TrimSpace(s)
		if validName(s) {
			name = formatName(s)
			break
		}
		fmt.Println(classes.BrightRed + "Nom invalide. Utilise uniquement des lettres." + classes.Reset)
	}
	fmt.Println(classes.Panel("CLASSES", "[1] Survivant — équilibré", "[2] Punk — critique", "[3] Médecin de fortune — soins", "[4] Sauveur — défense"))
	kind := ""
	for kind == "" {
		switch classes.ReadInt() {
		case 1:
			kind = "Survivant"
		case 2:
			kind = "Punk"
		case 3:
			kind = "Médecin de fortune"
		case 4:
			kind = "Sauveur"
		}
		if kind == "" {
			fmt.Print("Classe > ")
		}
	}
	c := classes.Classes[kind]
	c.Nom, c.Type = name, kind
	c.Level, c.PV, c.Mana = 1, c.PVBase, c.ManaMax
	c.Attacks = []string{"Attaque basique"}
	c.Inventory = []string{"Bandage"}
	c.SkillPoints = 1
	return c
}

func validName(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
			return false
		}
	}
	return true
}
func formatName(s string) string {
	r := []rune(strings.ToLower(s))
	if len(r) == 0 {
		return s
	}
	r[0] = []rune(strings.ToUpper(string(r[0])))[0]
	return string(r)
}

func (g *Game) loop() {
	for {
		classes.ClearScreen()
		fmt.Println(classes.Logo())
		fmt.Println(classes.TitleBox("☣ OUTBREAK — REFUGE DES SURVIVANTS ☣"))
		fmt.Println(classes.Panel("ÉTAT DU REFUGE",
			fmt.Sprintf("Jour %d — %s", g.Day, ternary(g.Night, "🌙 NUIT", "☀️ JOUR")),
			fmt.Sprintf("❤️ %d/%d PV   ⭐ Niveau %d   💰 %d", g.Player.PV, g.Player.PVBase, g.Player.Level, g.Player.Gold),
			fmt.Sprintf("🔩 %d   ⚙️ %d   💊 %d   ⛽ %d   🍖 %d", g.Scrap, g.Components, g.Medicine, g.Fuel, g.Food),
			fmt.Sprintf("🏚️ Refuge %d   🔨 Atelier %d   ⚡ Générateur %d", g.RefugeLevel, g.WorkshopLevel, g.GeneratorLevel),
			fmt.Sprintf("☣️ Menace %d%%   🙂 Moral %d%%   ⭐ Réputation %d (%s)", g.Threat, g.Morale, g.Reputation, g.reputationRank()),
		))
		fmt.Println(classes.Panel("ACTIONS",
			"[1] 🗺️ Partir en expédition",
			"[2] 🎒 Inventaire & équipement",
			"[3] 🛒 Marché noir",
			"[4] 🔨 Atelier & fabrication",
			"[5] 🏚️ Gérer le refuge",
			"[6] 👥 Survivants",
			"[7] 📻 Missions & événements",
			"[8] 👤 Fiche du personnage",
			"[9] 📋 Journal de survie",
			"[10] 🎰 Casino de la Zone Morte",
			"[0] Quitter",
		))
		fmt.Print("\nAction > ")
		switch classes.ReadInt() {
		case 1:
			g.expeditionMenu()
		case 2:
			g.inventoryMenu()
		case 3:
			g.market()
		case 4:
			g.workshop()
		case 5:
			g.refuge()
		case 6:
			g.survivorMenu()
		case 7:
			g.missions()
		case 8:
			classes.ClearScreen()
			classes.DisplayInfo(g.Player)
			classes.Pause()
		case 9:
			g.journal()
		case 10:
			casino.Casino(g.Player)
		case 0:
			fmt.Println("À bientôt, survivant.")
			return
		}
	}
}

func (g *Game) expeditionMenu() {
	for {
		classes.ClearScreen()
		fmt.Println(classes.TitleBox("🗺️ CARTE DE LA ZONE MORTE"))
		for i, z := range zones {
			unlocked := g.Unlocked[z.Name]
			status := "🔒"
			if unlocked {
				status = "☣️"
			}
			if g.Cleared[z.Name] {
				status = "✅"
			}
			fmt.Printf("[%d] %s %-22s Niveau %-2d Danger %-3d%% %s\n", i+1, z.Icon, z.Name, z.Level, z.Danger, status)
		}
		fmt.Println("[0] Retour")
		x := classes.ReadInt()
		if x == 0 {
			return
		}
		if x < 1 || x > len(zones) {
			continue
		}
		z := zones[x-1]
		if !g.Unlocked[z.Name] || g.Player.Level < z.Level {
			fmt.Println("Zone verrouillée.")
			classes.Pause()
			continue
		}
		g.runZone(z)
	}
}

func (g *Game) runZone(z Zone) {
	classes.ClearScreen()
	fmt.Println(classes.TitleBox(z.Icon + " EXPÉDITION — " + z.Name))
	fmt.Println(classes.Panel("BRIEFING", fmt.Sprintf("Niveau recommandé : %d", z.Level), fmt.Sprintf("Danger : %d%%", z.Danger), fmt.Sprintf("Vagues : %d + boss", z.Waves), "Objectif : survivre et ramener des ressources."))
	fmt.Println("[1] Partir   [2] Annuler")
	if classes.ReadInt() != 1 {
		return
	}
	g.Player.PV = g.Player.PVBase
	g.Player.Mana = g.Player.ManaMax
	for w := 1; w <= z.Waves; w++ {
		enemies := g.wave(z, w)
		if !g.combat(enemies, w, z.Name, false) {
			g.Threat += 4
			g.Morale -= 4
			return
		}
		g.scavenge(z, w)
		g.recoverBetween()
		if w < z.Waves && rand.Intn(100) < 45 {
			if !g.event() {
				return
			}
		}
	}
	boss := z.Boss
	boss.HP = boss.MaxHP + g.Player.Level*12
	boss.MaxHP = boss.HP
	boss.Attack += g.Player.Level / 2
	if !g.combat([]Enemy{boss}, z.Waves, z.Name, true) {
		return
	}
	g.Cleared[z.Name] = true
	g.Player.Gold += z.Reward
	g.Scrap += z.Waves*5 + boss.Scrap
	gainXP(g.Player, z.Reward/2)
	g.reputation(2)
	g.Threat -= 8
	if g.Threat < 5 {
		g.Threat = 5
	}
	g.Morale += 8
	if g.Morale > 100 {
		g.Morale = 100
	}
	g.ObjectiveDone = true
	g.unlockNext(z)
	fmt.Println(classes.Panel("🏆 EXPÉDITION RÉUSSIE", fmt.Sprintf("+%d 💰", z.Reward), fmt.Sprintf("+%d 🔩", z.Waves*5+boss.Scrap), "La zone est maintenant nettoyée."))
	g.advanceDay()
	classes.Pause()
}

func (g *Game) wave(z Zone, w int) []Enemy {
	n := 1 + (w+1)/2
	if n > 4 {
		n = 4
	}
	out := make([]Enemy, 0, n)
	for i := 0; i < n; i++ {
		e := z.Enemies[rand.Intn(len(z.Enemies))]
		scale := 1 + (w-1)/4
		e.MaxHP *= scale
		e.HP = e.MaxHP
		e.Attack += w * 2
		e.XP += w * 8
		e.Gold += w * 4
		out = append(out, e)
	}
	return out
}

func (g *Game) combat(enemies []Enemy, wave int, zone string, boss bool) bool {
	for len(enemies) > 0 && g.Player.PV > 0 {
		classes.ClearScreen()
		title := "⚔️ COMBAT"
		if boss {
			title = "👑 COMBAT DE BOSS"
		}
		fmt.Println(classes.TitleBox(title))
		fmt.Printf("%s — Vague %d | ❤️ %d/%d | Mana %d/%d | 🔫 %s\n\n", zone, wave, g.Player.PV, g.Player.PVBase, g.Player.Mana, g.Player.ManaMax, g.Weapon.Name)
		for i, e := range enemies {
			fmt.Printf("[%d] %s %d/%d HP  ATK:%d DEF:%d\n", i+1, e.Name, e.HP, e.MaxHP, e.Attack, e.Defense)
		}
		fmt.Println(classes.Panel("ACTIONS", "[1] 🔫 Tirer", "[2] ⚡ Frappe spéciale", "[3] 💊 Soin", "[4] 💣 Grenade", "[5] 🏃 Fuir"))
		action := classes.ReadInt()
		if action == 5 && !boss {
			return false
		}
		if action == 5 {
			fmt.Println("Impossible de fuir un boss.")
			classes.Pause()
			continue
		}
		target := 0
		if len(enemies) > 1 {
			fmt.Print("Cible > ")
			target = classes.ReadInt() - 1
			if target < 0 || target >= len(enemies) {
				continue
			}
		}
		switch action {
		case 1:
			g.attack(&enemies[target])
		case 2:
			g.skill(&enemies[target])
		case 3:
			g.heal()
		case 4:
			g.grenade(enemies)
		default:
			continue
		}
		for i := len(enemies) - 1; i >= 0; i-- {
			if enemies[i].HP <= 0 {
				g.kill(enemies[i])
				enemies = append(enemies[:i], enemies[i+1:]...)
			}
		}
		if len(enemies) == 0 {
			break
		}
		g.enemyTurn(enemies)
	}
	return g.Player.PV > 0
}

func (g *Game) attack(e *Enemy) {
	d := g.Weapon.Damage + g.Player.Attaque/2
	if rand.Intn(100) < g.Weapon.Crit {
		d *= 2
		fmt.Println("💥 COUP CRITIQUE !")
	}
	d -= e.Defense
	if d < 1 {
		d = 1
	}
	e.HP -= d
	fmt.Printf("🔫 %s inflige %d dégâts.\n", g.Weapon.Name, d)
	classes.Pause()
}
func (g *Game) skill(e *Enemy) {
	cost := 10
	if g.Player.Mana < cost {
		fmt.Println("Mana insuffisant.")
		classes.Pause()
		return
	}
	g.Player.Mana -= cost
	d := g.Player.Attaque + g.Weapon.Damage + 10 - e.Defense
	if d < 2 {
		d = 2
	}
	e.HP -= d
	fmt.Printf("⚡ Frappe spéciale : %d dégâts.\n", d)
	classes.Pause()
}
func (g *Game) heal() {
	if g.Medicine <= 0 {
		fmt.Println("Plus de médicaments.")
		classes.Pause()
		return
	}
	g.Medicine--
	h := 45
	if g.Player.Type == "Médecin de fortune" {
		h = 70
	}
	if g.SurvivorRole("Médecin") {
		h += 15
	}
	g.Player.PV += h
	if g.Player.PV > g.Player.PVBase {
		g.Player.PV = g.Player.PVBase
	}
	fmt.Printf("💊 +%d PV.\n", h)
	classes.Pause()
}
func (g *Game) grenade(es []Enemy) {
	if g.Components < 2 {
		fmt.Println("Il faut 2 composants pour fabriquer une grenade improvisée.")
		classes.Pause()
		return
	}
	g.Components -= 2
	for i := range es {
		d := 25 + rand.Intn(15) - es[i].Defense/2
		if d < 5 {
			d = 5
		}
		es[i].HP -= d
	}
	fmt.Println("💣 Explosion !")
	classes.Pause()
}
func (g *Game) enemyTurn(es []Enemy) {
	for _, e := range es {
		if e.HP <= 0 {
			continue
		}
		d := e.Attack - g.Player.Defense/2
		if g.Player.Type == "Sauveur" {
			d = d * 75 / 100
		}
		if g.SurvivorRole("Gardien") {
			d = d * 90 / 100
		}
		if d < 1 {
			d = 1
		}
		if rand.Intn(100) < 5 {
			fmt.Println(e.Name, "rate son attaque !")
			continue
		}
		g.Player.PV -= d
		fmt.Printf("🧟 %s inflige %d dégâts.\n", e.Name, d)
		if g.Player.PV <= 0 {
			g.Player.PV = 0
			fmt.Println("☠️ Vous êtes tombé.")
			classes.Pause()
			return
		}
	}
}
func (g *Game) kill(e Enemy) {
	g.Player.Kills++
	g.Player.Gold += e.Gold
	scrapGain := e.Scrap
	if g.SurvivorRole("Éclaireuse") && rand.Intn(100) < 35 {
		scrapGain += rand.Intn(5) + 1
	}
	g.Scrap += scrapGain
	gainXP(g.Player, e.XP)
	if e.Boss {
		g.Player.BossesDefeated++
		g.Components += rand.Intn(4) + 2
	}
	if rand.Intn(100) < 35+g.reputationBonus() {
		g.Medicine++
		fmt.Println("💊 Butin : médicament.")
	}
	if rand.Intn(100) < 30 {
		g.Components++
		fmt.Println("⚙️ Butin : composant électronique.")
	}
	fmt.Printf("☠ %s éliminé : +%d 💰 +%d 🔩\n", e.Name, e.Gold, scrapGain)
}

func gainXP(c *classes.Classe, x int) {
	c.Exp += x
	for c.Exp >= c.MaxExp {
		c.Exp -= c.MaxExp
		c.Level++
		c.MaxExp += 30
		c.PVBase += 12
		c.PV = c.PVBase
		c.ManaMax += 6
		c.Mana = c.ManaMax
		c.AttaqueBase += 2
		c.Attaque = c.AttaqueBase
		c.DefenseBase++
		c.Defense = c.DefenseBase
		c.SkillPoints++
		fmt.Printf("⭐ NIVEAU %d ! +1 point de compétence\n", c.Level)
	}
}

func (g *Game) inventoryMenu() {
	for {
		classes.ClearScreen()
		fmt.Println(classes.TitleBox("🎒 INVENTAIRE DU SURVIVANT"))
		fmt.Println(classes.Panel("RESSOURCES", fmt.Sprintf("🔩 Ferraille : %d", g.Scrap), fmt.Sprintf("⚙️ Composants : %d", g.Components), fmt.Sprintf("💊 Médicaments : %d", g.Medicine), fmt.Sprintf("⛽ Carburant : %d", g.Fuel), fmt.Sprintf("🍖 Nourriture : %d", g.Food)))
		fmt.Println(classes.Panel("ÉQUIPEMENT", fmt.Sprintf("🔫 %s — %d dégâts — %d%% critique", g.Weapon.Name, g.Weapon.Damage, g.Weapon.Crit), fmt.Sprintf("🎒 Objets : %d/%d", len(g.Player.Inventory), g.Player.MaxInventory)))
		fmt.Println("[1] Utiliser un objet  [2] Vendre un objet  [3] Armes  [0] Retour")
		switch classes.ReadInt() {
		case 1:
			g.useInventoryItem()
		case 2:
			g.sellInventoryItem()
		case 3:
			g.weaponShop()
		case 0:
			return
		}
	}
}

func (g *Game) useInventoryItem() {
	if len(g.Player.Inventory) == 0 {
		fmt.Println("Inventaire vide.")
		classes.Pause()
		return
	}
	for i, it := range g.Player.Inventory {
		fmt.Printf("[%d] %s\n", i+1, it)
	}
	fmt.Println("[0] Retour")
	x := classes.ReadInt()
	if x <= 0 || x > len(g.Player.Inventory) {
		return
	}
	it := g.Player.Inventory[x-1]
	switch it {
	case "Bandage":
		g.Player.PV += 20
		if g.Player.PV > g.Player.PVBase {
			g.Player.PV = g.Player.PVBase
		}
		classes.RemoveInventory(g.Player, it)
		fmt.Println("🩹 Bandage utilisé.")
	default:
		fmt.Println("Cet objet ne peut pas être utilisé ici.")
	}
	classes.Pause()
}
func (g *Game) sellInventoryItem() {
	if len(g.Player.Inventory) == 0 {
		fmt.Println("Inventaire vide.")
		classes.Pause()
		return
	}
	for i, it := range g.Player.Inventory {
		fmt.Printf("[%d] %s — 8 💰\n", i+1, it)
	}
	fmt.Println("[0] Retour")
	x := classes.ReadInt()
	if x <= 0 || x > len(g.Player.Inventory) {
		return
	}
	it := g.Player.Inventory[x-1]
	classes.RemoveInventory(g.Player, it)
	g.Player.Gold += 8
	fmt.Println("Vendu :", it, "pour 8 💰.")
	classes.Pause()
}

func (g *Game) weaponShop() {
	for {
		classes.ClearScreen()
		fmt.Println(classes.TitleBox("🔫 ARSENAL"))
		for i, w := range weapons {
			status := ""
			if w.Name == g.Weapon.Name {
				status = " ← ÉQUIPÉ"
			}
			fmt.Printf("[%d] %-22s Niveau %-2d Dégâts %-3d Crit %-2d%% Prix %-4d 💰%s\n", i+1, w.Name, w.Level, w.Damage, w.Crit, w.Price, status)
		}
		fmt.Println("[0] Retour")
		x := classes.ReadInt()
		if x == 0 {
			return
		}
		if x < 1 || x > len(weapons) {
			continue
		}
		w := weapons[x-1]
		if g.Player.Level < w.Level {
			fmt.Println("Niveau insuffisant.")
			classes.Pause()
			continue
		}
		if g.Player.Gold < w.Price {
			fmt.Println("Or insuffisant.")
			classes.Pause()
			continue
		}
		g.Player.Gold -= w.Price
		g.Weapon = w
		g.Player.WeaponName = w.Name
		g.Player.WeaponDamage = w.Damage
		g.Player.WeaponCrit = w.Crit
		fmt.Println("Nouvelle arme équipée :", w.Name)
		classes.Pause()
	}
}

func (g *Game) market() {
	for {
		classes.ClearScreen()
		fmt.Println(classes.TitleBox("🛒 MARCHÉ NOIR"))
		fmt.Println(classes.Panel("TON STOCK", fmt.Sprintf("💰 %d or", g.Player.Gold), fmt.Sprintf("🔩 %d ferraille", g.Scrap), fmt.Sprintf("⚙️ %d composants", g.Components)))
		fmt.Println("[1] Vendre 5 ferrailles → 15 💰", "\n[2] Vendre 1 composant → 12 💰", "\n[3] Acheter 1 médicament → 25 💰", "\n[4] Acheter 5 ferrailles → 25 💰", "\n[0] Retour")
		switch classes.ReadInt() {
		case 1:
			if g.Scrap >= 5 {
				g.Scrap -= 5
				g.Player.Gold += 15
				fmt.Println("Transaction effectuée.")
			} else {
				fmt.Println("Pas assez de ferraille.")
			}
			classes.Pause()
		case 2:
			if g.Components > 0 {
				g.Components--
				g.Player.Gold += 12
				fmt.Println("Composant vendu.")
			} else {
				fmt.Println("Aucun composant.")
			}
			classes.Pause()
		case 3:
			if g.Player.Gold >= 25 {
				g.Player.Gold -= 25
				g.Medicine++
				fmt.Println("Médicament acheté.")
			} else {
				fmt.Println("Or insuffisant.")
			}
			classes.Pause()
		case 4:
			if g.Player.Gold >= 25 {
				g.Player.Gold -= 25
				g.Scrap += 5
				fmt.Println("Ferraille achetée.")
			} else {
				fmt.Println("Or insuffisant.")
			}
			classes.Pause()
		case 0:
			return
		}
	}
}

func (g *Game) workshop() {
	for {
		classes.ClearScreen()
		fmt.Println(classes.TitleBox("🔨 ATELIER DU REFUGE"))
		fmt.Println(classes.Panel("NIVEAUX", fmt.Sprintf("🔨 Atelier : %d", g.WorkshopLevel), fmt.Sprintf("🔩 Ferraille : %d", g.Scrap), fmt.Sprintf("⚙️ Composants : %d", g.Components)))
		fmt.Println("[1] Améliorer l'arme", "\n[2] Fabriquer des médicaments", "\n[3] Fabriquer une grenade", "\n[4] Améliorer l'atelier", "\n[0] Retour")
		switch classes.ReadInt() {
		case 1:
			g.upgradeWeapon()
		case 2:
			if g.Scrap >= 4 && g.Components >= 1 {
				g.Scrap -= 4
				g.Components--
				g.Medicine++
				fmt.Println("💊 Médicament fabriqué.")
			} else {
				fmt.Println("Il faut 4 ferrailles et 1 composant.")
			}
			classes.Pause()
		case 3:
			if !classes.CanAddItem(g.Player) {
				fmt.Println("Inventaire plein.")
				classes.Pause()
				continue
			}
			if g.Scrap >= 3 && g.Components >= 2 {
				g.Scrap -= 3
				g.Components -= 2
				g.Player.Inventory = append(g.Player.Inventory, "Grenade improvisée")
				fmt.Println("💣 Grenade fabriquée.")
			} else {
				fmt.Println("Ressources insuffisantes.")
			}
			classes.Pause()
		case 4:
			g.upgradeWorkshop()
		case 0:
			return
		}
	}
}
func (g *Game) upgradeWeapon() {
	costS := 10*g.WeaponLevels[g.Weapon.Name] + 10
	costC := 2 + g.WeaponLevels[g.Weapon.Name]
	if g.SurvivorRole("Mécanicien") {
		costS = costS * 80 / 100
	}
	if g.Scrap < costS || g.Components < costC {
		fmt.Printf("Il faut %d 🔩 et %d ⚙️.\n", costS, costC)
		classes.Pause()
		return
	}
	g.Scrap -= costS
	g.Components -= costC
	g.WeaponLevels[g.Weapon.Name]++
	g.Weapon.Damage += 4
	g.Player.WeaponDamage = g.Weapon.Damage
	fmt.Printf("🔧 %s améliorée ! Dégâts : %d\n", g.Weapon.Name, g.Weapon.Damage)
	classes.Pause()
}
func (g *Game) upgradeWorkshop() {
	costS := 30 * g.WorkshopLevel
	costC := 5 * g.WorkshopLevel
	if g.Scrap < costS || g.Components < costC {
		fmt.Printf("Il faut %d 🔩 et %d ⚙️.\n", costS, costC)
		classes.Pause()
		return
	}
	g.Scrap -= costS
	g.Components -= costC
	g.WorkshopLevel++
	fmt.Println("🔨 Atelier amélioré au niveau", g.WorkshopLevel)
	classes.Pause()
}

func (g *Game) refuge() {
	for {
		classes.ClearScreen()
		fmt.Println(classes.TitleBox("🏚️ GESTION DU REFUGE"))
		fmt.Println(classes.Panel("INSTALLATIONS", fmt.Sprintf("🏚️ Refuge %d", g.RefugeLevel), fmt.Sprintf("⚡ Générateur %d", g.GeneratorLevel), fmt.Sprintf("📦 Stockage %d", g.StorageLevel), fmt.Sprintf("🛡️ Défenses %d", g.DefenseLevel)))
		fmt.Println("[1] Améliorer le générateur", "\n[2] Améliorer le stockage", "\n[3] Renforcer les défenses", "\n[4] Améliorer le refuge", "\n[5] Dormir jusqu'au lendemain", "\n[0] Retour")
		switch classes.ReadInt() {
		case 1:
			g.upgradeRefugePart("generator")
		case 2:
			g.upgradeRefugePart("storage")
		case 3:
			g.upgradeRefugePart("defense")
		case 4:
			g.upgradeRefugePart("refuge")
		case 5:
			g.sleep()
		case 6:
			g.barricade()
		case 0:
			return
		}
	}
}
func (g *Game) upgradeRefugePart(kind string) {
	lvl := 1
	name := ""
	switch kind {
	case "generator":
		lvl = g.GeneratorLevel
		name = "générateur"
	case "storage":
		lvl = g.StorageLevel
		name = "stockage"
	case "defense":
		lvl = g.DefenseLevel
		name = "défenses"
	case "refuge":
		lvl = g.RefugeLevel
		name = "refuge"
	}
	cs := 25 * lvl
	cc := 3 * lvl
	if g.SurvivorRole("Mécanicien") {
		cs -= 5
		if cs < 5 {
			cs = 5
		}
	}
	if g.Scrap < cs || g.Components < cc {
		fmt.Printf("Il faut %d 🔩 et %d ⚙️ pour améliorer le %s.\n", cs, cc, name)
		classes.Pause()
		return
	}
	g.Scrap -= cs
	g.Components -= cc
	switch kind {
	case "generator":
		g.GeneratorLevel++
		g.Fuel += 4
	case "storage":
		g.StorageLevel++
		g.Player.MaxInventory += 3
	case "defense":
		g.DefenseLevel++
	case "refuge":
		g.RefugeLevel++
	}
	fmt.Println("🏗️", name, "amélioré !")
	classes.Pause()
}

func (g *Game) survivorMenu() {
	for {
		classes.ClearScreen()
		fmt.Println(classes.TitleBox("👥 SURVIVANTS"))
		fmt.Printf("Réputation : %d\n\n", g.Reputation)
		if len(g.Survivors) == 0 {
			fmt.Println("Personne n'a encore rejoint votre refuge.")
		} else {
			for _, s := range g.Survivors {
				fmt.Printf("• %s — %s — niveau %d\n  %s\n", s.Name, s.Role, s.Level, s.Description)
			}
		}
		fmt.Println("\n[1] Recruter un survivant", "\n[2] Entraîner les survivants", "\n[0] Retour")
		switch classes.ReadInt() {
		case 1:
			g.recruit()
		case 2:
			g.trainSurvivors()
		case 0:
			return
		}
	}
}
func (g *Game) recruit() {
	if len(g.Survivors) >= 4 {
		fmt.Println("Votre refuge ne peut plus accueillir de survivants.")
		classes.Pause()
		return
	}
	available := []Survivor{}
	for _, s := range survivors {
		found := false
		for _, x := range g.Survivors {
			if x.Name == s.Name {
				found = true
			}
		}
		if !found {
			available = append(available, s)
		}
	}
	if len(available) == 0 {
		fmt.Println("Tous les survivants connus vous ont rejoint.")
		classes.Pause()
		return
	}
	s := available[rand.Intn(len(available))]
	cost := 20 + len(g.Survivors)*15
	if g.Player.Gold < cost {
		fmt.Printf("%s demande %d 💰.\n", s.Name, cost)
		classes.Pause()
		return
	}
	g.Player.Gold -= cost
	g.Survivors = append(g.Survivors, s)
	g.Reputation++
	fmt.Printf("👥 %s rejoint votre refuge !\n", s.Name)
	classes.Pause()
}
func (g *Game) trainSurvivors() {
	if len(g.Survivors) == 0 {
		fmt.Println("Aucun survivant à entraîner.")
		classes.Pause()
		return
	}
	cost := 20 + g.RefugeLevel*10
	if g.Scrap < cost {
		fmt.Println("Ferraille insuffisante.")
		classes.Pause()
		return
	}
	g.Scrap -= cost
	for i := range g.Survivors {
		g.Survivors[i].Level++
	}
	fmt.Println("👥 Votre communauté progresse.")
	classes.Pause()
}

func (g *Game) missions() {
	for {
		classes.ClearScreen()
		fmt.Println(classes.TitleBox("📻 MISSIONS & RADIO"))
		fmt.Println(classes.Panel("OBJECTIF DU JOUR", fmt.Sprintf("Jour %d", g.Day), g.DayObjective, fmt.Sprintf("État : %s", ternary(g.ObjectiveDone, "✓ TERMINÉ", "EN COURS"))))
		fmt.Println("[1] Écouter la radio", "\n[2] Faire une mission de ravitaillement", "\n[3] Voir les statistiques", "\n[0] Retour")
		switch classes.ReadInt() {
		case 1:
			g.radio()
		case 2:
			g.supplyMission()
		case 3:
			g.stats()
		case 0:
			return
		}
	}
}
func (g *Game) radio() {
	events := []string{"📻 Un survivant parle d'un convoi abandonné.", "📻 Une voix inconnue transmet les coordonnées d'un laboratoire.", "📻 Quelqu'un demande des médicaments.", "📻 Une station annonce une pénurie de carburant."}
	fmt.Println(events[rand.Intn(len(events))])
	g.Reputation++
	g.Morale += 2
	if g.Morale > 100 {
		g.Morale = 100
	}
	classes.Pause()
}
func (g *Game) supplyMission() {
	cost := 2
	if g.Fuel < cost {
		fmt.Println("Pas assez de carburant.")
		classes.Pause()
		return
	}
	g.Fuel -= cost
	fmt.Println("🚚 Mission de ravitaillement...")
	time.Sleep(500 * time.Millisecond)
	gain := rand.Intn(20) + 10
	g.Scrap += gain
	g.Food += rand.Intn(4) + 1
	g.MissionsDone++
	g.Morale += 5
	if g.Morale > 100 {
		g.Morale = 100
	}
	g.ObjectiveDone = true
	fmt.Printf("Mission réussie : +%d 🔩 et nourriture récupérée.\n", gain)
	classes.Pause()
}
func (g *Game) stats() {
	fmt.Println(classes.Panel("STATISTIQUES", fmt.Sprintf("☠️ Infectés éliminés : %d", g.Player.Kills), fmt.Sprintf("👑 Boss vaincus : %d", g.Player.BossesDefeated), fmt.Sprintf("🗺️ Zones nettoyées : %d/%d", len(g.Cleared), len(zones)), fmt.Sprintf("👥 Survivants : %d", len(g.Survivors)), fmt.Sprintf("📻 Missions : %d", g.MissionsDone), fmt.Sprintf("⭐ Niveau : %d", g.Player.Level), fmt.Sprintf("☣️ Menace : %d%%", g.Threat), fmt.Sprintf("🙂 Moral : %d%%", g.Morale), fmt.Sprintf("🏆 Réputation : %d (%s)", g.Reputation, g.reputationRank())))
	classes.Pause()
}

func (g *Game) event() bool {
	classes.ClearScreen()
	fmt.Println(classes.TitleBox("🎲 ÉVÉNEMENT"))
	switch rand.Intn(5) {
	case 0:
		fmt.Println("🚗 Une voiture abandonnée contient du carburant.")
		g.Fuel += 3
		g.Morale += 2
	case 1:
		fmt.Println("🏚️ Une cache contient des médicaments.")
		g.Medicine += 2
		if g.SurvivorRole("Éclaireuse") {
			g.Medicine++
		}
	case 2:
		fmt.Println("📦 Des débris cachent des composants.")
		g.Components += 3
	case 3:
		fmt.Println("🔩 Un atelier abandonné contient de la ferraille.")
		g.Scrap += 15
		if g.SurvivorRole("Éclaireuse") {
			g.Scrap += 5
		}
	case 4:
		fmt.Println("☣️ Embuscade ! Vous perdez quelques PV.")
		g.Player.PV -= 15
		g.Morale -= 5
		g.Threat += 3
		if g.Player.PV < 1 {
			g.Player.PV = 1
		}
	}
	classes.Pause()
	return true
}
func (g *Game) recoverBetween() {
	heal := g.Player.PVBase / 10
	if heal < 5 {
		heal = 5
	}
	g.Player.PV += heal
	if g.Player.PV > g.Player.PVBase {
		g.Player.PV = g.Player.PVBase
	}
	g.Player.Mana = g.Player.ManaMax
}
func (g *Game) unlockNext(z Zone) {
	for i := range zones {
		if zones[i].Name == z.Name && i+1 < len(zones) {
			n := zones[i+1]
			if g.Player.Level >= n.Level || g.Cleared[z.Name] {
				g.Unlocked[n.Name] = true
				fmt.Println("🔓 Nouvelle zone :", n.Name)
			}
		}
	}
}
func (g *Game) advanceDay() {
	g.Day++
	g.Night = true
	g.Threat += 2
	if g.Threat > 100 {
		g.Threat = 100
	}

	if g.Food > 0 {
		g.Food--
		g.Morale += 2
	} else {
		g.Morale -= 12
		fmt.Println("🍖 Le refuge manque de nourriture. Le moral chute.")
	}

	if g.Fuel > 0 {
		g.Fuel--
	} else {
		g.Player.PV -= 5
		if g.Player.PV < 1 {
			g.Player.PV = 1
		}
		fmt.Println("⚡ Le générateur est à sec. Le refuge manque d'énergie.")
	}

	if g.SurvivorRole("Médecin") && g.Medicine < 12 {
		g.Medicine++
	}
	if g.SurvivorRole("Gardien") {
		g.DefenseLevel++
		if g.DefenseLevel > 5 {
			g.DefenseLevel = 5
		}
	}
	if g.Morale < 0 {
		g.Morale = 0
	}
	if g.Morale > 100 {
		g.Morale = 100
	}

	g.nightRaid()
	g.DayObjective = randomObjective()
	g.ObjectiveDone = false
}

func (g *Game) sleep() {
	g.advanceDay()
	g.Night = false
	g.Player.PV = g.Player.PVBase
	g.Player.Mana = g.Player.ManaMax
	fmt.Printf("🌅 Jour %d. Vous êtes reposé.\n", g.Day)
	classes.Pause()
}
func (g *Game) scavenge(z Zone, wave int) {
	classes.ClearScreen()
	fmt.Println(classes.TitleBox("🔎 FOUILLE DES RUINES"))
	fmt.Println(classes.Panel("APRÈS LE COMBAT", "Les corps et les bâtiments abandonnés cachent peut-être encore quelque chose.", "Plus vous fouillez longtemps, plus le risque augmente."))
	fmt.Println("[1] Fouiller rapidement", "\n[2] Fouiller à fond", "\n[3] Ne rien risquer")
	choice := classes.ReadInt()
	if choice == 3 || choice < 1 || choice > 2 {
		return
	}

	chance := 65
	if choice == 2 {
		chance = 48
	}
	if g.SurvivorRole("Éclaireuse") {
		chance += 12
	}
	if rand.Intn(100) >= chance {
		loss := 4 + wave
		g.Player.PV -= loss
		if g.Player.PV < 1 {
			g.Player.PV = 1
		}
		g.Threat += 2
		fmt.Printf("☣️ Une embuscade ! Vous perdez %d PV.\n", loss)
		classes.Pause()
		return
	}

	mult := 1
	if choice == 2 {
		mult = 2
	}
	scrap := (rand.Intn(8) + 5) * mult
	components := rand.Intn(3)
	g.Scrap += scrap
	g.Components += components
	fmt.Printf("🔩 +%d ferraille\n", scrap)
	if components > 0 {
		fmt.Printf("⚙️ +%d composants\n", components)
	}
	if rand.Intn(100) < 25+g.reputationBonus()/2 {
		g.Fuel++
		fmt.Println("⛽ +1 carburant")
	}
	if rand.Intn(100) < 18 {
		g.Medicine++
		fmt.Println("💊 +1 médicament")
	}
	_ = z
	classes.Pause()
}

func (g *Game) nightRaid() {
	if rand.Intn(100) >= g.Threat/2 {
		return
	}
	classes.ClearScreen()
	fmt.Println(classes.TitleBox("🌙 ALERTE — ATTAQUE DU REFUGE"))
	strength := g.DefenseLevel*10 + len(g.Survivors)*4 + g.Morale/10
	danger := g.Threat/2 + rand.Intn(25)
	if g.SurvivorRole("Gardien") {
		strength += 15
	}
	if strength >= danger {
		reward := 5 + g.DefenseLevel*2
		g.Scrap += reward
		g.Reputation++
		g.Morale += 4
		if g.Morale > 100 {
			g.Morale = 100
		}
		fmt.Println("🛡️ Les défenses tiennent ! Les infectés sont repoussés.")
		fmt.Printf("🔩 Vous récupérez %d ferrailles sur les assaillants.\n", reward)
	} else {
		loss := 5 + rand.Intn(10)
		if g.Scrap >= loss {
			g.Scrap -= loss
		} else {
			g.Scrap = 0
		}
		if g.Food > 0 {
			g.Food--
		}
		g.Morale -= 10
		if g.Morale < 0 {
			g.Morale = 0
		}
		g.Threat += 5
		fmt.Println("🚨 Les défenses ont cédé pendant quelques minutes !")
		fmt.Printf("📦 Le refuge perd %d ferrailles et 1 nourriture.\n", loss)
	}
	classes.Pause()
}

func (g *Game) barricade() {
	cost := 8 + g.DefenseLevel*4
	if g.Scrap < cost {
		fmt.Printf("Il faut %d ferrailles.\n", cost)
		classes.Pause()
		return
	}
	g.Scrap -= cost
	g.DefenseLevel++
	if g.DefenseLevel > 10 {
		g.DefenseLevel = 10
	}
	g.Threat -= 3
	if g.Threat < 5 {
		g.Threat = 5
	}
	fmt.Printf("🧱 Barricades renforcées. Défense : %d.\n", g.DefenseLevel)
	classes.Pause()
}

func (g *Game) reputationRank() string {
	switch {
	case g.Reputation >= 20:
		return "Légende du refuge"
	case g.Reputation >= 12:
		return "Héros local"
	case g.Reputation >= 6:
		return "Allié fiable"
	case g.Reputation >= 2:
		return "Connu"
	default:
		return "Inconnu"
	}
}

func randomObjective() string {
	objectives := []string{
		"Nettoyer une zone infestée.",
		"Ramener au moins 20 ferrailles au refuge.",
		"Trouver 3 composants électroniques.",
		"Accomplir une mission de ravitaillement.",
		"Faire progresser la réputation du refuge.",
	}
	return objectives[rand.Intn(len(objectives))]
}

func (g *Game) journal() {
	classes.ClearScreen()
	fmt.Println(classes.TitleBox("📋 JOURNAL DE SURVIE"))
	fmt.Println(classes.Panel("SITUATION",
		fmt.Sprintf("Jour : %d", g.Day),
		fmt.Sprintf("Menace : %d%%", g.Threat),
		fmt.Sprintf("Moral : %d%%", g.Morale),
		fmt.Sprintf("Réputation : %d — %s", g.Reputation, g.reputationRank()),
		fmt.Sprintf("Objectif : %s", g.DayObjective),
	))
	fmt.Println(classes.Panel("COMMUNAUTÉ",
		fmt.Sprintf("Survivants : %d/4", len(g.Survivors)),
		fmt.Sprintf("Défense : %d", g.DefenseLevel),
		fmt.Sprintf("Zones nettoyées : %d/%d", len(g.Cleared), len(zones)),
	))
	classes.Pause()
}

func (g *Game) reputation(n int)     { g.Reputation += n }
func (g *Game) reputationBonus() int { return g.Reputation * 3 }
func (g *Game) SurvivorRole(role string) bool {
	for _, s := range g.Survivors {
		if s.Role == role {
			return true
		}
	}
	return false
}
func ternary(b bool, a, c string) string {
	if b {
		return a
	}
	return c
}

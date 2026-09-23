package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"outbreak/casino"
	"outbreak/classes"
	"outbreak/forgeron"
	"outbreak/inventaire"
	"outbreak/marchand"
	"outbreak/monstre"
)

type Weapon struct {
	Name, Rarity                       string
	Damage, Ammo, MaxAmmo, Crit, Price int
}
type Enemy struct {
	Name                                              string
	MaxHP, HP, Attack, Defense, Exp, Gold, Initiative int
	Boss                                              bool
	Phase                                             int
}
type Zone struct {
	Name, Icon           string
	Level, Waves, Reward int
	Enemies              []func() Enemy
	Boss                 func() Enemy
}

type Game struct {
	Player   *classes.Classe
	Weapon   Weapon
	Medkits  int
	Grenades int
	Scrap    int
	Unlocked map[string]bool
	Cleared  map[string]bool
	Day      int
}

var weapons = []Weapon{
	{"Couteau de survie", "Commun", 8, 0, 0, 5, 20},
	{"Pistolet M9", "Commun", 15, 12, 12, 10, 80},
	{"Shotgun", "Rare", 28, 6, 6, 8, 180},
	{"SMG Viper", "Rare", 24, 24, 24, 12, 260},
	{"AK-47", "Épique", 38, 30, 30, 15, 450},
	{"Fusil de précision", "Épique", 62, 5, 5, 25, 650},
	{"Lance-roquettes", "Légendaire", 100, 3, 3, 20, 1200},
}

func enemy(name string, hp, atk, def, exp, gold, ini int) func() Enemy {
	return func() Enemy {
		return Enemy{Name: name, MaxHP: hp, HP: hp, Attack: atk, Defense: def, Exp: exp, Gold: gold, Initiative: ini}
	}
}
func boss(name string, hp, atk, def, exp, gold int) func() Enemy {
	return func() Enemy {
		return Enemy{Name: name, MaxHP: hp, HP: hp, Attack: atk, Defense: def, Exp: exp, Gold: gold, Initiative: 5, Boss: true, Phase: 1}
	}
}

var zones = []Zone{
	{"Centre-ville", "🏙️", 1, 5, 120, []func() Enemy{enemy("Infecté errant", 45, 7, 1, 25, 12, 4), enemy("Coureur", 65, 11, 2, 35, 18, 8)}, boss("Mutant du centre", 260, 24, 5, 180, 120)},
	{"Hôpital", "🏥", 3, 6, 220, []func() Enemy{enemy("Infecté médical", 80, 13, 3, 45, 22, 5), enemy("Coureur malade", 95, 17, 2, 55, 28, 9), enemy("Infecté blindé", 125, 15, 7, 70, 35, 2)}, boss("Chirurgien infecté", 420, 30, 8, 300, 220)},
	{"Usine", "🏭", 6, 7, 360, []func() Enemy{enemy("Ouvrier infecté", 120, 20, 6, 70, 35, 4), enemy("Brute", 180, 25, 9, 100, 50, 2), enemy("Mutant toxique", 150, 23, 5, 110, 60, 6)}, boss("Titan industriel", 600, 38, 12, 500, 400)},
	{"Base militaire", "🪖", 9, 8, 550, []func() Enemy{enemy("Soldat contaminé", 170, 28, 10, 110, 65, 7), enemy("Lourd infecté", 240, 32, 14, 150, 80, 3), enemy("Traqueur", 130, 35, 7, 135, 75, 10)}, boss("Commandant zéro", 800, 45, 16, 750, 650)},
	{"Zone contaminée", "☣️", 13, 10, 900, []func() Enemy{enemy("Mutant alpha", 240, 36, 12, 180, 100, 7), enemy("Abomination", 320, 42, 18, 240, 140, 2), enemy("Prédateur", 200, 50, 10, 220, 130, 11)}, boss("Roi des infectés", 1200, 58, 20, 1500, 1200)},
}

func main() {
	rand.Seed(time.Now().UnixNano())
	c := characterCreation()
	g := &Game{Player: &c, Weapon: weapons[0], Medkits: 2, Grenades: 1, Scrap: 10, Unlocked: map[string]bool{}, Cleared: map[string]bool{}, Day: 1}
	g.Player.WeaponName = g.Weapon.Name
	g.Player.WeaponDamage = g.Weapon.Damage
	g.Player.WeaponCrit = g.Weapon.Crit
	g.Unlocked[zones[0].Name] = true
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
		fmt.Println(classes.BrightRed + "Nom invalide." + classes.Reset)
	}
	fmt.Println(classes.Panel("CLASSES", "[1] Survivant — équilibré", "[2] Punk — critique", "[3] Médecin de fortune — soins", "[4] Sauveur — défense"))
	var kind string
	for {
		x := classes.ReadInt()
		switch x {
		case 1:
			kind = "Survivant"
		case 2:
			kind = "Punk"
		case 3:
			kind = "Médecin de fortune"
		case 4:
			kind = "Sauveur"
		}
		if kind != "" {
			break
		}
		fmt.Print("Classe > ")
	}
	c := classes.Classes[kind]
	c.Nom = name
	c.Type = kind
	c.Level = 1
	c.PV = c.PVBase
	c.Mana = c.ManaMax
	c.Attacks = []string{"Attaque basique"}
	c.Inventory = []string{"Bandage"}
	c.SkillPoints = 1
	c.WeaponName = "Couteau de survie"
	c.WeaponDamage = 8
	c.WeaponCrit = 5
	classes.ClearScreen()
	fmt.Println(classes.TitleBox("✔ SURVIVANT CRÉÉ"))
	fmt.Println(classes.Panel("PERSONNAGE", "Nom : "+c.Nom, "Classe : "+c.Type, fmt.Sprintf("PV : %d | Mana : %d | Or : %d", c.PV, c.Mana, c.Gold)))
	classes.Pause()
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
		c := g.Player
		fmt.Println(classes.Logo())
		fmt.Println(classes.TitleBox("☣ REFUGE CENTRAL • JOUR " + fmt.Sprint(g.Day) + " ☣"))
		fmt.Println(classes.Panel(c.Nom+" • "+c.Type, fmt.Sprintf("PV   %s %d/%d", classes.HPBar(c.PV, c.PVBase), c.PV, c.PVBase), fmt.Sprintf("Mana %s %d/%d", classes.ManaBar(c.Mana, c.ManaMax), c.Mana, c.ManaMax), fmt.Sprintf("XP   %s %d/%d", classes.ExpBar(c.Exp, c.MaxExp), c.Exp, c.MaxExp), fmt.Sprintf("🔫 %s [%s]  Dégâts %d  Crit %d%%", g.Weapon.Name, g.Weapon.Rarity, g.Weapon.Damage, g.Weapon.Crit), fmt.Sprintf("💰 %d   🔩 %d   💉 %d   💣 %d", c.Gold, g.Scrap, g.Medkits, g.Grenades)))
		fmt.Println(classes.Panel("CENTRE DE COMMANDE", "[1] 🗺️ Partir en expédition", "[2] 🎒 Inventaire / équipement", "[3] 🛒 Marchand", "[4] 🔨 Forgeron", "[5] 👤 Fiche du survivant", "[6] ⚔️ Combat d'entraînement", "[7] 🎰 Casino", "[8] 🌳 Compétences", "[9] 💾 Statistiques", "[0] Quitter"))
		fmt.Print("\nAction > ")
		x := classes.ReadInt()
		switch x {
		case 1:
			g.mapMenu()
		case 2:
			g.inventoryMenu()
		case 3:
			marchand.Merchant(c)
		case 4:
			forgeron.BlacksmithTiers(c)
		case 5:
			classes.ClearScreen()
			classes.DisplayInfo(c)
			classes.Pause()
		case 6:
			monstre.SimulationFight(c)
		case 7:
			casino.Casino(c)
		case 8:
			g.skills()
		case 9:
			g.stats()
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
			classes.Pause()
		}
	}
}

func (g *Game) mapMenu() {
	for {
		classes.ClearScreen()
		fmt.Println(classes.TitleBox("🗺️ CARTE DES ZONES"))
		for i, z := range zones {
			status := "🔒"
			if g.Unlocked[z.Name] {
				status = "🟢"
			}
			if g.Cleared[z.Name] {
				status = "🏆"
			}
			fmt.Printf("[%d] %s %s — niveau %d — %d vagues %s\n", i+1, z.Icon, z.Name, z.Level, z.Waves, status)
		}
		fmt.Println("[0] Retour")
		fmt.Print("Zone > ")
		x := classes.ReadInt()
		if x == 0 {
			return
		}
		if x < 1 || x > len(zones) {
			continue
		}
		z := zones[x-1]
		if !g.Unlocked[z.Name] {
			fmt.Println("Zone verrouillée.")
			classes.Pause()
			continue
		}
		g.expedition(z)
	}
}

func (g *Game) expedition(z Zone) {
	c := g.Player
	for wave := 1; wave <= z.Waves; wave++ {
		classes.ClearScreen()
		fmt.Println(classes.TitleBox(z.Icon + " " + z.Name + " • VAGUE " + fmt.Sprint(wave) + "/" + fmt.Sprint(z.Waves)))
		enemies := g.spawnWave(z, wave)
		if !g.combat(enemies, wave, z) {
			return
		}
		if wave < z.Waves && rand.Intn(100) < 65 {
			g.event()
		}
	}
	classes.ClearScreen()
	fmt.Println(classes.TitleBox("🏆 ZONE NETTOYÉE"))
	fmt.Println(classes.Panel(z.Name, "Vous avez survécu aux "+fmt.Sprint(z.Waves)+" vagues.", fmt.Sprintf("Récompense : %d 💰 + %d 🔩", z.Reward, z.Waves*5)))
	c.Gold += z.Reward
	g.Scrap += z.Waves * 5
	g.Cleared[z.Name] = true
	g.unlockNext(z)
	g.healAfterRun()
	classes.Pause()
}
func (g *Game) spawnWave(z Zone, w int) []Enemy {
	n := 1 + (w+1)/2
	if n > 4 {
		n = 4
	}
	out := make([]Enemy, 0, n)
	for i := 0; i < n; i++ {
		e := z.Enemies[rand.Intn(len(z.Enemies))]()
		scale := 1 + (w-1)/4
		e.MaxHP *= scale
		e.HP = e.MaxHP
		e.Attack += w * 2
		e.Exp += w * 8
		e.Gold += w * 5
		out = append(out, e)
	}
	if w == z.Waves {
		e := z.Boss()
		out = append(out, e)
	}
	return out
}

func (g *Game) combat(enemies []Enemy, wave int, z Zone) bool {
	c := g.Player
	for len(enemies) > 0 && c.PV > 0 {
		classes.ClearScreen()
		fmt.Println(classes.TitleBox("⚔️ COMBAT • " + z.Name))
		fmt.Printf("%s  PV %d/%d   Mana %d/%d   🔫 %s (%d/%d)\n", c.Nom, c.PV, c.PVBase, c.Mana, c.ManaMax, g.Weapon.Name, g.Weapon.Ammo, g.Weapon.MaxAmmo)
		for i, e := range enemies {
			fmt.Printf("[%d] %s %s %d/%d HP", i+1, enemyIcon(e), e.Name, e.HP, e.MaxHP)
			if e.Boss {
				fmt.Print(" 👑")
			}
			fmt.Println()
		}
		fmt.Println(classes.Panel("ACTIONS", "[1] 🔫 Attaque avec arme", "[2] ⚡ Compétence", "[3] 💉 Soin", "[4] 💣 Grenade", "[5] 🔄 Recharger", "[6] 🏃 Fuir"))
		fmt.Print("Action > ")
		a := classes.ReadInt()
		if a == 6 && wave < z.Waves {
			fmt.Println("Vous fuyez l'expédition.")
			classes.Pause()
			return false
		}
		if a == 6 {
			fmt.Println("Impossible de fuir le boss.")
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
		switch a {
		case 1:
			g.shoot(&enemies[target])
		case 2:
			g.skillAttack(&enemies[target])
		case 3:
			g.useMedkit()
		case 4:
			g.grenade(enemies)
		case 5:
			g.reload()
		default:
			continue
		}
		if enemies[target].HP <= 0 {
			enemies = g.killEnemy(enemies, target)
			if len(enemies) == 0 {
				break
			}
			if target >= len(enemies) {
				target = len(enemies) - 1
			}
		}
		g.enemyTurn(enemies)
	}
	return c.PV > 0
}
func enemyIcon(e Enemy) string {
	if e.Boss {
		return "👑"
	}
	if e.Attack >= 40 {
		return "🧟"
	}
	if e.Defense >= 10 {
		return "🛡️"
	}
	return "🧟‍♂️"
}
func (g *Game) shoot(e *Enemy) {
	if g.Weapon.MaxAmmo > 0 && g.Weapon.Ammo <= 0 {
		fmt.Println("Plus de munitions !")
		classes.Pause()
		return
	}
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
	if g.Weapon.MaxAmmo > 0 {
		g.Weapon.Ammo--
	}
	fmt.Printf("🔫 %s inflige %d dégâts à %s.\n", g.Weapon.Name, d, e.Name)
	classes.Pause()
}
func (g *Game) skillAttack(e *Enemy) {
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
	fmt.Printf("⚡ Attaque spéciale : %d dégâts !\n", d)
	classes.Pause()
}
func (g *Game) useMedkit() {
	if g.Medkits <= 0 {
		fmt.Println("Aucun kit médical.")
		classes.Pause()
		return
	}
	g.Medkits--
	heal := 50
	if g.Player.Type == "Médecin de fortune" {
		heal = 75
	}
	g.Player.PV += heal
	if g.Player.PV > g.Player.PVBase {
		g.Player.PV = g.Player.PVBase
	}
	fmt.Printf("💉 +%d PV.\n", heal)
	classes.Pause()
}
func (g *Game) grenade(enemies []Enemy) {
	if g.Grenades <= 0 {
		fmt.Println("Plus de grenades.")
		classes.Pause()
		return
	}
	g.Grenades--
	for i := range enemies {
		d := 25 + rand.Intn(15) - enemies[i].Defense/2
		if d < 5 {
			d = 5
		}
		enemies[i].HP -= d
	}
	fmt.Println("💣 Explosion ! Tous les ennemis prennent des dégâts.")
	classes.Pause()
}
func (g *Game) reload() {
	if g.Weapon.MaxAmmo == 0 {
		fmt.Println("Cette arme ne nécessite pas de munitions.")
		classes.Pause()
		return
	}
	g.Weapon.Ammo = g.Weapon.MaxAmmo
	fmt.Println("🔄 Chargeur plein.")
	classes.Pause()
}
func (g *Game) enemyTurn(es []Enemy) {
	for i := range es {
		if es[i].HP <= 0 {
			continue
		}
		d := es[i].Attack - g.Player.Defense/2
		if d < 1 {
			d = 1
		}
		if g.Player.Type == "Sauveur" {
			d = d * 75 / 100
		}
		if rand.Intn(100) < 5 {
			fmt.Printf("%s rate son attaque !\n", es[i].Name)
			continue
		}
		g.Player.PV -= d
		fmt.Printf("%s inflige %d dégâts.\n", es[i].Name, d)
		if g.Player.PV <= 0 {
			g.Player.PV = 0
			fmt.Println("☠️ Vous êtes tombé au combat.")
			classes.Pause()
			return
		}
		time.Sleep(250 * time.Millisecond)
	}
}
func (g *Game) killEnemy(e []Enemy, i int) []Enemy {
	dead := e[i]
	g.Player.Kills++
	g.Player.Gold += dead.Gold
	gainXP(g.Player, dead.Exp)
	if dead.Boss {
		g.Player.BossesDefeated++
		g.Scrap += 30
		fmt.Println("👑 BOSS VAINCU ! +30 🔩")
	}
	g.loot(dead)
	e[i] = e[len(e)-1]
	return e[:len(e)-1]
}

func gainXP(c *classes.Classe, x int) {
	c.Exp += x
	for c.Exp >= c.MaxExp {
		c.Exp -= c.MaxExp
		c.Level++
		c.MaxExp += 25
		c.PVBase += 12
		c.PV = c.PVBase
		c.ManaMax += 6
		c.Mana = c.ManaMax
		c.AttaqueBase += 2
		c.Attaque = c.AttaqueBase
		c.DefenseBase += 1
		c.Defense = c.DefenseBase
		c.SkillPoints++
		fmt.Printf("⭐ NIVEAU %d ! +1 point de compétence\n", c.Level)
	}
}
func (g *Game) loot(e Enemy) {
	r := rand.Intn(100)
	if r < 35 {
		g.Scrap += 5 + rand.Intn(10)
		fmt.Println("🔩 Ferraille récupérée.")
	}
	if r < 18 {
		g.Medkits++
		fmt.Println("💉 Kit médical trouvé.")
	}
	if r < 10 {
		g.Grenades++
		fmt.Println("💣 Grenade trouvée.")
	}
	if e.Boss {
		fmt.Println("🎁 Le boss laisse tomber un équipement rare !")
		g.Scrap += 25
	}
}
func (g *Game) event() {
	classes.ClearScreen()
	events := []string{"🚗 Vous trouvez une voiture abandonnée.", "🏚️ Une cache de survivants est encore intacte.", "📦 Un sac de ravitaillement est coincé sous des débris.", "☣️ Vous entendez des grognements derrière une porte."}
	e := events[rand.Intn(len(events))]
	fmt.Println(classes.TitleBox("🎲 ÉVÉNEMENT"))
	fmt.Println(e)
	fmt.Println("[1] Fouiller   [2] Continuer")
	x := classes.ReadInt()
	if x == 1 {
		switch {
		case strings.HasPrefix(e, "🚗"):
			g.Player.Gold += rand.Intn(50) + 20
		case strings.HasPrefix(e, "🏚️"):
			g.Medkits++
			g.Scrap += 10
		case strings.HasPrefix(e, "📦"):
			g.Grenades++
			g.Scrap += 15
		default:
			fmt.Println("☣️ Une petite embuscade !")
			g.Player.PV -= 10
		}
	}
	classes.Pause()
}
func (g *Game) unlockNext(z Zone) {
	for i := range zones {
		if zones[i].Name == z.Name && i+1 < len(zones) {
			next := zones[i+1]
			if g.Player.Level >= next.Level || g.Cleared[z.Name] {
				g.Unlocked[next.Name] = true
				fmt.Println("🔓 Nouvelle zone :", next.Name)
			}
		}
	}
}
func (g *Game) healAfterRun() {
	g.Player.PV += g.Player.PVBase / 4
	if g.Player.PV > g.Player.PVBase {
		g.Player.PV = g.Player.PVBase
	}
	g.Player.Mana = g.Player.ManaMax
}
func (g *Game) inventoryMenu() {
	for {
		classes.ClearScreen()
		fmt.Println(classes.TitleBox("🎒 ARSENAL & INVENTAIRE"))
		fmt.Println(classes.Panel("ÉQUIPEMENT", fmt.Sprintf("🔫 %s [%s] — %d dégâts — critique %d%%", g.Weapon.Name, g.Weapon.Rarity, g.Weapon.Damage, g.Weapon.Crit), fmt.Sprintf("Munitions : %d/%d", g.Weapon.Ammo, g.Weapon.MaxAmmo), fmt.Sprintf("💉 Kits : %d   💣 Grenades : %d   🔩 Ferraille : %d", g.Medkits, g.Grenades, g.Scrap), fmt.Sprintf("Capacité : %d/%d objets", len(g.Player.Inventory), g.Player.MaxInventory)))
		fmt.Println("[1] Armes disponibles  [2] Équipement classique  [3] Consommer objet  [0] Retour")
		x := classes.ReadInt()
		switch x {
		case 1:
			g.weaponShop()
		case 2:
			inventaire.AccessInventory(g.Player)
		case 3:
			g.quickUse()
		case 0:
			return
		}
	}
}
func (g *Game) weaponShop() {
	for {
		classes.ClearScreen()
		fmt.Println(classes.TitleBox("🔫 ARSENAL"))
		for i, w := range weapons {
			owned := w.Name == g.Weapon.Name
			status := ""
			if owned {
				status = " ← ÉQUIPÉ"
			}
			fmt.Printf("[%d] %-24s %-11s Dégâts:%3d Crit:%2d%% Prix:%4d 💰%s\n", i+1, w.Name, w.Rarity, w.Damage, w.Crit, w.Price, status)
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
		if w.Name == g.Weapon.Name {
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
		fmt.Println("Équipé :", w.Name)
		classes.Pause()
	}
}
func (g *Game) quickUse() {
	if g.Medkits > 0 {
		g.useMedkit()
		return
	}
	fmt.Println("Aucun objet de soin rapide.")
	classes.Pause()
}
func (g *Game) skills() {
	classes.ClearScreen()
	fmt.Println(classes.TitleBox("🌳 ARBRE DE COMPÉTENCES"))
	fmt.Printf("Points disponibles : %d\n\n", g.Player.SkillPoints)
	skills := []struct {
		Name string
		Cost int
		Desc string
	}{{"Adrénaline", 1, "+3 attaque de base"}, {"Endurance", 1, "+20 PV max"}, {"Blindage", 1, "+2 défense"}, {"Maîtrise des armes", 2, "+5 dégâts d'arme"}, {"Médecine", 2, "+1 kit médical"}}
	for i, s := range skills {
		fmt.Printf("[%d] %s (%d pt) — %s\n", i+1, s.Name, s.Cost, s.Desc)
	}
	fmt.Println("[0] Retour")
	x := classes.ReadInt()
	if x < 1 || x > len(skills) {
		return
	}
	s := skills[x-1]
	if g.Player.SkillPoints < s.Cost {
		fmt.Println("Points insuffisants.")
		classes.Pause()
		return
	}
	g.Player.SkillPoints -= s.Cost
	switch x {
	case 1:
		g.Player.AttaqueBase += 3
		g.Player.Attaque = g.Player.AttaqueBase
	case 2:
		g.Player.PVBase += 20
		g.Player.PV = g.Player.PVBase
	case 3:
		g.Player.DefenseBase += 2
		g.Player.Defense = g.Player.DefenseBase
	case 4:
		g.Weapon.Damage += 5
		g.Player.WeaponDamage = g.Weapon.Damage
	case 5:
		g.Medkits++
	}
	fmt.Println("Compétence améliorée :", s.Name)
	classes.Pause()
}
func (g *Game) stats() {
	classes.ClearScreen()
	fmt.Println(classes.TitleBox("📊 DOSSIER DU SURVIVANT"))
	lines := []string{fmt.Sprintf("Niveau : %d", g.Player.Level), fmt.Sprintf("XP : %d/%d", g.Player.Exp, g.Player.MaxExp), fmt.Sprintf("Ennemis éliminés : %d", g.Player.Kills), fmt.Sprintf("Boss vaincus : %d", g.Player.BossesDefeated), fmt.Sprintf("Or : %d", g.Player.Gold), fmt.Sprintf("Ferraille : %d", g.Scrap), fmt.Sprintf("Arme : %s", g.Weapon.Name), fmt.Sprintf("Dégâts : %d", g.Weapon.Damage)}
	sort.Strings(lines)
	fmt.Println(classes.Panel("STATISTIQUES", lines...))
	classes.Pause()
}

package expedition

import (
	"fmt"
	"math/rand"
	"time"

	"outbreak/classes"
	"outbreak/monstre"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ============================================================
// ÉTAT DES EXPÉDITIONS
// ============================================================

var zonesNettoyees = make(map[string]bool)
var bossVaincus = make(map[string]bool)

// ============================================================
// ZONES
// ============================================================

type Zone struct {
	Nom          string
	Description  string
	NiveauMin    int
	Vagues       int
	Boss         func() monstre.Monster
	RecompenseXP int
	RecompenseOr int
}

var zones = []Zone{
	{
		Nom:          "Centre-ville",
		Description:  "Les rues sont infestées. Quelques ressources peuvent encore être récupérées.",
		NiveauMin:    1,
		Vagues:       3,
		Boss:         monstre.InitZombieBlinde,
		RecompenseXP: 150,
		RecompenseOr: 100,
	},
	{
		Nom:          "Hôpital abandonné",
		Description:  "Un ancien hôpital rempli d'infectés et de ressources médicales.",
		NiveauMin:    3,
		Vagues:       4,
		Boss:         monstre.InitCoureur,
		RecompenseXP: 250,
		RecompenseOr: 175,
	},
	{
		Nom:          "Usine",
		Description:  "Une usine dangereuse où les infectés ont établi leur repaire.",
		NiveauMin:    6,
		Vagues:       5,
		Boss:         monstre.InitBrute,
		RecompenseXP: 400,
		RecompenseOr: 275,
	},
	{
		Nom:          "Base militaire",
		Description:  "Les soldats ont disparu. Quelque chose de beaucoup plus dangereux vit ici.",
		NiveauMin:    9,
		Vagues:       6,
		Boss:         monstre.InitNecromancien,
		RecompenseXP: 650,
		RecompenseOr: 450,
	},
	{
		Nom:          "Zone contaminée",
		Description:  "La zone la plus dangereuse connue. Seuls les meilleurs survivants y entrent.",
		NiveauMin:    13,
		Vagues:       8,
		Boss:         monstre.InitRoiInfectes,
		RecompenseXP: 1200,
		RecompenseOr: 900,
	},
}

// ============================================================
// MENU PRINCIPAL DES EXPÉDITIONS
// ============================================================

func Expedition(c *classes.Classe) {

	for {

		classes.ClearScreen()

		fmt.Println(
			classes.TitleBox(
				"☣ EXPÉDITIONS — ZONE MORTE ☣",
			),
		)

		fmt.Println()

		fmt.Println(
			classes.Panel(
				"ÉTAT DU SURVIVANT",

				fmt.Sprintf(
					"%s%s%s  • Niveau %d",
					classes.BrightCyan,
					c.Nom,
					classes.Reset,
					c.Level,
				),

				fmt.Sprintf(
					"PV %s %d/%d",
					classes.HPBar(c.PV, c.PVBase),
					c.PV,
					c.PVBase,
				),

				fmt.Sprintf(
					"💰 %d or",
					c.Gold,
				),
			),
		)

		fmt.Println()

		fmt.Println(
			classes.Panel(
				"CARTE DES ZONES",

				zoneLine(0),
				zoneLine(1),
				zoneLine(2),
				zoneLine(3),
				zoneLine(4),

				"",
				classes.Dim+"[0] Retour au refuge"+classes.Reset,
			),
		)

		fmt.Print(
			"\n" +
				classes.BrightCyan +
				"Expédition > " +
				classes.Reset,
		)

		choix := classes.ReadInt()

		if choix == 0 {
			return
		}

		if choix < 1 || choix > len(zones) {

			fmt.Println(
				classes.BrightRed +
					"✖ Zone inconnue." +
					classes.Reset,
			)

			classes.Pause()
			continue
		}

		lancerZone(c, zones[choix-1])
	}
}

// ============================================================
// AFFICHAGE DES ZONES
// ============================================================

func zoneLine(index int) string {
	z := zones[index]
	num := fmt.Sprintf("[%d]", index+1)

	if zonesNettoyees[z.Nom] {
		return classes.BrightGreen +
			num +
			" ✓ " +
			z.Nom +
			classes.Reset +
			" • ZONE NETTOYÉE"
	}

	return classes.BrightYellow +
		num +
		" ☣ " +
		z.Nom +
		classes.Reset +
		fmt.Sprintf(" • Niveau %d+ • %d vagues", z.NiveauMin, z.Vagues)
}

// ============================================================
// LANCEMENT D'UNE ZONE
// ============================================================

func lancerZone(c *classes.Classe, zone Zone) {

	if c.Level < zone.NiveauMin {

		classes.ClearScreen()

		fmt.Println(classes.TitleBox("🔒 ZONE VERROUILLÉE"))

		fmt.Println()

		fmt.Printf(
			"Cette zone nécessite le niveau %d.\n",
			zone.NiveauMin,
		)

		fmt.Printf(
			"Votre niveau : %d\n\n",
			c.Level,
		)

		fmt.Println(
			classes.Dim +
				zone.Description +
				classes.Reset,
		)

		classes.Pause()
		return
	}

	classes.ClearScreen()

	fmt.Println(
		classes.TitleBox("☣ EXPÉDITION"),
	)

	fmt.Println()

	fmt.Println(
		classes.Panel(
			zone.Nom,
			zone.Description,
			fmt.Sprintf("Niveau requis : %d", zone.NiveauMin),
			fmt.Sprintf("Vagues : %d", zone.Vagues),
		),
	)

	fmt.Println()

	fmt.Println(
		classes.BrightYellow +
			"Préparation de l'expédition..." +
			classes.Reset,
	)

	time.Sleep(900 * time.Millisecond)

	c.PV = c.PVBase
	c.Mana = c.ManaMax

	for vague := 1; vague <= zone.Vagues; vague++ {

		if c.PV <= 0 {
			break
		}

		if !vagueEvenement(c, zone, vague) {
			break
		}

		monstres := creerVague(c, vague, zone)

		gagne := simulationCombat(
			c,
			monstres,
			vague,
			zone,
		)

		if !gagne {
			break
		}

		recuperation(c)

		if vague < zone.Vagues {

			fmt.Println()

			fmt.Printf(
				"%s✔ Vague %d/%d terminée !%s\n",
				classes.BrightGreen,
				vague,
				zone.Vagues,
				classes.Reset,
			)

			fmt.Println(
				classes.Dim +
					"Une nouvelle vague approche..." +
					classes.Reset,
			)

			time.Sleep(1000 * time.Millisecond)
		}
	}

	if c.PV <= 0 {
		return
	}

	fmt.Println()

	fmt.Println(
		classes.BrightRed +
			"☠☠☠ LE BOSS APPROCHE ☠☠☠" +
			classes.Reset,
	)

	time.Sleep(1200 * time.Millisecond)

	boss := zone.Boss()

	scaleBoss(&boss, c.Level)

	gagne := bossFight(c, &boss, zone)

	if !gagne {
		return
	}

	bossVaincus[zone.Nom] = true
	zonesNettoyees[zone.Nom] = true

	classes.ClearScreen()

	fmt.Println(
		classes.TitleBox("🏆 ZONE NETTOYÉE"),
	)

	fmt.Println()

	fmt.Printf(
		"%s☣ %s a été éliminé !%s\n\n",
		classes.BrightGreen,
		boss.Name,
		classes.Reset,
	)

	fmt.Println(
		classes.Panel(
			"RÉCOMPENSES",
			fmt.Sprintf(
				"%s+%d XP%s",
				classes.BrightMagenta,
				zone.RecompenseXP,
				classes.Reset,
			),
			fmt.Sprintf(
				"%s+%d 💰%s",
				classes.BrightYellow,
				zone.RecompenseOr,
				classes.Reset,
			),
			"☣ Zone définitivement nettoyée",
		),
	)

	monstre.GainExp(c, zone.RecompenseXP)
	monstre.GainGold(c, zone.RecompenseOr)

	donnerLoot(c)

	classes.Pause()
}

// ============================================================
// CRÉATION DES VAGUES
// ============================================================

func creerVague(
	c *classes.Classe,
	numero int,
	zone Zone,
) []monstre.Monster {

	nombre := 1 + numero/2

	if nombre > 5 {
		nombre = 5
	}

	var resultat []monstre.Monster

	for i := 0; i < nombre; i++ {

		m := monstre.SelectMonster(c.Level)

		// Bonus de difficulté selon la zone et la vague.
		bonus := (numero-1)*3 + zone.NiveauMin

		m.MaxHP += bonus * 3
		m.CurrentHP = m.MaxHP
		m.Attack += bonus

		m.ExpReward += bonus * 2
		m.GoldReward += bonus

		resultat = append(resultat, m)
	}

	return resultat
}

// ============================================================
// SIMULATION AUTOMATIQUE
// ============================================================

func simulationCombat(
	c *classes.Classe,
	monstres []monstre.Monster,
	vague int,
	zone Zone,
) bool {

	tour := 1

	for {

		vivants := compterVivants(monstres)

		if vivants == 0 {
			return true
		}

		if c.PV <= 0 {
			return false
		}

		classes.ClearScreen()

		fmt.Println(
			classes.TitleBox(
				"⚔ SIMULATION AUTOMATIQUE",
			),
		)

		fmt.Printf(
			"\n%s%s — Vague %d — Tour %d%s\n\n",
			classes.Bold,
			zone.Nom,
			vague,
			tour,
			classes.Reset,
		)

		fmt.Printf(
			"%s%s%s\n",
			classes.Bold,
			c.Nom,
			classes.Reset,
		)

		fmt.Printf(
			"PV    %s %d/%d\n",
			classes.HPBar(c.PV, c.PVBase),
			c.PV,
			c.PVBase,
		)

		fmt.Printf(
			"Mana  %s %d/%d\n\n",
			classes.ManaBar(c.Mana, c.ManaMax),
			c.Mana,
			c.ManaMax,
		)

		fmt.Println(
			classes.BrightRed +
				"──── ENNEMIS ────" +
				classes.Reset,
		)

		for i := range monstres {

			m := &monstres[i]

			if m.CurrentHP <= 0 {

				fmt.Printf(
					"%s☠ %s — VAINCU%s\n",
					classes.Dim,
					m.Name,
					classes.Reset,
				)

				continue
			}

			fmt.Printf(
				"%s[%d] %-24s%s %s %d/%d\n",
				classes.BrightRed,
				i+1,
				m.Name,
				classes.Reset,
				classes.HPBar(
					m.CurrentHP,
					m.MaxHP,
				),
				m.CurrentHP,
				m.MaxHP,
			)
		}

		fmt.Println()

		// ----------------------------------------------------
		// TOUR DU JOUEUR
		// ----------------------------------------------------

		cible := choisirCible(monstres)

		if cible >= 0 {

			attaque := choisirAttaque(c)

			damage := attaque.Damage + c.AttaqueBase

			// Critiques du Punk.
			if c.Type == "Punk" && rand.Intn(4) == 0 {

				damage = damage * 3 / 2

				fmt.Println(
					classes.BrightYellow +
						"⚡ COUP CRITIQUE ! ⚡" +
						classes.Reset,
				)
			}

			if attaque.ManaCost <= c.Mana {
				c.Mana -= attaque.ManaCost
			}

			monstres[cible].CurrentHP -= damage

			if monstres[cible].CurrentHP < 0 {
				monstres[cible].CurrentHP = 0
			}

			fmt.Printf(
				"%s🔫 %-28s%s → %s-%d PV%s\n",
				classes.BrightCyan,
				attaque.Name,
				classes.Reset,
				classes.BrightGreen,
				damage,
				classes.Reset,
			)

			if monstres[cible].CurrentHP == 0 {

				fmt.Printf(
					"%s☠ %s est éliminé !%s\n",
					classes.BrightRed,
					monstres[cible].Name,
					classes.Reset,
				)

				monstre.GainExp(
					c,
					monstres[cible].ExpReward,
				)

				monstre.GainGold(
					c,
					monstres[cible].GoldReward,
				)

				donnerLoot(c)
			}
		}

		// ----------------------------------------------------
		// TOUR DES ENNEMIS
		// ----------------------------------------------------

		for i := range monstres {

			m := &monstres[i]

			if m.CurrentHP <= 0 || c.PV <= 0 {
				continue
			}

			damage := m.Attack

			// Attaque spéciale tous les 3 tours.
			if tour%3 == 0 {
				damage *= 2

				fmt.Printf(
					"%s☠ %s lance une attaque spéciale !%s\n",
					classes.BrightRed,
					m.Name,
					classes.Reset,
				)
			}

			damage -= c.DefenseBase / 2

			// Bonus défensif du Sauveur.
			if c.Type == "Sauveur" {
				damage = damage * 80 / 100
			}

			if damage < 1 {
				damage = 1
			}

			c.PV -= damage

			if c.PV < 0 {
				c.PV = 0
			}

			fmt.Printf(
				"%s🧟 %-28s%s → %s-%d PV%s\n",
				classes.BrightRed,
				m.Name,
				classes.Reset,
				classes.BrightRed,
				damage,
				classes.Reset,
			)
		}

		// Régénération du Survivant.
		if c.Type == "Survivant" && c.PV > 0 {

			soin := 3

			c.PV += soin

			if c.PV > c.PVBase {
				c.PV = c.PVBase
			}

			fmt.Printf(
				"%s♥ Régénération : +%d PV%s\n",
				classes.BrightGreen,
				soin,
				classes.Reset,
			)
		}

		// Régénération légère de mana.
		c.Mana += 2

		if c.Mana > c.ManaMax {
			c.Mana = c.ManaMax
		}

		fmt.Println()

		if c.PV <= 0 {

			fmt.Println(
				classes.BrightRed +
					"☠ VOUS AVEZ ÉTÉ SUBMERGÉ ☠" +
					classes.Reset,
			)

			c.PV = c.PVBase / 2

			if c.PV < 1 {
				c.PV = 1
			}

			classes.Pause()

			return false
		}

		tour++

		time.Sleep(850 * time.Millisecond)
	}
}

// ============================================================
// COMBAT DU BOSS
// ============================================================

func bossFight(
	c *classes.Classe,
	boss *monstre.Monster,
	zone Zone,
) bool {

	tour := 1

	for boss.CurrentHP > 0 && c.PV > 0 {

		classes.ClearScreen()

		fmt.Println(
			classes.TitleBox(
				"☠ COMBAT DE BOSS ☠",
			),
		)

		fmt.Printf(
			"\n%s%s%s\n",
			classes.Bold+classes.BrightRed,
			boss.Name,
			classes.Reset,
		)

		fmt.Printf(
			"PV BOSS  %s %d/%d\n\n",
			classes.HPBar(
				boss.CurrentHP,
				boss.MaxHP,
			),
			boss.CurrentHP,
			boss.MaxHP,
		)

		fmt.Printf(
			"%s%s%s\n",
			classes.Bold,
			c.Nom,
			classes.Reset,
		)

		fmt.Printf(
			"PV       %s %d/%d\n",
			classes.HPBar(
				c.PV,
				c.PVBase,
			),
			c.PV,
			c.PVBase,
		)

		fmt.Printf(
			"Mana     %s %d/%d\n\n",
			classes.ManaBar(
				c.Mana,
				c.ManaMax,
			),
			c.Mana,
			c.ManaMax,
		)

		fmt.Printf(
			"%sTOUR %d%s\n\n",
			classes.BrightYellow,
			tour,
			classes.Reset,
		)

		// ----------------------------------------------------
		// PHASE DU BOSS
		// ----------------------------------------------------

		phase := 1

		if boss.CurrentHP <= boss.MaxHP/2 {
			phase = 2
		}

		if boss.CurrentHP <= boss.MaxHP/4 {
			phase = 3
		}

		fmt.Printf(
			"%s☣ PHASE %d%s\n\n",
			classes.BrightMagenta,
			phase,
			classes.Reset,
		)

		// ----------------------------------------------------
		// ATTAQUE DU JOUEUR
		// ----------------------------------------------------

		attaque := choisirAttaque(c)

		damage := attaque.Damage + c.AttaqueBase

		if phase >= 2 {
			damage += 3
		}

		if phase == 3 {
			damage += 5
		}

		if c.Type == "Punk" && rand.Intn(4) == 0 {

			damage = damage * 3 / 2

			fmt.Println(
				classes.BrightYellow +
					"⚡ COUP CRITIQUE ! ⚡" +
					classes.Reset,
			)
		}

		if attaque.ManaCost <= c.Mana {
			c.Mana -= attaque.ManaCost
		}

		boss.CurrentHP -= damage

		if boss.CurrentHP < 0 {
			boss.CurrentHP = 0
		}

		fmt.Printf(
			"%s🔫 %s%s → %s-%d PV%s\n",
			classes.BrightCyan,
			attaque.Name,
			classes.Reset,
			classes.BrightGreen,
			damage,
			classes.Reset,
		)

		if boss.CurrentHP <= 0 {
			break
		}

		// ----------------------------------------------------
		// ATTAQUE DU BOSS
		// ----------------------------------------------------

		damageBoss := boss.Attack

		if phase == 2 {
			damageBoss = damageBoss * 120 / 100
		}

		if phase == 3 {
			damageBoss = damageBoss * 150 / 100
		}

		if tour%3 == 0 {
			damageBoss *= 2

			fmt.Printf(
				"%s☠ ATTAQUE DÉVASTATRICE DU BOSS !%s\n",
				classes.BrightRed,
				classes.Reset,
			)
		}

		damageBoss -= c.DefenseBase / 2

		if c.Type == "Sauveur" {
			damageBoss = damageBoss * 80 / 100
		}

		if damageBoss < 1 {
			damageBoss = 1
		}

		c.PV -= damageBoss

		if c.PV < 0 {
			c.PV = 0
		}

		fmt.Printf(
			"%s☠ %s%s → %s-%d PV%s\n",
			classes.BrightRed,
			boss.Name,
			classes.Reset,
			classes.BrightRed,
			damageBoss,
			classes.Reset,
		)

		// Régénération.
		if c.Type == "Survivant" && c.PV > 0 {

			c.PV += 4

			if c.PV > c.PVBase {
				c.PV = c.PVBase
			}
		}

		c.Mana += 3

		if c.Mana > c.ManaMax {
			c.Mana = c.ManaMax
		}

		tour++

		time.Sleep(1000 * time.Millisecond)
	}

	return c.PV > 0 && boss.CurrentHP <= 0
}

// ============================================================
// CHOIX AUTOMATIQUES
// ============================================================

func choisirCible(monstres []monstre.Monster) int {

	meilleur := -1
	meilleurHP := 0

	for i := range monstres {

		if monstres[i].CurrentHP <= 0 {
			continue
		}

		if meilleur == -1 ||
			monstres[i].CurrentHP < meilleurHP {

			meilleur = i
			meilleurHP = monstres[i].CurrentHP
		}
	}

	return meilleur
}

func choisirAttaque(c *classes.Classe) *monstre.AttackInfo {

	var meilleure *monstre.AttackInfo

	for _, nom := range c.Attacks {

		attaque := monstre.GetAttackInfo(nom)

		if attaque == nil {
			continue
		}

		if attaque.ManaCost > c.Mana {
			continue
		}

		if meilleure == nil ||
			attaque.Damage > meilleure.Damage {

			meilleure = attaque
		}
	}

	if meilleure == nil {

		meilleure =
			monstre.GetAttackInfo(
				"Attaque basique",
			)
	}

	return meilleure
}

// ============================================================
// BOSS : MISE À L'ÉCHELLE
// ============================================================

func scaleBoss(
	boss *monstre.Monster,
	level int,
) {

	bonus := 1.0 + float64(level-1)*0.06

	boss.MaxHP =
		int(float64(boss.MaxHP) * bonus)

	boss.CurrentHP =
		boss.MaxHP

	boss.Attack =
		int(float64(boss.Attack) * bonus)

	boss.ExpReward =
		int(float64(boss.ExpReward) * bonus)

	boss.GoldReward =
		int(float64(boss.GoldReward) * bonus)

	if boss.MaxHP < 1 {
		boss.MaxHP = 1
	}

	if boss.Attack < 1 {
		boss.Attack = 1
	}
}

// ============================================================
// UTILITAIRES
// ============================================================

func compterVivants(monstres []monstre.Monster) int {

	n := 0

	for _, m := range monstres {

		if m.CurrentHP > 0 {
			n++
		}
	}

	return n
}

func recuperation(c *classes.Classe) {

	soin := c.PVBase / 10

	if soin < 5 {
		soin = 5
	}

	c.PV += soin

	if c.PV > c.PVBase {
		c.PV = c.PVBase
	}

	mana := c.ManaMax / 10

	if mana < 2 {
		mana = 2
	}

	c.Mana += mana

	if c.Mana > c.ManaMax {
		c.Mana = c.ManaMax
	}

	fmt.Printf(
		"%s♥ Récupération : +%d PV / +%d Mana%s\n",
		classes.BrightGreen,
		soin,
		mana,
		classes.Reset,
	)

	time.Sleep(700 * time.Millisecond)
}

// ============================================================
// ÉVÉNEMENTS ALÉATOIRES
// ============================================================

func vagueEvenement(
	c *classes.Classe,
	zone Zone,
	vague int,
) bool {

	if vague == 1 {
		return true
	}

	// 30 % de chance d'événement.
	if rand.Intn(100) >= 30 {
		return true
	}

	classes.ClearScreen()

	fmt.Println(
		classes.TitleBox(
			"❓ ÉVÉNEMENT",
		),
	)

	fmt.Println()

	switch rand.Intn(4) {

	case 0:

		fmt.Println(
			classes.Panel(
				"🚗 VÉHICULE ABANDONNÉ",

				"Vous trouvez une voiture abandonnée.",

				"[1] Fouiller",
				"[2] Continuer",
			),
		)

		choix := classes.ReadInt()

		if choix == 1 {

			gain := rand.Intn(70) + 20

			c.Gold += gain

			fmt.Printf(
				"\n%s✔ Vous trouvez %d 💰 !%s\n",
				classes.BrightYellow,
				gain,
				classes.Reset,
			)

			classes.Pause()
		}

	case 1:

		fmt.Println(
			classes.Panel(
				"📦 CACHE DE SURVIVANT",

				"Une ancienne cache contient des fournitures.",
			),
		)

		if classes.AddInventory(c, "Kit médical") {

			fmt.Println(
				classes.BrightGreen +
					"✔ Vous récupérez un kit médical." +
					classes.Reset,
			)
		}

		classes.Pause()

	case 2:

		fmt.Println(
			classes.Panel(
				"☣ EMBUSCADE",

				"Un petit groupe d'infectés vous attaque !",
			),
		)

		damage := rand.Intn(15) + 5

		c.PV -= damage

		if c.PV < 1 {
			c.PV = 1
		}

		fmt.Printf(
			"\n%s-%d PV%s\n",
			classes.BrightRed,
			damage,
			classes.Reset,
		)

		classes.Pause()

	case 3:

		fmt.Println(
			classes.Panel(
				"🧪 PRODUIT ÉTRANGE",

				"Vous trouvez une étrange injection.",
				"Elle semble encore utilisable.",
			),
		)

		soin := rand.Intn(30) + 15

		c.PV += soin

		if c.PV > c.PVBase {
			c.PV = c.PVBase
		}

		fmt.Printf(
			"\n%s✔ +%d PV%s\n",
			classes.BrightGreen,
			soin,
			classes.Reset,
		)

		classes.Pause()
	}

	_ = zone

	return c.PV > 0
}

// ============================================================
// LOOT
// ============================================================

func donnerLoot(c *classes.Classe) {

	if rand.Intn(100) >= 35 {
		return
	}

	loot := []string{
		"Kit médical",
		"Potion de vie",
		"Grenade",
		"Munition",
		"Pièces détachées",
	}

	objet := loot[rand.Intn(len(loot))]

	fmt.Printf(
		"%s📦 LOOT : %s%s\n",
		classes.BrightYellow,
		objet,
		classes.Reset,
	)

	classes.AddInventory(c, objet)
}
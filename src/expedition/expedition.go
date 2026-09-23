package expedition

import (
	"fmt"
	"math/rand"
	"time"

	"outbreak/classes"
	"outbreak/monstre"
)

type Zone struct {
	Name        string
	Description string
	Level       int
	Waves       int
	Boss        string
	RewardGold  int
	RewardXP    int
}

var zones = []Zone{
	{"Centre-ville", "Rues abandonnées, commerces pillés et premières hordes.", 1, 3, "", 80, 120},
	{"Hôpital", "Un ancien hôpital où les infectés semblent plus organisés.", 3, 3, "Nécromancien infecté", 140, 220},
	{"Usine", "Une usine pleine de machines encore alimentées et de mutants.", 6, 4, "Brute mutante", 220, 350},
	{"Base militaire", "Une base fortifiée où une ancienne escouade a disparu.", 9, 4, "Abomination", 350, 550},
	{"Zone contaminée", "Le cœur de l'épidémie. Personne ne sait ce qui attend au bout.", 13, 5, "Roi des infectés", 600, 900},
}

type event struct {
	Title string
	Text  string
	Run   func(*classes.Classe)
}

func Explore(c *classes.Classe) {
	for {
		classes.ClearScreen()
		fmt.Println(classes.TitleBox("🗺️ EXPÉDITIONS — OUTBREAK"))
		fmt.Println()
		fmt.Println(classes.Panel("CARTE", "Choisissez une zone à explorer.", "Les zones deviennent accessibles avec votre niveau."))
		fmt.Println()

		for i, z := range zones {
			unlocked := c.Level >= z.Level
			state := classes.BrightGreen + "DÉBLOQUÉE" + classes.Reset
			if !unlocked {
				state = classes.BrightRed + fmt.Sprintf("NIVEAU %d", z.Level) + classes.Reset
			}
			progress := c.CompletedZones[z.Name]
			fmt.Printf("%s[%d]%s %-20s %s  Vagues: %d  Terminées: %d\n", classes.BrightCyan, i+1, classes.Reset, z.Name, state, z.Waves, progress)
			fmt.Printf("    %s%s%s\n", classes.Dim, z.Description, classes.Reset)
			if z.Boss != "" {
				fmt.Printf("    ☣ Boss : %s%s%s\n", classes.BrightRed, z.Boss, classes.Reset)
			}
		}

		fmt.Println("\n" + classes.Dim + "[0] Retour" + classes.Reset)
		choice := classes.ReadIntPrompt("\nZone > ")
		if choice == 0 {
			return
		}
		if choice < 1 || choice > len(zones) {
			fmt.Println(classes.BrightRed + "✖ Zone invalide." + classes.Reset)
			classes.Pause()
			continue
		}
		z := zones[choice-1]
		if c.Level < z.Level {
			fmt.Println(classes.BrightRed + "✖ Cette zone est trop dangereuse pour votre niveau." + classes.Reset)
			classes.Pause()
			continue
		}
		runZone(c, z)
	}
}

func runZone(c *classes.Classe, z Zone) {
	classes.ClearScreen()
	fmt.Println(classes.TitleBox("🚨 EXPÉDITION : " + z.Name))
	fmt.Println()
	fmt.Println(classes.Panel("MISSION", z.Description, fmt.Sprintf("Vagues prévues : %d", z.Waves), fmt.Sprintf("Récompense de zone : %d 💰 / %d XP", z.RewardGold, z.RewardXP)))
	classes.Pause()

	for wave := 1; wave <= z.Waves; wave++ {
		classes.ClearScreen()
		fmt.Println(classes.TitleBox(fmt.Sprintf("☣ %s — VAGUE %d/%d", z.Name, wave, z.Waves)))
		fmt.Println()
		fmt.Println(classes.BrightYellow + "Une horde approche..." + classes.Reset)
		time.Sleep(500 * time.Millisecond)
		monstre.SimulationFight(c)

		if c.PV <= 0 {
			return
		}

		if wave < z.Waves {
			randomEvent(c)
		}
	}

	if z.Boss != "" {
		bossBattle(c, z)
		if c.PV <= 0 {
			return
		}
	}

	c.CompletedZones[z.Name]++
	c.Gold += z.RewardGold
	monstre.GainExp(c, z.RewardXP)

	loot(c, z)

	classes.ClearScreen()
	fmt.Println(classes.TitleBox("🏆 EXPÉDITION TERMINÉE"))
	fmt.Println()
	fmt.Printf("%s%s%s est revenu vivant de %s.\n\n", classes.BrightGreen, c.Nom, classes.Reset, z.Name)
	fmt.Printf("+%d 💰\n+%d XP\n", z.RewardGold, z.RewardXP)
	classes.Pause()
}

func bossBattle(c *classes.Classe, z Zone) {
	classes.ClearScreen()
	fmt.Println(classes.TitleBox("☣ BOSS — " + z.Boss))
	fmt.Println()
	fmt.Println(classes.BrightRed + "Le boss vous attend au cœur de la zone." + classes.Reset)
	classes.Pause()
	if monstre.BossFight(c, z.Boss) {
		c.BossesDefeated[z.Boss] = true
		bonus := 150 + c.Level*15
		c.Gold += bonus
		fmt.Printf("%s☠ BOSS VAINCU ! +%d 💰%s\n", classes.BrightYellow, bonus, classes.Reset)
		classes.Pause()
	}
}

func randomEvent(c *classes.Classe) {
	events := []event{
		{"🚗 Véhicule abandonné", "Vous trouvez une voiture encore pleine d'essence.", func(c *classes.Classe) {
			gain := 25 + rand.Intn(50)
			c.Gold += gain
			fmt.Printf("%s✔ Vous récupérez %d 💰.%s\n", classes.BrightYellow, gain, classes.Reset)
		}},
		{"📦 Cachette de survivant", "Une caisse contient du matériel médical.", func(c *classes.Classe) {
			c.PV += 35
			if c.PV > c.PVBase {
				c.PV = c.PVBase
			}
			fmt.Println(classes.BrightGreen + "✔ Vous récupérez 35 PV." + classes.Reset)
		}},
		{"☣ Embuscade", "Un groupe d'infectés surgit d'un bâtiment.", func(c *classes.Classe) {
			damage := 8 + rand.Intn(18)
			c.PV -= damage
			if c.PV < 1 {
				c.PV = 1
			}
			fmt.Printf("%s✖ Vous perdez %d PV.%s\n", classes.BrightRed, damage, classes.Reset)
		}},
		{"🎒 Sac oublié", "Vous trouvez quelques fournitures utiles.", func(c *classes.Classe) { loot(c, zones[0]) }},
	}
	e := events[rand.Intn(len(events))]
	classes.ClearScreen()
	fmt.Println(classes.TitleBox("❓ ÉVÉNEMENT"))
	fmt.Println()
	fmt.Println(classes.Panel(e.Title, e.Text))
	fmt.Println()
	e.Run(c)
	classes.Pause()
}

func loot(c *classes.Classe, z Zone) {
	items := []string{"Bandage renforcé", "Munitions récupérées", "Antidote artisanal", "Pièce d'arme", "Lampe tactique", "Sac de provisions"}
	item := items[rand.Intn(len(items))]
	if classes.AddInventory(c, item) {
		fmt.Printf("%s🎒 Loot : %s%s\n", classes.BrightGreen, item, classes.Reset)
	} else {
		bonus := 20 + rand.Intn(40)
		c.Gold += bonus
		fmt.Printf("%s🎒 Inventaire plein : le loot est revendu pour %d 💰.%s\n", classes.BrightYellow, bonus, classes.Reset)
	}
}

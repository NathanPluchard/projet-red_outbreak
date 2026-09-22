package monstre

import (
	"fmt"
	"math/rand"

	"outbreak/classes"
	"outbreak/inventaire"
)

type Monster struct {
	Name       string
	MaxHP      int
	CurrentHP  int
	Attack     int
	Initiative int
	ExpReward  int
	GoldReward int
}

func InitGoblin() Monster {
	return Monster{
		Name:       "Zombie errant",
		MaxHP:      40,
		CurrentHP:  40,
		Attack:     5,
		Initiative: 3,
		ExpReward:  30,
		GoldReward: 15,
	}
}

func InitZombieBlinde() Monster {
	return Monster{
		Name:       "Zombie blindé",
		MaxHP:      75,
		CurrentHP:  75,
		Attack:     8,
		Initiative: 2,
		ExpReward:  60,
		GoldReward: 30,
	}
}

func InitCoureur() Monster {
	return Monster{
		Name:       "Coureur infecté",
		MaxHP:      90,
		CurrentHP:  90,
		Attack:     13,
		Initiative: 8,
		ExpReward:  90,
		GoldReward: 45,
	}
}

func InitBrute() Monster {
	return Monster{
		Name:       "Brute mutante",
		MaxHP:      150,
		CurrentHP:  150,
		Attack:     18,
		Initiative: 2,
		ExpReward:  130,
		GoldReward: 70,
	}
}

func InitNecromancien() Monster {
	return Monster{
		Name:       "Nécromancien infecté",
		MaxHP:      190,
		CurrentHP:  190,
		Attack:     22,
		Initiative: 4,
		ExpReward:  200,
		GoldReward: 110,
	}
}

func InitAbomination() Monster {
	return Monster{
		Name:       "Abomination",
		MaxHP:      260,
		CurrentHP:  260,
		Attack:     28,
		Initiative: 2,
		ExpReward:  300,
		GoldReward: 180,
	}
}

func InitRoiInfectes() Monster {
	return Monster{
		Name:       "Roi des infectés",
		MaxHP:      400,
		CurrentHP:  400,
		Attack:     35,
		Initiative: 6,
		ExpReward:  500,
		GoldReward: 300,
	}
}

type MonsterEntry struct {
	Init        func() Monster
	UnlockLevel int
}

var AllMonsters = []MonsterEntry{
	{InitGoblin, 1},
	{InitZombieBlinde, 3},
	{InitCoureur, 5},
	{InitBrute, 7},
	{InitNecromancien, 9},
	{InitAbomination, 11},
	{InitRoiInfectes, 13},
}

func SelectMonster(level int) Monster {
	var pool []func() Monster
	for _, entry := range AllMonsters {
		if entry.UnlockLevel <= level {
			pool = append(pool, entry.Init)
		}
	}
	if len(pool) == 0 {
		pool = append(pool, InitGoblin)
	}
	choice := pool[rand.Intn(len(pool))]
	return choice()
}

type AttackInfo struct {
	Name        string
	Damage      int
	ManaCost    int
	UnlockLevel int
}

var AllAttacks = []AttackInfo{
	{"Attaque basique", 5, 0, 1},
	{"Coup de coude", 7, 0, 2},
	{"Crosse de fusil", 8, 0, 3},
	{"Lancer de couteau", 10, 3, 4},
	{"Cocktail Molotov", 18, 10, 5},
	{"Tir de pistolet", 14, 6, 6},
	{"Machette rouillée", 16, 4, 7},
	{"Rafale de fusil d'assaut", 20, 12, 8},
	{"Grenade artisanale", 24, 15, 9},
	{"Tir de fusil à pompe", 22, 9, 10},
	{"Batte cloutée", 19, 5, 11},
	{"Piège à mines improvisé", 28, 16, 12},
	{"Tronçonneuse", 26, 10, 13},
	{"Tir de sniper", 32, 18, 14},
	{"Lance-flammes de fortune", 30, 20, 15},
	{"Frappe explosive au C4", 42, 28, 16},
	{"Rafale de mitrailleuse lourde", 38, 24, 17},
	{"Tir groupé au lance-roquettes", 55, 35, 18},
}

func GetAttackInfo(name string) *AttackInfo {
	for i := range AllAttacks {
		if AllAttacks[i].Name == name {
			return &AllAttacks[i]
		}
	}
	return nil
}

func contains(list []string, item string) bool {
	for _, s := range list {
		if s == item {
			return true
		}
	}
	return false
}

func GainExp(c *classes.Classe, exp int) {
	c.Exp += exp
	fmt.Printf("%sVous gagnez %d points d'expérience.%s\n", classes.BrightMagenta, exp, classes.Reset)
	for c.Exp >= c.MaxExp {
		c.Exp -= c.MaxExp
		c.Level++
		c.MaxExp += 20
		c.PVBase += 10
		c.PV = c.PVBase
		c.ManaMax += 5
		c.Mana = c.ManaMax
		fmt.Println(classes.Bold + classes.BrightYellow + "★ Niveau supérieur ! Vous êtes maintenant niveau " + fmt.Sprint(c.Level) + " ★" + classes.Reset)
		for _, atk := range AllAttacks {
			if atk.UnlockLevel == c.Level && !contains(c.Attacks, atk.Name) {
				c.Attacks = append(c.Attacks, atk.Name)
				fmt.Println(classes.BrightGreen + "Nouvelle attaque débloquée : " + atk.Name + classes.Reset)
			}
		}
	}
}

func GainGold(c *classes.Classe, gold int) {
	c.Gold += gold
	fmt.Printf("%sVous récupérez %d pièces d'or sur le cadavre. 💰%s\n", classes.BrightYellow, gold, classes.Reset)
}

func IsDead(c *classes.Classe) {
	if c.PV <= 0 {
		fmt.Printf("%s%s%s s'effondre... mais se relève de justesse.%s\n", classes.Bold, classes.BrightRed, c.Nom, classes.Reset)
		c.PV = c.PVBase / 2
		fmt.Printf("PV : %s %d/%d\n", classes.HPBar(c.PV, c.PVBase), c.PV, c.PVBase)
	}
}

func CombatSpellMenu(c *classes.Classe, m *Monster) {
	fmt.Println(classes.BrightCyan + "--- Choisissez une attaque ---" + classes.Reset)
	for i, name := range c.Attacks {
		info := GetAttackInfo(name)
		if info.ManaCost > 0 {
			fmt.Printf("%s%d.%s %s (%s%d dégâts%s, %s%d mana%s)\n", classes.BrightCyan, i+1, classes.Reset, info.Name, classes.BrightRed, info.Damage, classes.Reset, classes.Cyan, info.ManaCost, classes.Reset)
		} else {
			fmt.Printf("%s%d.%s %s (%s%d dégâts%s)\n", classes.BrightCyan, i+1, classes.Reset, info.Name, classes.BrightRed, info.Damage, classes.Reset)
		}
	}
	fmt.Print(classes.Bold + "Choix : " + classes.Reset)

	choice := classes.ReadInt()
	if choice < 1 || choice > len(c.Attacks) {
		fmt.Println(classes.BrightRed + "Choix invalide, tour perdu." + classes.Reset)
		return
	}
	info := GetAttackInfo(c.Attacks[choice-1])

	if info.ManaCost > c.Mana {
		fmt.Println(classes.BrightYellow + "Mana insuffisant pour cette attaque." + classes.Reset)
		return
	}
	c.Mana -= info.ManaCost

	damage := info.Damage + c.AttaqueBase

	if c.Type == "Punk" && rand.Intn(4) == 0 {
		damage = damage * 3 / 2
		fmt.Println(classes.Bold + classes.BrightYellow + "⚡ Coup critique ! ⚡" + classes.Reset)
	}

	m.CurrentHP -= damage
	if m.CurrentHP < 0 {
		m.CurrentHP = 0
	}
	fmt.Printf("%s%s%s inflige %s%d dégâts%s à %s avec %s\n", classes.Bold, c.Nom, classes.Reset, classes.BrightGreen, damage, classes.Reset, m.Name, info.Name)
	fmt.Printf("%s : %s %d/%d\n", m.Name, classes.HPBar(m.CurrentHP, m.MaxHP), m.CurrentHP, m.MaxHP)
}

func GoblinPattern(m *Monster, c *classes.Classe, turn int) {
	baseDamage := m.Attack
	if turn%3 == 0 {
		baseDamage = m.Attack * 2
	}

	reduction := c.DefenseBase / 2
	damage := baseDamage - reduction
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
	fmt.Printf("%s%s%s inflige à %s %s%d de dégâts%s\n", classes.Bold, m.Name, classes.Reset, c.Nom, classes.BrightRed, damage, classes.Reset)
	fmt.Printf("%s : %s %d/%d\n", c.Nom, classes.HPBar(c.PV, c.PVBase), c.PV, c.PVBase)
}

func CharacterTurn(c *classes.Classe, m *Monster) {
	fmt.Println(classes.Bold + "--- Votre tour ---" + classes.Reset)
	fmt.Printf("%s1.%s Attaquer\n", classes.BrightCyan, classes.Reset)
	fmt.Printf("%s2.%s Inventaire\n", classes.BrightCyan, classes.Reset)
	fmt.Print(classes.Bold + "Choix : " + classes.Reset)

	switch classes.ReadInt() {
	case 1:
		CombatSpellMenu(c, m)
	case 2:
		inventaire.AccessInventory(c)
	default:
		fmt.Println(classes.BrightRed + "Choix invalide, tour perdu." + classes.Reset)
	}
}

func TrainingFight(c *classes.Classe) {
	monster := SelectMonster(c.Level)
	turn := 1
	fmt.Println()
	fmt.Println(classes.Bold + classes.BrightRed + "☣ Un " + monster.Name + " apparaît ! ☣" + classes.Reset)
	fmt.Printf("PV du monstre : %s %d/%d\n", classes.HPBar(monster.CurrentHP, monster.MaxHP), monster.CurrentHP, monster.MaxHP)

	c.Initiative = 10
	playerFirst := c.Initiative >= monster.Initiative

	for {
		fmt.Println()
		fmt.Println(classes.Dim + classes.Bold + "═══ Tour " + fmt.Sprint(turn) + " ═══" + classes.Reset)
		if c.Type == "Survivant" && c.PV > 0 {
			regen := c.PVBase / 25
			if regen < 1 {
				regen = 1
			}
			c.PV += regen
			if c.PV > c.PVBase {
				c.PV = c.PVBase
			}
			fmt.Printf("%s%s récupère %d PV grâce à sa résistance.%s (PV : %s %d/%d)\n", classes.BrightGreen, c.Nom, regen, classes.Reset, classes.HPBar(c.PV, c.PVBase), c.PV, c.PVBase)
		}

		if playerFirst {
			CharacterTurn(c, &monster)
			if monster.CurrentHP <= 0 {
				break
			}
			GoblinPattern(&monster, c, turn)
			IsDead(c)
		} else {
			GoblinPattern(&monster, c, turn)
			IsDead(c)
			CharacterTurn(c, &monster)
			if monster.CurrentHP <= 0 {
				break
			}
		}
		turn++
	}

	fmt.Println()
	fmt.Println(classes.Bold + classes.BrightGreen + "✔ " + monster.Name + " est vaincu ! ✔" + classes.Reset)
	GainExp(c, monster.ExpReward)
	GainGold(c, monster.GoldReward)

	c.Mana = c.ManaMax
	fmt.Printf("%sVous reprenez votre souffle, mana restauré.%s (Mana : %s %d/%d)\n", classes.BrightCyan, classes.Reset, classes.ManaBar(c.Mana, c.ManaMax), c.Mana, c.ManaMax)
}
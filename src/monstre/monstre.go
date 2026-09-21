package monstre

import (
	"fmt"

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
type AttackInfo struct {
	Name        string
	Damage      int
	ManaCost    int
	UnlockLevel int
}

var AllAttacks = []AttackInfo{
	{"Attaque basique", 5, 0, 1},
	{"Crosse de fusil", 8, 0, 3},
	{"Cocktail Molotov", 18, 10, 5},
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
	fmt.Printf("Vous gagnez %d points d'expérience.\n", exp)
	for c.Exp >= c.MaxExp {
		c.Exp -= c.MaxExp
		c.Level++
		c.MaxExp += 20
		c.PVBase += 10
		c.PV = c.PVBase
		fmt.Println("Niveau supérieur ! Vous êtes maintenant niveau", c.Level)
		for _, atk := range AllAttacks {
			if atk.UnlockLevel == c.Level && !contains(c.Attacks, atk.Name) {
				c.Attacks = append(c.Attacks, atk.Name)
				fmt.Println("Nouvelle attaque débloquée :", atk.Name)
			}
		}
	}
}

func GainGold(c *classes.Classe, gold int) {
	c.Gold += gold
	fmt.Printf("Vous récupérez %d pièces d'or sur le cadavre.\n", gold)
}

func IsDead(c *classes.Classe) {
	if c.PV <= 0 {
		fmt.Println(c.Nom, "est mort... mais se relève de justesse.")
		c.PV = c.PVBase / 2
		fmt.Printf("PV : %d / %d\n", c.PV, c.PVBase)
	}
}

func CombatSpellMenu(c *classes.Classe, m *Monster) {
	fmt.Println("--- Choisissez une attaque ---")
	for i, name := range c.Attacks {
		info := GetAttackInfo(name)
		if info.ManaCost > 0 {
			fmt.Printf("%d. %s (%d dégâts, %d mana)\n", i+1, info.Name, info.Damage, info.ManaCost)
		} else {
			fmt.Printf("%d. %s (%d dégâts)\n", i+1, info.Name, info.Damage)
		}
	}
	fmt.Print("Choix : ")

	choice := classes.ReadInt()
	if choice < 1 || choice > len(c.Attacks) {
		fmt.Println("Choix invalide, tour perdu.")
		return
	}
	info := GetAttackInfo(c.Attacks[choice-1])

	if info.ManaCost > c.Mana {
		fmt.Println("Mana insuffisant pour cette attaque.")
		return
	}
	c.Mana -= info.ManaCost

	m.CurrentHP -= info.Damage
	if m.CurrentHP < 0 {
		m.CurrentHP = 0
	}
	fmt.Printf("%s inflige %d dégâts à %s avec %s\n", c.Nom, info.Damage, m.Name, info.Name)
	fmt.Printf("%s : PV %d / %d\n", m.Name, m.CurrentHP, m.MaxHP)
}

func GoblinPattern(m *Monster, c *classes.Classe, turn int) {
	damage := m.Attack
	if turn%3 == 0 {
		damage = m.Attack * 2
	}
	c.PV -= damage
	if c.PV < 0 {
		c.PV = 0
	}
	fmt.Printf("%s inflige à %s %d de dégâts\n", m.Name, c.Nom, damage)
	fmt.Printf("%s : PV %d / %d\n", c.Nom, c.PV, c.PVBase)
}

func CharacterTurn(c *classes.Classe, m *Monster) {
	fmt.Println("--- Votre tour ---")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	fmt.Print("Choix : ")

	switch classes.ReadInt() {
	case 1:
		CombatSpellMenu(c, m)
	case 2:
		inventaire.AccessInventory(c)
	default:
		fmt.Println("Choix invalide, tour perdu.")
	}
}

func TrainingFight(c *classes.Classe) {
	monster := InitGoblin()
	turn := 1
	fmt.Println("Un", monster.Name, "apparaît !")

	c.Initiative = 10
	playerFirst := c.Initiative >= monster.Initiative

	for {
		fmt.Println("=== Tour", turn, "===")
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

	fmt.Println(monster.Name, "est vaincu !")
	GainExp(c, monster.ExpReward)
	GainGold(c, monster.GoldReward)
}
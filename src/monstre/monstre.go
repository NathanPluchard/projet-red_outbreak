package monstre

import (
"fmt"
"math/rand"
"time"


"outbreak/classes"


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

return pool[rand.Intn(len(pool))]()


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


fmt.Printf(
	"%s+%d XP%s\n",
	classes.BrightMagenta,
	exp,
	classes.Reset,
)

for c.Exp >= c.MaxExp {
	c.Exp -= c.MaxExp
	c.Level++

	c.MaxExp += 20
	c.PVBase += 10
	c.PV = c.PVBase

	c.ManaMax += 5
	c.Mana = c.ManaMax

	fmt.Printf(
		"%s★ NIVEAU %d ★%s\n",
		classes.BrightYellow,
		c.Level,
		classes.Reset,
	)

	for _, atk := range AllAttacks {
		if atk.UnlockLevel == c.Level &&
			!contains(c.Attacks, atk.Name) {

			c.Attacks = append(c.Attacks, atk.Name)

			fmt.Printf(
				"%s↳ Nouvelle attaque débloquée : %s%s\n",
				classes.BrightGreen,
				atk.Name,
				classes.Reset,
			)
		}
	}
}


}

func GainGold(c *classes.Classe, gold int) {
c.Gold += gold


fmt.Printf(
	"%s+%d 💰%s\n",
	classes.BrightYellow,
	gold,
	classes.Reset,
)


}

func IsDead(c *classes.Classe) {
if c.PV <= 0 {
c.PV = c.PVBase / 2

	fmt.Printf(
		"%s⚠ %s tombe mais se relève avec %d PV.%s\n",
		classes.BrightRed,
		c.Nom,
		c.PV,
		classes.Reset,
	)
}


}


func bestAttack(c *classes.Classe) *AttackInfo {
var best *AttackInfo


for _, name := range c.Attacks {
	attack := GetAttackInfo(name)

	if attack == nil {
		continue
	}

	if attack.ManaCost > c.Mana {
		continue
	}

	if best == nil || attack.Damage > best.Damage {
		best = attack
	}
}

if best == nil {
	best = GetAttackInfo("Attaque basique")
}

return best


}

// Le joueur attaque automatiquement.
func autoPlayerAttack(c *classes.Classe, m *Monster) {


attack := bestAttack(c)

if attack.ManaCost <= c.Mana {
	c.Mana -= attack.ManaCost
}

damage := attack.Damage + c.AttaqueBase


if c.Type == "Punk" && rand.Intn(4) == 0 {

	damage = damage * 3 / 2

	fmt.Println(
		classes.Bold +
			classes.BrightYellow +
			"⚡ COUP CRITIQUE ! ⚡" +
			classes.Reset,
	)
}

m.CurrentHP -= damage

if m.CurrentHP < 0 {
	m.CurrentHP = 0
}

fmt.Printf(
	"%s%-28s%s → %s-%d PV%s\n",
	classes.BrightCyan,
	attack.Name,
	classes.Reset,
	classes.BrightGreen,
	damage,
	classes.Reset,
)


}

// Attaque automatique d'un zombie.
func monsterAttack(
m *Monster,
c *classes.Classe,
turn int,
) {


damage := m.Attack


if turn%3 == 0 {
	damage *= 2
}

damage -= c.DefenseBase / 2


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
	"%s%-28s%s → %s-%d PV%s\n",
	classes.BrightRed,
	m.Name,
	classes.Reset,
	classes.BrightRed,
	damage,
	classes.Reset,
)


}

// Affichage du combat.
func showCombat(
c *classes.Classe,
monsters []Monster,
turn int,
) {


classes.ClearScreen()

fmt.Println(
	classes.TitleBox(
		"☣  SIMULATION DE COMBAT  ☣",
	),
)

fmt.Printf(
	"\n%sTOUR %02d%s\n",
	classes.Bold+classes.BrightYellow,
	turn,
	classes.Reset,
)

fmt.Printf(
	"\n%s%s%s\n",
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

for i := range monsters {

	if monsters[i].CurrentHP <= 0 {
		fmt.Printf(
			"[%d] %s☠ VAINCU%s\n",
			i+1,
			classes.Dim,
			classes.Reset,
		)

		continue
	}

	fmt.Printf(
		"[%d] %-22s %s %d/%d\n",
		i+1,
		monsters[i].Name,
		classes.HPBar(
			monsters[i].CurrentHP,
			monsters[i].MaxHP,
		),
		monsters[i].CurrentHP,
		monsters[i].MaxHP,
	)
}

fmt.Println()


}


func SimulationFight(c *classes.Classe) {


// Le nombre de zombies augmente avec le niveau.
numberOfMonsters := 2 + c.Level/4

if numberOfMonsters > 6 {
	numberOfMonsters = 6
}

monsters := make([]Monster, 0, numberOfMonsters)

for i := 0; i < numberOfMonsters; i++ {
	monsters = append(
		monsters,
		SelectMonster(c.Level),
	)
}

fmt.Printf(
	"%s☣ %d infectés détectés !%s\n",
	classes.BrightRed,
	numberOfMonsters,
	classes.Reset,
)

fmt.Println(
	classes.Dim +
		"Lancement de la simulation..." +
		classes.Reset,
)

time.Sleep(800 * time.Millisecond)

turn := 1

for {

	// Vérifie s'il reste des ennemis.
	alive := 0

	for _, monster := range monsters {
		if monster.CurrentHP > 0 {
			alive++
		}
	}

	if alive == 0 {
		break
	}

	showCombat(c, monsters, turn)

	
	if c.Type == "Survivant" && c.PV > 0 {

		regen := c.PVBase / 25

		if regen < 1 {
			regen = 1
		}

		c.PV += regen

		if c.PV > c.PVBase {
			c.PV = c.PVBase
		}

		fmt.Printf(
			"%s♥ Régénération : +%d PV%s\n",
			classes.BrightGreen,
			regen,
			classes.Reset,
		)
	}

	// Cible automatiquement le zombie ayant
	// le moins de PV.
	target := -1

	for i := range monsters {

		if monsters[i].CurrentHP <= 0 {
			continue
		}

		if target == -1 ||
			monsters[i].CurrentHP <
				monsters[target].CurrentHP {

			target = i
		}
	}

	if target >= 0 {

		fmt.Println(
			classes.BrightCyan +
				"⚔ VOTRE TOUR" +
				classes.Reset,
		)

		autoPlayerAttack(
			c,
			&monsters[target],
		)

		time.Sleep(500 * time.Millisecond)

		if monsters[target].CurrentHP <= 0 {

			fmt.Printf(
				"%s✔ %s éliminé !%s\n",
				classes.BrightGreen,
				monsters[target].Name,
				classes.Reset,
			)

			time.Sleep(300 * time.Millisecond)
		}
	}

	
	fmt.Println(
		"\n" +
			classes.BrightRed +
			"☠ TOUR DES INFECTÉS" +
			classes.Reset,
	)

	for i := range monsters {

		if monsters[i].CurrentHP <= 0 {
			continue
		}

		if c.PV <= 0 {
			break
		}

		monsterAttack(
			&monsters[i],
			c,
			turn,
		)

		time.Sleep(300 * time.Millisecond)
	}

	IsDead(c)

	turn++

	time.Sleep(500 * time.Millisecond)
}


totalXP := 0
totalGold := 0

for _, monster := range monsters {
	totalXP += monster.ExpReward
	totalGold += monster.GoldReward
}

classes.ClearScreen()

fmt.Println(
	classes.TitleBox(
		"✔  VAGUE ÉLIMINÉE  ✔",
	),
)

fmt.Printf(
	"\n%sRécompenses de la vague%s\n",
	classes.Bold,
	classes.Reset,
)

fmt.Printf(
	"  %s+%d XP%s\n",
	classes.BrightMagenta,
	totalXP,
	classes.Reset,
)

fmt.Printf(
	"  %s+%d 💰%s\n",
	classes.BrightYellow,
	totalGold,
	classes.Reset,
)

GainExp(c, totalXP)
GainGold(c, totalGold)

// Le mana revient à son maximum après le combat.
c.Mana = c.ManaMax

fmt.Printf(
	"\n%sMana restauré : %d/%d%s\n",
	classes.BrightCyan,
	c.Mana,
	c.ManaMax,
	classes.Reset,
)

classes.Pause()


}

func TrainingFight(c *classes.Classe) {
SimulationFight(c)
}


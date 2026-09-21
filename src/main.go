package src

import

func characterCreation() Character {
	fmt.Print("Entrez votre nom (lettres uniquement) : ")
	var name string
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if isAlpha(input) && input != "" {
			name = formatName(input)
			break
		}
		fmt.Print("Nom invalide, réessayez : ")
	}
}
 
	fmt.Println("Choisissez votre classe :")
	fmt.Println("1. Militaire (120 PV)")
	fmt.Println("2. Médecin (80 PV)")
	fmt.Println("3. Pillard (100 PV)")
 
	var class string
	var maxHP int
	for {
		choice := readInt()
		switch choice {
		case 1:
			class = "Survivanrt"
			maxHP = 200
		case 2:
			class = "Médecin de fortune"
			maxHP = 100
		case 3:
			class = "Punk"
			maxHP = 150
        case 4:
            class = "Sauveur"
            maxHP = 80
		default:
			fmt.Println("Choix invalide, réessayez.")
			continue
		}
		break
	}
func isAlpha(s string) bool {
	for _, r := range s {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
			return false
		}
	}
	return true
}
 
func formatName(s string) string {
	s = strings.ToLower(s)
	return strings.ToUpper(s[:1]) + s[1:]
}

func displayInfo(c *Classe) {
	fmt.Println("---------------------------------")
	fmt.Printf("Nom       : %s\n", c.Name)
	fmt.Printf("Classe    : %s\n", c.Class)
	fmt.Printf("Niveau    : %d\n", c.Level)
	fmt.Printf("PV        : %d / %d\n", c.CurrentHP, c.MaxHP)
	fmt.Printf("Mana      : %d / %d\n", c.Mana, c.MaxMana)
	fmt.Printf("Or        : %d\n", c.Gold)
	fmt.Printf("Exp       : %d / %d\n", c.Exp, c.MaxExp)
	fmt.Printf("Sorts     : %s\n", strings.Join(c.Skills, ", "))
	fmt.Printf("Casque    : %s\n", displayOrEmpty(c.Equip.Head))
	fmt.Printf("Torse     : %s\n", displayOrEmpty(c.Equip.Torso))
	fmt.Printf("Pieds     : %s\n", displayOrEmpty(c.Equip.Feet))
	fmt.Println("---------------------------------")
}
 
func displayOrEmpty(s string) string {
	if s == "" {
		return "Aucun"
	}
	return s
}

type AttackInfo struct {
	Name        string
	Damage      int
	ManaCost    int
	UnlockLevel int
}

var allAttacks = []AttackInfo{
	{"Attaque basique", 5, 0, 1},
	{"Crosse de fusil", 8, 0, 3},
	{"Cocktail Molotov", 18, 10, 5},
}

func getAttackInfo(name string) *AttackInfo {
	for i := range allAttacks {
		if allAttacks[i].Name == name {
			return &allAttacks[i]
		}
	}
	return nil
}

func gainExp(c *Classe, exp int) {
	c.Exp += exp
	fmt.Printf("Vous gagnez %d points d'expérience.\n", exp)

	for c.Exp >= c.MaxExp {
		c.Exp -= c.MaxExp
		c.Level++
		c.MaxExp += 20
		c.MaxHP += 10
		c.CurrentHP = c.MaxHP
		fmt.Println("Niveau supérieur ! Vous êtes maintenant niveau", c.Level)

		for _, atk := range allAttacks {
			if atk.UnlockLevel == c.Level && !contains(c.Attacks, atk.Name) {
				c.Attacks = append(c.Attacks, atk.Name)
				fmt.Println("Nouvelle attaque débloquée :", atk.Name)
			}
		}
	}
}
func combatSpellMenu(c *Classe, m *Monster) {
	fmt.Println("--- Choisissez une attaque ---")
	for i, name := range c.Attacks {
		info := getAttackInfo(name)
		if info.ManaCost > 0 {
			fmt.Printf("%d. %s (%d dégâts, %d mana)\n", i+1, info.Name, info.Damage, info.ManaCost)
		} else {
			fmt.Printf("%d. %s (%d dégâts)\n", i+1, info.Name, info.Damage)
		}
	}
	fmt.Print("Choix : ")

	choice := readInt()
	if choice < 1 || choice > len(c.Attacks) {
		fmt.Println("Choix invalide, tour perdu.")
		return
	}

	attackName := c.Attacks[choice-1]
	info := getAttackInfo(attackName)

	if info.ManaCost > c.Mana {
		fmt.Println("Mana insuffisant pour cette attaque.")
		return
	}
	c.Mana -= info.ManaCost

	m.CurrentHP -= info.Damage
	if m.CurrentHP < 0 {
		m.CurrentHP = 0
	}
	fmt.Printf("%s inflige %d dégâts à %s avec %s\n", c.Name, info.Damage, m.Name, info.Name)
	fmt.Printf("%s : PV %d / %d\n", m.Name, m.CurrentHP, m.MaxHP)
}

func mainMenu(c *Classe) {
	for {
		fmt.Println("\n===== MENU =====")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Marchand (Le Troqueur)")
		fmt.Println("4. Forgeron (Le Bricoleur)")
		fmt.Println("5. Entrainement")
		fmt.Println("6. Qui sont-ils")
		fmt.Println("0. Quitter")
		fmt.Print("Choix : ")
 
		choice := readInt()
		switch choice {
		case 1:
			displayInfo(c)
		case 2:
			accessInventory(c)
		case 3:
			merchant(c)
		case 4:
			blacksmith(c)
		case 5:
			trainingFight(c)
		case 6:
			showArtists()
		case 0:
			fmt.Println("À bientôt, survivant.")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
 
 
func readInt() int {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	n, err := strconv.Atoi(input)
	if err != nil {
		return -1
	}
	return n
}
 
 
func main() {
	fmt.Println("Bienvenue dans Outbreak")
	c1 := characterCreation()
	mainMenu(&c1)
}


package inventaire

import (
	"fmt"
)

func canAddItem(c *Classe) bool {
	return len(c.Inventory) < c.MaxInventory
}

func addInventory(c *Classe, item string) bool {
	if !canAddItem(c) {
		fmt.Println("Inventaire plein ! Impossible d'ajouter :", item)
		return false
	}
	c.Inventory = append(c.Inventory, item)
	fmt.Println(item, "ajouté à l'inventaire.")
	return true
}


func removeInventory(c *Classe, item string) bool {
	for i, it := range c.Inventory {
		if it == item {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return true
		}
	}
	return false // l'objet n'était pas dans l'inventaire
}

// ---------- Compter les exemplaires d'un objet ----------

func countItem(c *Classe, item string) int {
	count := 0
	for _, it := range c.Inventory {
		if it == item {
			count++
		}
	}
	return count
}


func accessInventory(c *Classe) {
	if len(c.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}

	fmt.Println("--- Inventaire ---")
	for i, item := range c.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("0. Retour")
	fmt.Print("Choisissez un objet à utiliser (numéro) : ")

	choice := readInt()
	if choice == 0 || choice < 0 || choice > len(c.Inventory) {
		return
	}

	item := c.Inventory[choice-1]
	useItem(c, item)
}

func useItem(c *Classe, item string) {
	switch item {
	case "Kit de premiers soins":
		takePot(c)
	case "Morsure infectée":
		poisonPot(c)
	case "Plan Molotov":
		spellBook(c)
	case "Casquette renforcée", "Veste de travail", "Chaussures de randonnée",
		"Casque de chantier renforcé", "Veste tactique légère", "Bottes de randonnée renforcées":
		equipItem(c, item)
	default:
		fmt.Println("Cet objet ne peut pas être utilisé ici.")
	}
}
func upgradeInventorySlot(c *Classe) {
	if c.InventoryUpgrades >= 3 {
		fmt.Println("Vous avez déjà atteint le nombre maximum d'améliorations d'inventaire.")
		return
	}
	removeInventory(c, "Sac à dos militaire supplémentaire")
	c.MaxInventory += 10
	c.InventoryUpgrades++
	fmt.Println("Capacité d'inventaire augmentée ! Nouvelle capacité :", c.MaxInventory)
}
func findTierSlot(item string) (tier *ArmorTier, slot string) {
	for i := range armorTiers {
		t := &armorTiers[i]
		if t.Head == item {
			return t, "head"
		}
		if t.Torso == item {
			return t, "torso"
		}
		if t.Feet == item {
			return t, "feet"
		}
	}
	return nil, ""
}

func hpBonus(slot string, level int) int {
	switch slot {
	case "head":
		return level * 5
	case "torso":
		return level * 10
	case "feet":
		return level * 7
	}
	return 0
}

func equipItem(c *Classe, item string) {
	tier, slot := findTierSlot(item)
	if tier == nil {
		fmt.Println("Cet objet ne peut pas être équipé.")
		return
	}

	removeInventory(c, item)
	bonus := hpBonus(slot, tier.Level)

	switch slot {
	case "head":
		if c.Equip.Head != "" {
			oldTier, _ := findTierSlot(c.Equip.Head)
			if oldTier != nil {
				c.MaxHP -= hpBonus("head", oldTier.Level)
			}
			addInventory(c, c.Equip.Head)
		}
		c.Equip.Head = item

	case "torso":
		if c.Equip.Torso != "" {
			oldTier, _ := findTierSlot(c.Equip.Torso)
			if oldTier != nil {
				c.MaxHP -= hpBonus("torso", oldTier.Level)
			}
			addInventory(c, c.Equip.Torso)
		}
		c.Equip.Torso = item

	case "feet":
		if c.Equip.Feet != "" {
			oldTier, _ := findTierSlot(c.Equip.Feet)
			if oldTier != nil {
				c.MaxHP -= hpBonus("feet", oldTier.Level)
			}
			addInventory(c, c.Equip.Feet)
		}
		c.Equip.Feet = item
	}

	c.MaxHP += bonus
	fmt.Println("Vous équipez :", item, "( +", bonus, "PV max )")
}
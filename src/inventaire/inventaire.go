package inventaire

import (
	"fmt"

	"outbreak/classes"
	"outbreak/forgeron"
	"outbreak/potion"
)

func AccessInventory(c *classes.Classe) {
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

	choice := classes.ReadInt()
	if choice == 0 || choice < 0 || choice > len(c.Inventory) {
		return
	}
	UseItem(c, c.Inventory[choice-1])
}

func UseItem(c *classes.Classe, item string) {
	for _, p := range potion.PotionsDeVieDisponibles {
		if p.Nom == item {
			potion.UtiliserPotion(c, item)
			return
		}
	}
	for _, p := range potion.PotionsDePoisonDisponibles {
		if p.Nom == item {
			potion.UtiliserPotion(c, item)
			return
		}
	}
	if tier, _ := forgeron.FindTierSlot(item); tier != nil {
		forgeron.EquipItem(c, item)
		return
	}
	if item == "Sac à dos militaire supplémentaire" {
		UpgradeInventorySlot(c)
		return
	}
	fmt.Println("Cet objet ne peut pas être utilisé ici.")
}

func UpgradeInventorySlot(c *classes.Classe) {
	if c.InventoryUpgrades >= 3 {
		fmt.Println("Vous avez déjà atteint le nombre maximum d'améliorations d'inventaire.")
		return
	}
	classes.RemoveInventory(c, "Sac à dos militaire supplémentaire")
	c.MaxInventory += 10
	c.InventoryUpgrades++
	fmt.Println("Capacité d'inventaire augmentée ! Nouvelle capacité :", c.MaxInventory)
}

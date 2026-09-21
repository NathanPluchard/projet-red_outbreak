package monstre

import (
	"fmt"
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

func initGoblin() Monster {
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

func gainGold(c *Classe, gold int) {
	c.Gold += gold
	fmt.Printf("Vous récupérez %d pièces d'or sur le cadavre.\n", gold)
}

func trainingFight(c *Classe) {
	monster := initGoblin()
	turn := 1

	fmt.Println("Un", monster.Name, "apparaît !")

	c.Initiative = 10
	playerFirst := c.Initiative >= monster.Initiative

	for {
		fmt.Println("=== Tour", turn, "===")

		if playerFirst {
			charaterTurn(c, &monster)
			if monster.CurrentHP <= 0 {
				break
			}
			goblinPattern(&monster, c, turn)
			isDead(c)
		} else {
			goblinPattern(&monster, c, turn)
			isDead(c)
			charaterTurn(c, &monster)
			if monster.CurrentHP <= 0 {
				break
			}
		}

		turn++
	}

	fmt.Println(monster.Name, "est vaincu !")
	gainExp(c, monster.ExpReward)
	gainGold(c, monster.GoldReward)
}
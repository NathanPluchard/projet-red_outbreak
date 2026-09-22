package casino

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"outbreak/classes"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Casino(c *classes.Classe) {

	for {

		classes.ClearScreen()

		fmt.Println(
			classes.TitleBox(
				"☠ LE CASINO DE LA ZONE MORTE ☠",
			),
		)

		fmt.Printf(
			"\n%sVotre portefeuille : %s%d 💰%s\n\n",
			classes.Bold,
			classes.BrightYellow,
			c.Gold,
			classes.Reset,
		)

		fmt.Println(
			classes.Panel(
				"TABLES OUVERTES",

				classes.BrightRed+
					"[1] "+
					classes.Reset+
					"🎡 Roulette",

				classes.BrightBlue+
					"[2] "+
					classes.Reset+
					"🃏 Blackjack",

				classes.Dim+
					"[0] "+
					classes.Reset+
					"Sortir du casino",
			),
		)

		fmt.Print(
			"\n" +
				classes.BrightCyan +
				"Table > " +
				classes.Reset,
		)

		choix :=
			classes.ReadInt()

		switch choix {

		case 1:
			roulette(c)

		case 2:
			blackjack(c)

		case 0:
			return

		default:

			fmt.Println(
				classes.BrightRed +
					"✖ Table inconnue." +
					classes.Reset,
			)

			classes.Pause()
		}
	}

}

func askBet(c *classes.Classe) int {

	if c.Gold <= 0 {

		fmt.Println(
			classes.BrightRed +
				"✖ Vous n'avez plus d'or." +
				classes.Reset,
		)

		classes.Pause()

		return 0
	}

	fmt.Printf(
		"\n%sOr disponible : %d 💰%s\n",
		classes.BrightYellow,
		c.Gold,
		classes.Reset,
	)

	fmt.Printf(
		"Mise (1-%d, 0 pour annuler) : ",
		c.Gold,
	)

	bet :=
		classes.ReadInt()

	if bet <= 0 {
		return 0
	}

	if bet > c.Gold {

		fmt.Println(
			classes.BrightRed +
				"✖ Vous ne pouvez pas miser plus que votre or." +
				classes.Reset,
		)

		classes.Pause()

		return 0
	}

	return bet

}

var rouletteNumbers = []int{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
	10,
	11,
	12,
	13,
	14,
	15,
	16,
	17,
	18,
	19,
	20,
	21,
	22,
	23,
	24,
	25,
	26,
	27,
	28,
	29,
	30,
	31,
	32,
	33,
	34,
	35,
	36,
}

// Détermine si un numéro est rouge.
func isRed(number int) bool {

	if number == 0 {
		return false
	}

	return ((number >= 1 && number <= 10 && number%2 == 1) ||
		(number >= 11 && number <= 18 && number%2 == 0) ||
		(number >= 19 && number <= 28 && number%2 == 1) ||
		(number >= 29 && number <= 36 && number%2 == 0))

}

func roulette(c *classes.Classe) {

	classes.ClearScreen()

	fmt.Println(
		classes.TitleBox(
			"🎡 ROULETTE — LE TOUR DU DESTIN",
		),
	)

	fmt.Println()

	bet :=
		askBet(c)

	if bet == 0 {
		return
	}

	fmt.Println()

	fmt.Println(
		classes.Panel(
			"CHOISISSEZ VOTRE PARI",

			"1  🔴 Rouge       • gain x1",

			"2  ⚫ Noir        • gain x1",

			"3  Pair          • gain x1",

			"4  Impair        • gain x1",

			"5  Numéro exact  • gain x35",
		),
	)

	fmt.Print(
		"\nPari > ",
	)

	kind :=
		classes.ReadInt()

	if kind < 1 || kind > 5 {

		fmt.Println(
			classes.BrightRed +
				"✖ Pari invalide." +
				classes.Reset,
		)

		classes.Pause()

		return
	}

	target := 0

	if kind == 5 {

		fmt.Print(
			"Numéro choisi (0-36) > ",
		)

		target =
			classes.ReadInt()

		if target < 0 || target > 36 {

			fmt.Println(
				classes.BrightRed +
					"✖ Numéro invalide." +
					classes.Reset,
			)

			classes.Pause()

			return
		}
	}

	classes.ClearScreen()

	fmt.Println(
		classes.TitleBox(
			"🎡 LA ROULETTE TOURNE",
		),
	)

	fmt.Println()

	fmt.Printf(
		"Mise : %d 💰\n\n",
		bet,
	)

	fmt.Print(
		classes.BrightYellow +
			"      ",
	)

	for i := 0; i < 24; i++ {

		number :=
			rouletteNumbers[rand.Intn(
				len(rouletteNumbers),
			)]

		fmt.Printf(
			"\r      🎡  %02d  🎡",
			number,
		)

		time.Sleep(
			70 * time.Millisecond,
		)
	}

	result :=
		rouletteNumbers[rand.Intn(
			len(rouletteNumbers),
		)]

	fmt.Printf(
		"\r      🎡  %02d  🎡\n\n",
		result,
	)

	win := false
	multiplier := 1

	switch kind {

	case 1:

		win = isRed(result)

	case 2:

		win =
			result != 0 &&
				!isRed(result)

	case 3:

		win =
			result != 0 &&
				result%2 == 0

	case 4:

		win =
			result%2 == 1

	case 5:

		win =
			result == target

		multiplier = 35
	}

	if win {

		gain :=
			bet * multiplier

		c.Gold += gain

		fmt.Println(
			classes.BrightGreen +
				"╔══════════════════════════════════════╗" +
				classes.Reset,
		)

		fmt.Println(
			classes.BrightGreen +
				"║           ✔ VOUS AVEZ GAGNÉ !       ║" +
				classes.Reset,
		)

		fmt.Println(
			classes.BrightGreen +
				"╚══════════════════════════════════════╝" +
				classes.Reset,
		)

		fmt.Printf(
			"\n%s+%d 💰%s\n",
			classes.BrightYellow,
			gain,
			classes.Reset,
		)

	} else {

		c.Gold -= bet

		fmt.Println(
			classes.BrightRed +
				"╔══════════════════════════════════════╗" +
				classes.Reset,
		)

		fmt.Println(
			classes.BrightRed +
				"║             ✖ VOUS AVEZ PERDU       ║" +
				classes.Reset,
		)

		fmt.Println(
			classes.BrightRed +
				"╚══════════════════════════════════════╝" +
				classes.Reset,
		)

		fmt.Printf(
			"\n%s-%d 💰%s\n",
			classes.BrightRed,
			bet,
			classes.Reset,
		)
	}

	fmt.Printf(
		"\nPortefeuille : %s%d 💰%s\n",
		classes.BrightYellow,
		c.Gold,
		classes.Reset,
	)

	classes.Pause()

}

type card struct {
	rank int
	suit string
}

// Valeur d'une carte.
func (c card) value() int {

	if c.rank >= 10 {
		return 10
	}

	if c.rank == 1 {
		return 11
	}

	return c.rank

}

func cardName(c card) string {

	names :=
		map[int]string{
			1:  "A",
			11: "J",
			12: "Q",
			13: "K",
		}

	rank :=
		fmt.Sprintf(
			"%d",
			c.rank,
		)

	if name, ok := names[c.rank]; ok {
		rank = name
	}

	return rank + c.suit

}

func score(hand []card) int {

	total := 0
	aces := 0

	for _, c := range hand {

		total += c.value()

		if c.rank == 1 {
			aces++
		}
	}

	for total > 21 && aces > 0 {

		total -= 10

		aces--
	}

	return total

}

// Pioche une carte.
func draw(deck *[]card) card {

	index :=
		rand.Intn(
			len(*deck),
		)

	card :=
		(*deck)[index]

	*deck =
		append(
			(*deck)[:index],
			(*deck)[index+1:]...,
		)

	return card

}

func cardNames(hand []card) string {

	names :=
		make(
			[]string,
			len(hand),
		)

	for i, c := range hand {

		names[i] =
			cardName(c)
	}

	return strings.Join(
		names,
		"  ",
	)

}

func blackjack(c *classes.Classe) {

	classes.ClearScreen()

	fmt.Println(
		classes.TitleBox(
			"🃏 BLACKJACK — TABLE 21",
		),
	)

	fmt.Println()

	bet :=
		askBet(c)

	if bet == 0 {
		return
	}

	suits :=
		[]string{
			"♠",
			"♥",
			"♦",
			"♣",
		}

	deck :=
		make(
			[]card,
			0,
			52,
		)

	for _, suit := range suits {

		for rank := 1; rank <= 13; rank++ {

			deck =
				append(
					deck,
					card{
						rank: rank,
						suit: suit,
					},
				)
		}
	}

	player :=
		[]card{
			draw(&deck),
			draw(&deck),
		}

	dealer :=
		[]card{
			draw(&deck),
			draw(&deck),
		}

	for {

		classes.ClearScreen()

		fmt.Println(
			classes.TitleBox(
				"🃏 BLACKJACK",
			),
		)

		fmt.Println()

		fmt.Printf(
			"%sVotre main%s\n",
			classes.Bold,
			classes.Reset,
		)

		fmt.Printf(
			"  %s\n",
			cardNames(player),
		)

		fmt.Printf(
			"  Score : %d\n\n",
			score(player),
		)

		fmt.Printf(
			"%sMain du croupier%s\n",
			classes.Bold,
			classes.Reset,
		)

		fmt.Printf(
			"  %s  🂠\n",
			cardName(dealer[0]),
		)

		fmt.Printf(
			"\n%sMise : %d 💰%s\n",
			classes.BrightYellow,
			bet,
			classes.Reset,
		)

		// Le joueur a dépassé 21.
		if score(player) > 21 {
			break
		}

		fmt.Println()

		fmt.Println(
			classes.Panel(
				"VOTRE CHOIX",

				"1  Tirer une carte",

				"2  Rester",
			),
		)

		fmt.Print(
			"\nAction > ",
		)

		choice :=
			classes.ReadInt()

		switch choice {

		case 1:

			player =
				append(
					player,
					draw(&deck),
				)

		case 2:

			goto dealerTurn

		default:

			fmt.Println(
				classes.BrightRed +
					"✖ Choix invalide." +
					classes.Reset,
			)

			classes.Pause()
		}
	}

dealerTurn:

	for score(player) <= 21 &&
		score(dealer) < 17 {

		dealer =
			append(
				dealer,
				draw(&deck),
			)
	}

	playerScore :=
		score(player)

	dealerScore :=
		score(dealer)

	classes.ClearScreen()

	fmt.Println(
		classes.TitleBox(
			"🃏 RÉSULTAT DU BLACKJACK",
		),
	)

	fmt.Println()

	fmt.Printf(
		"%sVous%s : %s  → %d\n",
		classes.BrightCyan,
		classes.Reset,
		cardNames(player),
		playerScore,
	)

	fmt.Printf(
		"%sCroupier%s : %s  → %d\n\n",
		classes.BrightYellow,
		classes.Reset,
		cardNames(dealer),
		dealerScore,
	)

	switch {

	case playerScore > 21:

		c.Gold -= bet

		fmt.Printf(
			"%s✖ Vous dépassez 21.%s\n",
			classes.BrightRed,
			classes.Reset,
		)

		fmt.Printf(
			"%s-%d 💰%s\n",
			classes.BrightRed,
			bet,
			classes.Reset,
		)

	case dealerScore > 21:

		c.Gold += bet

		fmt.Printf(
			"%s✔ Le croupier dépasse 21 !%s\n",
			classes.BrightGreen,
			classes.Reset,
		)

		fmt.Printf(
			"%s+%d 💰%s\n",
			classes.BrightGreen,
			bet,
			classes.Reset,
		)

	case playerScore > dealerScore:

		c.Gold += bet

		fmt.Printf(
			"%s✔ VOUS GAGNEZ !%s\n",
			classes.BrightGreen,
			classes.Reset,
		)

		fmt.Printf(
			"%s+%d 💰%s\n",
			classes.BrightGreen,
			bet,
			classes.Reset,
		)

	case playerScore == dealerScore:

		fmt.Println(
			classes.BrightYellow +
				"↔ ÉGALITÉ — votre mise est remboursée." +
				classes.Reset,
		)

	default:

		c.Gold -= bet

		fmt.Printf(
			"%s✖ Le croupier gagne.%s\n",
			classes.BrightRed,
			classes.Reset,
		)

		fmt.Printf(
			"%s-%d 💰%s\n",
			classes.BrightRed,
			bet,
			classes.Reset,
		)
	}

	fmt.Printf(
		"\nPortefeuille : %s%d 💰%s\n",
		classes.BrightYellow,
		c.Gold,
		classes.Reset,
	)

	classes.Pause()

}

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

	colorLabel := "🟢 Vert"

	if result != 0 {
		if isRed(result) {
			colorLabel = "🔴 Rouge"
		} else {
			colorLabel = "⚫ Noir"
		}
	}

	fmt.Printf(
		"%sC'est tombé sur le %d, de couleur %s !%s\n\n",
		classes.Bold,
		result,
		colorLabel,
		classes.Reset,
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

func totalWagered(hands []bjHand) int {

	total := 0

	for _, h := range hands {
		total += h.bet
	}

	return total

}

type bjHand struct {
	cards       []card
	bet         int
	doubled     bool
	busted      bool
	isSplitAce  bool
	surrendered bool
}

// Nombre maximum de mains simultanées (main initiale + 3 splits).
const maxBlackjackHands = 4

// Demande au joueur s'il souhaite prendre une assurance
// lorsque le croupier montre un As. Renvoie le montant misé
// (0 si refusée).
func askInsurance(c *classes.Classe, bet int) int {

	maxInsurance := bet / 2

	if maxInsurance <= 0 {
		return 0
	}

	if maxInsurance > c.Gold {
		maxInsurance = c.Gold
	}

	if maxInsurance <= 0 {
		return 0
	}

	classes.ClearScreen()

	fmt.Println(
		classes.TitleBox(
			"🛡️ ASSURANCE",
		),
	)

	fmt.Println()

	fmt.Println(
		classes.Panel(
			"LE CROUPIER MONTRE UN AS",

			"Vous pouvez vous assurer contre un Blackjack.",

			fmt.Sprintf(
				"Mise d'assurance : 0-%d 💰 (paiement 2:1)",
				maxInsurance,
			),
		),
	)

	fmt.Printf(
		"\nAssurance (0 pour refuser, max %d) > ",
		maxInsurance,
	)

	amount :=
		classes.ReadInt()

	if amount <= 0 {
		return 0
	}

	if amount > maxInsurance {
		amount = maxInsurance
	}

	return amount

}

// Crée un sabot composé de plusieurs paquets de 52 cartes,
// comme sur une vraie table de casino. Utiliser deux paquets
// (au lieu d'un seul) rend le comptage de cartes moins fiable
// et évite de manquer de cartes lorsque plusieurs mains sont
// ouvertes en même temps (splits multiples).
func newShoe(packs int) []card {

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
			52*packs,
		)

	for p := 0; p < packs; p++ {

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
	}

	return deck

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

	// Sabot à deux paquets de cartes (104 cartes).
	deck := newShoe(2)

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

	playerNatural := score(player) == 21
	dealerNatural := score(dealer) == 21

	// --------------------------------------------------------
	// ASSURANCE (si le croupier montre un As)
	// --------------------------------------------------------

	insurance := 0

	if dealer[0].rank == 1 {
		insurance = askInsurance(c, bet)
	}

	if insurance > 0 {

		classes.ClearScreen()

		fmt.Println(
			classes.TitleBox(
				"🛡️ RÉSULTAT DE L'ASSURANCE",
			),
		)

		fmt.Println()

		if dealerNatural {

			gain := insurance * 2

			c.Gold += gain

			fmt.Printf(
				"%sLe croupier a Blackjack — l'assurance paie 2:1.%s\n",
				classes.BrightGreen,
				classes.Reset,
			)

			fmt.Printf(
				"%s+%d 💰%s\n",
				classes.BrightGreen,
				gain,
				classes.Reset,
			)

		} else {

			c.Gold -= insurance

			fmt.Printf(
				"%sLe croupier n'a pas Blackjack — assurance perdue.%s\n",
				classes.BrightRed,
				classes.Reset,
			)

			fmt.Printf(
				"%s-%d 💰%s\n",
				classes.BrightRed,
				insurance,
				classes.Reset,
			)
		}

		classes.Pause()
	}

	// --------------------------------------------------------
	// BLACKJACK NATUREL (deux premières cartes = 21)
	// --------------------------------------------------------

	if playerNatural || dealerNatural {

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
			score(player),
		)

		fmt.Printf(
			"%sCroupier%s : %s  → %d\n\n",
			classes.BrightYellow,
			classes.Reset,
			cardNames(dealer),
			score(dealer),
		)

		switch {

		case playerNatural && dealerNatural:

			fmt.Println(
				classes.BrightYellow +
					"↔ Double Blackjack — égalité, votre mise est remboursée." +
					classes.Reset,
			)

		case playerNatural:

			gain := bet * 3 / 2

			c.Gold += gain

			fmt.Println(
				classes.Bold +
					classes.BrightGreen +
					"⚡ BLACKJACK ! ⚡" +
					classes.Reset,
			)

			fmt.Printf(
				"%s+%d 💰%s (paiement 3:2)\n",
				classes.BrightGreen,
				gain,
				classes.Reset,
			)

		default:

			c.Gold -= bet

			fmt.Printf(
				"%s✖ Le croupier a Blackjack.%s\n",
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

		return
	}

	// --------------------------------------------------------
	// TOUR DU JOUEUR (tirer / rester / doubler / split)
	// --------------------------------------------------------

	// Capacité fixée à maxBlackjackHands : jusqu'à 3 splits sont
	// autorisés, donc la tranche "hands" ne dépassera jamais ce
	// nombre d'éléments. Cela évite qu'un append() ne la réalloue
	// et n'invalide le pointeur "h" utilisé plus bas.
	hands := make([]bjHand, 1, maxBlackjackHands)
	hands[0] = bjHand{cards: player, bet: bet}

	i := 0

	for i < len(hands) {

		h := &hands[i]

	handLoop:
		for {

			classes.ClearScreen()

			fmt.Println(
				classes.TitleBox(
					"🃏 BLACKJACK",
				),
			)

			fmt.Println()

			if len(hands) > 1 {

				fmt.Printf(
					"%sMain %d/%d%s\n\n",
					classes.Bold,
					i+1,
					len(hands),
					classes.Reset,
				)
			}

			fmt.Printf(
				"%sVotre main%s\n",
				classes.Bold,
				classes.Reset,
			)

			fmt.Printf(
				"  %s\n",
				cardNames(h.cards),
			)

			fmt.Printf(
				"  Score : %d\n\n",
				score(h.cards),
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
				"\n%sMise sur cette main : %d 💰%s\n",
				classes.BrightYellow,
				h.bet,
				classes.Reset,
			)

			if score(h.cards) > 21 {
				h.busted = true
				break handLoop
			}

			// Une main issue d'un split d'As ne reçoit qu'une
			// seule carte supplémentaire et ne peut plus agir.
			if h.isSplitAce {
				break handLoop
			}

			canDouble :=
				len(h.cards) == 2 &&
					!h.doubled &&
					c.Gold >= totalWagered(hands)+h.bet

			canSplit :=
				len(hands) < maxBlackjackHands &&
					len(h.cards) == 2 &&
					h.cards[0].value() == h.cards[1].value() &&
					c.Gold >= totalWagered(hands)+h.bet

			canSurrender :=
				len(hands) == 1 &&
					len(h.cards) == 2 &&
					!h.doubled

			fmt.Println()

			options :=
				[]string{
					"1  Tirer une carte",
					"2  Rester",
				}

			if canDouble {
				options =
					append(
						options,
						"3  Doubler la mise",
					)
			}

			if canSplit {
				options =
					append(
						options,
						"4  Split",
					)
			}

			if canSurrender {
				options =
					append(
						options,
						"5  Abandonner (récupère la moitié de la mise)",
					)
			}

			fmt.Println(
				classes.Panel(
					"VOTRE CHOIX",
					options...,
				),
			)

			fmt.Print(
				"\nAction > ",
			)

			choice :=
				classes.ReadInt()

			switch choice {

			case 1:

				h.cards =
					append(
						h.cards,
						draw(&deck),
					)

			case 2:

				break handLoop

			case 3:

				if !canDouble {

					fmt.Println(
						classes.BrightRed +
							"✖ Choix invalide." +
							classes.Reset,
					)

					classes.Pause()

					continue handLoop
				}

				h.bet *= 2
				h.doubled = true

				h.cards =
					append(
						h.cards,
						draw(&deck),
					)

				if score(h.cards) > 21 {
					h.busted = true
				}

				break handLoop

			case 4:

				if !canSplit {

					fmt.Println(
						classes.BrightRed +
							"✖ Choix invalide." +
							classes.Reset,
					)

					classes.Pause()

					continue handLoop
				}

				isAceSplit :=
					h.cards[0].rank == 1

				second :=
					bjHand{
						cards:      []card{h.cards[1]},
						bet:        h.bet,
						isSplitAce: isAceSplit,
					}

				h.cards =
					[]card{h.cards[0]}

				h.isSplitAce = isAceSplit

				h.cards =
					append(
						h.cards,
						draw(&deck),
					)

				second.cards =
					append(
						second.cards,
						draw(&deck),
					)

				hands =
					append(
						hands,
						second,
					)

			case 5:

				if !canSurrender {

					fmt.Println(
						classes.BrightRed +
							"✖ Choix invalide." +
							classes.Reset,
					)

					classes.Pause()

					continue handLoop
				}

				h.surrendered = true

				break handLoop

			default:

				fmt.Println(
					classes.BrightRed +
						"✖ Choix invalide." +
						classes.Reset,
				)

				classes.Pause()
			}
		}

		i++
	}

	// --------------------------------------------------------
	// TOUR DU CROUPIER
	// --------------------------------------------------------

	anyoneAlive := false

	for _, h := range hands {
		if !h.busted && !h.surrendered {
			anyoneAlive = true
		}
	}

	if anyoneAlive {

		for score(dealer) < 17 {

			dealer =
				append(
					dealer,
					draw(&deck),
				)
		}
	}

	dealerScore :=
		score(dealer)

	// --------------------------------------------------------
	// RÉSULTATS
	// --------------------------------------------------------

	classes.ClearScreen()

	fmt.Println(
		classes.TitleBox(
			"🃏 RÉSULTAT DU BLACKJACK",
		),
	)

	fmt.Println()

	fmt.Printf(
		"%sCroupier%s : %s  → %d\n\n",
		classes.BrightYellow,
		classes.Reset,
		cardNames(dealer),
		dealerScore,
	)

	for idx, h := range hands {

		handScore := score(h.cards)

		if len(hands) > 1 {

			fmt.Printf(
				"%sMain %d%s : %s  → %d\n",
				classes.Bold,
				idx+1,
				classes.Reset,
				cardNames(h.cards),
				handScore,
			)

		} else {

			fmt.Printf(
				"%sVous%s : %s  → %d\n",
				classes.BrightCyan,
				classes.Reset,
				cardNames(h.cards),
				handScore,
			)
		}

		switch {

		case h.surrendered:

			loss := h.bet / 2

			c.Gold -= loss

			fmt.Printf(
				"  %s⚑ Abandonnée — -%d 💰%s\n",
				classes.BrightYellow,
				loss,
				classes.Reset,
			)

		case h.busted:

			c.Gold -= h.bet

			fmt.Printf(
				"  %s✖ Dépassé 21 — -%d 💰%s\n",
				classes.BrightRed,
				h.bet,
				classes.Reset,
			)

		case dealerScore > 21:

			c.Gold += h.bet

			fmt.Printf(
				"  %s✔ Le croupier dépasse 21 — +%d 💰%s\n",
				classes.BrightGreen,
				h.bet,
				classes.Reset,
			)

		case handScore > dealerScore:

			c.Gold += h.bet

			fmt.Printf(
				"  %s✔ Vous gagnez — +%d 💰%s\n",
				classes.BrightGreen,
				h.bet,
				classes.Reset,
			)

		case handScore == dealerScore:

			fmt.Printf(
				"  %s↔ Égalité — mise remboursée%s\n",
				classes.BrightYellow,
				classes.Reset,
			)

		default:

			c.Gold -= h.bet

			fmt.Printf(
				"  %s✖ Le croupier gagne — -%d 💰%s\n",
				classes.BrightRed,
				h.bet,
				classes.Reset,
			)
		}

		fmt.Println()
	}

	fmt.Printf(
		"Portefeuille : %s%d 💰%s\n",
		classes.BrightYellow,
		c.Gold,
		classes.Reset,
	)

	classes.Pause()

}

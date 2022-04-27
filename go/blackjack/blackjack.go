package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "king":
		return 10
	case "queen":
		return 10
	case "jack":
		return 10
	case "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case (ParseCard(card1)+ParseCard(card2) == 21):
		switch {
		case ParseCard(dealerCard) < 10:
			return "W"
		default:
			return "S"
		}
	case (ParseCard(card1)+ParseCard(card2)) >= 17 && (ParseCard(card1)+ParseCard(card2)) <= 20:
		return "S"
	case (ParseCard(card1)+ParseCard(card2)) >= 12 && (ParseCard(card1)+ParseCard(card2)) <= 16:
		switch {
		case ParseCard(dealerCard) >= 7:
			return "H"
		default:
			return "S"
		}
	default:
		return "H"
	}
}

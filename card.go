package pontifex

import (
	"fmt"
	"strconv"
)

/*
  Clubs     01 - 13
  Diamonds  14 - 26
  Hearts    27 - 39 ( 1 - 13)
  Spades    40 - 52 (14 - 26)
  Joker A   53
  Joker B   54
*/
type Card uint8

const JokerA Card = 53
const JokerB Card = 54

// Returns the mod 26 integer value of the card.
func (c Card) Value() int {
	switch {
	case c >= 27 && c <= 52:
		return int(c) - 26
	case c == JokerA:
		fallthrough
	case c == JokerB:
		return 53
	}
	return int(c)
}

// Output a card as a Unicode string.
// I chose not to use the actual unicode cards (e.g. 🂱)
// because they are hard to see at any reasonable font size.
func (c Card) String() string {
	switch {
	case c == JokerA:
		return " Ja"
	case c == JokerB:
		return " Jb"
	case c > 0 && c <= 13:
		return fmt.Sprintf("%s♣", faceForCard(c))
	case c <= 26:
		return fmt.Sprintf("%s♦", faceForCard(c))
	case c <= 39:
		return fmt.Sprintf("%s♥", faceForCard(c))
	case c <= 52:
		return fmt.Sprintf("%s♠", faceForCard(c))
	}
	return "???"
}

func faceForCard(c Card) string {
	switch c % 13 {
	case 1:
		return " A"
	case 10:
		return "10"
	case 11:
		return " J"
	case 12:
		return " Q"
	case 0:
		return " K"
	}
	return " " + strconv.Itoa(int(c%13))
}

package pontifex

import (
	"fmt"
	"strings"
)

type Deck [54]Card

// Create a new, sorted deck.
func NewDeck() *Deck {
	var d Deck
	for i := 0; i < 54; i++ {
		d[i] = Card(i + 1)
	}
	return &d
}

// Create a deck keyed by a passphrase
func NewDeckForPassphrase(phrase string) (*Deck, error) {
	d := NewDeck()
	for i, char := range phrase {
		c := int(char)
		if c < 65 || c > 90 {
			return nil, fmt.Errorf("Invalid passphrase character %q at offset %d. Must be from A-Z ascii (0x41-0x5A).", c, i)
		}
		d.advance()
		d.countCut(c - 64)
	}
	return d, nil
}

func (d *Deck) String() string {
	s := []string{""}
	for _, c := range d {
		s = append(s, c.String())
	}
	return strings.Join(s, " ")
}

// Find the card in the deck
func (d *Deck) findCard(c Card) int {
	for i, dc := range d {
		if c == dc {
			return i
		}
	}
	panic(fmt.Sprintf("Can't find card in deck: %d", c))
}

// Swap a card with the one below it, wrapping around if needed.
func (d *Deck) swapCardDown(index int) {
	if index == 53 {
		shift := make([]Card, 52)
		copy(shift, d[1:53])
		d[1] = d[index]
		for i, c := range shift {
			d[i+2] = c
		}
		return
	}
	below := d[index+1]
	d[index+1] = d[index]
	d[index] = below
}

// Move card two steps down the deck, wrapping around if needed
func (d *Deck) swapCardTwoDown(index int) {
	if index == 53 {
		shift := make([]Card, 51)
		copy(shift, d[2:53])
		d[2] = d[index]
		for i, c := range shift {
			d[i+3] = c
		}
		return
	}
	if index == 52 {
		shift := make([]Card, 51)
		copy(shift, d[1:52])
		d[1] = d[index]
		for i, c := range shift {
			d[i+2] = c
		}
		return
	}
	below := d[index+2]
	d[index+2] = d[index]
	d[index] = d[index+1]
	d[index+1] = below
}

// Find the jokers, swap the cards above the first
// with the cards below the second.
func (d *Deck) tripleCut() {
	a := d.findCard(JokerA)
	b := d.findCard(JokerB)
	if a > b {
		a, b = b, a
	}

	top := make([]Card, a)
	copy(top, d[0:a])

	middle := make([]Card, b-a+1)
	copy(middle, d[a:b+1])

	bottom := make([]Card, 53-b)
	copy(bottom, d[b+1:])

	for i, _ := range d {
		if i < len(bottom) {
			d[i] = bottom[i]
		} else if i-len(bottom) < len(middle) {
			d[i] = middle[i-len(bottom)]
		} else {
			d[i] = top[i-(len(bottom)+len(middle))]
		}
	}
}

// Cut the deck based on a number of cards down.
func (d *Deck) countCut(count int) {
	// Adjust for jokers
	if count > 53 {
		count = 53
	}

	top := make([]Card, count)
	copy(top, d[0:count])

	// Move the bottom up to the top
	for i := 0; i < 54-count; i++ {
		d[i] = d[i+count]
	}

	// Move the top down
	for i := 0; i < count; i++ {
		d[53-count+i] = top[i]
	}
}

// Mutate the deck to it's next state
func (d *Deck) advance() {
	i := d.findCard(JokerA)
	d.swapCardDown(i)

	i = d.findCard(JokerB)
	d.swapCardTwoDown(i)

	d.tripleCut()
	d.countCut(int(d[53]))
}

func (d *Deck) NextKey() Card {
	d.advance()
	i := int(d[0])
	if i > 53 {
		// Both Jokers are _valued_ at 53, but JokerB is 54 in this system for differentiation.
		i = 53
	}
	output := d[i]
	if output == JokerA || output == JokerB {
		return d.NextKey()
	}
	return output
}

func (d Deck) Encrypt(s string) (string, error) {
	out := []string{}
	for _, char := range s {
		// Don't mess with whitespace
		if char == ' ' {
			out = append(out, " ")
		} else {
			c := int(char)
			if c < 65 || c > 90 {
				return "", fmt.Errorf("Invalid input character %q. Must be from A-Z ascii (0x41-0x5A).", c)
			}
			key := d.NextKey()
			c = ((c-65)+int(key))%26 + 65
			out = append(out, string(c))
		}
	}
	return strings.Join(out, ""), nil
}

func (d Deck) Decrypt(s string) (string, error) {
	out := []string{}
	for _, char := range s {
		// Don't mess with whitespace
		if char == ' ' {
			out = append(out, " ")
		} else {
			c := int(char)
			if c < 65 || c > 90 {
				return "", fmt.Errorf("Invalid input character %q. Must be from A-Z ascii (0x41-0x5A).", c)
			}
			key := d.NextKey()

			// Shift down from ASCII range, and 1-indexed
			c = c%65 + 1
			if c < key.Value() {
				c = c + 26
			}

			// Subtract the key
			c = c - key.Value()
			if c == 0 {
				c = 26
			}
			// Shift back up to ASCII, and 0 indexed
			c = c + 64

			out = append(out, string(c))
		}
	}
	return strings.Join(out, ""), nil
}

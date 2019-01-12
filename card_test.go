package pontifex

import (
	"testing"
)

func TestCardString(t *testing.T) {
	tests := []struct {
		card   Card
		expect string
	}{
		{
			Card(1),
			" A♣",
		},
		{
			Card(2),
			" 2♣",
		},
		{
			Card(11),
			" J♣",
		},
		{
			Card(12),
			" Q♣",
		},
		{
			Card(13),
			" K♣",
		},
		{
			Card(14),
			" A♦",
		},
		{
			JokerA,
			" Ja",
		},
		{
			JokerB,
			" Jb",
		},
	}

	for _, test := range tests {
		if test.card.String() != test.expect {
			t.Errorf("Error converting card to string. Expected %s, got %s for %d", test.expect, test.card, test.card)
		}
	}
}

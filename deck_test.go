package pontifex

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestFindCard(t *testing.T) {
	deck := NewDeck()
	t.Log(deck)
	i := deck.findCard(JokerA)
	if i != 52 {
		t.Errorf("Did not find card. Expected %d, got %d", 52, i)
	}
}

type order struct {
	Value Card
	Index int
}

// Helper to check that a deck has cards in certain positions.
func assertDeckOrder(t *testing.T, d *Deck, assertions ...order) {
	for _, assertion := range assertions {
		if d[assertion.Index] != assertion.Value {
			t.Errorf("Bad deck order. Expected %s at index %d, got %s", assertion.Value, assertion.Index, d[assertion.Index])
			return
		}
	}
}

func TestSwapCardDown(t *testing.T) {
	deck := NewDeck()
	deck.swapCardDown(0)
	assertDeckOrder(t, deck, order{2, 0}, order{1, 1})

	deck = NewDeck()
	deck.swapCardDown(53)
	assertDeckOrder(t, deck, order{JokerA, 53}, order{1, 0}, order{JokerB, 1}, order{2, 2})
}

func TestSwapCardTwoDown(t *testing.T) {
	t.Run("Index 0", func(t *testing.T) {
		deck := NewDeck()
		deck.swapCardTwoDown(0)
		assertDeckOrder(t, deck, order{2, 0}, order{3, 1}, order{1, 2})
	})

	t.Run("Index 53", func(t *testing.T) {
		deck := NewDeck()
		deck.swapCardTwoDown(53)
		assertDeckOrder(t, deck, order{JokerA, 53}, order{1, 0}, order{2, 1}, order{JokerB, 2})
	})

	t.Run("Index 52", func(t *testing.T) {
		deck := NewDeck()
		deck.swapCardTwoDown(52)
		assertDeckOrder(t, deck, order{JokerB, 53}, order{1, 0}, order{JokerA, 1}, order{2, 2})
	})
}

func TestTripleCut(t *testing.T) {
	/*
	  0   1   2   3   4   5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  25  26  27  28  29  30  31  32  33  34  35  36  37  38  39  40  41  42  43  44  45  46  47  48  49  50  51  52  53
	 A♣  2♣  3♣  4♣  5♣  6♣  7♣  8♣  9♣ 10♣  Ja  Q♣  K♣  A♦  2♦  3♦  4♦  5♦  6♦  7♦  8♦  9♦ 10♦  J♦  Q♦  Jb  A♥  2♥  3♥  4♥  5♥  6♥  7♥  8♥  9♥ 10♥  J♥  Q♥  K♥  A♠  2♠  3♠  4♠  5♠  6♠  7♠  8♠  9♠ 10♠  J♠  Q♠  K♠  J♣  K♦
	 A♥  2♥  3♥  4♥  5♥  6♥  7♥  8♥  9♥ 10♥  J♥  Q♥  K♥  A♠  2♠  3♠  4♠  5♠  6♠  7♠  8♠  9♠ 10♠  J♠  Q♠  K♠  J♣  K♦  Ja  Q♣  K♣  A♦  2♦  3♦  4♦  5♦  6♦  7♦  8♦  9♦ 10♦  J♦  Q♦  Jb  A♣  2♣  3♣  4♣  5♣  6♣  7♣  8♣  9♣ 10♣
	*/
	t.Run("10-25", func(t *testing.T) {
		deck := NewDeck()
		deck[52] = deck[10]
		deck[53] = deck[25]
		deck[10] = JokerA
		deck[25] = JokerB
		deck.tripleCut()
		assertDeckOrder(t, deck, order{27, 0}, order{JokerA, 28}, order{JokerB, 43}, order{10, 53})

		// Joker order does not matter for shuffle
		deck = NewDeck()
		deck[52] = deck[10]
		deck[53] = deck[25]
		deck[10] = JokerB
		deck[25] = JokerA
		deck.tripleCut()
		assertDeckOrder(t, deck, order{27, 0}, order{JokerB, 28}, order{JokerA, 43}, order{10, 53})
	})

	/*
		Deck bottom bound
		  0   1   2   3   4   5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  25  26  27  28  29  30  31  32  33  34  35  36  37  38  39  40  41  42  43  44  45  46  47  48  49  50  51  52  53
		 A♣  2♣  3♣  4♣  5♣  6♣  7♣  8♣  9♣ 10♣  Ja  Q♣  K♣  A♦  2♦  3♦  4♦  5♦  6♦  7♦  8♦  9♦ 10♦  J♦  Q♦  K♦  A♥  2♥  3♥  4♥  5♥  6♥  7♥  8♥  9♥ 10♥  J♥  Q♥  K♥  A♠  2♠  3♠  4♠  5♠  6♠  7♠  8♠  9♠ 10♠  J♠  Q♠  K♠  J♣  Jb
		 Ja  Q♣  K♣  A♦  2♦  3♦  4♦  5♦  6♦  7♦  8♦  9♦ 10♦  J♦  Q♦  K♦  A♥  2♥  3♥  4♥  5♥  6♥  7♥  8♥  9♥ 10♥  J♥  Q♥  K♥  A♠  2♠  3♠  4♠  5♠  6♠  7♠  8♠  9♠ 10♠  J♠  Q♠  K♠  J♣  Jb  A♣  2♣  3♣  4♣  5♣  6♣  7♣  8♣  9♣ 10♣
	*/
	t.Run("10-53", func(t *testing.T) {
		deck := NewDeck()
		deck[52] = deck[10]
		deck[10] = JokerA
		deck.tripleCut()
		assertDeckOrder(t, deck, order{JokerA, 0}, order{JokerB, 43}, order{10, 53})
	})

	/*
		Deck top bound
		  0   1   2   3   4   5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  25  26  27  28  29  30  31  32  33  34  35  36  37  38  39  40  41  42  43  44  45  46  47  48  49  50  51  52  53
		 Ja  2♣  3♣  4♣  5♣  6♣  7♣  8♣  9♣ 10♣  J♣  Q♣  K♣  A♦  2♦  3♦  4♦  5♦  6♦  7♦  8♦  9♦ 10♦  J♦  Q♦  Jb  A♥  2♥  3♥  4♥  5♥  6♥  7♥  8♥  9♥ 10♥  J♥  Q♥  K♥  A♠  2♠  3♠  4♠  5♠  6♠  7♠  8♠  9♠ 10♠  J♠  Q♠  K♠  A♣  K♦
		 A♥  2♥  3♥  4♥  5♥  6♥  7♥  8♥  9♥ 10♥  J♥  Q♥  K♥  A♠  2♠  3♠  4♠  5♠  6♠  7♠  8♠  9♠ 10♠  J♠  Q♠  K♠  A♣  K♦  Ja  2♣  3♣  4♣  5♣  6♣  7♣  8♣  9♣ 10♣  J♣  Q♣  K♣  A♦  2♦  3♦  4♦  5♦  6♦  7♦  8♦  9♦ 10♦  J♦  Q♦  Jb
	*/
	t.Run("0-25", func(t *testing.T) {
		deck := NewDeck()
		deck[52] = deck[0]
		deck[53] = deck[25]
		deck[0] = JokerA
		deck[25] = JokerB
		deck.tripleCut()
		assertDeckOrder(t, deck, order{JokerA, 28}, order{JokerB, 53}, order{27, 0})
	})
}

func TestCountCut(t *testing.T) {
	/*
	  0   1   2   3   4   5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  25  26  27  28  29  30  31  32  33  34  35  36  37  38  39  40  41  42  43  44  45  46  47  48  49  50  51  52  53
	 A♣  2♣  3♣  4♣  5♣  6♣  7♣  8♣  9♣  Jb  J♣  Q♣  K♣  A♦  2♦  3♦  4♦  5♦  6♦  7♦  8♦  9♦ 10♦  J♦  Q♦  K♦  A♥  2♥  3♥  4♥  5♥  6♥  7♥  8♥  9♥ 10♥  J♥  Q♥  K♥  A♠  2♠  3♠  4♠  5♠  6♠  7♠  8♠  9♠ 10♠  J♠  Q♠  K♠  Ja 10♣
	 J♣  Q♣  K♣  A♦  2♦  3♦  4♦  5♦  6♦  7♦  8♦  9♦ 10♦  J♦  Q♦  K♦  A♥  2♥  3♥  4♥  5♥  6♥  7♥  8♥  9♥ 10♥  J♥  Q♥  K♥  A♠  2♠  3♠  4♠  5♠  6♠  7♠  8♠  9♠ 10♠  J♠  Q♠  K♠  Ja  A♣  2♣  3♣  4♣  5♣  6♣  7♣  8♣  9♣  Jb 10♣
	*/
	t.Run("10 On Bottom", func(t *testing.T) {
		deck := NewDeck()
		deck[53] = deck[9]
		deck[9] = JokerB
		deck.countCut(int(deck[53]))
		assertDeckOrder(t, deck, order{11, 0}, order{1, 43}, order{10, 53})
	})
}

func TestSampleKeyGeneration(t *testing.T) {
	// Samples from https://www.schneier.com/academic/solitaire/
	expected := []Card{4, 49, 10, 24, 8, 51, 44, 6, 4, 33}
	deck := NewDeck()

	for i, expect := range expected {
		key := deck.NextKey()
		if key != expect {
			t.Errorf("Generated incorrect key at step %d. Expected %s, got %s", i, expect, key)
		}
	}
}

func TestSampleEncryption(t *testing.T) {
	// Samples from https://www.schneier.com/academic/solitaire/
	t.Run("Ordered Deck", func(t *testing.T) {
		plaintext := "AAAAA  AAAAA"
		ciphertext := "EXKYI  ZSGEH"

		deck := NewDeck()
		out, err := deck.Encrypt(plaintext)
		if err != nil {
			t.Fatalf("Error in encrypt: %s", err)
		}
		if out != ciphertext {
			t.Errorf("Ciphertext did not match expected: %q != %q", out, ciphertext)
		}
	})

	t.Run("Passphrase Deck FOO", func(t *testing.T) {
		plaintext := "AAAAA  AAAAA  AAAAA"
		ciphertext := "ITHZU  JIWGR  FARMW"

		deck, err := NewDeckForPassphrase("FOO")
		if err != nil {
			t.Fatalf("Error in deck generation: %s", err)
		}
		out, err := deck.Encrypt(plaintext)
		if err != nil {
			t.Fatalf("Error in encrypt: %s", err)
		}
		if out != ciphertext {
			t.Errorf("Ciphertext did not match expected: %q != %q", out, ciphertext)
		}
	})

	t.Run("Passphrase Deck CRYPTONOMICON", func(t *testing.T) {
		plaintext := "SOLIT  AIREX"
		ciphertext := "KIRAK  SFJAN"

		deck, err := NewDeckForPassphrase("CRYPTONOMICON")
		if err != nil {
			t.Fatalf("Error in deck generation: %s", err)
		}
		out, err := deck.Encrypt(plaintext)
		if err != nil {
			t.Fatalf("Error in encrypt: %s", err)
		}
		if out != ciphertext {
			t.Errorf("Ciphertext did not match expected: %q != %q", out, ciphertext)
		}
	})
}

func TestSampleDecryption(t *testing.T) {
	// Samples from https://www.schneier.com/academic/solitaire/
	t.Run("Ordered Deck", func(t *testing.T) {
		plaintext := "AAAAA  AAAAA"
		ciphertext := "EXKYI  ZSGEH"

		deck := NewDeck()
		out, err := deck.Decrypt(ciphertext)
		if err != nil {
			t.Fatalf("Error in decrypt: %s", err)
		}
		if out != plaintext {
			t.Errorf("Plaintext did not match expected: %q != %q", out, plaintext)
		}
	})

	t.Run("Passphrase Deck FOO", func(t *testing.T) {
		plaintext := "AAAAA  AAAAA  AAAAA"
		ciphertext := "ITHZU  JIWGR  FARMW"

		deck, err := NewDeckForPassphrase("FOO")
		if err != nil {
			t.Fatalf("Error in deck generation: %s", err)
		}
		out, err := deck.Decrypt(ciphertext)
		if err != nil {
			t.Fatalf("Error in decrypt: %s", err)
		}
		if out != plaintext {
			t.Errorf("Plaintext did not match expected: %q != %q", out, plaintext)
		}
	})

	t.Run("Passphrase Deck CRYPTONOMICON", func(t *testing.T) {
		plaintext := "SOLIT  AIREX"
		ciphertext := "KIRAK  SFJAN"

		deck, err := NewDeckForPassphrase("CRYPTONOMICON")
		if err != nil {
			t.Fatalf("Error in deck generation: %s", err)
		}
		out, err := deck.Decrypt(ciphertext)
		if err != nil {
			t.Fatalf("Error in decrypt: %s", err)
		}
		if out != plaintext {
			t.Errorf("Plaintext did not match expected: %q != %q", out, plaintext)
		}
	})
}

func TestFuzzy(t *testing.T) {
	// Fuzzily test things, different every run unless you set the rand seed to debug errors
	var seed int64 = time.Now().UnixNano()
	// Uncomment below and set to a fixed int64 for reproducing tests.
	// seed = 1547273616491450000
	t.Logf("Random Test Seed: %v", seed)
	r := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		passphrase := randomLengthString(r)
		plaintext := randomLengthString(r)

		deck, err := NewDeckForPassphrase(passphrase)
		if err != nil {
			t.Fatal(err)
		}
		ciphertext, err := deck.Encrypt(plaintext)
		if err != nil {
			t.Fatal(err)
		}

		deck, _ = NewDeckForPassphrase(passphrase)
		result, err := deck.Decrypt(ciphertext)
		if err != nil {
			t.Fatal(err)
		}

		if result != plaintext {
			t.Fatalf("Iteration %d result does not match input.\nEncrypted: %q\nPlaintext: %q\nDecrypted: %q", i, ciphertext, plaintext, result)
		}
	}
}

// Generate a random length ASCII string A-Z only.
// Minimum of 5, maximum of 105 characters.
func randomLengthString(r *rand.Rand) string {
	length := r.Intn(100) + 5
	s := []string{}
	for i := 0; i < length; i++ {
		s = append(s, string(r.Intn(26)+65))
	}
	return strings.Join(s, "")
}

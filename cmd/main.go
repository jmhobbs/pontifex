package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jmhobbs/pontifex"
)

func main() {
	log.SetFlags(0)

	if len(os.Args) != 4 {
		usage(nil)
	}

	mode := os.Args[1]
	passphrase := strings.ToUpper(os.Args[2])
	input := strings.ToUpper(os.Args[3])

	deck, err := pontifex.NewDeckForPassphrase(passphrase)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	if mode == "encrypt" {
		out, err := deck.Encrypt(input)
		if err != nil {
			log.Println(err)
			os.Exit(2)
		}
		fmt.Println(out)
	} else if mode == "decrypt" {
		out, err := deck.Decrypt(os.Args[3])
		if err != nil {
			log.Println(err)
			os.Exit(2)
		}
		fmt.Println(out)
	} else {
		usage(fmt.Errorf("invalid mode %q; must be \"encrypt\" or \"decrypt\"", mode))
	}
}

func usage(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n\n", err)
	}
	fmt.Fprintln(os.Stderr, "usage: pontifex [encrypt|decrypt] <passphrase> <ciphertext|plaintext>")
	os.Exit(1)
}

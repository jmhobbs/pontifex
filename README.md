[![Build Status](https://travis-ci.org/jmhobbs/pontifex.svg?branch=master)](https://travis-ci.org/jmhobbs/pontifex) [![GoDoc](https://godoc.org/github.com/jmhobbs/pontifex?status.svg)](https://godoc.org/github.com/jmhobbs/pontifex)

# Solitaire / Pontifex

This is itended as a "clean-room" implementation of the Solitaire algorithm created by Bruce Schneier and used in the novel Cryptonomicon.

The only source material I will be using is the [description of the algorithm from Bruce's website](https://www.schneier.com/academic/solitaire/).  I _have_ read the Perl code in the book, but I have not taken the time to grok it in any meaningful sense.

# Usage

The basic usage is to create a deck, likely keyed by a passphrase. Then, encrypt or decrypt strings on that deck.

## API

    package main

    import (
      "fmt"

      "github.com/jmhobbs/pontifex"
    )

    func main() {
      deck, _ := pontifex.NewDeckForPassphrase("CRYPTONOMICON")
      ciphertext, _ := deck.Encrypt("SOLITAIRE")
      fmt.Println(ciphertext)

      plaintext, _ := deck.Decrypr(ciphertext)
      fmt.Println(plaintext)
    }

## CLI

    $ pontifex
    usage: pontifex [encrypt|decrypt] <passphrase> <ciphertext|plaintext>
    $ pontifex encrypt CRYPTONOMICON SOLITAIRE
    KIRAKSFJA
    $ pontifex decrypt CRYPTONOMICON KIRAKSFJA
    SOLITAIRE


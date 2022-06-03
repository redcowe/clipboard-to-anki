package main

import (
	"fmt"
	"log"

	"github.com/atotto/clipboard"
	"github.com/mattn/go-tty"
)

func main() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		if r == 24 {
			clipboard.WriteAll("日本語")
			text, err := clipboard.ReadAll()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(text)
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

func wordcount(f io.Reader) (int, error) {
	var numb_of_words int = 0
	var still_in_word bool = false
	r := bufio.NewReader(f)

	for {
		c, _, err := r.ReadRune()

		if unicode.IsSpace(c) {
			if still_in_word {
				numb_of_words++
			}
			still_in_word = false
		} else {
			still_in_word = true
		}

		if err == io.EOF {
			return numb_of_words, nil
		}
		if err != nil {
			return 0, err
		}
	}
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	words, err := wordcount(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The number of words: %v\n", words)
}

package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Ltime)

	var r = bufio.NewReader(os.Stdin)

	for {
		s, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("error in stdin:", err)
		}

		log.Print(s)
	}
}

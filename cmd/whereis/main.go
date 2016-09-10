package main

import (
	"github.com/zaquestion/whereis"
	"log"
)

func main() {
	if err := whereis.Run(); err != nil {
		log.Fatal(err)
	}
}

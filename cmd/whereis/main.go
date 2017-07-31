package main

import (
	"github.com/zaquestion/whereis.global"
	"log"
)

func main() {
	if err := whereis.Run(); err != nil {
		log.Fatal(err)
	}
}

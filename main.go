package main

import (
	"log"

	"github.com/julianbarrios/hexserver/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}

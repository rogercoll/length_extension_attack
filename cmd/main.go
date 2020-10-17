package main

import (
	"fmt"
	"github.com/rogercoll/merkledamgard"
	"log"
)

func main() {
	newHash, err := merkledamgard.LengthExtensionAttack(9, "d61f778d92e85f1cc9d43235a80ae95d", " world")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("New forged hash: %v\n", newHash)
}

package main

import (
	"fmt"
	"github.com/rogercoll/merkledamgard"
)

func main() {
	fmt.Println(merkledamgard.LengthExtensionAttack(9, "d61f778d92e85f1cc9d43235a80ae95d", " world"))
}

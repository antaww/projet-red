package main

import (
	"bufio"
	"os"
)

const (
	potionSoin = "Potion de soin"
)

var P1 Personnage
var BR *bufio.Reader

func main() {
	BR = bufio.NewReader(os.Stdin)
	charCreation()
	//InitCharacter()
	Menu()
}

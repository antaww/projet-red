package main

import (
	"bufio"
	"os"
)

var P1 Personnage
var BR *bufio.Reader

func main() {
	BR = bufio.NewReader(os.Stdin)
	InitCharacter()
	Menu()
}

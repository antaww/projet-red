package main

import (
	"bufio"
	"flag"
	"os"
)

const (
	potionSoin = "Potion de soin"

	Reset = "\033[0m"

	Red      = "\033[31m"
	Green    = "\033[32m"
	Yellow   = "\033[33m"
	Blue     = "\033[34m"
	Purple   = "\033[35m"
	Cyan     = "\033[36m"
	White    = "\033[37m"
	BIGreen  = "\033[1;92m"
	BIWhite  = "\033[1;97m"
	BIPurple = "\033[1;95m"
)

var P1 Personnage
var M1 Monster
var BR *bufio.Reader
var Debug bool

func main() {
	flag.BoolVar(&Debug, "debug", false, "boot game while skipping steps")
	flag.Parse()
	BR = bufio.NewReader(os.Stdin)
	if Debug {
		InitCharacter()
		InitGoblin()
		trainingFight()
	} else {
		charCreation()
		InitGoblin()
		Menu()
	}
}

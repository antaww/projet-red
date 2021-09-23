package main

var P1 Personnage

func main() {
	InitCharacter()
	P1.DisplayInfo()
	P1.accesInventory()
	P1.takePot()
	P1.DisplayInfo()
	P1.accesInventory()
}

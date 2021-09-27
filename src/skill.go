package main

import "fmt"

func (p *Personnage) spellBook() {
	if p.skill["Boule de feu"] == 0 {
		p.skill["Boule de feu"] += 1
		p.inventaire["Livre de sort : boule de feu"] -= 1
	} else {
		fmt.Println("Une seule boule de feu à la fois !")
	}
}

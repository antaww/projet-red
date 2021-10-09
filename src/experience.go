package main

import "fmt"

func (p *Personnage) giveXP(value int) {
	p.experience += value
	fmt.Println(">> Vous avez obtenu", value, "xp ! <<")
	p.levelCheck()
}

func (p *Personnage) levelCheck() {
	if p.experience >= 100 {
		p.niveau += 1
		fmt.Println(">>", p.nom, "vient de passer au niveau", p.niveau, "! <<")
		p.experience -= 100
	}
}

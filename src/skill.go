package main

import "fmt"

func (p *Personnage) spellBook() {
	if p.skill["Boule de feu"] == 0 {
		p.skill["Boule de feu"] += 1
		p.inventaire["Livre de sort : boule de feu"] -= 1
	} else {
		fmt.Println(">> Une seule boule de feu à la fois ! <<")
	}
}

func (p Personnage) useFireball() {
	ballDmg := 18
	if p.skill["Boule de feu"] >= 1 {
		p.skill["Boule de feu"] -= 1
		fmt.Println(p.nom, "attaque et fait", ballDmg, "dgt à", M1.name)
		M1.HP -= ballDmg
		if M1.HP > 0 {
			fmt.Println(M1.name, "a maintenant", M1.HP, "HP/", M1.maxHP, "HP")
		}
	} else {
		fmt.Println(">> Vous n'avez pas de boule de feu !<<")
		fmt.Println(">> Appuyez sur entrée pour continuer... ")
		Input()
		p.charAttack()
	}
}

func (p Personnage) usePunch() {
	punchDmg := 8
	if p.skill["Coup de poing"] >= 1 {
		fmt.Println(p.nom, "attaque et fait", punchDmg, "dgt à", M1.name)
		M1.HP -= punchDmg
		if M1.HP > 0 {
			fmt.Println(M1.name, "a maintenant", M1.HP, "HP/", M1.maxHP, "HP")
		}
	} else {
		fmt.Println(">> Vous n'avez pas de bras. <<")
		p.charAttack()
	}
}

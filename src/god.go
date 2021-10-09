package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func (p *Personnage) godTurn() {
	fmt.Println("1 >> Attaquer")
	fmt.Println("2 >> Inventaire")
	m, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	m = strings.Replace(m, "\r\n", "", -1)
	switch m {
	case "1":
		p.godAttack()
	case "2":
		ClearLog()
		p.accesInventory()
		p.useInventory(true)
	default:
		p.godTurn()
	}
}

func (p Personnage) godAttack() {
	ClearLog()
	fmt.Println(BIWhite + ">> Combat d'entraînement <<" + Reset)
	fmt.Println("1 >> Gear Four")
	fmt.Println("2 >> Boule de feu")
	fmt.Println("3 >> Retour")
	m, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	m = strings.Replace(m, "\r\n", "", -1)
	switch m {
	case "1":
		p.useGear()
	case "2":
		p.useFireball("Dieu")
	case "3":
		p.godTurn()
	default:
		p.godAttack()
	}
}

func (p Personnage) useGear() {
	rand.Seed(time.Now().UnixNano())
	gearDmg := 1000000
	if p.skill["Gear Four"] >= 1 {
		fmt.Println(p.nom, "utilise", BIYellow+"Gear Four"+Reset, "et inflige", gearDmg, "dégats à", M1.name)
		M1.HP -= (gearDmg - M1.maxHP)
		time.Sleep(400 * time.Millisecond)
		if M1.HP > 0 {
			fmt.Println(M1.name, "a maintenant", M1.HP, "HP/", M1.maxHP, "HP")
		}
	} else {
		fmt.Println(">> Impossible <<")
		p.godAttack()
	}
}

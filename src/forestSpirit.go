package main

import (
	"fmt"
	"strings"
	"time"
)

func (p *Personnage) spiritTurn() {
	fmt.Println("1 >> Attaquer")
	fmt.Println("2 >> Inventaire")
	m, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	m = strings.Replace(m, "\r\n", "", -1)
	switch m {
	case "1":
		p.spiritAttack()
	case "2":
		ClearLog()
		fmt.Println(BIWhite + ">> Appuyez sur entrée pour retourner au combat <<" + Reset)
		p.accesInventory()
		p.useInventory(true)
	default:
		p.charTurn()
	}
}

func (p Personnage) spiritAttack() {
	ClearLog()
	fmt.Println(">> Combat d'entraînement <<")
	fmt.Println("1 >> Tempête verte")
	fmt.Println("2 >> Boule de feu")
	fmt.Println("3 >> Retour")
	m, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	m = strings.Replace(m, "\r\n", "", -1)
	switch m {
	case "1":
		p.useStorm()
	case "2":
		p.useFireball()
	case "3":
		p.charTurn()
	}
}

func (p Personnage) useStorm() {
	stormDmg := 10
	if p.skill["Tempête verte"] >= 1 {
		fmt.Println(p.nom, "utilise", BIGreen+"Tempête verte"+Reset, "et inflige", stormDmg, "dégats à", M1.name)
		M1.HP -= stormDmg
		if M1.HP > 0 {
			fmt.Println(M1.name, "a maintenant", M1.HP, "HP/", M1.maxHP, "HP")
			time.Sleep(2 * time.Second)
			fmt.Println("La", BIGreen+"Tempête verte"+Reset, "surgit à nouveau et fait", stormDmg, "dégats à", M1.name)
			M1.HP -= stormDmg

			if M1.HP > 0 {
				fmt.Println(M1.name, "a maintenant", M1.HP, "HP/", M1.maxHP, "HP")
			}
		}

	} else {
		fmt.Println(">> Impossible <<")
		p.charAttack()
	}
}

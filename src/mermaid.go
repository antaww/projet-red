package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func (p *Personnage) mermaidTurn() {
	fmt.Println("1 >> Attaquer")
	fmt.Println("2 >> Inventaire")
	m, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	m = strings.Replace(m, "\r\n", "", -1)
	switch m {
	case "1":
		p.mermaidAttack()
	case "2":
		ClearLog()
		p.accesInventory()
		p.useInventory(true)
	default:
		p.mermaidTurn()
	}
}

func (p Personnage) mermaidAttack() {
	ClearLog()
	fmt.Println(BIWhite + ">> Combat d'entraînement <<" + Reset)
	fmt.Println("1 >> Danse pluie")
	fmt.Println("2 >> Boule de feu")
	fmt.Println("3 >> Retour")
	m, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	m = strings.Replace(m, "\r\n", "", -1)
	switch m {
	case "1":
		p.useDanse()
	case "2":
		p.useFireball("Sirène")
	case "3":
		p.mermaidTurn()
	default:
		p.mermaidAttack()
	}
}

func (p Personnage) useDanse() {
	rand.Seed(time.Now().UnixNano())
	danseDmg := rand.Intn(3)
	if p.skill["Danse pluie"] >= 1 {
		fmt.Println(p.nom, "utilise", BIBlue+"Danse pluie"+Reset, "et inflige", danseDmg, "dégats à", M1.name)
		M1.HP -= danseDmg
		time.Sleep(400 * time.Millisecond)
		if M1.HP > 0 {
			for i := 0; i < rand.Intn(11); i++ {
				if M1.HP > 0 {
					danseDmg = rand.Intn(3)
					if M1.HP-danseDmg >= 0 {
						fmt.Println("Les goûtes de la", BIBlue+"Danse pluie"+Reset, "s'abattent et infligent", danseDmg, "dégats à", M1.name)
						M1.HP -= danseDmg
					}
				}
			}
			fmt.Println(M1.name, "a maintenant", M1.HP, "HP/", M1.maxHP, "HP")
		}

	} else {
		fmt.Println(">> Impossible <<")
		p.mermaidAttack()
	}
}

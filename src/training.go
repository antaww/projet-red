package main

import (
	"fmt"
	"strings"
	"time"
)

func (m *Monster) goblinPattern(tour int) {
	attackCrit := m.attack * 2
	if tour%3 == 0 {
		fmt.Println("le gobelin attaque en critique et fait", attackCrit, "dgt à", P1.nom)
		P1.HP -= attackCrit
		fmt.Println(P1.nom, "a maintenant", P1.HP, "HP/", P1.maxHP, "HP")
	} else {
		fmt.Println("le gobelin attaque et fait", m.attack, "dgt à", P1.nom)
		P1.HP -= m.attack
		fmt.Println(P1.nom, "a maintenant", P1.HP, "HP/", P1.maxHP, "HP")
	}
}

func (p Personnage) charAttack() {
	fmt.Println(p.nom, "attaque et fait", p.attack, "dgt à", M1.name)
	M1.HP -= p.attack
	fmt.Println(M1.name, "a maintenant", M1.HP, "HP/", M1.maxHP, "HP")
}

func (p *Personnage) charTurn() {
	fmt.Println("1 >> Attaquer")
	fmt.Println("2 >> Inventaire")
	m, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	m = strings.Replace(m, "\r\n", "", -1)
	switch m {
	case "1":
		p.charAttack()
	case "2":
		p.accesInventory()
		p.useInventory(true)
	default:
		p.charTurn()
	}
}

func trainingFight() {
	InitGoblin()
	round := true
	ClearLog()
	fmt.Println(">> Combat d'entraînement <<")
	for i := 1; P1.HP > 0; i++ {
		if round == true {
			runeName := []rune(P1.nom)
			ClearLog()
			fmt.Println(">> Combat d'entraînement <<")
			fmt.Println(BIGreen+"Round ", i, Reset)
			if runeName[0] == 'A' || runeName[0] == 'E' || runeName[0] == 'I' || runeName[0] == 'O' || runeName[0] == 'U' {
				fmt.Print(">> C'est au tour d'", P1.nom, " ! <<\n")
			} else {
				fmt.Print(">> C'est au tour de ", P1.nom, " ! <<\n")
			}
			P1.charTurn()
			time.Sleep(1 * time.Second)
			round = false
		} else {
			fmt.Println(BIGreen+"Round ", i, Reset)
			fmt.Println(">> C'est au tour du", M1.name, " ! <<")
			M1.goblinPattern(i)
			fmt.Println(">> Appuyez sur entrée pour continuer... ")
			Input()
			round = true
		}
		if P1.HP == 0 {
			P1.dead()
			fmt.Println(">> Appuyez sur entrée pour continuer... ")
			Input()
			Menu()
		}
		if M1.HP == 0 {
			fmt.Println(">> ", M1.name, "est mort ! <<")
			fmt.Println(">> Appuyez sur entrée pour continuer... ")
			Input()
			Menu()
		}
	}
}

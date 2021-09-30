package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func (m *Monster) goblinPattern(tour int) {
	attackCrit := m.attack * 2
	rand.Seed(time.Now().UnixNano())
	if (rand.Intn(6)) == 5 {
		M1.HP += 10
		fmt.Println(M1.name, "mange une pomme en or et gagne 10HP, il a maintenant", M1.HP, "HP/", M1.maxHP, "HP")
	}
	if tour%3 == 0 {
		fmt.Println("Le gobelin attaque en ", Yellow+"critique"+Reset, "et fait", attackCrit, "dégats à", P1.nom)
		P1.HP -= attackCrit
		fmt.Println(P1.nom, "a maintenant", P1.HP, "HP/", P1.maxHP, "HP")
	} else {
		fmt.Println("Le gobelin attaque et fait", m.attack, "dégats à", P1.nom)
		P1.HP -= m.attack
		fmt.Println(P1.nom, "a maintenant", P1.HP, "HP/", P1.maxHP, "HP")
	}
}

func (p Personnage) charAttack() {
	ClearLog()
	fmt.Println(">> Combat d'entraînement <<")
	fmt.Println("1 >> Coup de poing")
	fmt.Println("2 >> Boule de feu")
	fmt.Println("3 >> Retour")
	m, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	m = strings.Replace(m, "\r\n", "", -1)
	switch m {
	case "1":
		p.usePunch()
	case "2":
		p.useFireball()
	case "3":
		p.charTurn()
	}
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
		ClearLog()
		fmt.Println(BIWhite + ">> Appuyez sur entrée pour retourner au combat <<" + Reset)
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
			if P1.classe == "Humain" {
				P1.charTurn()
			}
			if P1.classe == "Esprit de la forêt" {
				P1.spiritTurn()
			}
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
		if M1.HP <= 0 {
			min := 1
			max := 15
			fmt.Println(">> ", M1.name, "est mort ! <<")
			fmt.Println(">> Vous avez gagné", rand.Intn(max-min)+min, "! <<")
			fmt.Println(">> Appuyez sur entrée pour continuer... ")
			Input()
			Menu()
		}
	}
}

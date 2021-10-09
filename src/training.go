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
	if (rand.Intn(3)) == 2 {
		M1.maxHP += 10
		M1.HP += 10
		fmt.Println(M1.name, "mange une pomme en or et gagne 10HP, il a maintenant", M1.HP, "HP/", M1.maxHP, "HP")
	}
	if tour%3 == 0 {
		fmt.Println(M1.name, "attaque en", Yellow+"critique"+Reset, "et fait", attackCrit, "dégats à", P1.nom)
		P1.HP -= attackCrit
		fmt.Println(P1.nom, "a maintenant", P1.HP, "HP/", P1.maxHP, "HP")
	} else {
		fmt.Println(M1.name, "attaque et fait", m.attack, "dégats à", P1.nom)
		P1.HP -= m.attack
		fmt.Println(P1.nom, "a maintenant", P1.HP, "HP/", P1.maxHP, "HP")
	}
}

func (p Personnage) charAttack() {
	ClearLog()
	fmt.Println(BIWhite + ">> Combat d'entraînement <<" + Reset)
	fmt.Println("1 >> Coup de poing")
	fmt.Println("2 >> Boule de feu")
	fmt.Println("3 >> Retour")
	m, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	m = strings.Replace(m, "\r\n", "", -1)
	switch m {
	case "1":
		p.usePunch()
	case "2":
		p.useFireball("Humain")
	case "3":
		p.charTurn()
	default:
		p.charAttack()
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
	fmt.Println(BIWhite + ">> Combat d'entraînement <<" + Reset) // ne pas enlever sinon premier print non écrit
	for i := 1; P1.HP > 0; i++ {
		if round == true {
			runeName := []rune(P1.nom)
			ClearLog()
			fmt.Println(BIWhite + ">> Combat d'entraînement <<" + Reset)
			fmt.Println(BIGreen+"Round ", i, Reset)
			if runeName[0] == 'A' || runeName[0] == 'E' || runeName[0] == 'I' || runeName[0] == 'O' || runeName[0] == 'U' {
				fmt.Print(">> C'est au tour d'", P1.nom, " ! <<\n")
			} else {
				fmt.Print(">> C'est au tour de ", P1.nom, " ! <<\n")
			}
			if P1.classe == "Humain" || P1.classe == "Elfe" || P1.classe == "Nain" {
				P1.charTurn()
			}
			if P1.classe == "Esprit de la forêt" {
				P1.spiritTurn()
			}
			if P1.classe == "Sirène" {
				P1.mermaidTurn()
			}
			if P1.classe == "Dieu" {
				P1.godTurn()
			}
			time.Sleep(1 * time.Second)
			round = false
		} else {
			fmt.Println(BIGreen+"Round ", i, Reset)
			fmt.Println(">> C'est au tour du", M1.name, " ! <<")
			M1.goblinPattern(i)
			fmt.Println(BIWhite + ">> Appuyez sur entrée pour continuer... " + Reset)
			Input()
			round = true
		}
		if P1.HP == 0 {
			P1.dead()
			fmt.Println(BIWhite + ">> Appuyez sur entrée pour continuer... " + Reset)
			Input()
			Menu()
		}
		if M1.HP <= 0 {
			rand.Seed(time.Now().UnixNano())
			coins := rand.Intn(11) + 10 //coins min:10 & coins max:20
			xp := rand.Intn(28) + 10    //xp min:10 & xp max:37
			fmt.Println(">> ", M1.name, "est mort ! <<")
			fmt.Println(">> Vous avez gagné", coins, "coins ! <<")
			P1.money += coins
			P1.giveXP(xp)
			fmt.Println(BIWhite + ">> Appuyez sur entrée pour continuer... " + Reset)
			Input()
			Menu()
		}
	}
}

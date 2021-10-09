package main

import (
	"fmt"
	"strings"
)

func (p *Personnage) CraftCheck(item map[string]int) bool {
	count := 0
	item0 := 0
	allItem := 0
	for name, quantity := range item {
		if p.inventaire[name] < quantity {
			fmt.Println(">> Vous n'avez pas assez de", name, "! <<")
			fmt.Println(">> Dirigez-vous vers le marchand pour en acheter ! <<")
		}
		if _, isIn := p.inventaire[name]; !isIn || p.inventaire[name] == 0 {
			item0 += 1
			continue
		}
		if p.inventaire[name] >= quantity { //vérifier si le joueur possède le nombre d'item demandé
			if (p.inventaire[name] - quantity) >= 0 { //éviter d'avoir un nombre négatif d'item
				allItem++
				if count == len(item)-1 {
					for name, quantity := range item {
						p.inventaire[name] -= quantity //retirer la quantité d'item après le craft
					}
					return true
				}
			}
		}
		count++
	}
	if count == item0 {
		fmt.Println(">> Vous n'avez pas les objets requis... <<")
	}
	return false
}

func (p *Personnage) Forgeron() {
	//Affichage choix menu
	ClearLog()
	fmt.Println(BIWhite + ">> Forgeron <<" + Reset)
	fmt.Println(">> Vous avez", p.money, "coins ! <<")
	fmt.Println(">> Prix de craft : 5 coins <<")
	fmt.Println("1 >> Fabriquer casque en diamant")
	fmt.Println("\t ∟ 1 plume de corbeau & 1 cuir de sanglier")
	fmt.Println("2 >> Fabriquer jambières en diamant")
	fmt.Println("\t ∟ 2 fourrures de loup & 1 peau de troll")
	fmt.Println("3 >> Fabriquer bottes en diamant")
	fmt.Println("\t ∟ 1 fourrure de loup & 1 cuir de sanglier")
	fmt.Println(BIWhite + "e >> Options" + Reset)
	f, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	f = strings.Replace(f, "\r\n", "", -1)
	switch f {
	case "1":
		if p.CraftCheck(map[string]int{"Plume de corbeau": 1, "Cuir de sanglier": 1}) {
			if p.MoneyCheck(5) {
				p.GiveItem("Casque en diamant")
				fmt.Println(">>", p.nom, "a crafté un casque en diamant ! <<")
			}
		}
	case "2":
		if p.CraftCheck(map[string]int{"Fourrure de loup": 2, "Peau de troll": 1}) {
			if p.MoneyCheck(5) {
				p.GiveItem("Jambières en diamant")
				fmt.Println(">>", p.nom, "a crafté des jambières en diamant ! <<")
			}
		}
	case "3":
		if p.CraftCheck(map[string]int{"Fourrure de loup": 1, "Cuir de sanglier": 1}) {
			if p.MoneyCheck(5) {
				p.GiveItem("Bottes en diamant")
				fmt.Println(">>", p.nom, "a crafté des bottes en diamant ! <<")
			}
		}
	case "e":
		Options()
	default:
		p.Forgeron()
	}
	fmt.Println(BIWhite + ">> Appuyez sur entrée pour continuer... " + Reset)
	Input()
	p.Forgeron()
}

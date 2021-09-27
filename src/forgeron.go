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
		if p.inventaire[name] == 0 {
			item0 += 1
			count++
			continue
		}

		if p.inventaire[name] < quantity {
			fmt.Println(">> Vous n'avez pas assez de", name, "! <<")
			fmt.Println(">> Dirigez-vous vers le marchand pour en acheter!!!!! <<")
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
		fmt.Println(">> Vous n'avez aucun objet requis... <<")
	}
	return false
}

/*
	if p.inventaire[item] == 0 && p.inventaire[item2] == 0 {
		fmt.Println(">> Vous n'avez aucun objet requis... <<")
		fmt.Println(">> Dirigez-vous vers le marchand pour en acheter. <<")
	} else if p.inventaire[item] == 0 { //vérifier si le joueur possède l'item demandé
		fmt.Println(">> Vous n'avez pas assez de", item, "! <<")
		fmt.Println(">> Dirigez-vous vers le marchand pour en acheter. <<")
	} else if p.inventaire[item2] == 0 {
		fmt.Println(">> Vous n'avez pas assez de", item2, "! <<")
		fmt.Println(">> Dirigez-vous vers le marchand pour en acheter. <<")
	}
	if p.inventaire[item] >= quantity && p.inventaire[item2] >= quantity2 { //vérifier si le joueur possède le nombre d'item demandé
		if (p.inventaire[item]-quantity) >= 0 && (p.inventaire[item2]-quantity2) >= 0 { //éviter d'avoir un nombre négatif d'item
			p.inventaire[item] -= quantity //retirer la quantité d'item après le craft
			p.inventaire[item2] -= quantity2
			return true
		}
	}
	return false
}
*/

func (p *Personnage) Forgeron() {
	//Affichage choix menu
	fmt.Println(">> Prix de craft : 5 coins <<")
	fmt.Println("1 >> Fabriquer chapeau de l'aventurier")
	fmt.Println("\t ∟ 1 plume de corbeau & 1 cuir de sanglier")
	fmt.Println("2 >> Fabriquer tunique de l'aventurier")
	fmt.Println("3 >> Fabriquer bottes de l'aventurier")
	fmt.Println("4 >> Options")
	f, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	f = strings.Replace(f, "\r\n", "", -1)
	switch f {
	case "1":
		if p.CraftCheck(map[string]int{"Plume de corbeau": 1, "Cuir de sanglier": 1}) {
			if p.MoneyCheck(5) {
				p.GiveItem("Chapeau de l'aventurier")
				fmt.Println(">>", p.nom, "a crafté un chapeau de l'aventurier ! <<")
			}
		}
	case "4":
		Options()
	}
}

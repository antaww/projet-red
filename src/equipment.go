package main

import "fmt"

type Equipment struct {
	equiptete  string
	equiptorse string
	equippieds string
}

func (p *Personnage) EquipArmor(armor string, maxHPplus int) {
	if p.inventaire[armor] >= 1 {
		p.maxHP += maxHPplus
		p.inventaire[armor]--
		p.equipment.equiptete = armor
		fmt.Println(">>", p.nom, "a équipé", armor, "! <<")
		fmt.Println(">> Votre vie maximale est désormais de", p.maxHP, "HP <<")
	} else {
		fmt.Println(">> Vous ne possédez pas de", armor, "... <<")
	}
}

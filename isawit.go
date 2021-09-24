package main

type Equipement struct { //met le au début après la structure perso
	Equiptete  string
	Equiptorse string
	Equippieds string
}

type Personnage struct {
	nom        string
	classe     string
	niveau     int
	maxHP      int
	HP         int
	inventaire map[string]int
	alive      bool
	money      int
	equipement Equipement
}

func (p *Personnage) Init(Equiptete string, Equiptorse string, Equippieds string) { // fonction
	p.equipement.Equiptete = Equiptete
	p.equipement.Equiptorse = Equiptorse
	p.equipement.Equippieds = Equippieds
}

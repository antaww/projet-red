package main

import (
	"fmt"
	"time"
)

type personnage struct {
	HP int
}

func (p *personnage) poisonPot() {
	for i := 1; i <= 3; i++ {
		p.HP = p.HP - 10
		fmt.Println("salut")
		time.Sleep(1 * time.Second)
	}
}

package main

import (
	"fmt"
	"time"
)



func (p *personnage) PoisonPot() {
	for i := 1 ; i <= 3; i++ {
		p.HP - 10
		fmt.Println(p.HP)
	}

	time.Sleep(1 * time.Second)

}

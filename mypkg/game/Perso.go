package game

import (
	"fmt"
)
//import. "unsafe"

type Dammageable interface {
	Hurt(dammage int)
	Heal(heal int)
}
type Perso struct {
	Name string
	Lvl, Hp, hpMax int
	Wpn Weapon
}

func (p *Perso) Hurt(dammage int) {
	p.Hp -= dammage
	if(p.Hp < 0) {p.Hp = 0}
}
func (p *Perso) Heal (heal int) {
	p.Hp += heal
	if(p.Hp > p.hpMax) {p.Hp = p.hpMax}
}

func (p *Perso) Attack(enemy *Perso, log bool) (b bool){
	if p.Hp == 0 {b = false}
	if enemy.Hp > 0 {
		enemy.Hp -= p.Wpn.Damage
		if log{fmt.Println(p.Name, "attacks",enemy.Name,"with his",p.Wpn.Name)}
		if(enemy.Hp < 0) {
			enemy.Hp = 0
			if log{fmt.Println(enemy.Name,"is dead")}
			b = true

		} else {
			if log{fmt.Println(enemy.Name, "has now", enemy.Hp,"hp")}
		}
	} else {
		if log{fmt.Println(enemy.Name,"is already dead!")}
		b =  true
	}
	fmt.Scanln()
	Clear()
	b = false
	return

}



package game

import (
	"fmt"

	"time"
)
//import. "unsafe"

type Perso struct {

	Name string
	Lvl, Hp int
	Wpn Weapon
}

func (p *Perso) Attack(enemy *Perso) (b bool){
	if p.Hp == 0 {b = false}
	if enemy.Hp > 0 {
		enemy.Hp -= p.Wpn.Damage
		fmt.Println(p.Name, "attacks",enemy.Name,"with his",p.Wpn.Name)
		if(enemy.Hp < 0) {
			enemy.Hp = 0
			fmt.Println(enemy.Name,"is dead")
			//i := (Sizeof(enemy)-Sizeof(array))/Sizeof(array[0])
			//array = append(array[:i],array[i+1]...)
			b = true

		} else {
			fmt.Println(enemy.Name, "has now", enemy.Hp,"hp")
		}
	} else {
		fmt.Println(enemy.Name,"is already dead!")
		b =  true
	}

	b = false
	return

}
func removePerso(a []Perso,index int) []Perso {
	a = append(a[:index],a[index+1:]...)
	b := make([]Perso, len(a),cap(a)-1)
	copy(b, a)
	return b
}

func CleanDeads(a []Perso) {
	lenght := len(a)

	for i := 0 ; i < lenght ; i++ {
		if(a[i].Hp == 0) {
			a = removePerso(a,i)
			i--
			lenght--
		}
		fmt.Println(cap(a), len(a))
		time.Sleep(time.Second)
	}
}

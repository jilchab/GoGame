package main

import (
	"fmt"
	//"github.com/jilchab/GoGame/mypkg/game"
)
type Weapon struct {
	name string
	damage int
	//cooldown_ms int

}
type Perso struct {

	name string
	lvl, hp int
	weapon Weapon
}
func (p *Perso) Attack(enemy *Perso) {
	if p.hp == 0 {return}
	if enemy.hp > 0 {
		enemy.hp -= p.weapon.damage
		fmt.Printf("%s attacks %s with his %s\n",p.name,enemy.name,p.weapon.name)
		if(enemy.hp < 0) {
			enemy.hp = 0
			fmt.Printf("%s is dead\n", enemy.name)
		} else {
			fmt.Printf("%s has now %d hp\n",enemy.name, enemy.hp)
		}
	} else {fmt.Println(enemy.name,"is already dead!")}



}

func main () {

	//Player := Perso{"Hero",1,100,Weapon{"sword",20}}
	//Gobgob := Perso{"Gobgob",1,30,Weapon{"pickaxe",10}}


	fmt.Println("Hello World")


}

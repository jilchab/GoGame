package game

import "fmt"

type Perso struct {

	Name string
	Lvl, Hp int
	Wpn Weapon
}
func (p *Perso) Attack(enemy *Perso) {
	if p.Hp == 0 {return}
	if enemy.Hp > 0 {
		enemy.Hp -= p.Wpn.Damage
		fmt.Printf("%s attacks %s with his %s\n",p.Name,enemy.Name,p.Wpn.Name)
		if(enemy.Hp < 0) {
			enemy.Hp = 0
			fmt.Printf("%s is dead\n", enemy.Name)
		} else {
			fmt.Printf("%s has now %d hp\n",enemy.Name, enemy.Hp)
		}
	} else {fmt.Println(enemy.Name,"is already dead!")}



}
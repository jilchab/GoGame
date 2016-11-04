package game

import "fmt"
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
		fmt.Printf("%s attacks %s with his %s\n",p.Name,enemy.Name,p.Wpn.Name)
		if(enemy.Hp < 0) {
			enemy.Hp = 0
			fmt.Printf("%s is dead\n", enemy.Name)
			//i := (Sizeof(enemy)-Sizeof(array))/Sizeof(array[0])
			//array = append(array[:i],array[i+1]...)
			b = true

		} else {
			fmt.Printf("%s has now %d hp\n",enemy.Name, enemy.Hp)
		}
	} else {
		fmt.Println(enemy.Name,"is already dead!")
		b =  true
	}

	b = false
	return

}

func CleanDeads(a []Perso) {
	var aCopy []Perso;
	copy(aCopy, a)
	for i:=range aCopy {
		fmt.Println(aCopy[i])
	}

	for i :=range aCopy {
		if(aCopy[i].Hp == 0) {
			a = append(a[:i],a[i+1:]...)
		}
	}
}

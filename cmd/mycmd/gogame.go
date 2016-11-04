package main

import (
	. "github.com/jilchab/GoGame/mypkg/game"
	"fmt"
)




func main () {

	pickaxe := Weapon{Name:"pickaxe",Damage:10}
	sword := Weapon{Name:"sword",Damage:20}

	Player := Perso{Name:"Hero",Lvl:1,Hp:100,Wpn:sword}
	var Gobgobs []Perso



	Gobgobs = append(Gobgobs,
		Perso{Name:"Jean",Lvl:1,Hp:50,Wpn:pickaxe},
		Perso{Name:"Claude",Lvl:1,Hp:30,Wpn:pickaxe},
		Perso{Name:"Gobgob",Lvl:1,Hp:10,Wpn:pickaxe},
		Perso{Name:"Paul",Lvl:1,Hp:5,Wpn:pickaxe})



	for i :=range Gobgobs {
		Player.Attack(&Gobgobs[i])
	}
	CleanDeads(Gobgobs[:])

	for _,g:=range Gobgobs {
		fmt.Println(g)
	}


}

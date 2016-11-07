package game

import (
	"os/exec"
	"os"
)

func Clear() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func removePerso(a []Perso,index int) []Perso {
	a = append(a[:index],a[index+1:]...)
	b := make([]Perso, len(a),cap(a)-1)
	copy(b, a)
	return b
}

func CleanDeads(a []Perso) []Perso{
	for i := 0 ; i < len(a) ; i++ {
		if(a[i].Hp == 0) {
			a = removePerso(a,i)
			i--
		}
	}
	return a
}

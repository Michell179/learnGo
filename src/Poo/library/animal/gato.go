package animal

import "fmt"

type Gato struct {
	Nombre string
}

func (p *Gato) Sonido() {
	fmt.Println(p.Nombre, " Miau! Miau!")
}

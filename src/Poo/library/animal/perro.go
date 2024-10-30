package animal

import "fmt"

type Perro struct {
	Nombre string
}

func (p *Perro) Sonido() {
	fmt.Println(p.Nombre, " Guau! Guau!")
}

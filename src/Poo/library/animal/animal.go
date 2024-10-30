package animal

type Animal interface {
	Sonido()
}

func ReproducirSonido(animal Animal) {
	animal.Sonido()
}

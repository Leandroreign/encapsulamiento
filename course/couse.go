package course

type Course struct {
	Name        string
	Price       float64
	IsActive    bool
	ProffesorID uint
	Classes     map[uint]string
	UserIds     []uint
}

// si queremos actualizar datos de una estructura entonces debemos usar operadores
// de tipo puntero en los metodos mas sino queremos actualizar pues usamos el valor normal

// si esta trabajando con interfaces y uno de los metodos del struct es de tipo puntero
// por buena practica es mejor ponerlos todos de tipo puntero

// los parentesis antes del nombre de la funcion son para identificar a que struct
// esta vinculado el metodo, este ya no es una funcion sino que es un metodo de la
// estructura Course
func (c Course) printClasses() {
	print("\n")
	text := "Las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	print(text[:len(text)-2])
}

// el receptor es una copia del valor por lo que esto no
// cambiaria el valor de precio de un curso
func (c Course) ChangePrice1(newPrice float64) {
	c.Price = newPrice
}

// para que si funcione debo reibir un puntero y no el curso
func (c *Course) ChangePrice(newPrice float64) {
	c.Price = newPrice
}

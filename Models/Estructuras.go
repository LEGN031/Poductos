package Models

type Usuarios struct {
	Nombre string
	Edad   int
	Cedula string
}

type Compra struct {
	Producto Productos
	Cantidad int
}

func (c Compra) Precio() float64 {
	return float64(c.Cantidad) * c.Producto.PrecioUnitario
}

type Clientes struct {
	Usuarios
	ListaCompra []Compra
}

func (c *Clientes) AgregarCompra(nuevaCompra Compra) {
	c.ListaCompra = append(c.ListaCompra, nuevaCompra)
}

func (c *Clientes) TotalCompras() float64 {
	total := 0.0
	for _, compra := range c.ListaCompra {
		total += compra.Precio()
	}
	return total
}

type Empleados struct {
	Usuarios
	Cargo          string
	Salario        float64
	codigoEmpleado string
	Catalogo       Catalogos
}

type Productos struct {
	Nombre         string
	PrecioUnitario float64
}

type Catalogos struct {
	Producto5 []Productos
}

func (c *Catalogos) AgregarProducto(Nuevoproducto Productos) {
	c.Producto5 = append(c.Producto5, Nuevoproducto)
}

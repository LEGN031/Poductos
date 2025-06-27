package Models

import _ "fmt"

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
	CodigoEmpleado string
	Catalogo       Catalogos
}

type Productos struct {
	Id             int
	Nombre         string
	PrecioUnitario float64
	CodigoProducto string
}

type Catalogos struct {
	Producto5 []Productos
}

func (c *Catalogos) AgregarProducto(Nuevoproducto Productos) {
	c.Producto5 = append(c.Producto5, Nuevoproducto)
}

/*func (c *Catalogos) BuscarProducto (IdBusqueda int) *Productos {
	for_, producto := range c.Producto5 {
		if producto.Id == IdBusqueda {
			return &c.Producto5[producto.Id]
		}
	}
}*/

func (c *Catalogos) BuscarProducto(IdBusqueda int) *Productos {
	for i := range c.Producto5 {
		if c.Producto5[i].Id == IdBusqueda {
			return &c.Producto5[i]
		}
	}
	return nil
}

func (c *Catalogos) ModificarProducto(IdBusqueda int, nuevoProducto Productos) bool {
	producto := c.BuscarProducto(IdBusqueda)
	if producto != nil {
		*producto = nuevoProducto
		return true
	}
	return false
}

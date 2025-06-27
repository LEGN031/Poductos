package Models

import _ "fmt"

type Usuarios struct {
	Id     int
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
	Contraseña  string
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

func (c *Catalogos) BuscarProductoPorIndice(IdBusqueda int) int {
	for i := range c.Producto5 {
		if c.Producto5[i].Id == IdBusqueda {
			return i
		}
	}
	return -1
}

func (c *Catalogos) BuscarProductoPorCodigo(codigoBusqueda string) int {
	for i := range c.Producto5 {
		if c.Producto5[i].CodigoProducto == codigoBusqueda {
			return i
		}
	}
	return -1
}

func (c *Catalogos) ModificarProducto(IdBusqueda int, nuevoProducto Productos) bool {
	producto := c.BuscarProducto(IdBusqueda)
	if producto != nil {
		*producto = nuevoProducto
		return true
	}
	return false
}

func (c *Catalogos) EliminarProducto(IdBusqueda int) bool {
	indice := c.BuscarProductoPorIndice(IdBusqueda)
	if indice != -1 {
		c.Producto5 = append(c.Producto5[:indice], c.Producto5[indice+1:]...)
		return true
	}
	return false
}

//Clientes

func (c *Clientes) AñadirProductoALaCesta(catalogo *Catalogos, codigoBusqueda string, cantidad int) bool {
	indice := catalogo.BuscarProductoPorCodigo(codigoBusqueda)
	if indice != -1 {
		producto := catalogo.Producto5[indice]

		for i := range c.ListaCompra {
			if c.ListaCompra[i].Producto.Id == producto.Id {
				c.ListaCompra[i].Cantidad += cantidad
				return true
			}
		}
		c.ListaCompra = append(c.ListaCompra, Compra{Producto: producto, Cantidad: cantidad})
		return true
	}
	return false
}

func (c *Clientes) CrearNuevoRegistro(nombre string, edad int, cedula string, contraseña string) {
	c.Usuarios.Nombre = nombre
	c.Usuarios.Edad = edad
	c.Usuarios.Cedula = cedula
	c.Contraseña = contraseña
	c.ListaCompra = []Compra{}
}

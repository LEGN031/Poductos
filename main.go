package main

import (
	"Productos/Models"
	"fmt"
)

func main() {
	var respuesta string

	clientes := []Models.Clientes{
		{
			Usuarios: Models.Usuarios{Nombre: "Carlos", Edad: 40, Cedula: "111222", Id: 1},
			ListaCompra: []Models.Compra{
				{Producto: Models.Productos{Nombre: "Laptop", PrecioUnitario: 1200, CodigoProducto: "LAP123", Id: 1}, Cantidad: 1},
				{Producto: Models.Productos{Nombre: "Mouse", PrecioUnitario: 25, CodigoProducto: "MOU456", Id: 2}, Cantidad: 2},
			},
		},
		{
			Usuarios: Models.Usuarios{Nombre: "Maria", Edad: 35, Cedula: "333444", Id: 2},
			ListaCompra: []Models.Compra{
				{Producto: Models.Productos{Nombre: "Teclado", PrecioUnitario: 45}, Cantidad: 4},
			},
		},
	}

	laptop := Models.Productos{Nombre: "Laptop", PrecioUnitario: 1200, CodigoProducto: "LAP123", Id: 1}
	mouse := Models.Productos{Nombre: "Mouse", PrecioUnitario: 25, CodigoProducto: "MOU456", Id: 2}

	catalogo := Models.Catalogos{
		Producto5: []Models.Productos{laptop, mouse},
	}

	fmt.Println("Bienvenido al sistema de gestión de usuarios y Compras")
	fmt.Println("¿Eres empleado o cliente? (1 para empleado / 2 para cliente):")
	fmt.Scanln(&respuesta)

	if respuesta == "1" {
		for {
			var opcion int
			fmt.Println("Opciones de empleado:")
			fmt.Println("1. Agregar producto")
			fmt.Println("2. Modificar producto")
			fmt.Println("3. Eliminar producto")
			fmt.Println("4. Imprimir")
			fmt.Println("5. Salir")
			fmt.Print("Seleccione una opción: ")
			fmt.Scanln(&opcion)

			switch opcion {
			case 1:
				var nombre, codigo string
				var precio float64
				var id int
				fmt.Print("Ingrese el id del producto: ")
				fmt.Scanln(&id)
				fmt.Print("Ingrese el código del producto: ")
				fmt.Scanln(&codigo)
				fmt.Print("Ingrese el nombre del producto: ")
				fmt.Scanln(&nombre)
				fmt.Print("Ingrese el precio unitario: ")
				fmt.Scanln(&precio)
				catalogo.AgregarProducto(Models.Productos{Nombre: nombre, PrecioUnitario: precio, Id: id, CodigoProducto: codigo})
			case 2:
				var id, idBusqueda int
				var nombre, codigo string
				var precio float64
				fmt.Print("Ingrese el id del producto a modificar: ")
				fmt.Scanln(&idBusqueda)
				fmt.Print("Ingrese el nuevo id del producto: ")
				fmt.Scanln(&id)
				fmt.Print("Ingrese el nuevo nombre del producto: ")
				fmt.Scanln(&nombre)
				fmt.Print("Ingrese el nuevo precio unitario: ")
				fmt.Scanln(&precio)
				fmt.Print("Ingrese el nuevo código del producto: ")
				fmt.Scanln(&codigo)
				nuevoProducto := Models.Productos{Id: id, Nombre: nombre, PrecioUnitario: precio, CodigoProducto: codigo}
				catalogo.ModificarProducto(idBusqueda, nuevoProducto)
			case 3:
				var id int
				fmt.Print("Ingrese el id del producto a eliminar: ")
				fmt.Scanln(&id)
				if catalogo.EliminarProducto(id) {
					fmt.Println("Producto eliminado correctamente.")
				} else {
					fmt.Println("Producto no encontrado.")
				}
			case 4:
				fmt.Println("Productos en el catálogo:")
				for _, producto := range catalogo.Producto5 {
					fmt.Println("Id:", producto.Id, "Nombre:", producto.Nombre, "Precio Unitario:", producto.PrecioUnitario, "Código:", producto.CodigoProducto)
				}
			case 5:
				fmt.Println("Saliendo del sistema...")
				return
			default:
				fmt.Println("Opción no válida.")
			}
		}
	} else if respuesta == "2" {
		var cedulaIngresada string
		fmt.Println("Ingrese su cédula:")
		fmt.Scanln(&cedulaIngresada)

		var clienteEncontrado *Models.Clientes
		for i := range clientes {
			if clientes[i].Usuarios.Cedula == cedulaIngresada {
				clienteEncontrado = &clientes[i]
				break
			}
		}

		if clienteEncontrado == nil {
			var respuestaCedula string
			fmt.Println("Cédula no encontrada. ¿Es cliente nuevo? (1 para sí / 2 para no):")
			fmt.Scanln(&respuestaCedula)

			if respuestaCedula == "1" {
				var nombre, cedula, contraseña string
				var edad int
				fmt.Print("Ingrese su nombre: ")
				fmt.Scanln(&nombre)
				fmt.Print("Ingrese su edad: ")
				fmt.Scanln(&edad)
				fmt.Print("Ingrese su cédula: ")
				fmt.Scanln(&cedula)
				fmt.Print("Ingrese su contraseña: ")
				fmt.Scanln(&contraseña)

				nuevoCliente := Models.Clientes{
					Usuarios:    Models.Usuarios{Nombre: nombre, Edad: edad, Cedula: cedula, Id: len(clientes) + 1},
					ListaCompra: []Models.Compra{},
					Contraseña:  contraseña,
				}

				clientes = append(clientes, nuevoCliente)
				fmt.Println("Cliente registrado correctamente.")
				OpcionesCliente(&clientes[len(clientes)-1], &catalogo)
			} else if respuestaCedula == "2" {
				fmt.Println("Por favor, intente nuevamente con una cédula válida.")
				return
			} else {
				fmt.Println("Opción no válida. Intente nuevamente.")
				return
			}
		} else {
			OpcionesCliente(clienteEncontrado, &catalogo)
		}
	} else {
		fmt.Println("Opción no válida. Intente nuevamente.")
	}
}

func OpcionesCliente(cliente *Models.Clientes, catalogo *Models.Catalogos) {
	for {
		var opcion int
		fmt.Println("Opciones de Clientes:")
		fmt.Println("1. Agregar producto a la cesta")
		fmt.Println("2. Modificar producto")
		fmt.Println("3. Eliminar producto")
		fmt.Println("4. Imprimir")
		fmt.Println("5. Salir")
		fmt.Print("Seleccione una opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var codigoBusqueda string
			var cantidad int
			fmt.Print("Ingrese el código del producto a agregar a la cesta: ")
			fmt.Scanln(&codigoBusqueda)
			fmt.Print("Ingrese la cantidad: ")
			fmt.Scanln(&cantidad)
			if cliente.AñadirProductoALaCesta(catalogo, codigoBusqueda, cantidad) {
				fmt.Println("Producto agregado a la cesta correctamente.")
			} else {
				fmt.Println("Producto no encontrado o cantidad inválida.")
			}
		case 2:
			// Aquí puedes implementar la opción de modificar un producto en la cesta
		case 3:
			// Aquí puedes implementar la opción de eliminar un producto en la cesta
		case 4:
			fmt.Println("Productos en la cesta de compras:")
			for _, compra := range cliente.ListaCompra {
				fmt.Println("Producto:", compra.Producto.Nombre, "Cantidad:", compra.Cantidad, "Precio Unitario:", compra.Producto.PrecioUnitario)
			}
		case 5:
			fmt.Println("Saliendo del sistema...")
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}

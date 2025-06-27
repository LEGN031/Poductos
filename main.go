package main

import (
	"Productos/Models"
	"fmt"
)

func main() {
	//var compobacion string
	var respuesta string

	// empleados := []Empleados{a
	// 	{Usuarios: Usuarios{Nombre: "Ana", Edad: 30, Cedula: "123456"}, Cargo: "Gerente", Salario: 5000, codigoEmpleado: "EMP001"},
	// 	{Usuarios: Usuarios{Nombre: "Luis", Edad: 25, Cedula: "654321"}, Cargo: "Vendedor", Salario: 3000, codigoEmpleado: "EMP002"},
	// }

	/*clientes := []Clientes{
		{
			Usuarios: Usuarios{Nombre: "Carlos", Edad: 40, Cedula: "111222"},
			ListaCompra: []Compra{
				{Producto: Productos{Nombre: "Laptop", PrecioUnitario: 1200}, Cantidad: 1},
				{Producto: Models.Productos{Nombre: "Mouse", PrecioUnitario: 25}, Cantidad: 2},
			},

		},
		{
			Usuarios: Usuarios{Nombre: "Maria", Edad: 35, Cedula: "333444"},
			ListaCompra: []Compra{
				{Producto: Productos{Nombre: "Teclado", PrecioUnitario: 45}, Cantidad: 1},
			},
		},
	}*/

	laptop := Models.Productos{Nombre: "Laptop", PrecioUnitario: 1200, CodigoProducto: "LAP123", Id: 1}
	mouse := Models.Productos{Nombre: "Mouse", PrecioUnitario: 25, CodigoProducto: "MOU456", Id: 2}

	catalogo := Models.Catalogos{
		Producto5: []Models.Productos{laptop, mouse},
	}
	fmt.Println("Bienvenido al sistema de gestión de usuarios y Compras")

	fmt.Println("Eres empleado o cliente? (1 para empleado/2 para cliente):")
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
				var nombre string
				var precio float64
				var id int
				var codigo string
				fmt.Print("Ingrese el id del producto: ")
				fmt.Scanln(&id)
				fmt.Print("Ingrese el código del producto: ")
				fmt.Scanln(&codigo)
				fmt.Print("Ingrese el nombre del producto: ")
				fmt.Scanln(&nombre)
				fmt.Print("Ingrese el precio unitario: ")
				fmt.Scanln(&precio)
				catalogo.AgregarProducto(Models.Productos{Nombre: nombre, PrecioUnitario: precio})
			case 2:
				var id int
				var nombre string
				var precio float64
				var codigo string
				var Nuevoproducto Models.Productos
				var IdBusqueda int
				fmt.Print("Ingrese el id del producto a modificar (0 para el primero): ")
				fmt.Scanln(&IdBusqueda)
				fmt.Print("Ingrese el nuevo id del producto: ")
				fmt.Scanln(&id)
				fmt.Print("Ingrese el nuevo nombre del producto: ")
				fmt.Scanln(&nombre)
				fmt.Print("Ingrese el nuevo precio unitario: ")
				fmt.Scanln(&precio)
				fmt.Print("Ingrese el nuevo código del producto: ")
				fmt.Scanln(&codigo)
				Nuevoproducto = Models.Productos{Id: id, Nombre: nombre, PrecioUnitario: precio, CodigoProducto: codigo}
				catalogo.ModificarProducto(IdBusqueda, Nuevoproducto)

			case 3:
				// Aquí deberías mostrar la lista de productos y permitir eliminar uno
				fmt.Println("Funcionalidad de eliminar producto (por implementar)")
			case 4:
				fmt.Println("Productos en el catálogo:")
				for _, producto := range catalogo.Producto5 {
					fmt.Println("Nombre:", producto.Nombre, "Precio Unitario:", producto.PrecioUnitario)
				}
			case 5:
				fmt.Println("Saliendo del sistema...")
				return
			default:
				fmt.Println("Opción no válida")
			}

		}
	} else if respuesta == "2" {
		for {
			var opcion int
			fmt.Println("Opciones de Clientes:")
			fmt.Println("1. Agregar producto")
			fmt.Println("2. Modificar producto")
			fmt.Println("3. Eliminar producto")
			fmt.Println("4. Imprimir")
			fmt.Println("5. Salir")
			fmt.Print("Seleccione una opción: ")
			fmt.Scanln(&opcion)

		}
	} else {
		fmt.Println("Opción no válida. Intenta nuevamente.")
		return
	}
}

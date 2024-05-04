package main

import (
	"bufio"
	"fmt"
	"os"
)

// Estructura para representar todos los usuarios
type Usuario struct {
	id       int
	nick     string
	password string
	esAdmin  bool
}

// Funcion para crear usuarios
func crearUsuario(listaUsuario map[string]*Usuario, id int, nick, password string, esAdmin bool) {
	nuevoUsuario := Usuario{
		id:       id,
		nick:     nick,
		password: password,
		esAdmin:  esAdmin,
	}
	listaUsuario[nick] = &nuevoUsuario
}

// Funcion para iniciar sesion
func iniciarSesion(listaUsuario map[string]*Usuario, nick string, password string) {
	usuario, ok := listaUsuario[nick]
	if ok {
		if usuario.nick == nick && usuario.password == password {
			if usuario.esAdmin {
				menuAdmin(listaUsuario)
			} else {
				menuSupervisor(listaUsuario)
			}
		}
	} else {
		fmt.Println("Usuario no existe", nick)
	}
}

// Mostrar todos los usuarios
func mostrarUsuario(listaUsuario map[string]*Usuario) {
	fmt.Println("Todos los usuarios: ")
	for _, usuario := range listaUsuario {
		fmt.Printf("Identificador '%d' Usuario '%s' es Administrador '%t'.\n", usuario.id, usuario.nick, usuario.esAdmin)
	}
}

// Funcion para eliminar usuarios
func eliminarUsuario(listaUsuario map[string]*Usuario, nick string) {
	if _, encontrado := listaUsuario[nick]; encontrado {
		delete(listaUsuario, nick)
		fmt.Printf("El usuario '%s' ha sido eliminado.\n", nick)
	} else {
		fmt.Printf("El usuario '%s' no existe.\n", nick)
	}
}

// Funcion para buscar supervisores repetidos
func buscarRepetido(listaUsuario map[string]*Usuario) {
	idsVistos := make(map[int]bool) // Mapa para rastrear los Id de usuario vistos

	for _, usuario := range listaUsuario {
		// Verificar si el Id de usuario ya ha sido visto
		if _, encontrado := idsVistos[usuario.id]; encontrado {
			if usuario.esAdmin == false {
				for {
					fmt.Print("Desea eliminar el Supervisor repetido (s/n)?: ")
					var respuesta string
					fmt.Scanln(&respuesta)
					if respuesta == "s" || respuesta == "S" {
						eliminarUsuario(listaUsuario, usuario.nick)
						fmt.Println("El supervisor se ha eliminado correctamente.")
						break
					} else {
						if respuesta == "n" || respuesta == "N" {
							fmt.Println("El supervisor no ha sido eliminado.")
							break
						} else {
							fmt.Printf("La respuesta es incorrecta, vuelva a intentar.\n")
						}
					}
				}
			}
		} else {
			// El Id de usuario ya visto
			idsVistos[usuario.id] = true
		}
	}
}

// Funcion para contar supervisores
func contarSupervisor(listaUsuario map[string]*Usuario) int {
	contador := 0
	for _, usuario := range listaUsuario {
		if usuario.esAdmin == false {
			contador++
		}
	}
	return contador
}

// Funcion para generar un informe con la cantidad de supervisores
func generarInforme(listaUsuario map[string]*Usuario) {
	fmt.Print("Ingrese el nombre del informe que desea crear: ")
	var informe string
	fmt.Scanln(&informe)
	archivo, err := os.Create(informe)
	if err != nil {
		fmt.Println("Error al crear el archivo: ", err)
		return
	}
	defer archivo.Close()

	escritor := bufio.NewWriter(archivo)

	cantindadSupervisor := contarSupervisor(listaUsuario)

	texto := fmt.Sprintf("La cantidad actual de Supervisores es de: '%d'", cantindadSupervisor)

	_, err = escritor.WriteString(texto)
	if err != nil {
		fmt.Println("Error al escribir en el archivo: ", err)
		return
	}

	// Guardar los cambios en el archivo
	err = escritor.Flush()
	if err != nil {
		fmt.Println("Error al guardar los cambios en el archivo:", err)
		return
	}

	fmt.Println("Se ha escrito en el archivo correctamente.")
}

// Funcion para leer un informe por su nombre
func leerInforme() {
	// Nombre del archivo a leer

	fmt.Print("Ingrese el nombre del Informe que desea leer: ")
	var nombreArchivo string
	fmt.Scanln(&nombreArchivo)

	// Abrir el archivo en modo lectura
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		fmt.Printf("Error al abrir el archivo %s: %v\n", nombreArchivo, err)
		return
	}
	defer archivo.Close()

	// Crear un lector para el archivo
	lector := bufio.NewScanner(archivo)

	// Leer el contenido
	for lector.Scan() {
		linea := lector.Text()
		fmt.Println(linea)
	}

	// Verificar si ocurrieron errores durante la lectura
	if err := lector.Err(); err != nil {
		fmt.Printf("Error al leer el archivo %s: %v\n", nombreArchivo, err)
	}
}

// Funcion para eliminar un informe por su nombre
func eliminarInforme() {
	// Nombre del archivo a eliminar
	fmt.Print("Ingrese el nombre del Informe que desea eliminar: ")
	var nombreArchivo string
	fmt.Scanln(&nombreArchivo)

	// Intentar eliminar el archivo
	err := os.Remove(nombreArchivo)
	if err != nil {
		fmt.Printf("Error al eliminar el archivo %s: %v\n", nombreArchivo, err)
		return
	}

	fmt.Printf("El archivo %s ha sido eliminado exitosamente.\n", nombreArchivo)
}

// Menu de Administradores
func menuAdmin(listaUsuario map[string]*Usuario) {
	for {
		fmt.Println("*********************************")
		fmt.Println("Opciones de Administrador:")
		fmt.Println("*********************************")
		fmt.Println("1. Crear un nuevo Usuario")
		fmt.Println("2. Ver lista de Usuarios")
		fmt.Println("3. Eliminar un Usuario")
		fmt.Println("4. Leer Informes")
		fmt.Println("5. Eliminar informe")
		fmt.Println("6. Cerrar sesion")
		fmt.Println("*********************************")
		fmt.Print("Ingrese su opcion: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Crear un nuevo Usuario")
			fmt.Print("Ingrese el Id del nuevo Usuario: ")
			var id int
			fmt.Scanln(&id)
			fmt.Print("Ingrese el nuevo Usuario: ")
			var nick string
			fmt.Scanln(&nick)
			fmt.Print("Ingrese el password: ")
			var password string
			fmt.Scanln(&password)
			for {
				fmt.Print("El usuario es Administrador (s/n)?: ")
				var respuesta string
				fmt.Scanln(&respuesta)
				if respuesta == "s" || respuesta == "S" {
					crearUsuario(listaUsuario, id, nick, password, true)
					fmt.Println("Nuevo Administrador creado correctamente")
					break
				} else {
					if respuesta == "n" || respuesta == "N" {
						crearUsuario(listaUsuario, id, nick, password, false)
						fmt.Println("Nuevo Supervisor creado correctamente")
						break
					} else {
						fmt.Printf("La respuesta es incorrecta, vuelva a intentar.\n")
					}
				}
			}
		case 2:
			mostrarUsuario(listaUsuario)
		case 3:
			fmt.Print("Ingrese el nombre de Usuario que desea eliminar: ")
			var usuario string
			fmt.Scanln(&usuario)
			for {
				fmt.Print("Esta seguro que desea eliminar el Usuario (s/n)?: ")
				var respuesta string
				fmt.Scanln(&respuesta)
				if respuesta == "s" || respuesta == "S" {
					eliminarUsuario(listaUsuario, usuario)
					break
				} else {
					if respuesta == "n" || respuesta == "N" {
						break
					}
				}
			}
		case 4:
			leerInforme()
		case 5:
			eliminarInforme()
		case 6:
			return
		default:
			fmt.Println("Opcion invalida")
		}
	}
}

// Menu de Supervisores
func menuSupervisor(listaUsuario map[string]*Usuario) {
	for {
		fmt.Println("*********************************")
		fmt.Println("Opciones de Supervisor:")
		fmt.Println("*********************************")
		fmt.Println("1. Ver lista de Usuarios")
		fmt.Println("2. Crear un informe con la cantidad de Supervisores existentes")
		fmt.Println("3. Eliminar Supervisores repetidos")
		fmt.Println("4. Cerrar sesion")
		fmt.Println("*********************************")
		fmt.Print("Ingrese su opcion: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			mostrarUsuario(listaUsuario)
		case 2:
			generarInforme(listaUsuario)
		case 3:
			buscarRepetido(listaUsuario)
		case 4:
			return
		default:
			fmt.Println("Opcion invalida")
		}
	}
}
func main() {
	// Mapa para todos los usuario
	listaUsuario := make(map[string]*Usuario)
	// Creo un usuario Adiministrador que es el unico que no se puede crear desde el menu principal
	crearUsuario(listaUsuario, 1, "Admin", "root", true)

	// Menu principal
	for {
		fmt.Println("1. Iniciar sesion")
		fmt.Println("2. Crear usuario")
		fmt.Println("3. Salir")

		fmt.Print("Ingrese su opcion: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Ingrese su Usuario: ")
			var usuario string
			fmt.Scanln(&usuario)
			fmt.Print("Ingrese su Password: ")
			var password string
			fmt.Scanln(&password)
			iniciarSesion(listaUsuario, usuario, password)
		case 2:
			fmt.Println("Crear un nuevo Usuario")
			fmt.Print("Ingrese el Id del nuevo Usuario: ")
			var id int
			fmt.Scanln(&id)
			fmt.Print("Ingrese el nuevo Usuario: ")
			var nick string
			fmt.Scanln(&nick)
			fmt.Print("Ingrese el password: ")
			var password string
			fmt.Scanln(&password)
			crearUsuario(listaUsuario, id, nick, password, false)
		case 3:
			os.Exit(0)
		default:
			fmt.Println("Opcion invalida")
		}
	}
}

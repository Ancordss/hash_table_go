package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var opt int
	fmt.Println("leer o ingresar texto? 1/2")
	fmt.Scanln(&opt)
	if opt == 1 {

		// Declaración de variables
		archivo, err := os.Open("archivo.txt")
		if err != nil {
			fmt.Println("Error al abrir el archivo:", err)
			return
		}
		defer archivo.Close()

		scanner := bufio.NewScanner(archivo)
		var texto string

		for scanner.Scan() {
			texto += scanner.Text() + "\n"
		}

		if scanner.Err() != nil {
			fmt.Println("Error al leer el archivo:", scanner.Err())
			return
		}

		fmt.Println("Contenido del archivo:")
		fmt.Println(texto)

		showHash(texto)
		search(texto)

	} else if opt == 2 {
		var texto string

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Ingresa un texto con espacios: ")
		texto, _ = reader.ReadString('\n')
		texto = strings.TrimRight(texto, "\n")
		showHash(texto)
		search(texto)
	}

}

func search(texto string) {

	reader := bufio.NewReader(os.Stdin)

	var posicion string
	fmt.Print("Ingresa la posición del token: ")
	posicion, _ = reader.ReadString('\n')
	posicion = strings.TrimRight(posicion, "\n")

	buscarPorPosicion(posicion, texto)
}

func showHash(texto string) {
	// Tabla hash para almacenar los tokens con sus posiciones
	tablaHash := make(map[string]string)

	// Procesamiento del texto
	filas := strings.Split(texto, "\n")
	for i, fila := range filas {
		columnas := strings.Split(fila, " ")
		for j, token := range columnas {
			tablaHash[token] = fmt.Sprintf("(%d, %d)", i, j)
		}
	}

	// Imprimir la tabla hash
	for token, posicion := range tablaHash {
		fmt.Printf("Token: %s\tPosición: %s\n", token, posicion)
	}
}

func buscarPorPosicion(posicion string, texto string) {
	// Procesamiento del texto para encontrar el token según la posición
	filas := strings.Split(texto, "\n")
	for i, fila := range filas {
		columnas := strings.Split(fila, " ")
		for j, token := range columnas {
			if fmt.Sprintf("(%d, %d)", i, j) == posicion {
				fmt.Printf("Token encontrado: %s\n", token)
				return
			}
		}
	}

	fmt.Println("Token no encontrado para la posición:", posicion)
}

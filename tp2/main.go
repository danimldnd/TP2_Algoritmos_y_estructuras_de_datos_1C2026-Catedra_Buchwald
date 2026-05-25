package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// funcion para abrir archivo
func abrirArchivo(path string) (*bufio.Scanner, *os.File) {
	archivo, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error al abrir archivo:", err)
		os.Exit(1)
	}
	return bufio.NewScanner(archivo), archivo
}

// carga de doctores y pacientes ------------------------------------------------------------------
func cargaDoctores(path string) map[string]string {
	scanner, archivo := abrirArchivo(path)
	defer archivo.Close()

	doctores := make(map[string]string)
	for scanner.Scan() {
		partes := strings.SplitN(scanner.Text(), ",", 2)
		if len(partes) != 2 {
			continue
		}
		doctores[partes[0]] = partes[1]
	}
	return doctores
}

func cargaPacientes(path string) map[string]int {
	scanner, archivo := abrirArchivo(path)
	defer archivo.Close()

	pacientes := make(map[string]int)
	for scanner.Scan() {
		partes := strings.SplitN(scanner.Text(), ",", 2)
		if len(partes) != 2 {
			continue
		}
		anio, err := strconv.Atoi(strings.TrimSpace(partes[1]))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Año invalido para paciente:", partes[0])
			os.Exit(1)
		}
		pacientes[partes[0]] = anio
	}
	return pacientes
}

// mapa de especialidades----------------------------------------------------------------------
func construirEspecialidades(doctores map[string]string) map[string]*colaEspecialidad {
	especialidades := make(map[string]*colaEspecialidad)
	for _, especialidad := range doctores {
		if _, existe := especialidades[especialidad]; !existe {
			especialidades[especialidad] = nuevaColaEspecialidad()
		}
	}
	return especialidades
}

// loop de comandos ---------------------------------------------------------------------------
func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Uso: tp2 <doctores.csv> <pacientes.csv>")
		os.Exit(1)
	}

	doctores := cargaDoctores(os.Args[1])
	pacientes := cargaPacientes(os.Args[2])
	especialidades := construirEspecialidades(doctores)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linea := scanner.Text()
		partes := strings.SplitN(linea, ":", 2)
		if len(partes) != 2 {
			fmt.Printf("ERROR: formato de comando incorrecto ('%s')\n", linea)
			continue
		}

		comando := partes[0]
		args := strings.Split(partes[1], ",")

		switch comando {
		case "PEDIR_TURNO":
			pedirTurno(args, pacientes, especialidades)
		case "ATENDER_SIGUIENTE":
			atenderSiguiente(args, doctores)
		case "INFORME":
			informe(args, doctores)
		default:
			fmt.Printf("ERROR: no existe el comando '%s'\n", comando)
		}
	}
}

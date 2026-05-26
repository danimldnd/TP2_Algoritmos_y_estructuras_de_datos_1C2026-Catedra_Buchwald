package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	diccionario "tdas/diccionarios"
)

func cmpStrings(a, b string) int {
	return strings.Compare(a, b)
}

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
func cargaDoctores(path string) diccionario.DiccionarioOrdenado[string, *doctor] {
	scanner, archivo := abrirArchivo(path)
	defer archivo.Close()

	doctores := diccionario.CrearABB[string, *doctor](cmpStrings)

	for scanner.Scan() {
		partes := strings.SplitN(scanner.Text(), ",", 2)
		if len(partes) != 2 {
			continue
		}

		nombre := partes[0]
		especialidad := strings.TrimSpace(partes[1])

		doctores.Guardar(nombre, &doctor{
			nombre:       nombre,
			especialidad: especialidad,
			atendidos:    0,
		})
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
func construirEspecialidades(doctores diccionario.DiccionarioOrdenado[string, *doctor]) map[string]*colaEspecialidad {
	especialidades := make(map[string]*colaEspecialidad)

	doctores.Iterar(func(nombre string, doctor *doctor) bool {
		if _, existe := especialidades[doctor.especialidad]; !existe {
			especialidades[doctor.especialidad] = nuevaColaEspecialidad()
		}
		return true
	})

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
			fmt.Printf(ENOENT_FORMATO, linea)
			continue
		}

		comando := partes[0]
		args := strings.Split(partes[1], ",")

		switch comando {
		case "PEDIR_TURNO":
			pedirTurno(args, pacientes, especialidades)
		case "ATENDER_SIGUIENTE":
			atenderSiguiente(args, doctores, especialidades)
		case "INFORME":
			informe(args, doctores)
		default:
			fmt.Printf(ENOENT_CMD, comando)
		}
	}
}

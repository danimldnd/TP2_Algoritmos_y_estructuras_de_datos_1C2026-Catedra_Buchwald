package main

import (
	"fmt"
)

func pedirTurno(args []string, pacientes map[string]int, especialidades map[string]*colaEspecialidad) {
	if len(args) != 3 {
		fmt.Printf("ERROR: cantidad de parametros invalidos para comando 'PEDIR_TURNO'\n")
		return
	}

	nombre := args[0]
	especialidad := args[1]
	urgencia := args[2]

	hayError := false

	if _, existe := pacientes[nombre]; !existe {
		fmt.Printf("ERROR: no existe el/la paciente '%s'\n", nombre)
		hayError = true
	}
	if _, existe := especialidades[especialidad]; !existe {
		fmt.Printf("ERROR: no existe la especialidad '%s'\n", especialidad)
		hayError = true
	}
	if urgencia != "URGENTE" && urgencia != "REGULAR" {
		fmt.Printf("ERROR: grado de urgencia no identificado ('%s')\n", urgencia)
		hayError = true
	}

	if hayError {
		return
	}

	año := pacientes[nombre]
	urgente := urgencia == "URGENTE"
	especialidades[especialidad].encolar(nombre, año, urgente)

	fmt.Printf("Paciente %s encolado\n", nombre)
	fmt.Printf("%d paciente(s) en espera para %s\n", especialidades[especialidad].cantidad, especialidad)
}

func atenderSiguiente(args []string, doctores map[string]string) {
	//
}

func informe(args []string, doctores map[string]string) {
	//
}

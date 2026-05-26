package main

import (
	"fmt"
	diccionario "tdas/diccionarios"
)

type doctor struct {
	nombre       string
	especialidad string
	atendidos    int
}

func pedirTurno(args []string, pacientes map[string]int, especialidades map[string]*colaEspecialidad) {
	if len(args) != 3 {
		fmt.Printf(ENOENT_PARAMS, "PEDRI_TURNO")
		return
	}

	nombre := args[0]
	especialidad := args[1]
	urgencia := args[2]

	hayError := false

	if _, existe := pacientes[nombre]; !existe {
		fmt.Printf(ENOENT_PACIENTE, nombre)
		hayError = true
	}
	if _, existe := especialidades[especialidad]; !existe {
		fmt.Printf(ENOENT_ESPECIALIDAD, especialidad)
		hayError = true
	}
	if urgencia != "URGENTE" && urgencia != "REGULAR" {
		fmt.Printf(ENOENT_URGENCIA, urgencia)
		hayError = true
	}

	if hayError {
		return
	}

	año := pacientes[nombre]
	urgente := urgencia == "URGENTE"
	especialidades[especialidad].encolar(nombre, año, urgente)

	fmt.Printf(PACIENTE_ENCOLADO, nombre)
	fmt.Printf(CANT_PACIENTES_ENCOLADOS, especialidades[especialidad].cantidad, especialidad)
}

func atenderSiguiente(args []string, doctores diccionario.DiccionarioOrdenado[string, *doctor], especialidades map[string]*colaEspecialidad) {
	if len(args) != 1 {
		fmt.Printf(ENOENT_PARAMS, "ATENDER_SIGUIENTE")
		return
	}

	nombreDoctor := args[0]

	if !doctores.Pertenece(nombreDoctor) {
		fmt.Printf(ENOENT_DOCTOR, nombreDoctor)
		return
	}

	doctor := doctores.Obtener(nombreDoctor)
	especialidad := doctor.especialidad

	cola := especialidades[especialidad]

	if cola.estaVacia() {
		fmt.Printf(SIN_PACIENTES, especialidad)
		return
	}

	paciente := cola.desencolar()
	doctor.atendidos++

	fmt.Printf(PACIENTE_ATENDIDO, paciente)
	fmt.Printf(CANT_PACIENTES_ENCOLADOS, cola.cantidad, especialidad)
}

func informe(args []string, doctores diccionario.DiccionarioOrdenado[string, *doctor]) {
	if len(args) != 2 {
		fmt.Printf(ENOENT_PARAMS, "INFORME")
		return
	}

	inicio := args[0]
	fin := args[1]

	var desde *string = nil
	var hasta *string = nil

	if inicio != "" {
		desde = &inicio
	}
	if fin != "" {
		hasta = &fin
	}

	cantidad := 0

	doctores.IterarRango(desde, hasta, func(nombre string, doctor *doctor) bool {
		cantidad++
		return true
	})

	fmt.Printf("%d doctor(es) en el sistema\n", cantidad)

	i := 1
	doctores.IterarRango(desde, hasta, func(nombre string, doctor *doctor) bool {
		fmt.Printf(
			INFORME_DOCTOR,
			i,
			doctor.nombre,
			doctor.especialidad,
			doctor.atendidos,
		)
		i++
		return true
	})
}

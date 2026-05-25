package main

import (
	cola "tdas/cola"
	cola_prioridad "tdas/cola_prioridad"
)

// pacienteRegular representa un paciente en la cola de regulares
type pacienteRegular struct {
	nombre string
	año    int
	orden  int
}

// colaEspecialidad maneja la lista de espera de una especialidad
type colaEspecialidad struct {
	urgentes  cola.Cola[string]
	regulares cola_prioridad.ColaPrioridad[pacienteRegular]
	cantidad  int
	contador  int // para el orden de llegada
}

func cmpPacienteRegular(a, b pacienteRegular) int {
	if a.año != b.año {
		return b.año - a.año // menor año = mayor prioridad
	}
	return b.orden - a.orden // menor orden = llegó antes = mayor prioridad
}

func nuevaColaEspecialidad() *colaEspecialidad {
	return &colaEspecialidad{
		urgentes:  cola.CrearColaEnlazada[string](),
		regulares: cola_prioridad.CrearHeap[pacienteRegular](cmpPacienteRegular),
	}
}

func (cola *colaEspecialidad) encolar(nombre string, año int, urgente bool) {
	cola.contador++
	cola.cantidad++
	if urgente {
		cola.urgentes.Encolar(nombre)
	} else {
		cola.regulares.Encolar(pacienteRegular{nombre, año, cola.contador})
	}
}

func (cola *colaEspecialidad) desencolar() string {
	cola.cantidad--
	if !cola.urgentes.EstaVacia() {
		return cola.urgentes.Desencolar()
	}
	return cola.regulares.Desencolar().nombre
}

func (cola *colaEspecialidad) estaVacia() bool {
	return cola.cantidad == 0
}

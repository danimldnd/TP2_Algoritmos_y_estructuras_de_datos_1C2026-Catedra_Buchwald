package cola_test

import (
	"tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()

	require.True(t, c.EstaVacia())
	require.Panics(t, func() { c.VerPrimero() })
	require.Panics(t, func() { c.Desencolar() })
}

func TestEncolarYDesencolarFIFO(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()

	c.Encolar(1)
	c.Encolar(2)
	c.Encolar(3)

	require.Equal(t, 1, c.Desencolar())
	require.Equal(t, 2, c.Desencolar())
	require.Equal(t, 3, c.Desencolar())
}

func TestVolumen(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()

	n := 10000

	for i := 0; i < n; i++ {
		c.Encolar(i)
	}

	for i := 0; i < n; i++ {
		require.Equal(t, i, c.Desencolar())
	}

	require.True(t, c.EstaVacia())
}

func TestVerPrimero(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()

	c.Encolar(10)
	require.Equal(t, 10, c.VerPrimero())

	c.Encolar(20)
	require.Equal(t, 10, c.VerPrimero())
}

func TestCondicionesBorde(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()

	// cola recién creada
	require.True(t, c.EstaVacia())
	require.Panics(t, func() { c.VerPrimero() })
	require.Panics(t, func() { c.Desencolar() })

	// encolo y desencolo
	c.Encolar(1)
	require.Equal(t, 1, c.Desencolar())

	// vuelve a estar vacía
	require.True(t, c.EstaVacia())
	require.Panics(t, func() { c.VerPrimero() })
	require.Panics(t, func() { c.Desencolar() })
}

func TestIntercalado(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()

	c.Encolar(1)
	c.Encolar(2)
	require.Equal(t, 1, c.Desencolar())

	c.Encolar(3)
	require.Equal(t, 2, c.Desencolar())
	require.Equal(t, 3, c.Desencolar())

	require.True(t, c.EstaVacia())
}

func TestDistintosTipos(t *testing.T) {
	cadenas := cola.CrearColaEnlazada[string]()

	cadenas.Encolar("hola")
	cadenas.Encolar("chau")

	require.Equal(t, "hola", cadenas.Desencolar())
	require.Equal(t, "chau", cadenas.Desencolar())
}

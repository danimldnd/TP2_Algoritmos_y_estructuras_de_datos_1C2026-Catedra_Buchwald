package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	require.True(t, pila.EstaVacia())

	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.VerTope()
	})

	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.Desapilar()
	})
}

func TestApilarYDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)

	require.Equal(t, 3, pila.VerTope())

	require.Equal(t, 3, pila.Desapilar())
	require.Equal(t, 2, pila.VerTope())

	require.Equal(t, 2, pila.Desapilar())
	require.Equal(t, 1, pila.VerTope())

	require.Equal(t, 1, pila.Desapilar())

	require.True(t, pila.EstaVacia())
}

func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	n := 10000

	for i := 0; i < n; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())
	}

	for i := n - 1; i >= 0; i-- {
		require.Equal(t, i, pila.Desapilar())
	}

	require.True(t, pila.EstaVacia())
}

func TestPilaStrings(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()

	pila.Apilar("hola")
	pila.Apilar("chau")

	require.Equal(t, "chau", pila.VerTope())
	require.Equal(t, "chau", pila.Desapilar())
	require.Equal(t, "hola", pila.Desapilar())

	require.True(t, pila.EstaVacia())
}

func TestPilaVaciaDespuesDeUsar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	pila.Apilar(10)
	pila.Desapilar()

	require.True(t, pila.EstaVacia())

	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.VerTope()
	})

	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.Desapilar()
	})
}

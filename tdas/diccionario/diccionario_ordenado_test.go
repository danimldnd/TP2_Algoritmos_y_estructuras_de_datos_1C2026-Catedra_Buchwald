package diccionario_test

import (
	"fmt"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbdVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	assert.EqualValues(t, 0, abb.Cantidad())
	assert.False(t, abb.Pertenece("A"))
	assert.Panics(t, func() { abb.Obtener("A") })
	assert.Panics(t, func() { abb.Borrar("A") })
}

func TestAbdBorrados(t *testing.T) {
	t.Log("Prueba casos de borrado: hoja, un hijo, dos hijos")
	abb := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })

	claves := []int{10, 5, 15, 2, 7, 20}
	for _, k := range claves {
		abb.Guardar(k, fmt.Sprintf("valor %d", k))
	}

	assert.EqualValues(t, 6, abb.Cantidad())

	assert.EqualValues(t, "valor 2", abb.Borrar(2))
	assert.False(t, abb.Pertenece(2))
	assert.EqualValues(t, 5, abb.Cantidad())

	assert.EqualValues(t, "valor 15", abb.Borrar(15))
	assert.True(t, abb.Pertenece(20))
	assert.EqualValues(t, 4, abb.Cantidad())

	assert.EqualValues(t, "valor 10", abb.Borrar(10))
	assert.EqualValues(t, 3, abb.Cantidad())
	assert.True(t, abb.Pertenece(20))
	assert.True(t, abb.Pertenece(5))
	assert.True(t, abb.Pertenece(7))
}

func TestIteradorRango(t *testing.T) {
	t.Log("Prueba el iterador externo con rangos específicos")
	abb := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	for i := 1; i <= 10; i++ {
		abb.Guardar(i, i)
	}

	desde, hasta := 3, 7
	iter := abb.IteradorRango(&desde, &hasta)

	actual := desde
	for iter.HayAlgoMas() {
		k, _ := iter.VerActual()
		assert.EqualValues(t, actual, k)
		iter.Avanzar()
		actual++
	}
	assert.EqualValues(t, 8, actual)
}

func TestIteradorRangoNil(t *testing.T) {
	t.Log("Prueba rangos con extremos nil (infinitos)")
	abb := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	claves := []int{5, 2, 8, 1, 3, 7, 9}
	for _, k := range claves {
		abb.Guardar(k, k)
	}

	hasta := 3
	iter := abb.IteradorRango(nil, &hasta)
	esperados := []int{1, 2, 3}
	i := 0
	for iter.HayAlgoMas() {
		k, _ := iter.VerActual()
		assert.EqualValues(t, esperados[i], k)
		iter.Avanzar()
		i++
	}
	assert.EqualValues(t, len(esperados), i)
}

func TestIteradorInternoRango(t *testing.T) {
	t.Log("Prueba la función visitar con rangos")
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)

	abb.Guardar("B", 2)
	abb.Guardar("A", 1)
	abb.Guardar("D", 4)
	abb.Guardar("C", 3)
	abb.Guardar("E", 5)

	desde, hasta := "B", "D"
	contador := 0
	abb.IterarRango(&desde, &hasta, func(k string, v int) bool {
		contador++
		if k == "C" {
			return false
		}
		return true
	})

	assert.EqualValues(t, 2, contador)
}

func TestRangoInvertido(t *testing.T) {
	t.Log("Prueba que si desde > hasta, no se itera nada")
	abb := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	for i := 0; i < 10; i++ {
		abb.Guardar(i, i)
	}

	desde, hasta := 8, 2
	iter := abb.IteradorRango(&desde, &hasta)
	assert.False(t, iter.HayAlgoMas())

	contador := 0
	abb.IterarRango(&desde, &hasta, func(k int, v int) bool {
		contador++
		return true
	})
	assert.EqualValues(t, 0, contador)
}

func TestGuardarClaveExistente(t *testing.T) {
	t.Log("Prueba que guardar una clave existente actualiza el dato y no cambia la cantidad")
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)

	abb.Guardar("Gato", "Miau")
	assert.EqualValues(t, 1, abb.Cantidad())

	abb.Guardar("Gato", "Miau Reali")
	assert.EqualValues(t, 1, abb.Cantidad())
	assert.EqualValues(t, "Miau Reali", abb.Obtener("Gato"))
}

func TestIteradorCompletoNil(t *testing.T) {
	t.Log("Prueba que IteradorRango(nil, nil) es igual a Iterador()")
	abb := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	for i := 0; i < 5; i++ {
		abb.Guardar(i, i)
	}

	iterRango := abb.IteradorRango(nil, nil)
	iterNormal := abb.Iterador()

	for iterRango.HayAlgoMas() {
		k1, _ := iterRango.VerActual()
		k2, _ := iterNormal.VerActual()
		assert.EqualValues(t, k1, k2)
		iterRango.Avanzar()
		iterNormal.Avanzar()
	}
	assert.False(t, iterNormal.HayAlgoMas())
}

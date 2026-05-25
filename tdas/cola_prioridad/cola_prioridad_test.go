package cola_prioridad_test

import (
	cola_prioridad "tdas/cola_prioridad"

	"testing"

	"github.com/stretchr/testify/require"
)

const VOLUMEN_PRUEBA = 999999

func cmpInt(a, b int) int {
	if a > b {
		return 1
	}

	if a < b {
		return -1
	}
	return 0
}

func cmpFloat(a, b float64) int {
	if a > b {
		return 1
	}

	if a < b {
		return -1
	}
	return 0
}

func cmpStrings(a, b string) int { //orden alfabetico
	if a < b {
		return 1
	}

	if a > b {
		return -1
	}
	return 0
}

func crearHeapParaTestear() (cola_prioridad.ColaPrioridad[int], cola_prioridad.ColaPrioridad[string], cola_prioridad.ColaPrioridad[float64]) {
	heap_enteros := cola_prioridad.CrearHeap[int](cmpInt)
	heap_strings := cola_prioridad.CrearHeap[string](cmpStrings)
	heap_float := cola_prioridad.CrearHeap[float64](cmpFloat)

	return heap_enteros, heap_strings, heap_float
}
func TestHeapVacia(t *testing.T) {
	heap_enteros, heap_strings, heap_float := crearHeapParaTestear()

	require.True(t, heap_enteros.EstaVacia())
	require.True(t, heap_strings.EstaVacia())
	require.True(t, heap_float.EstaVacia())
}

func TestHeapEncolar(t *testing.T) {
	heap_enteros, heap_strings, heap_float := crearHeapParaTestear()

	heap_enteros.Encolar(900)
	require.Equal(t, 900, heap_enteros.VerMax())

	heap_strings.Encolar("frieren")
	require.Equal(t, "frieren", heap_strings.VerMax())

	heap_float.Encolar(1.23)
	require.Equal(t, 1.23, heap_float.VerMax())
}

func TestVerMax(t *testing.T) {
	heap_enteros, heap_strings, heap_float := crearHeapParaTestear()

	heap_enteros.Encolar(2)
	heap_enteros.Encolar(1)
	require.Equal(t, 2, heap_enteros.VerMax())

	heap_strings.Encolar("lol")
	heap_strings.Encolar("xd")
	require.Equal(t, "lol", heap_strings.VerMax())

	heap_float.Encolar(1.2)
	heap_float.Encolar(2.2)
	require.Equal(t, 2.2, heap_float.VerMax())

}

func TestVerMaxVacio(t *testing.T) {
	heap_enteros, heap_strings, heap_float := crearHeapParaTestear()

	require.True(t, heap_enteros.EstaVacia())
	require.Panics(t, func() { heap_enteros.VerMax() })

	require.True(t, heap_strings.EstaVacia())
	require.Panics(t, func() { heap_strings.VerMax() })

	require.True(t, heap_float.EstaVacia())
	require.Panics(t, func() { heap_float.VerMax() })

}

func TestHeapCantidad(t *testing.T) {
	heap_enteros, heap_strings, heap_float := crearHeapParaTestear()

	heap_enteros.Encolar(100)
	heap_enteros.Encolar(200)
	require.Equal(t, 2, heap_enteros.Cantidad())

	heap_strings.Encolar("dani")
	heap_strings.Encolar("moche")
	require.Equal(t, 2, heap_strings.Cantidad())

	heap_float.Encolar(1.12)
	heap_float.Encolar(2.32)
	require.Equal(t, 2, heap_float.Cantidad())
}

func TestDesencolar(t *testing.T) {
	heap_enteros, heap_strings, heap_float := crearHeapParaTestear()

	heap_enteros.Encolar(1000)
	heap_enteros.Encolar(200)
	require.Equal(t, 1000, heap_enteros.Desencolar())

	heap_strings.Encolar("dani")
	heap_strings.Encolar("moche")
	require.Equal(t, "dani", heap_strings.Desencolar())

	heap_float.Encolar(1.12)
	heap_float.Encolar(2.32)
	require.Equal(t, 2.32, heap_float.Desencolar())

}

func TestDesencolarVacio(t *testing.T) {
	heap_enteros, heap_strings, heap_float := crearHeapParaTestear()

	require.True(t, heap_enteros.EstaVacia())
	require.Panics(t, func() { heap_enteros.Desencolar() })

	require.True(t, heap_strings.EstaVacia())
	require.Panics(t, func() { heap_strings.Desencolar() })

	require.True(t, heap_float.EstaVacia())
	require.Panics(t, func() { heap_float.Desencolar() })
}

func TestDesencolarMantienePrioridad(t *testing.T) {
	heap_enteros, heap_strings, heap_float := crearHeapParaTestear()

	arreglo_enteros := []int{9, 22, 312, 1, 33, 4}
	for _, i := range arreglo_enteros {
		heap_enteros.Encolar(i)
	}
	require.Equal(t, 312, heap_enteros.Desencolar())
	require.Equal(t, 33, heap_enteros.Desencolar())
	require.Equal(t, 22, heap_enteros.Desencolar())
	require.Equal(t, 9, heap_enteros.Desencolar())
	require.Equal(t, 4, heap_enteros.Desencolar())
	require.Equal(t, 1, heap_enteros.Desencolar())

	arreglo_strings := []string{"golem", "pikachu", "zapdos", "rapidash", "sceptile", "magikarp"}
	for _, i := range arreglo_strings {
		heap_strings.Encolar(i)
	}
	require.Equal(t, "golem", heap_strings.Desencolar())
	require.Equal(t, "magikarp", heap_strings.Desencolar())
	require.Equal(t, "pikachu", heap_strings.Desencolar())
	require.Equal(t, "rapidash", heap_strings.Desencolar())
	require.Equal(t, "sceptile", heap_strings.Desencolar())
	require.Equal(t, "zapdos", heap_strings.Desencolar())

	arreglo_float := []float64{9.2, 22.2, 312.2, 1.2, 33.2, 4.2}
	for _, i := range arreglo_float {
		heap_float.Encolar(i)
	}
	require.Equal(t, 312.2, heap_float.Desencolar())
	require.Equal(t, 33.2, heap_float.Desencolar())
	require.Equal(t, 22.2, heap_float.Desencolar())
	require.Equal(t, 9.2, heap_float.Desencolar())
	require.Equal(t, 4.2, heap_float.Desencolar())
	require.Equal(t, 1.2, heap_float.Desencolar())

}

func TestDesencolarMismaPrioridad(t *testing.T) {
	heap_enteros, heap_strings, heap_float := crearHeapParaTestear()

	heap_enteros.Encolar(10)
	heap_enteros.Encolar(10)
	heap_enteros.Encolar(5)
	require.Equal(t, 10, heap_enteros.Desencolar())
	require.Equal(t, 10, heap_enteros.Desencolar())
	require.Equal(t, 5, heap_enteros.Desencolar())

	heap_strings.Encolar("gojo")
	heap_strings.Encolar("gojo")
	heap_strings.Encolar("sukuna")
	require.Equal(t, "gojo", heap_strings.Desencolar())
	require.Equal(t, "gojo", heap_strings.Desencolar())
	require.Equal(t, "sukuna", heap_strings.Desencolar())

	heap_float.Encolar(10.2)
	heap_float.Encolar(10.2)
	heap_float.Encolar(500.3)
	require.Equal(t, 500.3, heap_float.Desencolar())
	require.Equal(t, 10.2, heap_float.Desencolar())
	require.Equal(t, 10.2, heap_float.Desencolar())

}

func TestVolumen(t *testing.T) {
	heap_enteros := cola_prioridad.CrearHeap[int](cmpInt)

	for i := 0; i < VOLUMEN_PRUEBA; i++ {
		heap_enteros.Encolar(i)
		require.Equal(t, i, heap_enteros.VerMax())
	}

	for i := 0; i < VOLUMEN_PRUEBA; i++ {
		require.Equal(t, VOLUMEN_PRUEBA-i-1, heap_enteros.Desencolar())
	}
}

func TestUnHeapRecienCreadoSeComportaIgualAUnoVacio(t *testing.T) {
	heap_enteros_nuevo, heap_strings_nuevo, heap_float_nuevo := crearHeapParaTestear()
	heap_enteros_vaciado, heap_strings_vaciado, heap_float_vaciado := crearHeapParaTestear()

	heap_enteros_vaciado.Encolar(10)
	require.Equal(t, 10, heap_enteros_vaciado.VerMax())
	heap_enteros_vaciado.Desencolar()
	require.True(t, heap_enteros_vaciado.EstaVacia())
	require.True(t, heap_enteros_nuevo.EstaVacia())
	require.Panics(t, func() { heap_enteros_vaciado.Desencolar() })
	require.Panics(t, func() { heap_enteros_vaciado.VerMax() })
	require.Panics(t, func() { heap_enteros_nuevo.Desencolar() })
	require.Panics(t, func() { heap_enteros_nuevo.VerMax() })

	heap_strings_vaciado.Encolar("jojo")
	require.Equal(t, "jojo", heap_strings_vaciado.VerMax())
	heap_strings_vaciado.Desencolar()
	require.True(t, heap_strings_vaciado.EstaVacia())
	require.True(t, heap_strings_nuevo.EstaVacia())
	require.Panics(t, func() { heap_strings_vaciado.Desencolar() })
	require.Panics(t, func() { heap_strings_vaciado.VerMax() })
	require.Panics(t, func() { heap_strings_nuevo.Desencolar() })
	require.Panics(t, func() { heap_strings_nuevo.VerMax() })

	heap_float_vaciado.Encolar(10.23)
	require.Equal(t, 10.23, heap_float_vaciado.VerMax())
	heap_float_vaciado.Desencolar()
	require.True(t, heap_float_vaciado.EstaVacia())
	require.True(t, heap_float_nuevo.EstaVacia())
	require.Panics(t, func() { heap_float_vaciado.Desencolar() })
	require.Panics(t, func() { heap_float_vaciado.VerMax() })
	require.Panics(t, func() { heap_float_nuevo.Desencolar() })
	require.Panics(t, func() { heap_float_nuevo.VerMax() })
}

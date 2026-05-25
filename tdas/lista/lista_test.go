package lista_test

import (
	"tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const _VOLUMEN_TEST = 100000

func crearListasParaTestear() (lista.Lista[int], lista.Lista[string], lista.Lista[float64]) {
	lista_enteros := lista.CrearListaEnlazada[int]()
	lista_strings := lista.CrearListaEnlazada[string]()
	lista_float64 := lista.CrearListaEnlazada[float64]()

	return lista_enteros, lista_strings, lista_float64
}

func TestListaVacia(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	require.True(t, lista_enteros.EstaVacia())
	require.Equal(t, 0, lista_enteros.Largo())

	require.Panics(t, func() { lista_enteros.VerPrimero() })
	require.Panics(t, func() { lista_enteros.VerUltimo() })
	require.Panics(t, func() { lista_enteros.BorrarPrimero() })

	it := lista_enteros.Iterador()
	require.False(t, it.HayAlgoMas())

	require.True(t, lista_strings.EstaVacia())
	require.True(t, lista_float64.EstaVacia())
}
func TestInsertarPrimero(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarPrimero(10)
	require.False(t, lista_enteros.EstaVacia())

	lista_strings.InsertarPrimero("martu")
	require.False(t, lista_strings.EstaVacia())

	lista_float64.InsertarPrimero(10.2)
	require.False(t, lista_float64.EstaVacia())
}

func TestInsertarUltimo(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarUltimo(10)
	require.False(t, lista_enteros.EstaVacia())

	lista_strings.InsertarUltimo("martu")
	require.False(t, lista_strings.EstaVacia())

	lista_float64.InsertarUltimo(10.2)
	require.False(t, lista_float64.EstaVacia())
}

func TestVerPrimero(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarPrimero(10)
	lista_enteros.InsertarUltimo(110)
	require.Equal(t, 10, lista_enteros.VerPrimero())

	lista_strings.InsertarPrimero("martu")
	lista_strings.InsertarUltimo("paloma")
	require.Equal(t, "martu", lista_strings.VerPrimero())

	lista_float64.InsertarPrimero(1.23)
	lista_float64.InsertarUltimo(1.55)
	require.Equal(t, 1.23, lista_float64.VerPrimero())

}

func TestVerUltimo(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarUltimo(10)
	lista_enteros.InsertarPrimero(110)
	require.Equal(t, 10, lista_enteros.VerUltimo())

	lista_strings.InsertarUltimo("martu")
	lista_strings.InsertarPrimero("paloma")
	require.Equal(t, "martu", lista_strings.VerUltimo())

	lista_float64.InsertarUltimo(1.23)
	lista_float64.InsertarPrimero(1.55)
	require.Equal(t, 1.23, lista_float64.VerUltimo())
}

func TestVerPrimeroEnUnaListaVacia(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarPrimero(10)
	lista_enteros.BorrarPrimero()
	require.True(t, lista_enteros.EstaVacia())
	require.Panics(t, func() { lista_enteros.VerPrimero() })

	lista_strings.InsertarPrimero("moche")
	lista_strings.BorrarPrimero()
	require.True(t, lista_strings.EstaVacia())
	require.Panics(t, func() { lista_strings.VerPrimero() })

	lista_float64.InsertarPrimero(10.2)
	lista_float64.BorrarPrimero()
	require.True(t, lista_float64.EstaVacia())
	require.Panics(t, func() { lista_float64.VerPrimero() })
}

func TestVerUltimoEnUnaListaVacia(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarPrimero(10)
	lista_enteros.BorrarPrimero()
	require.True(t, lista_enteros.EstaVacia())
	require.Panics(t, func() { lista_enteros.VerUltimo() })

	lista_strings.InsertarPrimero("moche")
	lista_strings.BorrarPrimero()
	require.True(t, lista_strings.EstaVacia())
	require.Panics(t, func() { lista_strings.VerUltimo() })

	lista_float64.InsertarPrimero(10.2)
	lista_float64.BorrarPrimero()
	require.True(t, lista_float64.EstaVacia())
	require.Panics(t, func() { lista_float64.VerUltimo() })
}

func TestBorrarPrimero(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarPrimero(10)
	lista_enteros.BorrarPrimero()

	require.True(t, lista_enteros.EstaVacia())

	lista_strings.InsertarPrimero("martu")
	lista_strings.BorrarPrimero()
	require.True(t, lista_strings.EstaVacia())

	lista_float64.InsertarPrimero(10.2)
	lista_float64.BorrarPrimero()
	require.True(t, lista_float64.EstaVacia())
}

func TestBorrarPrimeroDevuelveElPrimerElemento(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarPrimero(10)
	lista_enteros.InsertarUltimo(15)
	require.Equal(t, 10, lista_enteros.BorrarPrimero())

	lista_strings.InsertarPrimero("martu")
	lista_strings.InsertarPrimero("dani")
	require.Equal(t, "dani", lista_strings.BorrarPrimero())

	lista_float64.InsertarPrimero(10.2)
	lista_float64.InsertarPrimero(44.3)
	require.Equal(t, 44.3, lista_float64.BorrarPrimero())

}

func TestBorrarPrimeroEnUnaListaVacia(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarPrimero(10)
	lista_enteros.BorrarPrimero()
	require.True(t, lista_enteros.EstaVacia())
	require.Panics(t, func() { lista_enteros.BorrarPrimero() })

	lista_strings.InsertarPrimero("tengosueño")
	lista_strings.BorrarPrimero()
	require.True(t, lista_strings.EstaVacia())
	require.Panics(t, func() { lista_strings.BorrarPrimero() })

	lista_float64.InsertarPrimero(10.2)
	lista_float64.BorrarPrimero()
	require.True(t, lista_float64.EstaVacia())
	require.Panics(t, func() { lista_float64.BorrarPrimero() })
}

func TestBorrarPrimeroConvierteAlSegundoEnPrimero(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarPrimero(10)
	lista_enteros.InsertarUltimo(100)
	lista_enteros.InsertarUltimo(10000)
	lista_enteros.BorrarPrimero()
	require.Equal(t, 100, lista_enteros.VerPrimero())

	lista_strings.InsertarPrimero("coscu")
	lista_strings.InsertarUltimo("nachoide")
	lista_strings.InsertarUltimo("nanoide")
	lista_strings.BorrarPrimero()
	require.Equal(t, "nachoide", lista_strings.VerPrimero())

	lista_float64.InsertarPrimero(1.23)
	lista_float64.InsertarUltimo(1.2343)
	lista_float64.InsertarUltimo(2.34)
	lista_float64.BorrarPrimero()
	require.Equal(t, 1.2343, lista_float64.VerPrimero())
}

func TestLargoIncremento(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarPrimero(10)
	require.Equal(t, 1, lista_enteros.Largo())
	lista_enteros.InsertarUltimo(100)
	require.Equal(t, 2, lista_enteros.Largo())
	lista_enteros.InsertarUltimo(10000)
	require.Equal(t, 3, lista_enteros.Largo())

	lista_strings.InsertarPrimero("coscu")
	require.Equal(t, 1, lista_strings.Largo())
	lista_strings.InsertarUltimo("nachoide")
	require.Equal(t, 2, lista_strings.Largo())
	lista_strings.InsertarUltimo("nanoide")
	require.Equal(t, 3, lista_strings.Largo())

	lista_float64.InsertarPrimero(1.23)
	require.Equal(t, 1, lista_float64.Largo())
	lista_float64.InsertarUltimo(1.2343)
	require.Equal(t, 2, lista_float64.Largo())
	lista_float64.InsertarUltimo(2.34)
	require.Equal(t, 3, lista_float64.Largo())
}

func TestLargoDecrecimiento(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarPrimero(40)
	lista_enteros.InsertarPrimero(67)
	lista_enteros.InsertarPrimero(5)
	require.Equal(t, 3, lista_enteros.Largo())
	lista_enteros.BorrarPrimero()
	require.Equal(t, 2, lista_enteros.Largo())
	lista_enteros.BorrarPrimero()
	require.Equal(t, 1, lista_enteros.Largo())
	lista_enteros.BorrarPrimero()
	require.Equal(t, 0, lista_enteros.Largo())

	lista_strings.InsertarPrimero("gradiente")
	lista_strings.InsertarPrimero("mli")
	lista_strings.InsertarPrimero("proyecto")
	require.Equal(t, 3, lista_strings.Largo())
	lista_strings.BorrarPrimero()
	require.Equal(t, 2, lista_strings.Largo())
	lista_strings.BorrarPrimero()
	require.Equal(t, 1, lista_strings.Largo())
	lista_strings.BorrarPrimero()
	require.Equal(t, 0, lista_strings.Largo())

	lista_float64.InsertarPrimero(40.2)
	lista_float64.InsertarPrimero(67.67)
	lista_float64.InsertarPrimero(5.3)
	require.Equal(t, 3, lista_float64.Largo())
	lista_float64.BorrarPrimero()
	require.Equal(t, 2, lista_float64.Largo())
	lista_float64.BorrarPrimero()
	require.Equal(t, 1, lista_float64.Largo())
	lista_float64.BorrarPrimero()
	require.Equal(t, 0, lista_float64.Largo())

}

func TestPrimeroYUltimoSonElMismoElementoEnUnaListaDeLargoUno(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarPrimero(10)
	require.Equal(t, lista_enteros.VerPrimero(), lista_enteros.VerUltimo())

	lista_strings.InsertarPrimero("jordiwild")
	require.Equal(t, lista_strings.VerPrimero(), lista_strings.VerUltimo())

	lista_float64.InsertarPrimero(10.89)
	require.Equal(t, lista_float64.VerPrimero(), lista_float64.VerUltimo())

}

func TestExtremosVerPrimeroYVerUltimo(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarPrimero(10)
	lista_enteros.InsertarPrimero(15)
	lista_enteros.InsertarUltimo(55)
	require.Equal(t, 15, lista_enteros.VerPrimero())
	require.Equal(t, 55, lista_enteros.VerUltimo())

	lista_strings.InsertarPrimero("+")
	lista_strings.InsertarPrimero("{}")
	lista_strings.InsertarUltimo("*")
	require.Equal(t, "{}", lista_strings.VerPrimero())
	require.Equal(t, "*", lista_strings.VerUltimo())

	lista_float64.InsertarPrimero(10.44)
	lista_float64.InsertarPrimero(15.44)
	lista_float64.InsertarUltimo(55.44)
	require.Equal(t, 15.44, lista_float64.VerPrimero())
	require.Equal(t, 55.44, lista_float64.VerUltimo())

}

func TestVolumen(t *testing.T) {
	lista_enteros := lista.CrearListaEnlazada[int]()

	for i := 0; i < _VOLUMEN_TEST; i++ {
		lista_enteros.InsertarUltimo(i)
		require.Equal(t, i+1, lista_enteros.Largo())
		require.Equal(t, i, lista_enteros.VerUltimo())
	}

	require.Equal(t, _VOLUMEN_TEST, lista_enteros.Largo())

	for i := 0; i < _VOLUMEN_TEST; i++ {
		valor := lista_enteros.BorrarPrimero()
		require.Equal(t, i, valor)
	}

	require.True(t, lista_enteros.EstaVacia())
	require.Equal(t, 0, lista_enteros.Largo())
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

func TestIteradorInsertarAlInicio(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(2)
	l.InsertarUltimo(3)

	it := l.Iterador()
	it.Insertar(1)

	require.Equal(t, 1, l.VerPrimero())
	require.Equal(t, 3, l.VerUltimo())
	require.Equal(t, 3, l.Largo())
}

func TestIteradorInsertarEnMedio(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	l.InsertarUltimo(3)

	it := l.Iterador()
	it.Avanzar()
	it.Insertar(2)

	require.Equal(t, 3, l.Largo())

	require.Equal(t, 1, l.BorrarPrimero())
	require.Equal(t, 2, l.BorrarPrimero())
	require.Equal(t, 3, l.BorrarPrimero())
}

func TestIteradorInsertarAlFinal(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	l.InsertarUltimo(2)

	it := l.Iterador()
	for it.HayAlgoMas() {
		it.Avanzar()
	}

	it.Insertar(3)

	require.Equal(t, 3, l.VerUltimo())
	require.Equal(t, 3, l.Largo())
}

func TestIteradorBorrarPrimero(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	l.InsertarUltimo(2)

	it := l.Iterador()
	borrado := it.Borrar()

	require.Equal(t, 1, borrado)
	require.Equal(t, 2, l.VerPrimero())
}

func TestIteradorBorrarMedio(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	l.InsertarUltimo(2)
	l.InsertarUltimo(3)

	it := l.Iterador()
	it.Avanzar() // apunta a 2

	borrado := it.Borrar()

	require.Equal(t, 2, borrado)

	require.Equal(t, 1, l.BorrarPrimero())
	require.Equal(t, 3, l.BorrarPrimero())
}

func TestIteradorBorrarUltimo(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	l.InsertarUltimo(2)

	it := l.Iterador()
	it.Avanzar() // 2

	borrado := it.Borrar()

	require.Equal(t, 2, borrado)
	require.Equal(t, 1, l.VerUltimo())
}

func TestIteradorPanics(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)

	it := l.Iterador()
	it.Avanzar() // ahora está al final

	require.False(t, it.HayAlgoMas())

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		it.VerActual()
	})

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		it.Avanzar()
	})

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		it.Borrar()
	})
}

func TestIteradorListaVacia(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	it := l.Iterador()

	require.False(t, it.HayAlgoMas())

	it.Insertar(10)

	require.Equal(t, 10, l.VerPrimero())
	require.Equal(t, 10, l.VerUltimo())
}

func TestIteradorInternoCompleto(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	for i := 1; i <= 5; i++ {
		l.InsertarUltimo(i)
	}

	suma := 0
	l.Iterar(func(v int) bool {
		suma += v
		return true
	})

	require.Equal(t, 15, suma)
}

func TestIteradorInternoCorte(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	for i := 1; i <= 5; i++ {
		l.InsertarUltimo(i)
	}

	contador := 0
	l.Iterar(func(v int) bool {
		contador++
		return contador < 3
	})

	require.Equal(t, 3, contador)
}

func TestIteradorInsertarEnListaVacia(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	it := l.Iterador()

	it.Insertar(10)

	require.Equal(t, 10, l.VerPrimero())
	require.Equal(t, 10, l.VerUltimo())
	require.Equal(t, 1, l.Largo())
}

func TestIteradorBorrarUnicoElemento(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)

	it := l.Iterador()
	borrado := it.Borrar()

	require.Equal(t, 1, borrado)
	require.True(t, l.EstaVacia())
}

func TestIteradorInsertarMultiplesSinAvanzar(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	it := l.Iterador()

	it.Insertar(1)
	it.Insertar(2)
	it.Insertar(3)

	require.Equal(t, 3, l.BorrarPrimero())
	require.Equal(t, 2, l.BorrarPrimero())
	require.Equal(t, 1, l.BorrarPrimero())
}

func TestIteradorBorrarHastaVacio(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	l.InsertarUltimo(2)

	it := l.Iterador()
	it.Borrar()
	it.Borrar()

	require.True(t, l.EstaVacia())
}

func TestIteradorInsertarLuegoDeTerminar(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)

	it := l.Iterador()
	it.Avanzar() // queda al final

	it.Insertar(2)

	require.Equal(t, 2, l.VerUltimo())
}

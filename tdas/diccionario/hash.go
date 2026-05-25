package diccionario

import (
	"fmt"

	"tdas/lista"
)

const (
	CAPACIDAD_INICIAL   = 10
	FACTOR_CARGA_MAXIMO = 0.75
	FACTOR_CARGA_MINIMO = 0.25
	FACTOR_REDIMENSION  = 2
)

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []lista.Lista[parClaveValor[K, V]]
	tam      int
	cantidad int
}

type iterHash[K comparable, V any] struct {
	hash         *hashAbierto[K, V]
	iterLista    lista.IteradorLista[parClaveValor[K, V]]
	indiceActual int
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func funcionHashing(bytes []byte) int {
	hash := 0

	for _, b := range bytes {
		hash = hash*31 + int(b)
	}

	if hash < 0 {
		hash = -hash
	}

	return hash
}

func (hash *hashAbierto[K, V]) indice(clave K) int {
	return funcionHashing(convertirABytes(clave)) % hash.tam
}

func crearTabla[K comparable, V any](tam int) []lista.Lista[parClaveValor[K, V]] {
	tabla := make([]lista.Lista[parClaveValor[K, V]], tam)

	for i := range tabla {
		tabla[i] = lista.CrearListaEnlazada[parClaveValor[K, V]]()
	}

	return tabla
}

func (hash *hashAbierto[K, V]) redimensionar(nuevoTam int) {
	tablaVieja := hash.tabla
	hash.tabla = crearTabla[K, V](nuevoTam)
	hash.tam = nuevoTam
	hash.cantidad = 0

	for _, l := range tablaVieja {
		l.Iterar(func(par parClaveValor[K, V]) bool {
			indice := hash.indice(par.clave)
			hash.tabla[indice].InsertarUltimo(par)
			hash.cantidad++
			return true
		})
	}
}

func (hash *hashAbierto[K, V]) buscarClave(clave K) lista.IteradorLista[parClaveValor[K, V]] {
	iter := hash.tabla[hash.indice(clave)].Iterador()
	for iter.HayAlgoMas() {
		if iter.VerActual().clave == clave {
			return iter
		}
		iter.Avanzar()
	}
	return iter
}

func (hash *hashAbierto[K, V]) buscarClaveExistente(clave K) lista.IteradorLista[parClaveValor[K, V]] {
	iter := hash.buscarClave(clave)
	if !iter.HayAlgoMas() {
		panic("La clave no pertenece al diccionario")
	}
	return iter
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {

	tabla := crearTabla[K, V](CAPACIDAD_INICIAL)

	return &hashAbierto[K, V]{tabla: tabla, tam: CAPACIDAD_INICIAL, cantidad: 0}
}

func (hash *hashAbierto[K, V]) Guardar(clave K, dato V) {

	iter := hash.buscarClave(clave)

	if !iter.HayAlgoMas() {

		if float64(hash.cantidad+1)/float64(hash.tam) > FACTOR_CARGA_MAXIMO {
			hash.redimensionar(hash.tam * FACTOR_REDIMENSION)

			iter = hash.buscarClave(clave)
		}

		hash.cantidad++
	} else {
		iter.Borrar()
	}

	iter.Insertar(parClaveValor[K, V]{clave, dato})
}

func (hash *hashAbierto[K, V]) Obtener(clave K) V {
	return hash.buscarClaveExistente(clave).VerActual().dato
}

func (hash *hashAbierto[K, V]) Pertenece(clave K) bool {
	return hash.buscarClave(clave).HayAlgoMas()
}

func (hash *hashAbierto[K, V]) Borrar(clave K) V {
	iter := hash.buscarClaveExistente(clave)
	dato := iter.VerActual().dato
	iter.Borrar()
	hash.cantidad--

	if float64(hash.cantidad)/float64(hash.tam) < FACTOR_CARGA_MINIMO &&
		hash.tam > CAPACIDAD_INICIAL {

		nuevoTam := hash.tam / FACTOR_REDIMENSION

		if nuevoTam < CAPACIDAD_INICIAL {
			nuevoTam = CAPACIDAD_INICIAL
		}

		hash.redimensionar(nuevoTam)
	}

	return dato
}

func (hash *hashAbierto[K, V]) Cantidad() int {
	return hash.cantidad
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (iter *iterHash[K, V]) saltarVacios() {
	for iter.indiceActual < iter.hash.tam && iter.hash.tabla[iter.indiceActual].EstaVacia() {
		iter.indiceActual++
	}
}

func (hash *hashAbierto[K, V]) Iterar(f func(clave K, dato V) bool) {
	for _, l := range hash.tabla {
		salir := false
		l.Iterar(func(par parClaveValor[K, V]) bool {
			if !f(par.clave, par.dato) {
				salir = true
				return false
			}
			return true
		})
		if salir {
			break
		}
	}
}

func (hash *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iterHash[K, V]{hash: hash, indiceActual: 0}

	iter.saltarVacios()

	if iter.indiceActual < iter.hash.tam {
		iter.iterLista = iter.hash.tabla[iter.indiceActual].Iterador()
	}

	return iter
}

func (iter *iterHash[K, V]) HayAlgoMas() bool {
	return iter.iterLista != nil && iter.iterLista.HayAlgoMas()
}

func (iter *iterHash[K, V]) VerActual() (K, V) {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	par := iter.iterLista.VerActual()
	return par.clave, par.dato
}

func (iter *iterHash[K, V]) Avanzar() {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}

	iter.iterLista.Avanzar()

	if !iter.iterLista.HayAlgoMas() {
		iter.indiceActual++

		iter.saltarVacios()

		if iter.indiceActual < iter.hash.tam {
			iter.iterLista = iter.hash.tabla[iter.indiceActual].Iterador()
		} else {
			iter.iterLista = nil
		}
	}
}

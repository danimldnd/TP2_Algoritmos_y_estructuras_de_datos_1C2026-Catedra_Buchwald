package diccionario

import "tdas/pila"

type funcCmp[K comparable] func(K, K) int

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

type iterAbb[K comparable, V any] struct {
	pila  pila.Pila[*nodoAbb[K, V]]
	abb   *abb[K, V]
	desde *K
	hasta *K
}

func crearNodoAbb[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{
		clave: clave,
		dato:  dato,
	}
}

func (a *abb[K, V]) guardarNodo(actual *nodoAbb[K, V], clave K, dato V) *nodoAbb[K, V] {
	if actual == nil {
		a.cantidad++
		return crearNodoAbb(clave, dato)
	}
	comparacion := a.cmp(clave, actual.clave)

	if comparacion < 0 {

		actual.izquierdo = a.guardarNodo(actual.izquierdo, clave, dato)

	} else if comparacion > 0 {

		actual.derecho = a.guardarNodo(actual.derecho, clave, dato)

	} else {

		actual.dato = dato
	}

	return actual
}

func (a *abb[K, V]) obtenerNodo(actual *nodoAbb[K, V], clave K) *nodoAbb[K, V] {
	if actual == nil {
		return nil
	}

	comparacion := a.cmp(clave, actual.clave)

	if comparacion < 0 {
		return a.obtenerNodo(actual.izquierdo, clave)
	}

	if comparacion > 0 {
		return a.obtenerNodo(actual.derecho, clave)
	}

	return actual
}

func buscarMinimo[K comparable, V any](actual *nodoAbb[K, V]) *nodoAbb[K, V] {
	if actual.izquierdo == nil {
		return actual
	}

	return buscarMinimo(actual.izquierdo)
}

func (a *abb[K, V]) borrarNodo(actual **nodoAbb[K, V], clave K) (V, bool) {

	var cero V

	if *actual == nil {
		return cero, false
	}

	comparacion := a.cmp(clave, (*actual).clave)

	if comparacion < 0 {
		return a.borrarNodo(&(*actual).izquierdo, clave)
	}

	if comparacion > 0 {
		return a.borrarNodo(&(*actual).derecho, clave)
	}

	datoBorrado := (*actual).dato

	if (*actual).izquierdo == nil {
		*actual = (*actual).derecho
		return datoBorrado, true
	}

	if (*actual).derecho == nil {
		*actual = (*actual).izquierdo
		return datoBorrado, true
	}

	sucesor := buscarMinimo((*actual).derecho)

	(*actual).clave = sucesor.clave
	(*actual).dato = sucesor.dato

	a.borrarNodo(&(*actual).derecho, sucesor.clave)

	return datoBorrado, true
}

func (a *abb[K, V]) iterarRangoNodo(actual *nodoAbb[K, V], desde *K, hasta *K, visitar func(K, V) bool) bool {

	if actual == nil {
		return true
	}

	if desde == nil || a.cmp(actual.clave, *desde) > 0 {
		if !a.iterarRangoNodo(actual.izquierdo, desde, hasta, visitar) {
			return false
		}
	}

	if (desde == nil || a.cmp(actual.clave, *desde) >= 0) &&
		(hasta == nil || a.cmp(actual.clave, *hasta) <= 0) {

		if !visitar(actual.clave, actual.dato) {
			return false
		}
	}

	if hasta == nil || a.cmp(actual.clave, *hasta) < 0 {
		if !a.iterarRangoNodo(actual.derecho, desde, hasta, visitar) {
			return false
		}
	}

	return true
}

func (iter *iterAbb[K, V]) apilarRamaIzquierdaRango(actual *nodoAbb[K, V]) {

	for actual != nil {

		if iter.desde != nil &&
			iter.abb.cmp(actual.clave, *iter.desde) < 0 {

			actual = actual.derecho

		} else {

			iter.pila.Apilar(actual)
			actual = actual.izquierdo
		}
	}
}

func (iter *iterAbb[K, V]) HayAlgoMas() bool {

	if iter.pila.EstaVacia() {
		return false
	}

	if iter.hasta != nil {

		actual := iter.pila.VerTope()

		if iter.abb.cmp(actual.clave, *iter.hasta) > 0 {
			return false
		}
	}

	return true
}

func (iter *iterAbb[K, V]) VerActual() (K, V) {

	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}

	actual := iter.pila.VerTope()

	return actual.clave, actual.dato
}

func (iter *iterAbb[K, V]) Avanzar() {

	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}

	actual := iter.pila.Desapilar()

	if actual.derecho != nil {
		iter.apilarRamaIzquierdaRango(actual.derecho)
	}
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: 0, cmp: funcion_cmp}
}

func (a *abb[K, V]) Guardar(clave K, dato V) {
	a.raiz = a.guardarNodo(a.raiz, clave, dato)
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	return a.obtenerNodo(a.raiz, clave) != nil
}

func (a *abb[K, V]) Obtener(clave K) V {
	elemento := a.obtenerNodo(a.raiz, clave)
	if elemento == nil {
		panic("La clave no pertenece al diccionario")
	}
	return elemento.dato
}

func (a *abb[K, V]) Borrar(clave K) V {

	dato, ok := a.borrarNodo(&a.raiz, clave)

	if !ok {
		panic("La clave no pertenece al diccionario")
	}

	a.cantidad--

	return dato
}

func (a *abb[K, V]) Cantidad() int {
	return a.cantidad
}

func (a *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {

	if desde != nil && hasta != nil &&
		a.cmp(*desde, *hasta) > 0 {
		return
	}

	a.iterarRangoNodo(a.raiz, desde, hasta, visitar)
}

func (a *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {

	iter := &iterAbb[K, V]{
		pila:  pila.CrearPilaDinamica[*nodoAbb[K, V]](),
		abb:   a,
		desde: desde,
		hasta: hasta,
	}

	if desde != nil && hasta != nil &&
		a.cmp(*desde, *hasta) > 0 {

		return iter
	}

	iter.apilarRamaIzquierdaRango(a.raiz)

	return iter
}

func (a *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return a.IteradorRango(nil, nil)
}

func (a *abb[K, V]) Iterar(visitar func(K, V) bool) {
	a.IterarRango(nil, nil, visitar)
}

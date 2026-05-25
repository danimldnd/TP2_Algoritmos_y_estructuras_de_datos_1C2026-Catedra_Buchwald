package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func crearNodo[T any](dato T, siguiente *nodoLista[T]) *nodoLista[T] {
	return &nodoLista[T]{dato: dato, siguiente: siguiente}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(elementoParaInsertar T) {

	nuevo_nodo := crearNodo(elementoParaInsertar, lista.primero)

	if lista.EstaVacia() {
		lista.ultimo = nuevo_nodo
	}

	lista.primero = nuevo_nodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elementoParaInsertar T) {

	nuevo_nodo := crearNodo(elementoParaInsertar, nil)

	if lista.EstaVacia() {
		lista.primero = nuevo_nodo
	} else {
		lista.ultimo.siguiente = nuevo_nodo
	}

	lista.ultimo = nuevo_nodo
	lista.largo++
}
func (lista *listaEnlazada[T]) BorrarPrimero() T {

	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	dato := lista.primero.dato
	lista.primero = lista.primero.siguiente

	if lista.primero == nil {
		lista.ultimo = nil
	}
	lista.largo--

	return dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {

	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero

	for actual != nil {
		if !visitar(actual.dato) {
			return
		}
		actual = actual.siguiente
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type iteradorLista[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorLista[T]{
		actual:   lista.primero,
		anterior: nil,
		lista:    lista,
	}
}

func (iter *iteradorLista[T]) HayAlgoMas() bool {
	return iter.actual != nil
}

func (iter *iteradorLista[T]) VerActual() T {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.dato
}

func (iter *iteradorLista[T]) Avanzar() {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iteradorLista[T]) Insertar(valor T) {
	nuevo := crearNodo(valor, iter.actual)

	if iter.anterior == nil {
		iter.lista.primero = nuevo
	} else {
		iter.anterior.siguiente = nuevo
	}

	if iter.actual == nil {
		iter.lista.ultimo = nuevo
	}

	iter.actual = nuevo
	iter.lista.largo++
}

func (iter *iteradorLista[T]) Borrar() T {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}

	borrado := iter.actual

	if iter.anterior == nil {
		iter.lista.primero = borrado.siguiente
	} else {
		iter.anterior.siguiente = borrado.siguiente
	}

	if borrado.siguiente == nil {
		iter.lista.ultimo = iter.anterior
	}

	iter.actual = borrado.siguiente
	iter.lista.largo--

	return borrado.dato
}

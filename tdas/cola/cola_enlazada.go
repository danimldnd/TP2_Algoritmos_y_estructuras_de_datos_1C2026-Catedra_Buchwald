package cola

type nodo[T any] struct {
	valor T
	sig   *nodo[T]
}

type colaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
}

func crearNodo[T any](valor T) *nodo[T] {
	return &nodo[T]{valor: valor}
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{}
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.valor
}

func (cola *colaEnlazada[T]) Encolar(valor T) {
	nuevo := crearNodo(valor)

	if cola.ultimo != nil {
		cola.ultimo.sig = nuevo
	} else {
		cola.primero = nuevo
	}

	cola.ultimo = nuevo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}

	valor := cola.primero.valor
	cola.primero = cola.primero.sig

	if cola.primero == nil {
		cola.ultimo = nil
	}

	return valor
}

package lista

type Lista[T any] interface {

	//EstaVacia devuelve verdadero si la lista esta vacia
	EstaVacia() bool

	//InsertarPrimero inserta un elemento al principio de la lista
	InsertarPrimero(T)

	//InsertarUltimo inserta un elemento al final de la lista
	InsertarUltimo(T)

	//BorrarPrimero elimina el primer elemento de la lista y devuelve ese valor. Si la lista esta vacia
	//Entra en panico con un mensaje "la lista esta vacia"
	BorrarPrimero() T

	//VerPrimero devuelve el primer elemento de la lista. Si la lista esta vacia
	//Entra en panico con un mensaje "la lista esta vacia"
	VerPrimero() T

	//VerUltimo devuelve el ultimo elemento de la lista. Si la lista esta vacia
	//Entra en panico con un mensaje "la lista esta vacia"
	VerUltimo() T

	//Largo devuelve el largo de la lista.
	Largo() int

	//Iterar itera de manera interna a la lista
	Iterar(visitar func(T) bool)

	//Iterador devuelve un iterador externo
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual devuelve el elemento actual del iterador.
	// Debe existir un elemento actual válido; en caso contrario entra en pánico
	// con el mensaje "El iterador termino de iterar".
	// No modifica el estado del iterador.
	VerActual() T

	// HayAlgoMas indica si el iterador tiene un elemento actual válido.
	// Devuelve true si el iterador no llegó al final, false en caso contrario.
	// No modifica el estado del iterador.
	HayAlgoMas() bool

	// Avanzar mueve el iterador al siguiente elemento de la lista.
	// Debe existir un elemento actual válido; en caso contrario entra en pánico
	// con el mensaje "El iterador termino de iterar".
	// Luego de avanzar, el iterador apunta al siguiente elemento o queda al final.
	Avanzar()

	// Insertar agrega un nuevo elemento en la posición actual del iterador.
	// Si el iterador está al inicio, el elemento se inserta como primero.
	// Si el iterador está en el medio, se inserta antes del elemento actual.
	// Si el iterador está al final, se inserta como último.
	// Luego de la operación, el iterador pasa a apuntar al elemento insertado.
	Insertar(T)

	// Borrar elimina el elemento actual del iterador y lo devuelve.
	// Debe existir un elemento actual válido; en caso contrario entra en pánico
	// con el mensaje "El iterador termino de iterar".
	// Luego de borrar, el iterador pasa a apuntar al siguiente elemento o queda al final.
	Borrar() T
}

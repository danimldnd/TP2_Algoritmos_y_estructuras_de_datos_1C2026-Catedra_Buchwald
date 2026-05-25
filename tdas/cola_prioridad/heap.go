package cola_prioridad

const (
	FACTOR_REDIMENSION = 2
	CAPACIDAD_INICIAL  = 10
	CRITERIO_REDUCCION = 4
)

type heap[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func heapify[T any](datos []T, tam int, cmp func(T, T) int) {
	for i := (tam / 2) - 1; i >= 0; i-- {
		downheap(datos, i, tam, cmp)
	}
}

func CrearHeapArr[T any](arreglo []T, cmp func(T, T) int) ColaPrioridad[T] {
	n := len(arreglo)
	capacidad := n

	if capacidad < CAPACIDAD_INICIAL {
		capacidad = CAPACIDAD_INICIAL
	}

	datos := make([]T, capacidad)
	copy(datos, arreglo)

	h := &heap[T]{
		datos: datos,
		cant:  n,
		cmp:   cmp,
	}

	heapify(h.datos, h.cant, h.cmp)

	return h
}

func HeapSort[T any](elementos []T, cmp func(T, T) int) {
	n := len(elementos)

	heapify(elementos, n, cmp)

	for tam := n - 1; tam > 0; tam-- {
		elementos[0], elementos[tam] = elementos[tam], elementos[0]
		downheap(elementos, 0, tam, cmp)
	}
}

func downheap[T any](datos []T, pos, tam int, cmp func(T, T) int) {
	for {
		izq := 2*pos + 1
		der := 2*pos + 2
		mayor := pos

		if izq < tam && cmp(datos[izq], datos[mayor]) > 0 {
			mayor = izq
		}
		if der < tam && cmp(datos[der], datos[mayor]) > 0 {
			mayor = der
		}
		if mayor == pos {
			return
		}

		datos[pos], datos[mayor] = datos[mayor], datos[pos]
		pos = mayor
	}
}

func (h *heap[T]) redimensionar(nuevoTam int) {
	nuevo := make([]T, nuevoTam)
	copy(nuevo, h.datos)
	h.datos = nuevo
}

func (h *heap[T]) upheap(pos int) {
	if pos == 0 {
		return
	}

	padre := (pos - 1) / 2

	if h.cmp(h.datos[pos], h.datos[padre]) <= 0 {
		return
	}

	h.datos[pos], h.datos[padre] = h.datos[padre], h.datos[pos]

	h.upheap(padre)
}

func (h *heap[T]) Cantidad() int {
	return h.cant
}

func (h *heap[T]) EstaVacia() bool {
	return h.cant == 0
}

func CrearHeap[T any](cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{
		datos: make([]T, CAPACIDAD_INICIAL),
		cmp:   cmp,
	}
}

func (h *heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	return h.datos[0]
}

func (h *heap[T]) Encolar(elem T) {

	if h.cant == len(h.datos) {
		h.redimensionar(len(h.datos) * FACTOR_REDIMENSION)
	}
	h.datos[h.cant] = elem
	h.cant++

	h.upheap(h.cant - 1)
}

func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}

	h.cant--

	h.datos[0], h.datos[h.cant] = h.datos[h.cant], h.datos[0]

	max := h.datos[h.cant]

	var cero T
	h.datos[h.cant] = cero

	if h.cant > 0 {
		downheap(h.datos, 0, h.cant, h.cmp)
	}

	if h.cant*CRITERIO_REDUCCION <= len(h.datos) &&
		len(h.datos) > CAPACIDAD_INICIAL {
		h.redimensionar(len(h.datos) / FACTOR_REDIMENSION)
	}

	return max
}

package pila

const (
	TAM_INICIAL        = 1
	FACTOR_CRECIMIENTO = 2
	FACTOR_REDUCCION   = 4
)

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func (p *pilaDinamica[T]) redimensionar(nuevoTam int) {
	nuevosDatos := make([]T, nuevoTam)
	copy(nuevosDatos, p.datos[:p.cantidad])
	p.datos = nuevosDatos
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(elemento T) {
	if p.cantidad == len(p.datos) {
		p.redimensionar(len(p.datos) * FACTOR_CRECIMIENTO)
	}
	p.datos[p.cantidad] = elemento
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}

	tope := p.datos[p.cantidad-1]
	p.cantidad--

	if p.cantidad <= len(p.datos)/FACTOR_REDUCCION && len(p.datos)/FACTOR_CRECIMIENTO >= TAM_INICIAL {
		p.redimensionar(len(p.datos) / FACTOR_CRECIMIENTO)
	}

	return tope
}
func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{
		datos:    make([]T, TAM_INICIAL),
		cantidad: 0,
	}
}

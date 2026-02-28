package set

type Set[T comparable] map[T]T

func New[T comparable](values ...T) Set[T] {
	s := make(Set[T])

	for _, value := range values {
		s.Add(value)
	}

	return s
}

func (s Set[T]) Add(value T) {
	s[value] = value
}

func (s Set[T]) Remove(value T) {
	delete(s, value)
}

package list

type list[T comparable] struct {
	items []T
}

// Add implements List.
func (l *list[T]) Add(item T) {
	l.items = append(l.items, item)
}

// Contains implements List.
func (l *list[T]) Contains(item T) bool {
	for _, i := range l.items {
		if i == item {
			return true
		}
	}

	return false
}

// Get implements List.
func (l *list[T]) Get(item T) (T, bool) {
	has := false
	for _, i := range l.items {
		if i == item {
			item = i
			has = true
		}
	}
	return item, has
}

// Len implements List.
func (l *list[T]) Len() int {
	return len(l.items)
}

// Remove implements List.
func (l *list[T]) Remove(item T) {
	if !l.Contains(item) {
		return
	}

	for i, v := range l.items {
		if v == item {
			l.items = l.items[:i+copy(l.items[i:], l.items[i+1:])]
			break
		}
	}
}

func (l *list[T]) copy(a, b []T) []T {
	b = make([]T, len(a))
	copy(b, a)
	return b
}

// Values implements List.
func (l *list[T]) Values() []T {
	return l.items
}

type List[T any] interface {
	Add(T)
	Remove(T)
	Contains(T) bool
	Len() int
	Values() []T
	Get(T) (T, bool)
}

func New[T comparable](size int) List[T] {
	return &list[T]{items: make([]T, size)}
}

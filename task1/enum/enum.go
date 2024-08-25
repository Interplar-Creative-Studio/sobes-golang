package enum

type Enum[K comparable, V any] struct {
	// TODO implement
}

func New[K comparable, V any]() *Enum[K, V] {
	return &Enum[K, V]{}
}

// TODO нужен функционал для добавления допустимых значений перечисления
// TODO нужен функционал для проверки существования значения

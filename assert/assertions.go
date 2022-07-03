package assert

func That[T any](t T) FluentAssertion[T] {
	return nil
}

type FluentAssertion[T any] interface {
}

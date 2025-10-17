package helpers

import "iter"

func EmptyIter[T any]() iter.Seq[T] {
	return func(_ func(T) bool) {}
}

func ErrorFunc(err error) func() error {
	return func() error {
		return err
	}
}

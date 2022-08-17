package asserter

// Stack We're saying here that the type of the stack constrains what values it can work with.
//If you intiailize an int stack, you can only use ints, whereas if we used interface{} instead,
//We could throw anything onto the stack and lead to some ugly behavior.
type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) IsEmpty() bool {
	if len(s.values) == 0 {
		return true
	}
	return false
}

func (s *Stack[T]) Push(val T) {
	s.values = append(s.values, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	index := len(s.values) - 1
	val := s.values[index]
	s.values = s.values[:index]
	return val, true
}

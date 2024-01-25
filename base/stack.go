type Stack []interface{}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Top() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		element := (*s)[index]
		return element, true
	}
}

func main() {
	s := NewStack()
	s.Push(1)
	s.Push('a')
	s.Push("hello")
}
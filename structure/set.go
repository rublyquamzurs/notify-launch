package structure

type key = string

type Set struct {
	mem map[key]bool
}

func NewSet() *Set {
	return &Set{make(map[key]bool)}
}

func (s *Set) Add(value key) {
	s.mem[value] = true
}

func (s *Set) Contains(value key) bool {
	return s.mem[value]
}

func (s *Set) Remove(value key) {
	delete(s.mem, value)
}

func (s *Set) Len() int {
	return len(s.mem)
}

func (s *Set) Clear() {
	s.mem = make(map[key]bool)
}

func (s *Set) IsEmpty() bool {
	return len(s.mem) == 0
}

func (s *Set) Values() []key {
	values := make([]key, 0)
	for value, _ := range s.mem {
		values = append(values, value)
	}
	return values
}

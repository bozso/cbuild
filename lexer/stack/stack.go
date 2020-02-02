package stack

import (
    "github.com/bozso/cbuild/lexer/item"
)

type T struct {
    items []item.T
    pos int
}


func New(count int) (s T) {
	s.items = make([]item.T, count)
	s.pos = 0

	return
}
    
func (s *Stack) Push(it T) {
    if s.pos > len(s.items) {
        s.items = append(s.items, it)
    } else {
        s.items[s.pos++] = it
    }
}

func (s *Stack) Pop() (it T, e error) {
	if s.pos == 0 {
        e = fmt.Errorf("No more items")
        return
	}

	s.pos--
	return s.items[s.pos], nil
}


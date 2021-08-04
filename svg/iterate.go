package svg

import (
	"io"
)

// IterFunc is the function called upon each iteration of the iterator.
type IterFunc func(io.Writer, *iterator, *Node) *Node

type stack []string

func (s *stack) add(str string) {
	*s = append(*s, str)
}

func (s *stack) pop() (str string) {
	st := *s
	if len(st) == 0 {
		return ""
	}
	str = st[len(st)-1]
	*s = st[:len(st)-1]
	return
}

type iterator struct {
	stack
	depth int
	fn    IterFunc
}

func (i *iterator) iterate(w io.Writer, n *Node) {

	// Work to be done.
	if i.fn != nil {
		n = i.fn(w, i, n)
	}

	// Nested elements.
	if n.FirstChild != nil {
		i.iterate(w, n.FirstChild)
	}

	// Siblings.
	if n.NextSibling != nil {
		i.iterate(w, n.NextSibling)
	}
}

func (n *Node) Iterate(w io.Writer, fn IterFunc) {
	i := iterator{make(stack, 0, 10), 0, fn}
	i.iterate(w, n)
}

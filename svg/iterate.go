package svg

import (
	"io"
)

// IterFunc is the function called upon each iteration of the iteratior.
type IterFunc func(io.Writer, *Node, int) (*Node, int)

type stack []*Node

func (s *stack) add(n *Node) {
	*s = append(*s, n)
}

func (s *stack) pop() (n *Node) {
	st := *s
	if len(st) == 0 {
		return nil
	}
	n = st[len(st)-1]
	*s = st[:len(st)-1]
	return
}

type iterator struct {
	stack
}

func (i iterator) iterate(w io.Writer, n *Node, fn IterFunc, d int) {

	// Work to be done.
	if fn != nil {
		n, d = fn(w, n, d)
	}

	// Nested elements.
	if n.FirstChild != nil {
		i.iterate(w, n.FirstChild, fn, d)
	}

	// Siblings.
	if n.NextSibling != nil {
		i.iterate(w, n.NextSibling, fn, d)
	}
}

func (n *Node) Iterate(w io.Writer, fn IterFunc) {
	i := iterator{make(stack, 0, 10)}
	i.iterate(w, n, fn, 0)
}

package xml

import (
	"io"
)

// IterFunc is the function called upon each iteration of the iteratior.
type IterFunc func(io.Writer, *Node, int, interface{}) (*Node, int)

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

func (i iterator) iterate(w io.Writer, n *Node, depth int, fn IterFunc) {

	// End of the branch.
	if n == nil {
		return
	}

	// Work to be done.
	if fn != nil {
		n, depth = fn(w, n, depth, nil)
	}

	// Nested elements.
	if n.FirstChild != nil {
		i.iterate(w, n.FirstChild, depth, fn)
	}

	// Siblings.
	if n.NextSibling != nil {
		i.iterate(w, n.NextSibling, depth, fn)
	}
}

func (n *Node) Iterate(w io.Writer, fn IterFunc) {
	var depth int
	i := iterator{make(stack, 0, 10)}
	i.iterate(w, n, depth, fn)
}

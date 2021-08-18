package svg

import (
	"io"
)

// IterFunc is the function called upon each iteration.
type IterFunc func(*iterator, *Node) *Node

type stack []string

func (s *stack) add(str string) {
	*s = append(*s, str)
}

func (s *stack) pop() (str string) {
	if len(*s) == 0 {
		return ""
	}
	str = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}

func (s *stack) peek() (str string) {
	if len(*s) == 0 {
		return ""
	}
	return (*s)[len(*s)-1]
}

type iterator struct {
	// the iterator writes to this output.
	w io.Writer
	// stack maintains state by preserving a last in first out
	// record of the tag names when nesting.
	stack
	// depth holds the current indentation depth of the iterator.
	depth int
	// fn is called upon each iteration.
	fn IterFunc
	// buffer for indentation characters.
	indent []byte
	// the indentation that is to be used.
	ichar string
}

func (i *iterator) iterate(n *Node) {
	// Something to be done?
	if i.fn != nil {
		n = i.fn(i, n)
	}
	// Nested elements, depth first.
	if n.FirstChild != nil {
		i.iterate(n.FirstChild)
	}
	// Siblings, linked list.
	if n.NextSibling != nil {
		i.iterate(n.NextSibling)
	}
}

func (n *Node) Iterate(w io.Writer, fn IterFunc) {
	i := iterator{w, make(stack, 10), 0, fn, make([]byte, 10), "\t"}
	i.iterate(n)
}

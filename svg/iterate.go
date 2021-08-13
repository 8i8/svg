package svg

import (
	"io"
)

// IterFunc is the function called upon each iteration of the iterator.
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
	// the iterator writes to this ouput.
	w io.Writer
	// stack maintains state by preserving a last in first out
	// record of the tag names whist nesting.
	stack
	// depth holds the current indentation depth of the iterator.
	depth int
	// Fn is called upon each iteration.
	fn IterFunc
	// buffer for indentation characters.
	indent []byte
	// the indentation that is to be used.
	ichar string
}

func (i *iterator) iterate(n *Node) {

	if n == nil {
		panic("expected to recieve a valid node")
	}

	// Work to be done.
	if i.fn != nil {
		n = i.fn(i, n)
	}

	// Nested elements.
	if n.FirstChild != nil {
		i.iterate(n.FirstChild)
	}

	// Siblings.
	if n.NextSibling != nil {
		i.iterate(n.NextSibling)
	}
}

func (n *Node) Iterate(w io.Writer, fn IterFunc) {
	i := iterator{w, make(stack, 10), 0, fn, make([]byte, 10), "\t"}
	i.iterate(n)
}

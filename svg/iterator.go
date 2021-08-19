package svg

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"svg/svg/lex"
	"svg/svg/xml"
)

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Iterator
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

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

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  pretty print
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// nameSpaceToken converts the xml name space of nameSpaceToken generated
// xml.Tokens into the required abbreviation for the nameSpaceToken program and
// for browsers.
func nameSpaceToken(str string) string {
	switch str {
	case "http://www.w3.org/XML/1998/namespace":
		return "xml"
	case "http://www.w3.org/1999/xlink":
		return "xlink"
	case "http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd":
		return "sodipodi"
	case "http://www.inkscape.org/namespaces/inkscape":
		return "inkscape"
	case "http://www.w3.org/1999/02/22-rdf-syntax-ns#":
		return "rdf"
	case "ttp://www.w3.org/2000/01/rdf-schema#":
		return "rdfs"
	case "http://xmlns.com/foaf/0.1/":
		return "foaf"
	default:
		return str
	}
}

func (i iterator) indentation() {
	if str := i.peek(); !processingText(str) {
		if cap(i.indent) < i.depth*len(i.ichar) {
			i.indent = make([]byte, i.depth*len(i.ichar))
		}
		i.indent = i.indent[:0]
		for n := 0; n < i.depth; n++ {
			i.indent = append(i.indent, i.ichar...)
		}
		i.w.Write(i.indent)
	}
}

func processingText(str string) bool {
	switch str {
	case "text", "tspan":
		return true
	default:
		return false
	}
}

func insideToken(str string) bool {
	if strings.Contains(str, "-token") {
		return true
	}
	return false
}

func (i iterator) newLineInToken() {
	if str := i.peek(); insideToken(str) {
		io.WriteString(i.w, "\n")
	}
}

func (i iterator) copyFormatNewLine(c xml.CharData) {
	n := bytes.Count([]byte(c), []byte("\n"))
	for n > 0 {
		io.WriteString(i.w, "\n")
		n--
	}
}

func PrettyPrint(i *iterator, n *Node) *Node {
	switch n.Elem.typ {
	case Root:
	case StartElement:
		v := n.Elem.xml.(xml.StartElement)
		i.indentation()

		// Name
		io.WriteString(i.w, "<"+v.Name.Local)
		i.add(v.Name.Local)
		i.add(v.Name.Local + "-token")

		// Attributes.
		i.depth++ // Augment nesting.
		for _, a := range v.Attr {
			i.newLineInToken()
			i.indentation()
			// Name spaces.
			if a.Name.Space != "" {
				io.WriteString(i.w, nameSpaceToken(a.Name.Space)+":"+
					a.Name.Local+"=\"")
			} else {
				io.WriteString(i.w, a.Name.Local+"=\"")
			}
			io.WriteString(i.w, a.Value+"\"")
		}
		i.pop()
		i.depth-- // Decrement nesting.

		// If there is no next sibling we need to close the tag
		// and indent.
		if n.NextSibling == nil {
			io.WriteString(i.w, ">")
			i.depth++
			return n
		}

		// If there is a next sibling which is an EndElement
		// then this must be an element with an intigrated end
		// tag, remove the tag from the stack closing its state.
		if n.NextSibling.Elem.typ == EndElement {
			io.WriteString(i.w, " />")
			i.pop()
			return n.NextSibling
		}

		// We have arrived end of the open element tag and need
		// to close it as the next tag will be nested, if the
		// next tag is CharData then skip the new line char.
		io.WriteString(i.w, ">")
		i.depth++

	case EndElement:
		v := n.Elem.xml.(xml.EndElement)
		i.depth--
		i.indentation()
		io.WriteString(i.w, "</"+v.Name.Local+">")
		i.pop()

	case CharData:
		v := n.Elem.xml.(xml.CharData)
		if processingText(i.peek()) {
			i.w.Write([]byte(v))
		} else {
			i.copyFormatNewLine(v)
		}

	case Comment:
		v := n.Elem.xml.(xml.Comment)
		i.indentation()
		io.WriteString(i.w, "<!--")
		i.w.Write([]byte(v))
		io.WriteString(i.w, "-->")

	case ProcInst:
		v := n.Elem.xml.(xml.ProcInst)
		i.indentation()
		io.WriteString(i.w, "<?"+v.Target+" ")
		i.w.Write(v.Inst)
		io.WriteString(i.w, "?>")

	case Directive:
		v := n.Elem.xml.(xml.Directive)
		i.indentation()
		i.w.Write(v)
	default:
		fmt.Printf("svg/xml: printNode: unknown type: %#v\n", n.Elem)
	}
	return n
}

func tabIndent(n int) []byte {
	buf := make([]byte, n)
	for n > 0 {
		buf[n-1] = '\t'
		n--
	}
	return buf
}

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Functions
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

func retrieveElemByID(node *Node, id string) IterFunc {
	return func(i *iterator, n *Node) *Node {
		if n == nil {
			panic("expected a valid node")
		}
		if n.Elem.typ != StartElement {
			return n
		}
		elem := n.Elem.xml.(xml.StartElement)
		for _, attr := range elem.Attr {
			if attr.Name.Local == "id" && attr.Value == id {
				*node = *n
			}
		}
		return n
	}
}

func insertElemById(node *Node, id string) IterFunc {
	return func(i *iterator, n *Node) *Node {
		if n == nil {
			panic("expected a valid node")
		}
		if n.Elem.typ != StartElement {
			return n
		}
		elem := n.Elem.xml.(xml.StartElement)
		for _, a := range elem.Attr {
			if a.Name.Local == "id" && a.Value == id {
				*n = *node
			}
		}
		return n
	}
}

type replace struct {
	attr, old, new string
}

func replaceAttr(l *lex.Lexer, e xml.StartElement, repl ...replace) {
	for _, r := range repl {
		for _, a := range e.Attr {
			if a.Name.Local == r.attr {
				l.Lex(e.Name.Local, r.attr, a.Value)
				for {
					item := l.Next()
					if item.EOF() {
						break
					}
					if err := item.Error(); err != nil {
						fmt.Println(err)
						break
					}
					fmt.Println("item:", item)
				}
			}
		}
	}
}

func checkAttrID(l *lex.Lexer, e xml.StartElement, id string, r ...replace) {
	for _, a := range e.Attr {
		if a.Value == id {
			replaceAttr(l, e, r...)
		}
	}
}

func setAttrStyle(l *lex.Lexer, n *Node, id string, repl ...replace) error {
	node := &Node{}
	n.Iterate(os.Stdout, retrieveElemByID(node, id))
	if node == nil || node.Elem.xml == nil {
		return fmt.Errorf("id %q not found", id)
	}
	if node.Elem.typ != StartElement {
		return fmt.Errorf("not a StartElement: %q", node.Elem.typ)
	}
	e := node.Elem.xml.(xml.StartElement)
	checkAttrID(l, e, id, repl...)
	node.Elem.xml = e
	return nil

}

// pprint pretty prints the parse tree.
func pprint(n *Node) {
	n.Iterate(os.Stdout, PrettyPrint)
}

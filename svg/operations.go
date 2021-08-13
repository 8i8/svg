package svg

import (
	"fmt"
	"os"
	"svg/svg/attr/lex"
	"svg/svg/xml"
)

type Lexer interface {
	Reset(elem, name, value string)
	Next() lex.Item
}

func retrieveElemByID(node *Node, id string) IterFunc {
	return func(i *iterator, n *Node) *Node {
		if n == nil {
			panic("expected a valid node")
		}
		elem, ok := n.Elem.(xml.StartElement)
		if !ok {
			return n
		}
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
		elem, ok := n.Elem.(xml.StartElement)
		if !ok {
			return n
		}
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
	if node == nil || node.Elem == nil {
		return fmt.Errorf("id %q not found", id)
	}
	if e, ok := node.Elem.(xml.StartElement); ok {
		checkAttrID(l, e, id, repl...)
		node.Elem = e
	}
	return nil

}

func pprint(n *Node) {
	n.Iterate(os.Stdout, PrettyPrint)
}

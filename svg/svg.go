package svg

import (
	"fmt"
	"io"
	"os"
	"svg/svg/xml"
)

func getElemByID(node *Node, id string) IterFunc {
	return func(w io.Writer, i *iterator, n *Node) *Node {
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
	return func(w io.Writer, i *iterator, n *Node) *Node {
		elem, ok := n.Elem.(xml.StartElement)
		if !ok {
			return n
		}
		for _, attr := range elem.Attr {
			if attr.Name.Local == "id" && attr.Value == id {
				*n = *node
			}
		}
		return n
	}
}

func Default(in io.ReadCloser) {
	//_ = Parse(in)
	root := Parse(in)
	// node := &Node{}
	// root.Iterate(os.Stdout, getElemByID(node, "path40096"))
	// if e, ok := node.Elem.(xml.StartElement); ok {
	// 	e := e
	// 	for i, a := range e.Attr {
	// 		if a.Name.Local == "style" {
	// 			e.Attr[i].Value = "display:inline;fill:#ffffff;fill-opacity:1;stroke:#ff00ff;stroke-opacity:1"
	// 		}
	// 	}
	// 	node.Elem = e
	// }
	root.Iterate(os.Stdout, PrettyPrint)
}

func Open(in io.Reader) {
	fmt.Println("Open mode")
}

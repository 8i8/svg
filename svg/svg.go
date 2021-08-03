package svg

import (
	"fmt"
	"io"
	"os"
	"svg/svg/xml"
)

func getElemByID(node *Node, id string) IterFunc {
	return func(w io.Writer, n *Node, d int) (*Node, int) {
		elem, ok := n.Elem.(xml.StartElement)
		if !ok {
			return n, d
		}
		for _, attr := range elem.Attr {
			if attr.Name.Local == "id" && attr.Value == id {
				*node = *n
			}
		}
		return n, d
	}
}

func insertElemById(node *Node, id string) IterFunc {
	return func(w io.Writer, n *Node, d int) (*Node, int) {
		elem, ok := n.Elem.(xml.StartElement)
		if !ok {
			return n, d
		}
		for _, attr := range elem.Attr {
			if attr.Name.Local == "id" && attr.Value == id {
				*n = *node
			}
		}
		return n, d
	}
}

func Default(in io.ReadCloser) {
	//_ = Parse(in)
	root := Parse(in)
	// node := &Node{}
	// root.Iterate(os.Stdout, getElemByID(node, "test"))
	// if node != nil {
	// 	elem, ok := node.Elem.(xml.StartElement)
	// 	if !ok {
	// 		log.Fatal("svg.Default: xml.StartElement !ok")
	// 	}
	// 	var found bool
	// 	for i, a := range elem.Attr {
	// 		if a.Name.Local == "fill" {
	// 			elem.Attr[i].Value = "blue"
	// 		}
	// 		if !found {
	// 			elem.Attr = append(elem.Attr, xml.Attr{
	// 				Name:  xml.Name{Local: "fill"},
	// 				Value: "blue"})
	// 		}
	// 	}
	// 	root.Iterate(os.Stdout, insertElemById(node, "test"))
	// 	fmt.Printf("%v\n", node)
	// }
	root.Iterate(os.Stdout, PrettyPrint)
}

func Open(in io.Reader) {
	fmt.Println("Open mode")
}

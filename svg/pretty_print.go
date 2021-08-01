package svg

import (
	"fmt"
	"io"
	"svg/svg/xml"
)

func PrettyPrint(w io.Writer, n *Node) *Node {
	switch v := n.Elem.(type) {
	case xml.StartElement:
		io.WriteString(w, "<"+v.Name.Local)
		for _, attr := range v.Attr {
			io.WriteString(w, " "+attr.Name.Local+"=\"")
			io.WriteString(w, attr.Value+"\"")
		}
		if _, ok := n.NextSibling.Elem.(xml.EndElement); ok {
			io.WriteString(w, " />")
			n = n.NextSibling
		} else {
			io.WriteString(w, ">")
		}
	case xml.EndElement:
		io.WriteString(w, "</"+v.Name.Local+">")
	case xml.CharData:
		w.Write([]byte(v))
	case xml.Comment:
		io.WriteString(w, "<!--")
		w.Write([]byte(v))
		io.WriteString(w, "-->")
	case xml.ProcInst:
		io.WriteString(w, "<?"+v.Target+" ")
		w.Write(v.Inst)
		io.WriteString(w, "?>")
	case xml.Directive:
		w.Write(v)
	default:
		fmt.Printf("svg/xml: printNode: unknown type: %#v\n", v)
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

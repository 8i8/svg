package svg

import (
	"bytes"
	"fmt"
	"io"
	"svg/svg/xml"
)

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

func stripIndenting(v xml.CharData) []byte {
	return bytes.Trim([]byte(v), "\v\r\t ")
}

func checkElem(name string) bool {
	switch name {
	case "text", "tspan", "/tspan":
		return true
	default:
		return false
	}
}

func prevNotCharData(w io.Writer, n *Node) bool {
	if n.PrevSibling != nil {
		if _, ok := n.PrevSibling.Elem.(xml.CharData); ok {
			return false
		}
	}
	return true
}

func indentation(w io.Writer, n int) {
	buf := make([]byte, n, n)
	for i := 0; i < n; i++ {
		buf[i] = '\t'
	}
	w.Write(buf)
}

func newLine(w io.Writer) {
	io.WriteString(w, "\n")
}

func PrettyPrint(w io.Writer, i *iterator, n *Node) *Node {
	switch v := n.Elem.(type) {
	case xml.StartElement:
		// TODO Display bug is caused by this indentation call.
		//indentation(w, d)
		// Name
		io.WriteString(w, "<"+v.Name.Local)

		// Attributes.
		i.depth++ // Augment nesting.
		for _, a := range v.Attr {
			newLine(w)
			indentation(w, i.depth)
			// Name spaces.
			if a.Name.Space != "" {
				io.WriteString(w, nameSpaceToken(a.Name.Space)+":"+
					a.Name.Local+"=\"")
			} else {
				io.WriteString(w, a.Name.Local+"=\"")
			}
			io.WriteString(w, a.Value+"\"")
		}
		i.depth-- // Decrement nesting.

		// If there is no next sibling we need to close the tag
		// and indent.
		if n.NextSibling == nil {
			io.WriteString(w, ">")
			newLine(w)
			i.depth++
			return n
		}

		// If there is a next sibling then this must be an
		// element with an intigrated end tag.
		if _, ok := n.NextSibling.Elem.(xml.EndElement); ok {
			io.WriteString(w, " />")
			newLine(w)
			return n.NextSibling
		}

		// We have arrived end of the open element tag and need
		// to close it as the next tag will be nested, if the
		// next tag is CharData then skip the new line char.
		i.depth++
		if n.NextSibling != nil {
			if _, ok := n.NextSibling.Elem.(xml.CharData); ok {
				io.WriteString(w, ">")
				return n
			}
		}
		io.WriteString(w, ">")
		newLine(w)
	case xml.EndElement:
		i.depth--
		if prevNotCharData(w, n) {
			// TODO Display bug is caused by this indentation call.
			indentation(w, i.depth)
		}
		io.WriteString(w, "</"+v.Name.Local+">")
		newLine(w)
	case xml.CharData:
		w.Write([]byte(v))
	case xml.Comment:
		indentation(w, i.depth)
		io.WriteString(w, "<!--")
		w.Write([]byte(v))
		io.WriteString(w, "-->")
		newLine(w)
	case xml.ProcInst:
		indentation(w, i.depth)
		io.WriteString(w, "<?"+v.Target+" ")
		w.Write(v.Inst)
		io.WriteString(w, "?>")
		newLine(w)
	case xml.Directive:
		indentation(w, i.depth)
		w.Write(v)
		newLine(w)
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

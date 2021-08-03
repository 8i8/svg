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

func exceptions(w io.Writer, n *Node) bool {
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

func PrettyPrint(w io.Writer, n *Node, d int) (*Node, int) {
	switch v := n.Elem.(type) {
	case xml.StartElement:
		indentation(w, d)
		// Name
		io.WriteString(w, "<"+v.Name.Local)

		// Attributes.
		d++ // Augment nesting.
		for _, a := range v.Attr {
			newLine(w)
			indentation(w, d)
			// Name spaces.
			if a.Name.Space != "" {
				io.WriteString(w, nameSpaceToken(a.Name.Space)+":"+
					a.Name.Local+"=\"")
			} else {
				io.WriteString(w, a.Name.Local+"=\"")
			}
			io.WriteString(w, a.Value+"\"")
		}
		d-- // Decrement nesting.

		// If there is no next sibling then this is the end of a
		// StartElement mid flow and we need to close it.
		if n.NextSibling == nil {
			io.WriteString(w, ">")
			newLine(w)
			d++
			return n, d
		}

		// If there is a next sibling then this must be an
		// element with an intigrated end tag.
		if _, ok := n.NextSibling.Elem.(xml.EndElement); ok {
			io.WriteString(w, " />")
			newLine(w)
			return n.NextSibling, d
		}

		// We have arrived end of the open element tag and need
		// to close it as the next tag will be nested, if the
		// next tag is CharData then skip the new line char.
		d++
		if n.NextSibling != nil {
			if _, ok := n.NextSibling.Elem.(xml.CharData); ok {
				io.WriteString(w, ">")
				return n, d
			}
		}
		io.WriteString(w, ">")
		newLine(w)
	case xml.EndElement:
		d--
		if exceptions(w, n) {
			indentation(w, d)
		}
		io.WriteString(w, "</"+v.Name.Local+">")
		newLine(w)
	case xml.CharData:
		w.Write([]byte(v))
	case xml.Comment:
		indentation(w, d)
		io.WriteString(w, "<!--")
		w.Write([]byte(v))
		io.WriteString(w, "-->")
		newLine(w)
	case xml.ProcInst:
		indentation(w, d)
		io.WriteString(w, "<?"+v.Target+" ")
		w.Write(v.Inst)
		io.WriteString(w, "?>")
		newLine(w)
	case xml.Directive:
		indentation(w, d)
		w.Write(v)
		newLine(w)
	default:
		fmt.Printf("svg/xml: printNode: unknown type: %#v\n", v)
	}
	return n, d
}

func tabIndent(n int) []byte {
	buf := make([]byte, n)
	for n > 0 {
		buf[n-1] = '\t'
		n--
	}
	return buf
}

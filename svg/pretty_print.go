package svg

import (
	"bytes"
	"fmt"
	"io"
	"strings"
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
	switch v := n.Elem.(type) {
	case xml.StartElement:

		// TODO Display bug is caused by this indentation call.
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

		// If there is a next sibling then this must be an
		// element with an intigrated end tag, remove the tag
		// from the stack closing its state.
		if _, ok := n.NextSibling.Elem.(xml.EndElement); ok {
			io.WriteString(i.w, " />")
			i.pop()
			return n.NextSibling
		}

		// We have arrived end of the open element tag and need
		// to close it as the next tag will be nested, if the
		// next tag is CharData then skip the new line char.
		i.depth++
		io.WriteString(i.w, ">")

	case xml.EndElement:
		i.depth--
		i.indentation()
		io.WriteString(i.w, "</"+v.Name.Local+">")
		i.pop()

	case xml.CharData:
		if processingText(i.peek()) {
			i.w.Write([]byte(v))
		} else {
			i.copyFormatNewLine(v)
		}

	case xml.Comment:
		i.indentation()
		io.WriteString(i.w, "<!--")
		i.w.Write([]byte(v))
		io.WriteString(i.w, "-->")

	case xml.ProcInst:
		i.indentation()
		io.WriteString(i.w, "<?"+v.Target+" ")
		i.w.Write(v.Inst)
		io.WriteString(i.w, "?>")

	case xml.Directive:
		i.indentation()
		i.w.Write(v)
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

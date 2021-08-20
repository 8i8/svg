package svg

import (
	"fmt"
	"io"
	"log"
	"svg/svg/xml"
)

const verbose = false

// Type defines the type of xml node.
type Type int

const (
	Error Type = iota
	Root
	StartElement
	EndElement
	CharData
	Comment
	ProcInst
	Directive
	End
)

type parser struct {
	head    *Node
	current **Node
	depth   int
	endelem bool // endelem tracks when an xml StartElement closes ritht away
}

type token struct {
	typ Type
	xml xml.Token
}

type Node struct {
	Parent, FirstChild, LastChild, PrevSibling, NextSibling *Node

	Elem token
}

func (p *parser) init(n *Node) {
	p.head = n
	p.head.Parent = new(Node)
	p.head.Parent.Elem.typ = Root
	p.head.Parent.FirstChild = p.head
	p.head.Parent.LastChild = p.head
	p.current = &(p.head)
}

// makeNode returns a copy of the given node.
func makeNode(t xml.Token) (node *Node) {
	node = new(Node)
	switch v := t.(type) {
	case xml.StartElement:
		node.Elem = token{StartElement, v.Copy()}
	case xml.EndElement:
		node.Elem = token{EndElement, t}
	case xml.CharData:
		node.Elem = token{CharData, v.Copy()}
	case xml.Comment:
		node.Elem = token{Comment, v.Copy()}
	case xml.ProcInst:
		node.Elem = token{ProcInst, v.Copy()}
	case xml.Directive:
		node.Elem = token{Directive, v.Copy()}
	}
	return
}

// addLastSibling adds an EndElement, linking back up to the orginating
// parent.
func (p *parser) addLastSibling(n *Node) {
	if (*p.current) == nil {
		log.Fatalf("busted last sibling: %#v\n", n.Elem)
	}
	if (*p.current).Parent == nil {
		log.Fatalf("busted last sibling has no parent: %#v\n", n.Elem)
	}
	n.Parent = (*p.current).Parent
	n.PrevSibling = (*p.current)
	(*p.current).NextSibling = n
	(*p.current).Parent.LastChild = n
	p.current = &(*p.current).Parent
}

// addSibling sets the node for many of the xml.Token types for the
// parseToken function.
func (p *parser) addSibling(n *Node) {
	if (*p.current) == nil {
		log.Fatalf("busted sibling: %#v\n", n.Elem)
	}
	n.PrevSibling = (*p.current)
	n.Parent = (*p.current).Parent
	(*p.current).NextSibling = n
	p.current = &(*p.current).NextSibling
}

func (p *parser) addChild(n *Node) {
	if (*p.current) == nil {
		log.Fatalf("busted child: %#v\n", n.Elem)
	}
	if (*p.current).Parent == nil {
		log.Fatalf("busted last sibling has no parent: %#v\n", n.Elem)
	}
	if (*p.current).FirstChild == nil {
		(*p.current).FirstChild = n
		(*p.current).FirstChild.Parent = (*p.current)
		(*p.current).LastChild = (*p.current).FirstChild
		p.current = &(*p.current).FirstChild
		return
	}
	n.PrevSibling = (*p.current).LastChild
	(*p.current).LastChild.NextSibling = n
	n.Parent = (*p.current)
	(*p.current).LastChild = n
	p.current = &(*p.current).LastChild
}

// addToParseTree progressively constructs a parse tree.
//
// element types:
//	StartElement
//	EndElement
//	CharData
//	Comment
//	ProcInst
//	Directive
//
func (p *parser) addToParseTree(current, next *Node) {
	switch current.Elem.typ {
	case StartElement:
		if next.Elem.typ == EndElement {
			// The next token is an EndElement so we need to
			// add this as a sibling rather than as a child.
			p.addSibling(current)
			p.endelem = true
			if verbose {
				v := current.Elem.xml.(xml.StartElement)
				fmt.Println("StartElement: sibling", v.Name.Local)
			}
		} else {
			p.addChild(current)
			if verbose {
				v := current.Elem.xml.(xml.StartElement)
				fmt.Println("StartElement: nested:", v.Name.Local)
			}
		}
		if verbose {
			v := current.Elem.xml.(xml.StartElement)
			for _, atr := range v.Attr {
				fmt.Printf("\t%s: %s\n", atr.Name.Local, atr.Value)
			}
		}
	case EndElement:
		// TODO now this function needs to know whether or not
		// the EndElement has been nesting or not.
		if p.endelem {
			p.addSibling(current)
			p.endelem = false
		} else {
			p.addLastSibling(current)
		}
		if verbose {
			v := current.Elem.xml.(xml.EndElement)
			fmt.Println("EndElement:", v.Name.Local)
		}
	case CharData:
		p.addSibling(current)
		if verbose {
			v := current.Elem.xml.(xml.CharData)
			fmt.Printf("CharData: %#v\n", v)
		}
	case Comment:
		p.addSibling(current)
		if verbose {
			v := current.Elem.xml.(xml.Comment)
			fmt.Printf("Comment: %s\n", string(v))
		}
	case ProcInst:
		p.addSibling(current)
		if verbose {
			v := current.Elem.xml.(xml.ProcInst)
			fmt.Printf("ProcInst: %v\n", v)
		}
	case Directive:
		p.addSibling(current)
		if verbose {
			v := current.Elem.xml.(xml.Directive)
			fmt.Printf("Directive: %s\n", string(v))
		}
	default:
		fmt.Printf("Unknown token: %#v\n", current.Elem.xml)
	}
	return
}

// parse the input stream top down, with a lookahead of one token.
func (p *parser) parse(in io.Reader) *Node {
	d := xml.NewDecoder(in)

	var current, next *Node

	// Set the first token into the parser, the parsing loop can not
	// do this as simply as we can from here.
	xmltoken, err := d.Token()
	if err != nil {
		fmt.Println("svg/svg: parse:", err)
		return nil
	}
	p.init(makeNode(xmltoken))

	// We need a lookahead of one token, for this we prepare a token
	// in advance.
	xmltoken, err = d.Token()
	if err != nil {
		fmt.Println("svg/svg: parse:", err)
		return nil
	}
	current = makeNode(xmltoken)

	for {
		xmltoken, err = d.Token()
		if err != nil && err != io.EOF {
			fmt.Println("error: svg/svg: parse:", err)
			continue
		}
		// Lookahead in the parse tree function uses the next
		// token, whilst acting upon the current.
		next = makeNode(xmltoken)
		if err != nil {
			p.addToParseTree(current, next)
			break
		}
		p.addToParseTree(current, next)
		current = next
	}
	return p.head
}

func Parse(in io.ReadCloser) *Node {
	defer in.Close()
	p := new(parser)
	return p.parse(in)
}

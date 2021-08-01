package xml

import (
	"encoding/xml"
	"fmt"
	"io"
)

const verbose = false

type NodeType int

const (
	Head NodeType = iota
	StartElement
	EndElement
	CharData
	Comment
	ProcInst
	Directive
)

type parser struct {
	*head
	current **Node
	depth   int
}

type head struct {
	Type NodeType

	node, FirstChild, LastChild, PrevSibling, NextSibling *Node
}

type Node struct {
	Parent, FirstChild, LastChild, PrevSibling, NextSibling *Node

	Type NodeType
	Elem xml.Token
}

func (p *parser) init() {
	p.head = new(head)
	p.head.Type = Head
}

// WhichType returns the type of an xml.Token as one of the packages
// mirroring constants.
func WhichType(t xml.Token) (typ NodeType) {
	switch t.(type) {
	case xml.StartElement:
		return StartElement
	case xml.EndElement:
		return EndElement
	case xml.CharData:
		return CharData
	case xml.Comment:
		return Comment
	case xml.ProcInst:
		return ProcInst
	case xml.Directive:
		return Directive
	}
	return
}

// CopyNode returns a copy of the given node.
func CopyNode(t xml.Token) (node xml.Token) {
	switch v := t.(type) {
	case xml.StartElement:
		return v.Copy()
	case xml.EndElement:
		return t
	case xml.CharData:
		return v.Copy()
	case xml.Comment:
		return v.Copy()
	case xml.ProcInst:
		return v.Copy()
	case xml.Directive:
		return v.Copy()
	}
	return
}

// addNode sets the node for many of the xml.Token types for the
// parseToken function.
func (p *parser) addNode(t NodeType, v xml.Token) {
	if p.current == nil {
		p.node = &Node{Type: t, Elem: v}
		p.current = &p.node
		return
	}
	(*p.current).NextSibling = &Node{Type: t, Elem: v}
	(*p.current).NextSibling.Parent = (*p.current).Parent
	(*p.current).NextSibling.PrevSibling = (*p.current)
	p.current = &(*p.current).NextSibling
	if (*p.current).Parent != nil {
		(*p.current).Parent.LastChild = (*p.current)
	}
}

func (p *parser) nestNode(t NodeType, v xml.Token) {
	if p.current == nil {
		p.node = &Node{Type: t, Elem: v}
		p.current = &p.node
		return
	}
	(*p.current).FirstChild = &Node{Type: t, Elem: v}
	(*p.current).FirstChild.Parent = (*p.current)
	(*p.current).LastChild = (*p.current).FirstChild
	p.current = &(*p.current).FirstChild
}

// parseToken generates a parse tree from the tokens that it recieves.
//
// element types:
//	StartElement
//	EndElement
//	CharData
//	Comment
//	ProcInst
//	Directive
//
func (p *parser) parseToken(t, n xml.Token) {
	switch v := t.(type) {
	case xml.StartElement:
		next := WhichType(n)
		if next == EndElement {
			p.addNode(StartElement, v.Copy())
			if verbose {
				fmt.Println("StartElement:", v.Name.Local)
			}
		} else {
			p.nestNode(StartElement, v.Copy())
			if verbose {
				fmt.Println("StartElement: nested:", v.Name.Local)
			}
		}
		if verbose {
			for _, atr := range v.Attr {
				fmt.Printf("\t%s: %s\n", atr.Name.Local, atr.Value)
			}
		}
	case xml.EndElement:
		p.addNode(EndElement, v)
		if verbose {
			fmt.Println("EndElement:", v.Name.Local)
		}
	case xml.CharData:
		p.addNode(CharData, v.Copy())
		if verbose {
			fmt.Printf("CharData: %#v\n", v)
		}
	case xml.Comment:
		p.addNode(Comment, v.Copy())
		if verbose {
			fmt.Printf("Comment: %s\n", string(v))
		}
	case xml.ProcInst:
		p.addNode(ProcInst, v.Copy())
		if verbose {
			fmt.Printf("ProcInst: %v\n", v)
		}
	case xml.Directive:
		p.addNode(Directive, v.Copy())
		if verbose {
			fmt.Printf("Directive: %s\n", string(v))
		}
	default:
		fmt.Printf("Unknown token: %#v\n", v)
	}
	return
}

func (p *parser) parse(in io.Reader) *Node {
	p.init()
	d := xml.NewDecoder(in)

	var token, next xml.Token
	var err error

	token, err = d.Token()
	if err != nil {
		fmt.Println("svg/xml:", err)
		return nil
	}
	for {
		next, err = d.Token()
		if err == io.EOF {
			p.parseToken(token, next)
			break
		} else if err != nil {
			fmt.Println("svg/xml:", err)
			continue
		}
		p.parseToken(token, next)
		token = CopyNode(next)
	}
	return p.head.node
}

func Parse(in io.ReadCloser) *Node {
	defer in.Close()
	p := new(parser)
	return p.parse(in)
}

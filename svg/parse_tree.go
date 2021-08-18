package svg

import (
	"fmt"
	"io"
	"svg/svg/xml"
)

const verbose = false

type parser struct {
	*head
	current **Node
	depth   int
}

type head struct {
	node, FirstChild, LastChild, PrevSibling, NextSibling *Node
}

type Node struct {
	Parent, FirstChild, LastChild, PrevSibling, NextSibling *Node

	Elem xml.Token
}

func (p *parser) init() {
	p.head = new(head)
}

// copyNode returns a copy of the given node.
func copyNode(t xml.Token) (node xml.Token) {
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

// addSibling sets the node for many of the xml.Token types for the
// parseToken function.
func (p *parser) addSibling(v xml.Token) {
	if p.current == nil {
		p.node = &Node{Elem: v}
		p.current = &p.node
		return
	}
	(*p.current).NextSibling = &Node{Elem: v}
	(*p.current).NextSibling.Parent = (*p.current).Parent
	(*p.current).NextSibling.PrevSibling = (*p.current)
	p.current = &(*p.current).NextSibling
	if (*p.current).Parent != nil {
		(*p.current).Parent.LastChild = (*p.current)
	}
}

func (p *parser) addChild(v xml.Token) {
	if p.current == nil {
		p.node = &Node{Elem: v}
		p.current = &p.node
		return
	}
	(*p.current).FirstChild = &Node{Elem: v}
	(*p.current).FirstChild.Parent = (*p.current)
	(*p.current).LastChild = (*p.current).FirstChild
	p.current = &(*p.current).FirstChild
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
func (p *parser) addToParseTree(t, n xml.Token) {
	switch v := t.(type) {
	case xml.StartElement:
		if _, ok := n.(xml.EndElement); ok {
			p.addSibling(v.Copy())
			if verbose {
				fmt.Println("StartElement: sibling", v.Name.Local)
			}
		} else {
			p.addChild(v.Copy())
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
		p.addSibling(v)
		if verbose {
			fmt.Println("EndElement:", v.Name.Local)
		}
	case xml.CharData:
		p.addSibling(v.Copy())
		if verbose {
			fmt.Printf("CharData: %#v\n", v)
		}
	case xml.Comment:
		p.addSibling(v.Copy())
		if verbose {
			fmt.Printf("Comment: %s\n", string(v))
		}
	case xml.ProcInst:
		p.addSibling(v.Copy())
		if verbose {
			fmt.Printf("ProcInst: %v\n", v)
		}
	case xml.Directive:
		p.addSibling(v.Copy())
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
		fmt.Println("svg/svg: parse:", err)
		return nil
	}
	token = copyNode(token)

	// Scan tokens.
	for {
		next, err = d.Token()
		if err == io.EOF {
			p.addToParseTree(token, next)
			break
		} else if err != nil {
			fmt.Println("svg/svg: parse:", err)
			continue
		}
		p.addToParseTree(token, next)
		token = copyNode(next)
	}
	return p.head.node
}

func Parse(in io.ReadCloser) *Node {
	defer in.Close()
	p := new(parser)
	return p.parse(in)
}
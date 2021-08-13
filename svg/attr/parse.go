package attr

import (
	"fmt"
	"svg/svg/attr/lex"
	"svg/svg/xml"
)

type Lexer interface {
	Reset(elem, name, value string)
	Next() string
}

func AttributeBuilder(l *lex.Lexer, elem string, a xml.Attr) interface{} {
	switch elem {
	case "path":
		return path(l, elem, a.Name.Local, a.Value)
	default:
		return nil
	}
}

func path(l *lex.Lexer, elem, name, value string) (p Path) {
	l.Lex(elem, name, value)
	for {
		a := l.Next()
		fmt.Println(a)
	}
	return p
}

package lex

import (
	"fmt"
	"svg/svg/attr/lex/set"
)

const verbose = false
const (
	rgbVarCount = 3
)

// analyze is the syntactic analysis listening loop that advances the
// function states one by one, that can be canceled by sending to the
// stop channel.
func (l *Lexer) analyze(fn stateFn) {
	l.anlz.run = true
	for state := fn; state != nil; {
		select {
		case <-l.anlz.stop:
			state = nil
		default:
			state = state(l)
		}
	}
	l.anlz.run = false
}

// startAnalysis is the entry point for each analysis path, specific to
// each SVG element type.
func startAnalysis(l *Lexer) stateFn {
	switch l.elem {
	case "path":
		switch l.attr {
		case "style":
			return anlzAttrKey
		}
	case "text":
		switch l.attr {
		case "style":
			return anlzAttrKey
		}
	}
	return nil
}

// attributeType sets an attribute specific expression en route.
func attributeType(l *Lexer, key string) (fn stateFn) {
	switch l.elem {
	case "path":
		return pathAttr(l, key)
	case "text":
		return textAttr(l, key)
	}
	const msg = "unknown type"
	return l.syntaxErrorf("%s: %s: %s: %q", l.elem, l.attr, msg, key)
}

func pathAttr(l *Lexer, key string) (fn stateFn) {
	switch key {
	case "fill":
		return anlzColon(anlzWord(anlzSemiColon(anlzAttrKey)))
	case "stroke-width":
		return anlzColon(anlzNumber(anlzUnit(anlzSemiColon(anlzAttrKey))))
	case "stroke":
		return anlzColon(anlzColour(anlzSemiColon(anlzAttrKey)))
	}
	const msg = "unknown attribute"
	return l.syntaxErrorf("%s: %s: %s: %q", l.elem, l.attr, msg, key)
}

func textAttr(l *Lexer, key string) (fn stateFn) {
	switch key {
	case "font-family":
		//return anlzFontFamily(nil)
		return nil
	case "font":
		return
	}
	const msg = "unknown attribute"
	return l.syntaxErrorf("%s: %s: %s: %q", l.elem, l.attr, msg, key)
}

// anlzAttrKey is the root expression of the style attribute.
func anlzAttrKey(l *Lexer) stateFn {
	t := <-l.anlz.send
	if verbose {
		const fname = "anlzAttrKey"
		fmt.Printf("%s: %s: %s: %s(%s)\n",
			fname, l.elem, l.attr, t.Type, t.Value)
	}
	if t.Type == ItemWhitespace {
		l.anlz.send <- t
		return anlzAttrKey
	}
	if t.Type == ItemEOF {
		l.anlz.send <- t
		return anlzEOF
	}
	if t.Type == ItemText {
		l.anlz.send <- t
		return attributeType(l, t.Value)
	}
	const msg = "attribute key error"
	return l.syntaxErrorf("%s: %s: %s: %s(%s)",
		l.elem, l.attr, msg, t.Type, t.Value)
}

func anlzColon(next stateFn) stateFn {
	return func(l *Lexer) stateFn {
		t := <-l.anlz.send
		if verbose {
			const fname = "anlzColon"
			fmt.Printf("%s: %s: %s: %s(%s)\n",
				fname, l.elem, l.attr, t.Type, t.Value)
		}
		if t.Type == ItemColon {
			l.anlz.send <- t
			return next
		}
		const msg = "expected a colon"
		return l.syntaxErrorf("%s: %s: %s: %s(%s)",
			l.elem, l.attr, msg, t.Type, t.Value)
	}
}

func anlzSemiColon(next stateFn) stateFn {
	return func(l *Lexer) stateFn {
		t := <-l.anlz.send
		if t.Type == ItemWhitespace {
			l.anlz.send <- t
			return anlzSemiColon(next)
		}
		if verbose {
			const fname = "anlzSemiColon"
			fmt.Printf("%s: %s: %s: %s(%s)\n",
				fname, l.elem, l.attr, t.Type, t.Value)
		}
		// Type or EOF, a semicolon may always be replaced by
		// an end of file token if it is the last attribute
		// within a semicolon seperated string.
		if t.Type == ItemSemiColon || t.Type == ItemEOF {
			l.anlz.send <- t
			return next
		}
		const msg = "expected a semicolon"
		return l.syntaxErrorf("%s: %s: %s: %s(%s)",
			l.elem, l.attr, msg, t.Type, t.Value)
	}
}

func anlzWord(next stateFn) stateFn {
	return func(l *Lexer) stateFn {
		t := <-l.anlz.send
		if t.Type == ItemWhitespace {
			l.anlz.send <- t
			return anlzWord(next)
		}
		if verbose {
			const fname = "anlzWord"
			fmt.Printf("%s: %s: %s: %s(%s)\n",
				fname, l.elem, l.attr, t.Type, t.Value)
		}
		if t.Type == ItemText {
			l.anlz.send <- t
			return next
		}
		const msg = "expected a word"
		return l.syntaxErrorf("%s: %s: %s: %s(%s)",
			l.elem, l.attr, msg, t.Type, t.Value)
	}
}

func anlzNumber(next stateFn) stateFn {
	return func(l *Lexer) stateFn {
		t := <-l.anlz.send
		if t.Type == ItemComma {
			l.anlz.send <- t
			return anlzNumber(next)
		}
		if t.Type == ItemWhitespace {
			l.anlz.send <- t
			return anlzNumber(next)
		}
		if verbose {
			const fname = "anlzNumber"
			fmt.Printf("%s: %s: %s: %s(%s)\n",
				fname, l.elem, l.attr, t.Type, t.Value)
		}
		if t.Type == ItemNumber {
			l.anlz.send <- t
			return next
		}
		const msg = "expected a number"
		return l.syntaxErrorf("%s: %s: %s: %s(%s)",
			l.elem, l.attr, msg, t.Type, t.Value)
	}
}

func anlzUnit(next stateFn) stateFn {
	return func(l *Lexer) stateFn {
		t := <-l.anlz.send
		if t.Type == ItemWhitespace {
			l.anlz.send <- t
			return anlzUnit(next)
		}
		// If there is a unit that is fine, account for it.
		if t.Type == ItemUnit {
			if t.Value == "" {
				l.anlz.send <- t
				if verbose {
					const fname = "no unit"
					fmt.Printf("%s: %s: %s: %s(%s)\n",
						fname, l.elem, l.attr, t.Type, t.Value)
				}
				return next
			}
			if set.UnitCheck[t.Value] {
				l.anlz.send <- t
				if verbose {
					const fname = "anlzUnit"
					fmt.Printf("%s: %s: %s: %s(%s)\n",
						fname, l.elem, l.attr, t.Type, t.Value)
				}
				return next
			} else {
				const msg = "unknown unit"
				return l.syntaxErrorf("%s: %s: %s: %s(%s)",
					l.elem, l.attr, msg, t.Type, t.Value)
			}
		}
		const msg = "expected a unit"
		return l.syntaxErrorf("%s: %s: %s: %s(%s)",
			l.elem, l.attr, msg, t.Type, t.Value)
	}
}

// https://www.w3.org/TR/1998/REC-CSS2-19980512/syndata.html#value-def-color
// https://www.w3.org/Graphics/SVG/1.1/types.html#DataTypeColor
// https://www.w3.org/Graphics/SVG/1.1/types.html#ColorKeywords
func anlzColour(next stateFn) stateFn {
	return func(l *Lexer) stateFn {
		t := <-l.anlz.send
		if t.Type == ItemWhitespace {
			l.anlz.send <- t
			return anlzColour(next)
		}
		if verbose {
			const fname = "anlzColour"
			fmt.Printf("%s: %s: %s: %s(%s)\n",
				fname, l.elem, l.attr, t.Type, t.Value)
		}
		if t.Type == ItemHEXColour {
			l.anlz.send <- t
			return next
		}
		if t.Type == ItemRGBColour {
			l.anlz.send <- t
			return anlzOpenBracket(next)
		}
		if t.Type == ItemText {
			if _, ok := set.ColourCheck[t.Value]; ok {
				l.anlz.send <- t
				return next
			} else {
				const msg = "unknown colour"
				return l.syntaxErrorf("%s: %s: %s: %s(%s)",
					l.elem, l.attr, msg, t.Type, t.Value)
			}
		}
		const msg = "expected a colour"
		return l.syntaxErrorf("%s: %s: %s: %s(%s)",
			l.elem, l.attr, msg, t.Type, t.Value)
	}
}

func anlzOpenBracket(next stateFn) stateFn {
	return func(l *Lexer) stateFn {
		t := <-l.anlz.send
		if verbose {
			const fname = "anlzOpenBracket"
			fmt.Printf("%s: %s: %s: %s(%s)\n",
				fname, l.elem, l.attr, t.Type, t.Value)
		}
		if t.Type == ItemOpenBracket {
			l.anlz.send <- t
			return anlzRGB(next)
		}
		const msg = "expected an open bracket"
		return l.syntaxErrorf("%s: %s: %s: %s(%s)",
			l.elem, l.attr, msg, t.Type, t.Value)
	}
}

func anlzRGB(next stateFn) stateFn {
	return func(l *Lexer) stateFn {
		t := <-l.anlz.send
		if verbose {
			const fname = "anlzRGB"
			fmt.Printf("%s: %s: %s: %s(%s)\n",
				fname, l.elem, l.attr, t.Type, t.Value)
		}
		if t.Type == ItemWhitespace {
			l.anlz.send <- t
			return anlzRGB(next)
		}
		if t.Type == ItemText {
			const msg = "expected a number"
			return l.syntaxErrorf("%s: %s: %s: %s(%s)",
				l.elem, l.attr, msg, t.Type, t.Value)
		}
		if t.Type == ItemNumber && l.count < rgbVarCount-1 {
			l.anlz.send <- t
			l.count++
			return anlzUnit(anlzComma(anlzRGB(next)))
		}
		if t.Type >= ItemNumber {
			l.anlz.send <- t
			l.count++
			return anlzUnit(anlzRGB(next))
		}
		if t.Type == ItemCloseBracket && l.count == rgbVarCount {
			l.anlz.send <- t
			l.count = 0
			return anlzRGB(next)
		} else {
			const msg = "expected 3 variables"
			return l.syntaxErrorf("%s: %s: %s: %s(%s)",
				l.elem, l.attr, msg, t.Type, t.Value)
		}
		const msg = "expected an rgb colour"
		return l.syntaxErrorf("%s: %s: %s: %s(%s)",
			l.elem, l.attr, msg, t.Type, t.Value)
	}
}

func anlzComma(next stateFn) stateFn {
	return func(l *Lexer) stateFn {
		t := <-l.anlz.send
		if verbose {
			const fname = "anlzComma"
			fmt.Printf("%s: %s: %s: %s(%s)\n",
				fname, l.elem, l.attr, t.Type, t.Value)
		}
		if t.Type == ItemComma {
			l.anlz.send <- t
			return next
		}
		const msg = "expected a comma"
		return l.syntaxErrorf("%s: %s: %s: %s(%s)",
			l.elem, l.attr, msg, t.Type, t.Value)
	}
}

func anlzEOF(l *Lexer) (fn stateFn) {
	// Pause and wait for next state function.
	if verbose {
		const fname = "anlzEOF"
		fmt.Printf("%s: %s: %s\n", fname, l.elem, l.attr)
	}
	return fn
}

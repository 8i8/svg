package lex

import "fmt"

const verbose = true

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

func startAnalysis(l *Lexer) stateFn {
	switch l.elem {
	case "path":
		switch l.attr {
		case "style":
			return anlzStyle
		}
	case "text":
	}
	return nil
}

func attrType(l *Lexer, key string) (fn stateFn) {
	switch l.attr {
	case "style":
		switch key {
		case "fill":
			return anlzColon(anlzWord(anlzSemiColon(anlzStyle)))
		case "stroke-width":
			return anlzColon(anlzNumber(anlzUnit(anlzSemiColon(anlzStyle))))
		case "stroke":
			return anlzColon(anlzColour(anlzSemiColon(anlzStyle)))
		}
	}
	const msg = "unknown type"
	return l.syntaxErrorf("%s: %s: %s: %q", l.elem, l.attr, msg, key)
}

func anlzStyleWrap(next stateFn) stateFn {
	return func(l *Lexer) stateFn {
		return next
		t := <-l.anlz.send
		if t.Type == ItemEOF {
			l.anlz.send <- t
			return anlzEOF
		}
		if t.Type == ItemWhitespace {
			l.anlz.send <- t
			return anlzStyle
		}
		if t.Type == ItemText {
			l.anlz.send <- t
			return attrType(l, t.Value)
		}
		const msg = "attribute key error"
		return l.syntaxErrorf("%s: %s: %s: %s(%s)",
			l.elem, l.attr, msg, t.Type, t.Value)
	}
}

func anlzStyle(l *Lexer) stateFn {
	t := <-l.anlz.send
	if verbose {
		const fname = "anlzStyle"
		fmt.Printf("%s: %s: %s: %s(%s)\n",
			fname, l.elem, l.attr, t.Type, t.Value)
	}
	if t.Type == ItemWhitespace {
		l.anlz.send <- t
		return anlzStyle
	}
	if t.Type == ItemEOF {
		l.anlz.send <- t
		return anlzEOF
	}
	if t.Type == ItemText {
		l.anlz.send <- t
		return attrType(l, t.Value)
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
			if unitCheck[t.Value] {
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

// func anlzWhitespace(next stateFn) stateFn {
// 	return func(l *Lexer) stateFn {
// 		t := l.anlz.get()
// 		// If there is any whitespace, account for it.
// 		if t.Type == ItemWhitespace {
// 			if verbose {
// 				const fname = "anlzWhitespace"
// 				fmt.Printf("%s: %s: %s: %s(%s)\n",
// 					fname, l.elem, l.attr, t.Type, t.Value)
// 			}
// 			l.anlz.send <- t
// 			return next
// 		}
// 		// If there was no whitespace, put the token back onto
// 		// the channel.
// 		l.anlz.back <- t
// 		if verbose {
// 			const fname = "anlzWhitespace sending back"
// 			fmt.Printf("%s: %s: %s: %s(%s)\n",
// 				fname, l.elem, l.attr, t.Type, t.Value)
// 		}
// 		return next
// 	}
// }

func anlzColour(next stateFn) stateFn {
	return func(l *Lexer) stateFn {
		t := <-l.anlz.send
		if verbose {
			const fname = "anlzColour"
			fmt.Printf("%s: %s: %s: %s(%s)\n",
				fname, l.elem, l.attr, t.Type, t.Value)
		}
		if t.Type == ItemColour {
			l.anlz.send <- t
			return next
		}
		const msg = "expected a colour"
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

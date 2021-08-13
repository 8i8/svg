package lex

import "unicode"

// lexPathD consumes the string that follows the path data attribute.
// d="(path data)"
// https://www.w3.org/TR/SVG11/paths.html#PathData
func lexPathD(l *Lexer) stateFn {
	return nil
}

// lexNubger consumes a range of different number types.
func lexNumber(l *Lexer, next stateFn) stateFn {
	return func(l *Lexer) stateFn {
		// Optional leading sign.
		l.accept("+-")
		digits := "0123456789"
		// Is it hex?
		if l.accept("0") && l.accept("xX") {
			digits = "0123456789abcdefABCDEF"
		}
		l.acceptRun(digits)
		// Is there a decimal part?
		if l.accept(".") {
			l.acceptRun(digits)
		}
		// Scientific notation?
		if l.accept("eE") {
			l.accept("+-")
			l.acceptRun("0123456789")
		}
		// Is it imaginary?
		l.accept("i")
		// Next thing must not be alphanumeric.
		if unicode.IsLetter(l.peek()) || unicode.IsNumber(l.peek()) {
			l.next()
			return l.errorf("%s: bad number syntax: %q",
				l.elem, l.input[l.start:l.pos])
		}
		l.emit(ItemNumber)
		return next
	}
}

// lexText scans in text mode.
func lexText(l *Lexer) stateFn {
	for {
		if l.next() == eof {
			if l.pos > l.start {
				l.emit(ItemText)
			}
			break
		}
	}
	if l.pos > l.start {
		l.emit(ItemText)
	}
	l.emit(ItemEOF) // Useful to send and EOF token.
	return nil
}

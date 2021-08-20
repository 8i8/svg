package lex

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	// -1 cannot possible arrive in the input string, as such it is
	// safe for use as an EOF character.
	eof        = -1
	leftMeta   = "<" // Left and right meta are the tags that will.
	rightMeta  = ">" // tell the scanner that it is within a token.
	alphabet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"
	wspace     = " \t\v"
	hex        = "0123456789abcdefABCDEF"
	digits     = "0123456789."
	units      = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz%"
	whitespace = "\t\n\v\f\r "
)

type messager struct {
	send chan Item
	stop chan Item
	run  bool
}

// Lexer holds the state of the scanner.
type Lexer struct {
	elem   string   // used for error reports and selecting attribute sets.
	attr   string   // the elements attribute that is being scanned.
	input  string   // string being scanned.
	start  int      // start position of this item.
	pos    int      // current position in the input.
	width  int      // width of last rune read.
	count  int      // counts variables between brackets.
	lex    messager // channels for messaging the lexer.
	anlz   messager // channels for messaging the anylyzer.
	Output []Item
}

// NewLexer returns a lexer set with the given name and input string if
// they are provided, however none are required to create the lexer.
func NewLexer() *Lexer {
	l := &Lexer{
		lex:    messager{make(chan Item, 1), make(chan Item), false},
		anlz:   messager{make(chan Item, 1), make(chan Item), false},
		Output: make([]Item, 0, 100),
	}
	return l
}

// Lex reruns the lexer for the given element attribute.
func (l *Lexer) Lex(elem, attr, input string) {
	l.Stop()
	l.elem = elem
	l.attr = attr
	l.input = input
	l.start = 0
	l.pos = 0
	l.width = 0
	go l.lexLoop(startLexer(l))
	go l.analyze(startAnalysis(l))
}

// Next returns the next item from the scanner.
func (l *Lexer) Next() Item {
	return <-l.lex.send
}

// Stop halts the lexers running loop.
func (l *Lexer) Stop() {
	if l.lex.run {
		l.lex.stop <- Item{}
	}
	if l.anlz.run {
		l.anlz.stop <- Item{}
	}
}

// Close shuts down the lexers channel.
func (l *Lexer) Close() {
	close(l.lex.send)
	close(l.lex.stop)
	close(l.anlz.send)
	close(l.anlz.stop)
}

// lexLoop starts the lexers scanning loop, we start in identifier mode
// as this parser is performing the relay and taking over from the
// encoding/xml parser as such we expect to receive only those stings
// from inside xml tokens.
func (l *Lexer) lexLoop(fn stateFn) {
	l.lex.run = true
	for state := fn; state != nil; {
		select {
		case <-l.lex.stop:
			state = nil
		default:
			state = state(l)
		}
	}
	l.lex.run = false
}

// emit sends and item down the items channel for the given itemType.
func (l *Lexer) emit(t itemType) {
	l.anlz.send <- Item{t, l.input[l.start:l.pos]}
	if typ := <-l.anlz.send; typ.Type == ItemError {
		l.lex.send <- typ
	}
	l.lex.send <- Item{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

// next consumes the next rune and moves the l.pos on by the runes
// width.
func (l *Lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

// ignore steps over the pending input prior to this point.
func (l *Lexer) ignore() {
	l.start = l.pos
}

// backup steps back one rune.
func (l *Lexer) backup() {
	l.pos -= l.width
}

// peek returns but does not consume the next rune from the input.
func (l *Lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// accept consumes the next rune if it is from the valid set.
func (l *Lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of valid runes from the valid set.
func (l *Lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

// errorf returns an error token and terminates the scan by passing back
// a nil pointer that will be the next state, terminating l.run.
func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
	l.lex.send <- Item{
		ItemError,
		fmt.Sprintf(format, args...),
	}
	return nil
}

// serrorf returns an syntax error token and terminates the scan by
// passing back a nil pointer that will be the next state, terminating
// l.run
func (l *Lexer) syntaxErrorf(format string, args ...interface{}) stateFn {
	l.lex.send <- Item{
		ItemSyntaxError,
		fmt.Sprintf(format, args...),
	}
	return nil
}

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Item type tokens
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// itemType describe all the types that we expect to find within our
// lexicon.
//go:generate stringer -type=itemType
type itemType int

const (
	ItemError itemType = iota
	ItemSyntaxError
	ItemColon
	ItemDot
	ItemOpenBracket
	ItemCloseBracket
	ItemSemiColon
	ItemComma
	ItemHEXColour
	ItemRGBColour
	ItemAttribute
	ItemNumber
	ItemUnit
	ItemText
	ItemWhitespace
	ItemEOF
)

// Item contains the data of the token along with its type for further
// treatment.
type Item struct {
	Type  itemType
	Value string
}

func (i Item) EOF() bool {
	if i.Type == ItemEOF {
		return true
	}
	return false
}

func (i Item) Error() error {
	if i.Type == ItemError {
		return fmt.Errorf("lexical error: %s", i.Value)
	} else if i.Type == ItemSyntaxError {
		return fmt.Errorf("syntax error: %s", i.Value)
	}
	return nil
}

func (i Item) String() string {
	switch i.Type {
	case ItemEOF:
		return "EOF"
	case ItemError:
		return i.Value
	}
	if len(i.Value) > 20 {
		return fmt.Sprintf("%.20q...", i.Value)
	}
	return fmt.Sprintf("%q", i.Value)
}

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  state functions
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// stateFn represents the state of the scanner as a function that
// returns the next state.
type stateFn func(*Lexer) stateFn

func startLexer(l *Lexer) stateFn {
	const fname = "startLexer"
	switch l.elem {
	case "path":
		return startPath(l)
	case "text":
		return startPath(l)
	}
	return l.errorf("%s: %s: %q: unknown element",
		fname, l.elem, l.attr)
}

func startPath(l *Lexer) stateFn {
	const fname = "startPath"
	switch l.attr {
	case "d":
		return lexPathD
	case "style":
		return lexStyle
	}
	return l.errorf("%s: %s: %q: unknown attribute",
		fname, l.elem, l.attr)
}

func startText(l *Lexer) stateFn {
	const fname = "startText"
	switch l.attr {
	case "font-family":
		return lexStyle
	}
	return l.errorf("%s: %s: %q: unknown attribute",
		fname, l.elem, l.attr)
}

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  which state?
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

func (l *Lexer) whatIs() stateFn {
	const fname = "whatIs"
	switch l.attr {
	case "style":
		return l.readStyle()
	case "id":
	}
	return l.errorf("error: %s: %s", fname, l.attr)
}

func (l *Lexer) readStyle() stateFn {
	const fname = "readStyle"
	style := l.input[l.start:l.pos]
	switch style {
	case "fill":
		return lexStyle
	case "stroke":
		return lexStyle
	case "stroke-width":
		return lexStyle
	}
	return l.errorf("error: %s: %s", fname, style)
}

// lexStyle starts scanning a style attribute string.
func lexStyle(l *Lexer) stateFn {
	for {
		switch r := l.next(); {
		case r == eof:
			l.emit(ItemEOF)
			return nil
		case unicode.IsSpace(r):
			return lexSpace(lexStyle)
		case r == ':':
			l.emit(ItemColon)
			return lexStyle
		case r == '(':
			l.emit(ItemOpenBracket)
			return lexStyle
		case r == ')':
			l.emit(ItemCloseBracket)
			return lexStyle
		case r == ',':
			l.emit(ItemComma)
			return lexStyle
		case r == ';':
			l.emit(ItemSemiColon)
			return lexStyle
		case r == '#':
			l.acceptRun(hex)
			l.emit(ItemHEXColour)
			return lexStyle
		case unicode.IsNumber(r):
			l.backup()
			return lexValue(lexStyle)
		case r == '\'':
			lexQuotes(lexStyle)
		case r == '"':
			lexDoubleQuotes(lexStyle)
		case r == 'r':
			l.backup()
			return lexRGB(lexStyle)
		case unicode.IsLetter(r):
			l.backup()
			return lexWord(lexStyle)
		case unicode.IsControl(r) || unicode.IsPrint(r):
			l.emit(ItemError)
			return l.errorf("unexpected character: %q", l.input[l.start:l.pos])
		}
	}
	l.emit(ItemEOF)
	return nil
}

func lexWord(next stateFn) stateFn {
	return func(l *Lexer) stateFn {
		l.acceptRun(alphabet)
		l.emit(ItemText)
		return next
	}
}

func lexQuotes(next stateFn) stateFn {
	return func(l *Lexer) stateFn {
		l.acceptRun(alphabet + " ")
		if !l.accept("'") {
			return l.errorf("expected a closing quote")
		}
		l.emit(ItemText)
		return next
	}
}

func lexDoubleQuotes(next stateFn) stateFn {
	return func(l *Lexer) stateFn {
		l.acceptRun(alphabet + " ")
		if !l.accept("\"") {
			return l.errorf("expected a closing quote")
		}
		l.emit(ItemText)
		return next
	}
}

func lexValue(next stateFn) stateFn {
	return func(l *Lexer) stateFn {
		l.acceptRun(digits)
		l.emit(ItemNumber)
		// We send units even when empty so that the analysis
		// code reamins sane, negating the need for look ahead.
		l.acceptRun(units)
		l.emit(ItemUnit)
		return next
	}
}

func lexSpace(next stateFn) stateFn {
	return func(l *Lexer) stateFn {
		l.acceptRun(whitespace)
		l.emit(ItemWhitespace)
		return next
	}
}

func lexRGB(next stateFn) stateFn {
	return func(l *Lexer) stateFn {
		start := l.start
		l.acceptRun("rgbRGB")
		if l.input[l.start:l.pos] == "rgb" ||
			l.input[l.start:l.pos] == "RGB" {
			l.emit(ItemRGBColour)
			return next
		}
		l.start, l.pos, l.width = start, start, 0
		return lexWord(next)
	}
}

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Lex functions
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

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

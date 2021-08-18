package lex

// https://www.w3.org/Graphics/SVG/1.1/text.html#FontPropertiesUsedBySVG
func anlzFontFamily(next stateFn) stateFn {
	return func(l *Lexer) (fn stateFn) {
		return fn
	}
}

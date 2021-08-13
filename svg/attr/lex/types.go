package lex

// cm 	centimeters
// mm 	millimeters
// in 	inches (1in = 96px = 2.54cm)
// px  	pixels (1px = 1/96th of 1in)
// pt 	points (1pt = 1/72 of 1in)
// pc 	picas (1pc = 12 pt)

// Relative length units specify a length relative to another length
// property. Relative length units scale better between different
// rendering medium.

// em 	Relative to the font-size of the element (2em means 2 times the size of the current font)
// ex 	Relative to the x-height of the current font (rarely used)
// ch 	Relative to the width of the "0" (zero)
// rem 	Relative to font-size of the root element
// vw 	Relative to 1% of the width of the viewport*
// vh 	Relative to 1% of the height of the viewport*
// vmin 	Relative to 1% of viewport's* smaller dimension
// vmax 	Relative to 1% of viewport's* larger dimension
// % 	Relative to the parent element

var unitCheck = map[string]bool{
	"cm":   true,
	"mm":   true,
	"in":   true,
	"px":   true,
	"pt":   true,
	"pc":   true,
	"em":   true,
	"ex":   true,
	"ch":   true,
	"rem":  true,
	"vw":   true,
	"vh":   true,
	"vmin": true,
	"vmax": true,
	"%":    true,
}

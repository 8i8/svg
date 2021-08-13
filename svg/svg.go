package svg

import (
	"fmt"
	"io"
	"svg/svg/attr/lex"

	"github.com/8i8/conf"
)

func Default(in io.ReadCloser, c *conf.Config) {
	root := Parse(in)
	l := lex.NewLexer()
	defer l.Close()
	err := setAttrStyle(l, root, "flame", replace{"style", "stroke", "blue"})
	if err != nil {
		fmt.Println(err)
		return
	}
	//pprint(root)
}

func Open(in io.Reader, c *conf.Config) {
	fmt.Println("Open mode")
}

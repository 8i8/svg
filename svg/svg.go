package svg

import (
	"fmt"
	"io"
	"os"
	"svg/svg/xml"
)

func Default(in io.ReadCloser) {
	root := xml.Parse(in)
	root.Iterate(os.Stdout, xml.PrettyPrint)
}

func Open(in io.Reader) {
	fmt.Println("Open mode")
}

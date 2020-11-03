package svg

import (
	"bytes"
	"testing"
)

func TestViewBoxToBytes(t *testing.T) {
	vb := viewBox{minX: 234, minY: 232345092459, width: 3456, height: 3456}
	byt := vb.viewBoxToBytes()
	str1 := string(byt)
	str2 := "234 232345092459 3456 3456"
	if str1 != str2 {
		//t.Errorf("len: %s", byt[0:len(byt)])
		t.Errorf("error: viewBoxToByte: expected %s recieved %s", str2, str1)
	}
}

func TestViewBox(t *testing.T) {
	fname := "ViewBox"
	buf := bytes.Buffer{}
	i := Image{}
	i.ViewBox(&buf, 234, 232345092459, 3456, 3456)
	str1 := buf.String()
	str2 := `<svg version="1.1" viewBox="234 232345092459 3456 3456" xmlns="http://www.w3.org/2000/svg">`
	if str1 != str2 {
		t.Errorf(fname+": expected %s recieved %s", str2, str1)
	}
}

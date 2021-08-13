package attr

type Composer interface {
	String() string
}

type Path struct {
	id    string
	d     []Composer
	style string
}

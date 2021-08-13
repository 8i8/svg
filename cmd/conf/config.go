package conf

import (
	"github.com/8i8/conf"
)

func Configure() (*conf.Config, error) {
	_, err := c.Compose(opts...)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

var (
	c       = conf.Config{}
	Default = c.Command(helpBase, helpDef)
	Open    = c.Command("open", helpOpen)
)

// {
// 	Type:     conf.Int,
// 	Flag:     "n",
// 	Default:  0,
// 	Usage:    "test int",
// 	Commands: Default | Open,
// 	Check: func(v interface{}) (interface{}, error) {
// 		i := *v.(*int)
// 		if i != 0 {
// 			return v, fmt.Errorf("-n is in testing")
// 		}
// 		return v, nil
// 	},
// },
var opts = []conf.Option{
	{
		Type:     conf.String,
		Flag:     "id",
		Default:  "",
		Usage:    "specify the id of the element to be targeted.",
		Commands: Default | Open,
	},
	{
		Type:     conf.Bool,
		Flag:     "handle",
		Default:  false,
		Usage:    "toggle handles on the target element.",
		Commands: Default | Open,
	},
}

var helpBase = `NAME
        svg

SYNOPSIS
        svg | [mode] | -[flag] | -[flag] <value> | -[flag] <'value,value,value'>

EXAMPLE
	svg open <fileName>
`

var helpDef = `MODES
        svg [mode] -[flag] <file>

	open     opens the given file in the default browser.

	Further information pertaining to the use of each mode can be
	had by running the following command.

	svg [mode] -help  or svg [mode] -h

FLAGS
`

var helpOpen = `MODE
	svg open
		svg open <filename>

FLAGS
`

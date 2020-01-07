package cycle

import (
	"github.com/hngo4619/leftpad"
)

var DEFAULT_CHAR = ' '

func FormatDouble(s string, i int) string {
	return leftpad.Format(s+s, i)
}

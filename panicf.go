package blank_identifier_alias_this_package

import (
	"fmt"
)

func Panicf(msg string, args ...any) {
	panic(fmt.Sprintf(msg, args...))
}

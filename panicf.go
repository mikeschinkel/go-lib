package lib

import (
	"fmt"
)

func Panicf(msg string, args ...any) {
	panic(fmt.Sprintf(msg, args...))
}

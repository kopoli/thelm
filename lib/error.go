package thelm

import (
	"fmt"
)

func ErrAnnotate(err error, a ...interface{}) error {
	s := fmt.Sprint(a...)
	return fmt.Errorf("%s: %w", s, err)
}

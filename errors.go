package cnv

import (
	"errors"
)

var (
	ErrNotAPointer     = errors.New("not a pointer")
	ErrUnsupportedType = errors.New("unsupported type")
)

package cnv

import (
	"errors"
)

var (
	// ErrNotAPointer variable is not a pointer.
	ErrNotAPointer = errors.New("not a pointer")

	// ErrUnsupportedType variable is of unsupported type.
	ErrUnsupportedType = errors.New("unsupported type")
)

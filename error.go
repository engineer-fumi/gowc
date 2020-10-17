package gowc

import "errors"

var (
	ErrInvalidType    = errors.New("the type is not defined or the wrong type is specified")
	ErrInvalidProgram = errors.New("A invalid program detected")
)

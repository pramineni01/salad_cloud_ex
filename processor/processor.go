package processor

import (
	"errors"
)

type Processor interface {
	Process()
}

func NewProcessor(source string, port uint) (Processor, error) {
	// not implemented
	return nil, errors.New("Not implemented")
}

package processor

import (
	"context"
	"errors"
)

type Processor interface {
	Process(context.Context)
}

func NewProcessor(source string, port uint) (Processor, error) {
	// not implemented
	return nil, errors.New("Not implemented")
}

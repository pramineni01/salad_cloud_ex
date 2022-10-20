package processor

import (
	"context"
	"errors"
)

type Processor interface {
	Process(context.Context)
}

func NewProcessor(source string, port uint) (Processor, error) {

	if (len(source) < 1) || (port <= 0) {
		return nil, errors.New("Invalid input")
	}

	return &processor{
		source: source,
		port: port,
	}, nil
}

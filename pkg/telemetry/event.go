package telemetry

import (
	"fmt"
)

type layer string

var (
	Transport layer = "Transport"
)

type operation string

var (
	Marshal   operation = "Marshal"
	Unmarshal operation = "Unmarshal"
	Envelop   operation = "Envelop"
	Request   operation = "Request"
	Response  operation = "Response"
	Publish   operation = "Publish"
	Subscribe operation = "Subscribe"
)

type event struct {
	Layer     layer
	Operation operation
	Err       error
}

func Event(layer layer, operation operation, err error) event {
	return event{layer, operation, err}
}

func (m event) Error() string {
	return fmt.Sprintf("%v [%v]: %v", m.Layer, m.Operation, m.Err)
}

func (m event) Unwrap() error {
	return m.Err
}

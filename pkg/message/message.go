package message

import (
	"encoding/json"
	"errors"

	"github.com/adamjohnston/agent/pkg/telemetry"
)

type message[T any] struct {
	Payload T      `json:"payload"`
	Error   string `json:"error,omitempty"`
}

func (m message[T]) Err() error {
	return errors.New(m.Error)
}

func marshal[T any](v T) ([]byte, error) {
	b, err := json.Marshal(&message[T]{Payload: v})
	if err != nil {
		return nil, telemetry.Event(telemetry.Transport, telemetry.Marshal, err)
	}
	return b, nil
}

func unmarshal[T any](b []byte) (T, error) {
	var m message[T]
	if err := json.Unmarshal(b, &m); err != nil {
		var zero T
		return zero, telemetry.Event(telemetry.Transport, telemetry.Marshal, err)
	}
	if err := m.Err(); err != nil {
		var zero T
		return zero, telemetry.Event(telemetry.Transport, telemetry.Envelop, err)
	}

	return m.Payload, nil
}

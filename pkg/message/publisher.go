package message

import (
	"context"

	"github.com/adamjohnston/agent/pkg/telemetry"
	"github.com/nats-io/nats.go"
)

type Publisher[I, O any] interface {
	Request(ctx context.Context, req I) (res O, err error)
	Publish(req I) (err error)
}

type publisher[I any, O any] struct {
	conn *nats.Conn
	subj string
}

func NewPublisher[I any, O any](conn *nats.Conn, subj string) Publisher[I, O] {
	return &publisher[I, O]{conn: conn, subj: subj}
}

func (p *publisher[I, O]) Request(ctx context.Context, req I) (O, error) {
	b, err := marshal(req)
	if err != nil {
		var res O
		return res, telemetry.Event(telemetry.Transport, telemetry.Request, err)
	}

	m, err := p.conn.RequestWithContext(ctx, p.subj, b)
	if err != nil {
		var res O
		return res, telemetry.Event(telemetry.Transport, telemetry.Request, err)
	}

	res, err := unmarshal[O](m.Data)
	if err != nil {
		var res O
		return res, telemetry.Event(telemetry.Transport, telemetry.Request, err)
	}

	return res, nil
}

func (p *publisher[I, O]) Publish(req I) error {
	b, err := marshal(req)
	if err != nil {
		return telemetry.Event(telemetry.Transport, telemetry.Publish, err)
	}
	if err := p.conn.Publish(p.subj, b); err != nil {
		return telemetry.Event(telemetry.Transport, telemetry.Publish, err)
	}
	return nil
}

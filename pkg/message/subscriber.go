package message

import (
	"encoding/json"

	"github.com/adamjohnston/agents/pkg/telemetry"
	"github.com/nats-io/nats.go"
)

type Subscriber[I, O any] interface {
	Subscribe(h Handler[I, O]) (*nats.Subscription, error)
	SubscribeWithMiddleware(h Handler[I, O], s ...Middleware[I, O]) (*nats.Subscription, error)
}

type subscriber[I any, O any] struct {
	conn *nats.Conn
	subj string
}

func NewSubscriber[I any, O any](conn *nats.Conn, subj string) Subscriber[I, O] {
	return &subscriber[I, O]{conn: conn, subj: subj}
}

func (p *subscriber[I, O]) Subscribe(h Handler[I, O]) (*nats.Subscription, error) {
	return p.conn.Subscribe(p.subj, func(msg *nats.Msg) {
		req, err := unmarshal[I](msg.Data)
		if err != nil {
			b, _ := json.Marshal(message[O]{Error: telemetry.Event(telemetry.Transport, telemetry.Subscribe, err).Error()})
			msg.Respond(b)
			return
		}

		res, err := h(req)
		if err != nil {
			b, _ := json.Marshal(message[O]{Error: telemetry.Event(telemetry.Transport, telemetry.Subscribe, err).Error()})
			msg.Respond(b)
			return
		}

		b, err := marshal(res)
		if err != nil {
			b, _ := json.Marshal(message[O]{Error: telemetry.Event(telemetry.Transport, telemetry.Subscribe, err).Error()})
			msg.Respond(b)
			return
		}

		msg.Respond(b)
	})
}

func (p *subscriber[I, O]) SubscribeWithMiddleware(h Handler[I, O], mws ...Middleware[I, O]) (*nats.Subscription, error) {
	return p.conn.Subscribe(p.subj, func(msg *nats.Msg) {
		req, err := unmarshal[I](msg.Data)
		if err != nil {
			b, _ := json.Marshal(message[O]{Error: telemetry.Event(telemetry.Transport, telemetry.Subscribe, err).Error()})
			msg.Respond(b)
			return
		}

		for i := len(mws) - 1; i >= 0; i-- {
			h = mws[i](h)
		}

		res, err := h(req)
		if err != nil {
			b, _ := json.Marshal(message[O]{Error: telemetry.Event(telemetry.Transport, telemetry.Subscribe, err).Error()})
			msg.Respond(b)
			return
		}

		b, err := marshal(res)
		if err != nil {
			b, _ := json.Marshal(message[O]{Error: telemetry.Event(telemetry.Transport, telemetry.Subscribe, err).Error()})
			msg.Respond(b)
			return
		}

		msg.Respond(b)
	})
}

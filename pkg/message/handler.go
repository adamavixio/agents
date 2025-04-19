package message

type Handler[I, O any] func(I) (O, error)
type Middleware[I, O any] func(Handler[I, O]) Handler[I, O]

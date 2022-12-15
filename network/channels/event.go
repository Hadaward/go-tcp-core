package channels

type Emitter struct {
	handlers map[string][]chan []any
}

func (e *Emitter) CallbackListen(name string, callback func([]any)) {
	channel := make(chan []any)

	go func() {
		for args := range channel {
			callback(args)
		}
	}()

	e.Listen(name, channel)
}

func (e *Emitter) CallbackListenChannel(name string, callback func([]any), channel chan []any) {
	go func() {
		for args := range channel {
			callback(args)
		}
	}()

	e.Listen(name, channel)
}

func (e *Emitter) Listen(name string, channel chan []any) {
	if e.handlers == nil {
		e.handlers = make(map[string][]chan []any)
	}

	e.handlers[name] = append(e.handlers[name], channel)
}

func (e *Emitter) Trigger(name string, args ...any) {
	if _, ok := e.handlers[name]; ok {
		for _, handler := range e.handlers[name] {
			go func(handler chan []any) {
				handler <- args
			}(handler)
		}
	}
}

package application

type Environment struct {
	Status  int
	Headers map[string]string
	Content []string
}

type Middleware interface {
	Name() string
	OnRequest(Environment) Environment
	OnResponse(Environment) Environment
}

type MiddlewareFunc func(Environment) Environment

type middleware struct {
	name       string
	onRequest  MiddlewareFunc
	onResponse MiddlewareFunc
}

func (middleware *middleware) Name() string {
	return middleware.name
}

func (middleware *middleware) OnRequest(env Environment) Environment {
	return middleware.onRequest(env)
}

func (middleware *middleware) OnResponse(env Environment) Environment {
	return middleware.onResponse(env)
}

func NewMiddleware(name string, onRequest MiddlewareFunc, onResponse MiddlewareFunc) Middleware {
	middleware := &middleware{name, onRequest, onResponse}
	return middleware
}

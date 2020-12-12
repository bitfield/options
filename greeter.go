package greeter

import "fmt"

type Greeter struct {
	Greeting string
	Message  string
}

type Option func(*Greeter)

func WithGreeting(greeting string) Option {
	return func(g *Greeter) {
		g.Greeting = greeting
	}
}

func WithMessage(message string) Option {
	return func(g *Greeter) {
		g.Message = message
	}
}

func New(opts ...Option) *Greeter {
	g := &Greeter{
		Greeting: "Hello",
		Message:  "Nice to meet you!",
	}
	for _, o := range opts {
		o(g)
	}
	return g
}

func (g Greeter) Greet(name string) {
	fmt.Printf("%s %s! %s\n", g.Greeting, name, g.Message)
}

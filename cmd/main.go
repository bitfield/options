package main

import "greeter"

func main() {
	g := greeter.New(
		greeter.WithGreeting("Hola"),
		greeter.WithMessage("Delighted to meet you."),
	)
	g.Greet("Rakesh")
}
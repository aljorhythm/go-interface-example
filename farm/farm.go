package farm

import "fmt"

type SomethingThatQuacks interface {
	quack() string
}

type Farm struct {
	somethingThatQuacks SomethingThatQuacks
}

// quackEverything makes the thing that can quack quack
func (f Farm) quackEverything() string {
	return fmt.Sprintf("something is quacking: %s", f.somethingThatQuacks.quack())
}

package farm

type RealDuck struct {
}

func (RealDuck) quack() string {
	return "QUACK!"
}

func (RealDuck) name() string {
	return "Duckie"
}

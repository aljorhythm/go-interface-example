package farm

import "testing"

func TestFarm_RealDuck(t *testing.T) {
	farm := Farm{somethingThatQuacks: RealDuck{}}
	res := farm.quackEverything()

	if res != "something is quacking: QUACK!" {
		t.Errorf(res)
	}

	t.Logf("res: %s", res)
}

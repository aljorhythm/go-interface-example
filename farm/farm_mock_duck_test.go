package farm

import "testing"

type MockDuck struct {
	mockQuack func() string
}

func (m MockDuck) quack() string {
	return m.mockQuack()
}

func TestFarm_MockDuck(t *testing.T) {
	mockDuck := MockDuck{
		mockQuack: func() string {
			return "fake quackkkkkk"
		},
	}
	farm := Farm{somethingThatQuacks: mockDuck}
	res := farm.quackEverything()

	if res != "something is quacking: fake quackkkkkk" {
		t.Errorf(res)
	}

	t.Logf("res: %s", res)
}

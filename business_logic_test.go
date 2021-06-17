package main

import "testing"

type MockDataRetriever struct {
	mockPeople []Person
}

func (m MockDataRetriever) GetAllPeople() []Person {
	return m.mockPeople
}

func Test_PrintAllPeople(t *testing.T) {
	mockRepo := MockDataRetriever{
		mockPeople: []Person{{
			"kenny",
		}, {
			"kelly",
		}},
	}
	logic := BusinessLogic{
		dataRetriever: mockRepo,
	}

	if logic.PrintAllPeople() != "kenny, kelly" {
		t.Errorf("Error")
	}
}

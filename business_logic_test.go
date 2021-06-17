package main

import "testing"

type MockDataRetriever struct {
}

func (m MockDataRetriever) GetAllPeople() []Person {
	return []Person{{
		"kenny",
	},{
		"kelly",
	}}
}

func Test_PrintAllPeople(t *testing.T) {
	mockRepo := MockDataRetriever {
	}
	logic := BusinessLogic{
		dataRetriever: mockRepo,
	}

	if logic.PrintAllPeople() != "kenny, kelly" {
		t.Errorf("Error")
	}
}
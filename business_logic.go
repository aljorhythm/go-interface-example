package main

import "strings"

type DataRetriever interface {
	GetAllPeople() []Person
}

type BusinessLogic struct {
	dataRetriever DataRetriever
}

func (logic BusinessLogic) PrintAllPeople() string {
	people := logic.dataRetriever.GetAllPeople()
	personStrings := []string{}
	for _, person := range people {
		personStrings = append(personStrings, person.Name)
	}
	return strings.Join(personStrings, ", ")
}

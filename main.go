package main

import "fmt"

func main() {
	repo := PersonRepo{}
	logic := BusinessLogic{
		dataRetriever: repo,
	}
	fmt.Println(logic.PrintAllPeople())
}

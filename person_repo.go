package main

type PersonRepo struct {
}

func (PersonRepo) FindById(id string) Person {
	return Person{}
}

func (PersonRepo) GetAllPeople() []Person {
	return []Person{}
}

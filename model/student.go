package model

type Student struct {
	Id int  `json:"id"`
	Name string  `json:"name"`
}

type UpdateStudent struct {
	Name string  `json:"name"`
}

type DeleteStudent struct {
	Name string  `json:"name"`
}
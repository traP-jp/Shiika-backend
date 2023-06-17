package main

type User struct {
	Id       string `json:"id,omitempty" db:"id"`
	Name     string `json:"name,omitempty" db:"name"`
	Password string `json:"password,omitempty"  db:"password"`
}
type Kaminoku struct {
	Id      string `json:"id,omitempty" db:"id"`
	Content string `json:"name,omitempty" db:"content"`
	Userid  int    `json:"userid,omitempty" db:"userid"`
}
type Simonoku struct {
	Id         string `json:"id,omitempty" db:"id"`
	Content    string `json:"name,omitempty" db:"content"`
	KaminokuId int    `json:"kaminokuid,omitempty" db:"kaminokuid"`
	Userid     int    `json:"userid,omitempty" db:"userid"`
}

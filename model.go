package main

type User struct {
	Id       int    `json:"id,omitempty" db:"id"`
	Name     string `json:"name,omitempty" db:"name"`
	Password string `json:"password,omitempty"  db:"password"`
}
type Kaminoku struct {
	Id      int    `json:"id,omitempty" db:"id"`
	Content string `json:"name,omitempty" db:"name"`
	Userid  int    `json:"userid,omitempty" db:"userid"`
}
type Shimonoku struct {
	Id         int    `json:"id,omitempty" db:"id"`
	Content    string `json:"name,omitempty" db:"name"`
	KaminokuId int    `json:"kaminokuid,omitempty" db:"kaminokuid"`
	Userid     int    `json:"userid,omitempty" db:"userid"`
}

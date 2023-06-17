package main

type User struct {
	Name     string `json:"name,omitempty" db:"name"`
	Password string `json:"password,omitempty"  db:"password"`
}
type Kaminoku struct {
	Id      string `json:"id,omitempty" db:"id"`
	Content string `json:"name,omitempty" db:"content"`
	Userid  string `json:"userid,omitempty" db:"userid"`
}
type LoginRequestBody struct {
	Username string `json:"userid,omitempty" form:"username"`
	Password string `json:"password,omitempty" form:"password"`
}

// for post kaminoku's request body's binding
type KaminokuReq struct {
	Content string `json:"content,omitempty"`
}

// for post simonoku's request body's binding
type SimonokuReq struct {
	Content string `json:"content,omitempty"`
}

type TankaRes struct {
	Id           string `json:"id,omitempty" db:"id"`
	Kaminoku     string `json:"kaminoku,omitempty" db:"kaminoku"`
	KaminokuUser string `json:"kaminokuuser,omitempty" db:"kaminokuuser"`
	Simonoku     string `json:"simonoku,omitempty"  db:"simonoku"`
	SimonokuUser string `json:"simonokuuser,omitempty" db:"simonokuuser"`
}

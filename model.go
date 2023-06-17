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

// for post kaminoku's request body's binding
type KaminokuReq struct {
	Content string `json:"content,omitempty"`
}

// kaminoku for response
type KaminokuRes struct {
	Id       string `json:"id,omitempty"`
	Content  string `json:"name,omitempty"`
	Userid   int    `json:"userid,omitempty"`
	UserName string `json:"username,omitempty"`
}

// for post simonoku's request body's binding
type SimonokuReq struct {
	Content string `json:"content,omitempty"`
}

//
type SimonokuRes struct {
	Id         string `json:"id,omitempty"`
	Content    string `json:"name,omitempty"`
	KaminokuId int    `json:"kaminokuid,omitempty"`
	Userid     int    `json:"userid,omitempty"`
	UserName   string `json:"username,omitempty"`
}

type TankaRes struct {
	Kaminoku KaminokuRes
	Simonoku SimonokuRes
}

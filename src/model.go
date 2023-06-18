package main

type User struct {
	Name     string `json:"name,omitempty" db:"name"`
	Password string `json:"password,omitempty"  db:"password"`
}
type Kaminoku struct {
	Id      string `json:"id,omitempty" db:"id"`
	Content Fsf    `json:"content,omitempty" db:"content"`
	Userid  string `json:"userid,omitempty" db:"userid"`
}
type Simonoku struct {
	Id      string `json:"id,omitempty" db:"id"`
	Content Ss     `json:"content,omitempty" db:"content"`
	Userid  string `json:"userid,omitempty" db:"userid"`
}
type LoginRequestBody struct {
	Username string `json:"userid,omitempty" form:"userid"`
	Password string `json:"password,omitempty" form:"password"`
}

type KaminokuReq struct {
	Content Fsf `json:"content,omitempty"`
}

type SimonokuReq struct {
	Content Ss `json:"content,omitempty"`
}

type Kaminokudb struct {
	Id     string `db:"id"`
	First  string `db:"first"`
	Second string `db:"second"`
	Third  string `db:"third"`
	Userid string `db:"userid"`
}
type Simonokudb struct {
	Simonokuid     string `db:"simonokuid"`
	Kaminokuid     string `db:"kaminokuid"`
	First          string `db:"first"`
	Second         string `db:"second"`
	Third          string `db:"third"`
	Fourth         string `db:"fourth"`
	Fifth          string `db:"fifth"`
	Kaminokuuserid string `db:"kaminokuuser"`
	Simonokuuserid string `db:"simonokuuser"`
}

type Fsf struct {
	First  string `json:"first,omitempty"`
	Second string `json:"second,omitempty"`
	Third  string `json:"third,omitempty"`
}

type Ss struct {
	Fourth string `json:"fourth,omitempty"`
	Fifth  string `json:"fifth,omitempty"`
}

type TankaRes struct {
	Kaminoku Kaminoku `json:"kaminoku,omitempty"`
	Simonoku Simonoku `json:"simonoku,omitempty"`
}

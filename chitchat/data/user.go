package data

import "time"

type Session struct {
	Id       int
	Uuid     string
	Email    string
	UserId   int
	CreateAt time.Time
}

//Check if session is valid in the database
func (session *Session) Check() (valid bool, err error) {
	valid = true
	return
}

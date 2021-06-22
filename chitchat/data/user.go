package data

import (
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

//Create a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
	session = Session{}
	statement := "insert into sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) return id, uuid, email, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	//FIXME: Session is not saved to the database
	err = stmt.QueryRow(createUUID(), user.Email, user.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

//Create a new user and save it into the database
func (user *User) Create() (err error) {
	statement := "insert into users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	//return a row and scan the returned id into the User struct
	err = stmt.QueryRow(createUUID(), user.Name, user.Email, Encrypt(user.Password), time.Now()).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return
}

//Check if session is valid in the database
func (session *Session) Check() (valid bool, err error) {
	valid = true
	return
}

func DeleteSessionsFromDatabase() (err error) {
	_, err = Db.Exec("delete from threads")
	return
}

func DeleteUsersFromDatabase() (err error) {
	_, err = Db.Exec("delete from users")
	return
}

func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email=$1", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

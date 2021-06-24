package data

import (
	"database/sql"
	"testing"
)

func Test_UserCreate(t *testing.T) {
	setup()

	user := users[0]

	if err := user.Create(); err != nil {
		t.Error(err, "Cannot create user")
	}

	if user.Id == 0 {
		t.Errorf("No id in user")
	}

	fetchedUser, err := UserByEmail(user.Email)
	if err != nil {
		t.Error(err, "User not created")
	}

	if user.Email != fetchedUser.Email {
		t.Errorf("Retrieved user is not the same as the one created")
	}
}

func Test_UserDelete(t *testing.T) {
	setup()
	if err := testUser.Create(); err != nil {
		t.Error(err, "Cannot create user")
	}
	if err := testUser.Delete(); err != nil {
		t.Error(err, "- Cannot detele user")
	}
	_, err := UserByEmail(testUser.Email)
	if err != sql.ErrNoRows {
		t.Error(err, "- User not deleted")
	}
}

func Test_CreateSession(t *testing.T) {
	setup()
	if err := testUser.Create(); err != nil {
		t.Error(err, "Cannot create user")
	}
	session, err := testUser.CreateSession()
	if err != nil {
		t.Error(err, "Cannot create session")
	}
	if session.UserId != testUser.Id {
		t.Error("User not linked with session")
	}
}

func Test_GetSession(t *testing.T) {
	setup()
	if err := testUser.Create(); err != nil {
		t.Error(err, "Cannot create user")
	}
	session, err := testUser.CreateSession()
	if err != nil {
		t.Error(err, "Cannot create session")
	}

	s, err := testUser.Session()
	if err != nil {
		t.Error(err, "Cannot get session")
	}
	if s.Id == 0 {
		t.Error("No session retrieved")
	}
	if s.Id != session.Id {
		t.Error("Different session retrieved")
	}
}

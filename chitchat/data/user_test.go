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

func Test_UserUpdate(t *testing.T) {
	setup()
	if err := testUser.Create(); err != nil {
		t.Error(err, "Cannot create user")
	}
	testUser.Name = "Random"
	if err := testUser.Update(); err != nil {
		t.Error(err, "- Cannot update user")
	}
	u, err := UserByEmail(testUser.Email)
	if err != nil {
		t.Error(err, "- Cannot get user")
	}
	if u.Name != "Random" {
		t.Error(err, "- User not updated")
	}
}

func Test_UserByUUID(t *testing.T) {
	setup()
	if err := testUser.Create(); err != nil {
		t.Error(err, "Cannot create user")
	}
	u, err := UserByUUID(testUser.Uuid)
	if err != nil {
		t.Error(err, "User not created")
	}
	if testUser.Email != u.Email {
		t.Errorf("User retrieved is not the same as the one created")
	}
}

func Test_Users(t *testing.T) {
	setup()
	for _, user := range users {
		if err := user.Create(); err != nil {
			t.Error(err, "Cannot create user")
		}
	}
	u, err := Users()
	if err != nil {
		t.Error(err, "Cannot retrieve users")
	}
	if len(u) != 2 {
		t.Error(err, "Wrong number of users retrieved")
	}
	if u[0].Email != users[0].Email {
		t.Error(u[0], users[0], "Wrong user retrieved")
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

//More test here

func Test_DeleteSession(t *testing.T) {
	setup()
	if err := testUser.Create(); err != nil {
		t.Error(err, "Cannot create user")
	}
	session, err := testUser.CreateSession()
	if err != nil {
		t.Error(err, "Cannot create session")
	}
	err = session.DeleteByUUID()
	if err != nil {
		t.Error(err, "Cannot delete session")
	}
	s := Session{Uuid: session.Uuid}
	valid, err := s.Check()
	if err == nil {
		t.Error(err, "Session is valid even though deleted")
	}
	if valid == true {
		t.Error(err, "Session is not deleted")
	}
}

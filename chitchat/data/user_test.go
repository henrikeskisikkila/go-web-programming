package data

import "testing"

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

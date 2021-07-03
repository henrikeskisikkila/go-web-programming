package data

import "testing"

func DeleteThreadsFromDatabase() (err error) {
	_, err = Db.Exec("delete from threads")
	return
}

func Test_CreateThread(t *testing.T) {
	setup()
	if err := testUser.Create(); err != nil {
		t.Error(err, "Cannot create user")
	}

	conv, err := testUser.CreateThread("My first thread")
	if err != nil {
		t.Error(err, "Cannot create thread")
	}
	if conv.UserId != testUser.Id {
		t.Error("User not linked with thread")
	}
}

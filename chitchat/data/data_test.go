package data

//Test data
var users = []User{
	{
		Name:     "Peter Jones",
		Email:    "peter@jones.com",
		Password: "secret",
	},
	{
		Name:     "Lisa Wood",
		Email:    "lisa@wood.com",
		Password: "secret",
	},
}

var testUser = users[0]

func setup() {
	DeleteThreadsFromDatabase()
	DeleteSessionsFromDatabase()
	DeleteUsersFromDatabase()
}

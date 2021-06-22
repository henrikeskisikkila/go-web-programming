package data

func DeleteThreadsFromDatabase() (err error) {
	_, err = Db.Exec("delete from threads")
	return
}

package data

import "time"

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

func Threads() (threads []Thread, err error) {
	var results []Thread
	thread := Thread{1, "uiid", "Topic", 2, time.Now()}
	results = append(results, thread)
	return results, nil
}

func Hello() string {
	return "hello"
}

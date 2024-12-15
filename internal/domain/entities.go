package domain

type User struct {
	Uuid     string
	Username string
	Password string
}

type Thread struct {
	Uuid   string
	Topic  string
	UserId int
}

type Post struct {
	Uuid     string
	Body     string
	ThreadId int
	UserId   int
}

package proto

type User struct {
	ID       int64
	Nickname string
}

type DrawTab struct {
	ID   int64
	User User
}

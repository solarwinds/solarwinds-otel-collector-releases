package loggedusers

type User struct {
	Name        string // Name used to log in to the system
	DisplayName string // Windows display name of the user
	TTY         string // Linux tty from which session is initiated
}

type Data struct {
	Users []User
	Error error
}

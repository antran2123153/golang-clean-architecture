package presenter

type User struct {
	ID    string
	Name  string
	Email string
}

type CreateUser struct {
	Name  string
	Email string
}

type UpdateUser struct {
	Name  string
	Email string
}

type ListUser struct {
	Users []User
	Count int
}

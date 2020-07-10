package main

type user struct {
	name  string
	email string
}

func (u *user) changeEmail(email string) {
	u.email = email
}

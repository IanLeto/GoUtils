package testsuit

import "GoUtils/utils"

type Person interface {
	Get() error
}

type User struct {
	id int
}

func (u *User) Get() error {
	return nil
}

func Invoke(person Person) {
	utils.NoErr(person.Get())
}

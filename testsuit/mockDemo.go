package testsuit

type User struct {
	id int
}

func (u *User) Get() {

}

func (u *User) Invoke() {
	u.Get()
}

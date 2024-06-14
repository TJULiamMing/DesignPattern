package iterator

type Iterator interface {
	HasNext() bool
	Next() *User
}

type UserIterator struct {
	Index int
	Users []*User
}

func (u *UserIterator) HasNext() bool {
	if u.Index >= len(u.Users) {
		return false
	}
	return true
}

func (u *UserIterator) Next() *User {
	if u.HasNext() {
		user := u.Users[u.Index]
		u.Index++
		return user
	}
	return nil
}

package iterator

type Collection interface {
	NewIterator() Iterator
}

type User struct {
	Name string
}

type UserCollection struct {
	Users []*User
}

func (u *UserCollection) NewIterator() Iterator {
	return &UserIterator{
		Users: u.Users,
	}
}

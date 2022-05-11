package users

type UserRepo struct {
	isActive bool
	Users    []*User
}

type IUserRepo interface {
	CreateUser(user User)
	GetUser(name string)
	UpdateUser()
}

var usrDB *UserRepo

func GetUserRepo() *UserRepo {
	if usrDB == nil || usrDB.isActive == false {
		usrDB = &UserRepo{
			isActive: true,
		}
		return usrDB

	}

	return usrDB

}

func (usrR *UserRepo) CreateUser(user *User) {
	usrR.Users = append(usrR.Users, user)
}

func (usrR *UserRepo) GetUser(name string) *User {
	for _, u := range usrR.Users {
		if u.Name == name {
			return u
		}
	}
	return nil
}
func (usrR *UserRepo) UpdateUser(user *User) {
	for _, u := range usrR.Users {
		if u.Name == user.Name {
			u.Playlist = user.Playlist
		}
	}
}

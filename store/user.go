package store

type User struct {
	Username       string
	HashedPassword string
	Role           string
}

func (user *User) Clone() *User {
	return &User{
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
		Role:           user.Role,
	}
}

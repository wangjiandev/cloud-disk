package model

type UserBasic struct {
	Id       int
	Identity string
	Name     string
	Password string
	Email    string
}

func (u *UserBasic) TabelName() string {
	return "user_basic"
}
